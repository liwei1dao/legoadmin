package gateway

import (
	"context"
	"fmt"
	"legoadmin/comm"
	"legoadmin/pb"
	"sync"
	"sync/atomic"

	"github.com/liwei1dao/lego/base"
	"github.com/liwei1dao/lego/core"
	"github.com/liwei1dao/lego/core/cbase"
	"github.com/liwei1dao/lego/sys/log"

	"google.golang.org/protobuf/proto"
)

/*
用户代理对象管理组件
*/

type AgentMgrComp struct {
	cbase.ModuleCompBase
	options    *Options
	service    base.IRPCXService
	module     *Gateway
	agents     *sync.Map
	users      *sync.Map
	onlineuser int32
}

func (this *AgentMgrComp) Init(service core.IService, module core.IModule, comp core.IModuleComp, options core.IModuleOptions) (err error) {
	err = this.ModuleCompBase.Init(service, module, comp, options)
	this.options = options.(*Options)
	this.service = service.(base.IRPCXService)
	this.module = module.(*Gateway)
	this.agents = new(sync.Map)
	this.users = new(sync.Map)
	return
}

func (this *AgentMgrComp) Start() (err error) {
	err = this.ModuleCompBase.Start()
	return
}

func (this *AgentMgrComp) getAgent(sid string) (agent IAgent) {
	var (
		a  any
		ok bool
	)
	if a, ok = this.agents.Load(sid); !ok {
		return
	}
	agent = a.(IAgent)
	return
}

func (this *AgentMgrComp) getonlineNum() int32 {
	return atomic.LoadInt32(&this.onlineuser)
}

// Connect 加入新的用户
func (this *AgentMgrComp) Connect(a IAgent) {
	this.agents.Store(a.SessionId(), a)
}

// DisConnect 移除断开的用户
func (this *AgentMgrComp) DisConnect(a IAgent) {
	if a.(IAgent).UserId() != "" {
		this.users.Delete(a.(IAgent).UserId())
	}

	this.agents.Delete(a.SessionId())
	if a.UserId() != "" { //登录用户 通知业务服务处理玩家离线相关
		atomic.AddInt32(&this.onlineuser, -1)
		if _, err := this.service.RpcGo(context.Background(), a.ServicePath(), string(comm.Rpc_GatewayNoticeUserClose), &pb.NoticeUserCloseReq{
			UserSession: a.GetSessionData(),
		}, nil); err != nil {
			log.Errorf("uId:%s Rpc_NoticeUserClose err:%v", a.UserId(), err)
		}
	}
}

func (this *AgentMgrComp) Logined(a IAgent) {
	this.users.Store(a.UserId(), a)
	atomic.AddInt32(&this.onlineuser, 1)
}

// SendMsgToAgent 向用户发送消息
func (this *AgentMgrComp) SendMsgToAgent(ctx context.Context, args *pb.AgentSendMessageReq, reply *pb.RPCMessageReply) error {
	this.module.Debugf("SendMsgToAgent: agent:%s msg:%v", args.UserSessionId, args.Reply)
	if a, ok := this.agents.Load(args.UserSessionId); ok {
		a.(IAgent).WriteMsgs(args.Reply...)
	} else {
		reply.ErrorData = &pb.ErrorData{
			Code:    pb.ErrorCode_UserSessionNobeing,
			Message: fmt.Sprintf("解绑SessionId:%s失败!", args.UserSessionId),
		}
	}
	return nil
}

// SendMsgToAgents 向多个户发送消息
func (this *AgentMgrComp) SendMsgToAgents(ctx context.Context, args *pb.BatchMessageReq, reply *pb.RPCMessageReply) (err error) {
	var (
		data []byte
	)
	msg := &pb.MessagePackage{
		Messages: make([]*pb.UserMessage, 0),
	}
	msg.Messages = append(msg.Messages, &pb.UserMessage{
		MsgName: args.MsgName,
		Data:    args.Data,
	})
	this.module.Debugf("SendMsgToAgents: agents:%v msg:%v", args.UserSessionIds, msg)
	if data, err = proto.Marshal(msg); err != nil {
		return
	}
	for _, v := range args.UserSessionIds {
		if a, ok := this.agents.Load(v); ok {
			agent := a.(IAgent)
			if agent.UserId() != "" { //自发送登录用户
				if err = agent.WriteBytes(data); err != nil {
					this.module.Errorln(err)
				}
			}
		}
	}
	return nil
}

