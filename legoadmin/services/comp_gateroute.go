package services

import (
	"context"

	"fmt"
	"legoadmin/comm"
	"legoadmin/pb"
	"reflect"
	"sync"
	"time"

	"github.com/liwei1dao/lego/core"
	"github.com/liwei1dao/lego/core/cbase"
	"github.com/liwei1dao/lego/sys/event"
	"github.com/liwei1dao/lego/sys/log"
	"github.com/liwei1dao/lego/sys/pools"
	"github.com/liwei1dao/lego/utils/mapstructure"
)

// 组件参数
type CompOptions struct {
	MaxTime int32
}

func (this *CompOptions) LoadConfig(settings map[string]interface{}) (err error) {
	if settings != nil {
		err = mapstructure.Decode(settings, this)
	}
	return
}

/*
	服务网关组件 用于接收网关服务发送过来的消息
*/

func NewGateRouteComp() comm.ISC_GateRouteComp {
	comp := new(SCompGateRoute)
	return comp
}

// 服务网关组件
type SCompGateRoute struct {
	cbase.ServiceCompBase
	options    *CompOptions
	service    comm.IService                //rpc服务对象 通过这个对象可以发布服务和调用其他服务的接口
	msghandles map[string]*msghandle        //处理函数的管理对象
	slock      sync.RWMutex                 //回话锁
	sessions   map[string]comm.IUserSession //用户会话对象管理
}

// 设置服务组件名称 方便业务模块中获取此组件对象
func (this *SCompGateRoute) GetName() core.S_Comps {
	return comm.SC_ServiceGateRouteComp
}

func (this *SCompGateRoute) NewOptions() (options core.ICompOptions) {
	return new(CompOptions)
}

// 组件初始化函数
func (this *SCompGateRoute) Init(service core.IService, comp core.IServiceComp, options core.ICompOptions) (err error) {
	err = this.ServiceCompBase.Init(service, comp, options)
	this.options = options.(*CompOptions)
	this.service = service.(comm.IService)
	this.msghandles = make(map[string]*msghandle)
	this.sessions = make(map[string]comm.IUserSession)
	return err
}

// 组件启动时注册rpc服务监听
func (this *SCompGateRoute) Start() (err error) {
	this.service.RegisterFunctionName(string(comm.Rpc_GatewayRoute), this.Rpc_GatewayRoute)                     //注册网关路由接收接口
	this.service.RegisterFunctionName(string(comm.Rpc_GatewayNoticeUserClose), this.Rpc_GatewayNoticeUserClose) //注册网关路由接收接口
	err = this.ServiceCompBase.Start()
	return
}

// 业务模块注册用户消息处理路由
func (this *SCompGateRoute) RegisterRoute(methodName string, comp reflect.Value, msg reflect.Type, handele reflect.Method) {
	//log.Debugf("注册用户路由【%s】", methodName)
	_, ok := this.msghandles[methodName]
	if ok {
		log.Errorf("重复 注册网关消息【%s】", methodName)
		return
	}
	this.msghandles[methodName] = &msghandle{
		rcvr:    comp,
		msgType: msg,
		handle:  handele,
	}
	//注册类型池
	pools.InitType(msg)
}

