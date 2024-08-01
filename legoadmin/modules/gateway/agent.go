package gateway

import (
	"context"
	"fmt"
	"legoadmin/pb"
	"sync"
	"sync/atomic"
	"time"

	"github.com/liwei1dao/lego/sys/log"

	"github.com/gorilla/websocket"
	"github.com/liwei1dao/lego/utils/container/id"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

/*
用户代理对象
封装用户socket 对象 处理用户的消息读取 写入 关闭等操作
*/

func newAgent(gateway IGateway, conn *websocket.Conn) *Agent {
	agent := &Agent{
		gateway:     gateway,
		wsConn:      conn,
		sessionId:   id.NewXId(),
		uId:         "",
		writeChan:   make(chan []byte, 2),
		closeSignal: make(chan bool),
		state:       1,
	}
	agent.wg.Add(2)
	go agent.readLoop()
	go agent.writeLoop()
	return agent
}

// 用户代理
type Agent struct {
	gateway      IGateway
	wsConn       *websocket.Conn
	sessionId    string
	uId          string
	spath        string
	writeChan    chan []byte
	closeSignal  chan bool
	state        int32 //状态 0 关闭 1 运行 2 关闭中
	wg           sync.WaitGroup
	queueIndex   int32     //排队编号
	lastpushtime time.Time //上次推送时间
}

func (this *Agent) readLoop() {
	defer this.wg.Done()
	var (
		data       []byte
		msgpackage *pb.MessagePackage = &pb.MessagePackage{}
		err        error
	)
locp:
	for {
		if _, data, err = this.wsConn.ReadMessage(); err != nil || len(data) == 0 {
			this.gateway.Errorf("agent:%s uId:%s ReadMessage err:%v", this.sessionId, this.uId, err)
			go this.Close()
			break locp
		}

		if err = proto.Unmarshal(data, msgpackage); err != nil {
			this.gateway.Errorf("agent:%s uId:%s Unmarshal err:%v", this.sessionId, this.uId, err)
			go this.Close()
			break locp
		} else {
			this.wsConn.SetReadDeadline(time.Now().Add(time.Second * 60))
			for _, msg := range msgpackage.Messages {
				if msg.MsgName == comm.Msg_GatewayHeartbeat { //心跳消息 无需校验秘钥
					data, _ := anypb.New(&pb.GatewayHeartbeatResp{
						Timestamp: time.Now().Unix(),
					})
					this.WriteMsgs(&pb.UserMessage{
						MsgName: comm.Msg_GatewayHeartbeat,
						Data:    data,
					})
					continue
				}

				if this.uId == "" && msg.MsgName != "user/login" {
					data, _ := anypb.New(&pb.GatewayErrorNotifyPush{
						MsgName:     msg.MsgName,
						ServicePath: msg.ServicePath,
						Req:         msg.Data,
						Err:         &pb.ErrorData{Code: pb.ErrorCode_NoLogin, Message: pb.ErrorCode_NoLogin.String()},
					})
					err = this.WriteMsgs(&pb.UserMessage{
						MsgName:     "gateway/errornotify",
						ServicePath: msg.ServicePath,
						Data:        data,
					})
					continue
				}

				if this.gateway.IsGameOpen(msg.ServicePath) { //游戏是否开启
					if err = this.messageDistribution(msg); err != nil {
						this.gateway.Errorf("messageDistribution err:%v", err)
						data, _ := anypb.New(&pb.GatewayErrorNotifyPush{
							MsgName:     msg.MsgName,
							ServicePath: msg.ServicePath,
							Req:         msg.Data,
							Err:         &pb.ErrorData{Code: pb.ErrorCode_GatewayException, Message: err.Error()},
						})
						err = this.WriteMsgs(&pb.UserMessage{
							MsgName:     comm.Msg_GatewayHeartbeat,
							ServicePath: msg.ServicePath,
							Data:        data,
						})
						go this.Close()
						break locp
					}
				} else {
					data, _ := anypb.New(&pb.GatewayErrorNotifyPush{
						MsgName:     msg.MsgName,
						ServicePath: msg.ServicePath,
						Req:         msg.Data,
						Err:         &pb.ErrorData{Code: pb.ErrorCode_GameInMaintenance, Message: pb.ErrorCode_GameInMaintenance.String()},
					})
					err = this.WriteMsgs(&pb.UserMessage{
						MsgName:     "gateway/errornotify",
						ServicePath: msg.ServicePath,
						Data:        data,
					})
				}
			}
		}
	}
	this.gateway.Debugf("agent:%s uId:%s readLoop end!", this.sessionId, this.uId)
}

func (this *Agent) writeLoop() {
	defer this.wg.Done()
	var (
		// data []byte
		err error
	)
locp:
	for {
		select {
		case <-this.closeSignal:
			break locp
		case msgs, ok := <-this.writeChan:
			if ok {
				//data, err = proto.Marshal(msg)
				if err = this.wsConn.WriteMessage(websocket.BinaryMessage, msgs); err != nil {
					this.gateway.Errorf("agent:%s uId:%s WriteMessage err:%v", this.sessionId, this.uId, err)
					go this.Close()
				}
			} else {
				go this.Close()
			}
		}
	}
	this.gateway.Debugf("agent:%s uId:%s writeLoop end!", this.sessionId, this.uId)
}

func (this *Agent) SessionId() string {
	return this.sessionId
}

func (this *Agent) IP() string {
	return this.wsConn.RemoteAddr().String()
}
func (this *Agent) UserId() string {
	return this.uId
}
func (this *Agent) AgentId() string {
	return this.aId
}
func (this *Agent) GameId() string {
	return this.gid
}
func (this *Agent) ServicePath() string {
	return this.spath
}

func (this *Agent) GetSessionData() *pb.UserSessionData {
	return &pb.UserSessionData{
		SessionId:  this.sessionId,
		AgentId:    this.aId,
		UserId:     this.uId,
		Ip:         this.IP(),
		ServiceTag: this.gateway.Service().GetTag(),
		GatewayId:  this.gateway.Service().GetId(),
	}
}

func (this *Agent) WriteMsgs(msgs ...*pb.UserMessage) (err error) {
	if atomic.LoadInt32(&this.state) != 1 {
		return
	}
	var (
		msgpackage *pb.MessagePackage = &pb.MessagePackage{}
		data       []byte
	)
	msgpackage.Messages = msgs
	if data, err = proto.Marshal(msgpackage); err != nil {
		return
	}
	this.writeChan <- data
	return
}

func (this *Agent) WriteBytes(data []byte) (err error) {
	if atomic.LoadInt32(&this.state) != 1 {
		err = fmt.Errorf("Uid%s Staet:%d", this.uId, this.state)
		return
	}
	this.writeChan <- data
	return
}

// 外部代用关闭
func (this *Agent) Close() {
	if !atomic.CompareAndSwapInt32(&this.state, 1, 2) {
		return
	}
	this.wsConn.Close()
	this.closeSignal <- true
	this.wg.Wait()
	atomic.StoreInt32(&this.state, 0)
	this.gateway.DisConnect(this)
}

// 处理用户消息 提供给外部使用 比如 登录等待逻辑
func (this *Agent) HandleMessage(msg *pb.UserMessage) (err error) {
	if err = this.messageDistribution(msg); err != nil {
		this.gateway.Errorf("messageDistribution err:%v", err)
		data, _ := anypb.New(&pb.GatewayErrorNotifyPush{
			MsgName:     msg.MsgName,
			ServicePath: msg.ServicePath,
			Req:         msg.Data,
			Err:         &pb.ErrorData{Code: pb.ErrorCode_GatewayException, Message: err.Error()},
		})
		err = this.WriteMsgs(&pb.UserMessage{
			MsgName: comm.Msg_GatewayHeartbeat,
			Data:    data,
		})
	}
	return
}

// 分发用户消息
func (this *Agent) messageDistribution(msg *pb.UserMessage) (err error) {
	var (
		spath string              = this.spath
		req   *pb.AgentMessage    = getmsg()
		reply *pb.RPCMessageReply = getmsgreply()
	)
	defer func() {
		putmsg(req)
		putmsgreply(reply)
	}()
	req.UserSession.Ip = this.IP()
	req.UserSession.SessionId = this.sessionId
	req.UserSession.AgentId = this.aId
	req.UserSession.UserId = this.uId
	req.UserSession.ServiceTag = this.gateway.Service().GetTag()
	req.UserSession.GatewayId = this.gateway.Service().GetId()
	req.MsgName = msg.MsgName
	req.Message = msg.Data
	stime := time.Now()
	if spath == "" {
		spath = msg.ServicePath
	}
	if req.MsgName == "" {
		err = fmt.Errorf("no MsgName!")
		this.gateway.Error("[UserResponse]",
			log.Field{Key: "uid", Value: this.uId},
			log.Field{Key: "serviceTag", Value: req.UserSession.ServiceTag},
			log.Field{Key: "servicePath", Value: msg.ServicePath},
			log.Field{Key: "req", Value: fmt.Sprintf("%s:%v", req.MsgName, req.Message.String())},
			log.Field{Key: "err", Value: err.Error()},
		)
		return
	}

	if err = this.gateway.Service().RpcCall(context.Background(), spath, string(comm.Rpc_GatewayRoute), req, reply); err != nil {
		this.gateway.Error("[UserResponse]",
			log.Field{Key: "uid", Value: this.uId},
			log.Field{Key: "serviceTag", Value: req.UserSession.ServiceTag},
			log.Field{Key: "servicePath", Value: msg.ServicePath},
			log.Field{Key: "req", Value: fmt.Sprintf("%s:%v", req.MsgName, req.Message.String())},
			log.Field{Key: "err", Value: err.Error()},
		)
		return
	}

	this.gateway.Debug("[UserResponse]",
		log.Field{Key: "t", Value: time.Since(stime).Milliseconds()},
		log.Field{Key: "uid", Value: this.uId},
		log.Field{Key: "req", Value: fmt.Sprintf("%s:%v", req.MsgName, req.Message.String())},
		log.Field{Key: "reply", Value: reply.String()},
	)

	if reply.ErrorData != nil {
		data, _ := anypb.New(&pb.GatewayErrorNotifyPush{
			MsgName: msg.MsgName,
			Req:     msg.Data,
			Err:     reply.ErrorData})
		err = this.WriteMsgs(&pb.UserMessage{
			MsgName: comm.Msg_GatewayErrornotify,
			Data:    data,
		})
		return
	} else {
		for _, v := range reply.Reply {
			if v.MsgName == msg.MsgName && v.MsgName == comm.Msg_UserLogin {
				var (
					resp      proto.Message
					loginresp *pb.UserLoginResp
				)
				if resp, err = v.Data.UnmarshalNew(); err != nil {
					return
				}
				loginresp = resp.(*pb.UserLoginResp)
				this.uId = loginresp.Uid
				this.aId = loginresp.Agentid
				this.gid = loginresp.Gameid
				this.spath = reply.ServicePath
				this.gateway.LoginNotice(this)
			}
		}
		if err = this.WriteMsgs(reply.Reply...); err != nil {
			return
		}
	}
	return nil
}
