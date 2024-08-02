package comm

import (
	"context"
	"fmt"
	"legoadmin/pb"
	"sync"

	"github.com/liwei1dao/lego/sys/log"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

func NewUserSession() IUserSession {
	return &UserContext{
		userCache: &pb.UserSessionData{},
		msgqueue:  make([]*pb.UserMessage, 0),
	}
}

// 用户会话
type UserContext struct {
	context.Context
	userCache *pb.UserSessionData
	service   IService
	msgqueue  []*pb.UserMessage
	lock      sync.RWMutex
	mate      map[string]interface{}
	offline   bool
}

// 重置
func (this *UserContext) SetSession(ctx context.Context, service IService, cache *pb.UserSessionData) {
	this.Context = ctx
	this.service = service
	this.userCache.SessionId = cache.SessionId
	this.userCache.UserId = cache.UserId
	this.userCache.Ip = cache.Ip
	this.userCache.GatewayId = cache.GatewayId
	this.msgqueue = this.msgqueue[:0]
	this.mate = make(map[string]interface{})
}

// 重置
func (this *UserContext) Reset() {
	this.userCache.Ip = ""
	this.userCache.SessionId = ""
	this.userCache.GatewayId = ""
	this.userCache.UserId = ""
	this.msgqueue = this.msgqueue[:0]
	this.mate = make(map[string]interface{})
}

// 获取用户的会话id
func (this *UserContext) GetCache() *pb.UserSessionData {
	return this.userCache
}

// 获取用户id
func (this *UserContext) GetUserId() string {
	return this.userCache.UserId
}

func (this *UserContext) IsOnline() bool {
	if this.userCache.UserId == "" {
		return false
	}
	return true
}

func (this *UserContext) SetOffline(offline bool) {
	this.offline = offline
}
func (this *UserContext) GetOffline() bool {
	return this.offline
}

// 解绑uid 注销和切换账号是处理
func (this *UserContext) UnBind() (err error) {
	if err = this.service.RpcCall(context.Background(), fmt.Sprintf("%s/%s", Service_Gateway, this.userCache.GatewayId), string(Rpc_GatewayAgentUnBind), &pb.Rpc_GatewayAgentUnBindReq{
		UserSessionId: this.userCache.SessionId,
	}, nil); err != nil {
		log.Errorf("UnBuild UserSession:%s UserId:%s err:%v", this.userCache.SessionId, this.userCache.UserId, err)
		return
	}
	this.userCache.UserId = ""
	return
}

// 向用户发送消息
func (this *UserContext) SendMsg(mainType, subType string, msg proto.Message) (err error) {
	// log.Debugf("SendMsg to UserId:[%s] Data: %v", this.UserId, msg)
	data, _ := anypb.New(msg)
	this.msgqueue = append(this.msgqueue, &pb.UserMessage{
		MsgName: fmt.Sprintf("%s/%s", mainType, subType),
		Data:    data,
	})
	return
}

// 关闭用户连接对象
func (this *UserContext) Close() (err error) {
	if _, err = this.service.RpcGo(context.Background(), fmt.Sprintf("%s/%s", Service_Gateway, this.userCache.GatewayId), string(Rpc_GatewayAgentClose), &pb.Rpc_GatewayAgentCloseReq{
		UserSessionId: this.userCache.SessionId,
	}, nil); err != nil {
		log.Errorf("Close UserSession:%s UserId:%s err:%v", this.userCache.SessionId, this.userCache.UserId, err)
	}
	return
}

// 清空消息队列
func (this *UserContext) Polls() []*pb.UserMessage {
	msgs := this.msgqueue
	this.msgqueue = this.msgqueue[:0]
	return msgs
}

// 推送消息到用户
func (this *UserContext) Push() (err error) {
	if len(this.msgqueue) > 0 {
		if _, err = this.service.RpcGo(context.Background(), fmt.Sprintf("%s/%s", Service_Gateway, this.userCache.GatewayId), string(Rpc_GatewayAgentSendMsg), &pb.Rpc_GatewayAgentSendMsgReq{
			UserSessionId: this.userCache.SessionId,
			Reply:         this.msgqueue,
		}, nil); err != nil {
			log.Errorf("Push:%v err:%s", this, err.Error())
		}
	}
	this.msgqueue = this.msgqueue[:0]
	return
}

func (this *UserContext) SyncPush() (err error) {
	if len(this.msgqueue) > 0 {
		if err = this.service.RpcCall(context.Background(), fmt.Sprintf("%s/%s", Service_Gateway, this.userCache.GatewayId), string(Rpc_GatewayAgentSendMsg), &pb.Rpc_GatewayAgentSendMsgReq{
			UserSessionId: this.userCache.SessionId,
			Reply:         this.msgqueue,
		}, &pb.Rpc_GatewayAgentSendMsgResp{}); err != nil {
			log.Errorf("SendMsgToUsers:%v err:%v", this, err)
		}
	}
	this.msgqueue = this.msgqueue[:0]
	return
}

// 写入元数据
func (this *UserContext) SetMate(name string, value interface{}) {
	this.lock.Lock()
	this.mate[name] = value
	this.lock.Unlock()
}

// 写入元数据
func (this *UserContext) GetMate(name string) (ok bool, value interface{}) {
	this.lock.RLock()
	value, ok = this.mate[name]
	this.lock.RUnlock()
	return
}

// 克隆
func (this *UserContext) Clone() (session IUserSession) {
	session = this.service.GetUserSession(this.userCache)
	this.lock.RLock()
	for k, v := range this.mate {
		session.SetMate(k, v)
	}
	this.lock.RUnlock()
	return
}
