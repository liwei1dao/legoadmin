package comm

import "github.com/liwei1dao/lego/core"

//服务定义
const (
	Service_Gateway = "gateway" //网关服务 可多开
	Service_Timer   = "timer"   //定时任务
	Service_Agent   = "agent"   //渠道接口服
	Service_Api     = "api"     //渠道接口服
)

// 模块名定义处
const (
	ModuleGate  core.M_Modules = "gateway" //gate模块 网关服务模块
	ModuleTimer core.M_Modules = "timer"   //定时任务模块
	ModuleApi   core.M_Modules = "api"     //后台
	ModuleUser  core.M_Modules = "user"    //用户模块
)

// 服务组件名称
const (
	SC_ServiceGateRouteComp core.S_Comps = "SC_GateRouteComp" //服务组件 消息路由组件
	SC_ServiceHttpRouteComp core.S_Comps = "SC_HttpRouteComp" //服务组件 消息路由组件
)

// RPC服务接口定义处
const ( //Rpc
	//Gateway 网关消息
	Rpc_GatewayRoute               core.Rpc_Key = "Rpc_GatewayRoute"               //网关路由
	Rpc_GatewayHttpRoute           core.Rpc_Key = "Rpc_GatewayHttpRoute"           //Http网关路由
	Rpc_GatewayAgentSendMsg        core.Rpc_Key = "Rpc_GatewayAgentSendMsg"        //代理发送消息 向用户发送消息
	Rpc_GatewaySendBatchMsg        core.Rpc_Key = "Rpc_GatewaySendBatchMsg"        //向多个用户发送消息
	Rpc_GatewaySendBatchMsgs       core.Rpc_Key = "Rpc_GatewaySendBatchMsg"        //向多个用户发送不同的消息
	Rpc_GatewaySendBatchMsgByUid   core.Rpc_Key = "Rpc_GatewaySendBatchMsgByUid"   //向多个用户发送消息 查询uid
	Rpc_GatewaySendRadioMsg        core.Rpc_Key = "Rpc_GatewaySendRadioMsg"        //广播消息
	Rpc_GatewaySendRadioMsgByGroup core.Rpc_Key = "Rpc_GatewaySendRadioMsgByGroup" //广播消息到组
	Rpc_GatewayAgentUnBind         core.Rpc_Key = "Rpc_GatewayAgentUnBind"         //代理解绑 解绑用户Id
	Rpc_GatewayAgentClose          core.Rpc_Key = "Rpc_GatewayAgentClose"          //代理关闭 关闭用户连接
	Rpc_GatewayNoticeUserLogin     core.Rpc_Key = "Rpc_NoticeUserLogin"            //通知用户登录
	Rpc_GatewayNoticeUserCreate    core.Rpc_Key = "Rpc_NoticeUserCreate"           //通知用户创角
	Rpc_GatewayNoticeUserClose     core.Rpc_Key = "Rpc_NoticeUserClose"            //通知用户离线

)

// 事件类型定义处
const (
	EventUserLogin   core.Event_Key = "Event_UserLogin"   //登录事件
	EventUserOffline core.Event_Key = "Event_UserOffline" //用户离线事件
	EventGameOff     core.Event_Key = "Event_GameOff"     //游戏关闭
)

//数据库表定义
const (
	TableWebUser = "webuser" //后台用户表
	TableUser    = "user"    // 用户表
)

//session 缓存数据Key
const (
	Session_User = "user"
)

//session 缓存数据Key
const (
	HttpContext_UserId   = "userid"
	HttpContext_Identity = "identity"
)

//对象池
const (
	Pool_HttpSession = "Pool_HttpSession"
	Pool_UserSession = "Pool_UserSession"
)
