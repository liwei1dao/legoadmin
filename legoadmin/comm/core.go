package comm

import "github.com/liwei1dao/lego/core"

// 模块名定义处
const (
	ModuleGate core.M_Modules = "gateway" //gate模块 网关服务模块

)

// RPC服务接口定义处
const ( //Rpc
	//Gateway 网关消息
	Rpc_GatewayRoute               core.Rpc_Key = "Rpc_GatewayRoute"               //网关路由
	Rpc_GatewayHttpRoute           core.Rpc_Key = "Rpc_GatewayHttpRoute"           //Http网关路由
	Rpc_GatewayAgentUnBind         core.Rpc_Key = "Rpc_GatewayAgentUnBind"         //代理解绑 解绑用户Id
	Rpc_GatewayAgentSendMsg        core.Rpc_Key = "Rpc_GatewayAgentSendMsg"        //代理发送消息 向用户发送消息
	Rpc_GatewaySendBatchMsg        core.Rpc_Key = "Rpc_GatewaySendBatchMsg"        //向多个用户发送消息
	Rpc_GatewaySendBatchMsgByUid   core.Rpc_Key = "Rpc_GatewaySendBatchMsgByUid"   //向多个用户发送消息 查询uid
	Rpc_GatewaySendRadioMsg        core.Rpc_Key = "Rpc_GatewaySendRadioMsg"        //广播消息
	Rpc_GatewaySendRadioMsgByGroup core.Rpc_Key = "Rpc_GatewaySendRadioMsgByGroup" //广播消息到组
	Rpc_GatewayAgentClose          core.Rpc_Key = "Rpc_GatewayAgentClose"          //代理关闭 关闭用户连接
	Rpc_GatewayNoticeUserLogin     core.Rpc_Key = "Rpc_NoticeUserLogin"            //通知用户登录
	Rpc_GatewayNoticeUserCreate    core.Rpc_Key = "Rpc_NoticeUserCreate"           //通知用户创角
	Rpc_GatewayNoticeUserClose     core.Rpc_Key = "Rpc_NoticeUserClose"            //通知用户离线

)