func (this *AgentMgrComp) SendMsgsToAgents(ctx context.Context, args *pb.BatchMessagesReq, reply *pb.RPCMessageReply) (err error) {
	var (
		data []byte
	)
	for k, v := range args.Data {
		if a, ok := this.agents.Load(k); ok {
			agent := a.(IAgent)
			if agent.UserId() != "" { //自发送登录用户

				msg := &pb.MessagePackage{
					Messages: make([]*pb.UserMessage, 0),
				}
				msg.Messages = append(msg.Messages, &pb.UserMessage{
					MsgName: args.MsgName,
					Data:    v,
				})
				this.module.Debugf("SendMsgsToAgents: agent:%v msg:%v", k, msg)
				if data, err = proto.Marshal(msg); err != nil {
					this.module.Errorln(err)
					return
				}
				if err = agent.WriteBytes(data); err != nil {
					this.module.Errorln(err)
					return
				}
			}
		}
	}
	return nil
}

// SendMsgToAllAgent 向所有户发送消息
func (this *AgentMgrComp) SendMsgToAllAgent(ctx context.Context, args *pb.BroadCastMessageReq, reply *pb.RPCMessageReply) (err error) {
	var (
		data []byte
	)
	msg := &pb.UserMessage{
		MsgName: args.MsgName,
		Data:    args.Data,
	}
	this.module.Debugf("SendMsgToAllAgent: msg:%v", msg)
	if data, err = proto.Marshal(msg); err != nil {
		return
	}
	this.agents.Range(func(key, value any) bool {
		agent := value.(IAgent)
		if agent.UserId() != "" { //只发送登录用户
			agent.WriteBytes(data)
		}
		return true
	})
	return
}

// SendMsgToAllAgent 向所有户发送消息
func (this *AgentMgrComp) SendMsgToUsers(ctx context.Context, args *pb.BatchUsersMessageReq, reply *pb.RPCMessageReply) (err error) {
	var (
		data []byte
	)
	msg := &pb.UserMessage{
		MsgName: args.MsgName,
		Data:    args.Data,
	}
	this.module.Debugf("SendMsgToAgents: agents:%v msg:%v", args.Uids, msg)
	if data, err = proto.Marshal(msg); err != nil {
		return
	}
	for _, v := range args.Uids {
		if a, ok := this.users.Load(v); ok {
			agent := a.(IAgent)
			if err = agent.WriteBytes(data); err != nil {
				this.module.Errorln(err)
			}
		}
	}
	return nil
}

// CloseAgent 关闭某个用户
func (this *AgentMgrComp) CloseAgent(ctx context.Context, args *pb.AgentCloseeReq, reply *pb.RPCMessageReply) error {
	if a, ok := this.agents.Load(args.UserSessionId); ok {
		if a.(IAgent).UserId() != "" {
			this.users.Delete(a.(IAgent).UserId())
		}
		a.(IAgent).Close()
		this.agents.Delete(args.UserSessionId)
	} else {
		reply.ErrorData = &pb.ErrorData{
			Code:    pb.ErrorCode_UserSessionNobeing,
			Message: fmt.Sprintf("解绑SessionId:%s失败!", args.UserSessionId),
		}
	}
	return nil
}

// 关闭目标游戏用户链接
func (this *AgentMgrComp) CloseAgentsForGamdId(ctx context.Context, args *pb.AgentCloseForGameIdReq, reply *pb.RPCMessageReply) error {
	this.agents.Range(func(key, value any) bool {
		agent := value.(IAgent)
		if agent.GameId() == args.Gameid { //只发送登录用户
			agent.(IAgent).Close()
		}
		return true
	})
	return nil
}