// RPC----------------------------------------------------------------------------------------------------------------------
// Rpc_GatewayRoute服务接口的接收函数
func (this *SCompGateRoute) Rpc_GatewayRoute(ctx context.Context, args *pb.Rpc_GatewayRouteReq, reply *pb.Rpc_GatewayRouteResp) (err error) {
	var (
		msghandle *msghandle
		session   comm.IUserSession
		ok        bool
	)
	reply.ServicePath = fmt.Sprintf("%s/%s", this.service.GetType(), this.service.GetId())
	msghandle, ok = this.msghandles[args.MsgName]
	if ok {
		if args.UserSession.UserId != "" {
			this.slock.RLock()
			session, ok = this.sessions[args.UserSession.UserId]
			this.slock.RUnlock()
			if !ok {
				session = this.service.GetUserSession(args.UserSession)
			}
		} else {
			session = this.service.GetUserSession(args.UserSession)

		}

		//序列化用户消息对象
		var msg interface{}

		if msg, err = args.Message.UnmarshalNew(); err != nil {
			log.Errorf("[Handle Api] UserMessage:%s Unmarshal err:%v", args.MsgName, err)
			return err
		}

		//执行处理流
		stime := time.Now()
		handlereturn := msghandle.handle.Func.Call([]reflect.Value{msghandle.rcvr, reflect.ValueOf(session), reflect.ValueOf(msg)})
		errdata := handlereturn[0]
		if !errdata.IsNil() { //处理返货错误码 返回用户错误信息
			//data, _ := anypb.New(errdata.(proto.Message))
			reply.ErrorData = errdata.Interface().(*pb.ErrorData)
			// log.Errorf("[Handle Api] t:%v m:%s req:%v reply:%v", time.Since(stime), method, msg, reply)
			log.Error("[Handle Api]",
				log.Field{Key: "t", Value: time.Since(stime).Milliseconds()},
				log.Field{Key: "m", Value: args.MsgName},
				log.Field{Key: "uid", Value: args.UserSession.UserId},
				log.Field{Key: "req", Value: msg},
				log.Field{Key: "reply", Value: reply.String()},
			)
		} else {
			reply.Reply = session.Polls()
			// log.Debugf("[Handle Api] t:%v m:%s uid:%s req:%v reply:%v", time.Since(stime), method, args.UserId, msg, reply)
			nt := time.Since(stime).Milliseconds()
			if this.options.MaxTime == 0 || nt < int64(this.options.MaxTime) {
				log.Debug("[Handle Api]",
					log.Field{Key: "t", Value: time.Since(stime).Milliseconds()},
					log.Field{Key: "m", Value: args.MsgName},
					log.Field{Key: "uid", Value: args.UserSession.UserId},
					log.Field{Key: "req", Value: msg},
					log.Field{Key: "reply", Value: reply.String()},
				)
			} else {
				log.Error("[Handle Api]",
					log.Field{Key: "t", Value: time.Since(stime).Milliseconds()},
					log.Field{Key: "m", Value: args.MsgName},
					log.Field{Key: "uid", Value: args.UserSession.UserId},
					log.Field{Key: "req", Value: msg},
					log.Field{Key: "reply", Value: reply.String()},
				)
			}
		}
	} else { //未找到消息处理函数
		log.Errorf("[Handle Api] no found handle %s", args.MsgName)
		reply.ErrorData = &pb.ErrorData{
			Code:    pb.ErrorCode_ReqParameterError,
			Message: fmt.Sprintf("[Handle Api] no found handle %s", args.MsgName),
		}
	}
	return nil
}

// RPC_NoticeUserClose 接收用户离线通知
func (this *SCompGateRoute) Rpc_GatewayNoticeUserClose(ctx context.Context, args *pb.RPC_Gateway_NoticeUserCloseReq, reply *pb.RPC_Gateway_NoticeUserCloseResp) error {
	var (
		session comm.IUserSession
		ok      bool
	)
	this.slock.RLock()
	session, ok = this.sessions[args.UserSession.UserId]
	this.slock.RUnlock()
	if ok {
		this.slock.Lock()
		delete(this.sessions, args.UserSession.UserId)
		this.slock.Unlock()
		event.TriggerEvent(comm.EventUserOffline, session)
	} else {
		log.Errorf("[NoticeUserClose] err:no fund session:%s", args.String())
	}
	return nil
}

// 对外接口----------------------------------------------------------------------------------------------------------------------------------
func (this *SCompGateRoute) AddUserSession(session comm.IUserSession) {
	this.slock.Lock()
	this.sessions[session.GetUserId()] = session
	this.slock.Unlock()
}
