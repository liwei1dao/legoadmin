/*
网关模块 用于调度所有外部的访问数据
*/
package gateway

import (
	"legoadmin/comm"
	"legoadmin/modules"

	"github.com/liwei1dao/lego/base"
	"github.com/liwei1dao/lego/core"
	"github.com/liwei1dao/lego/sys/log"
)

func NewModule() core.IModule {
	m := new(Gateway)
	return m
}

type Gateway struct {
	modules.ModuleBase
	options   *Options
	service   base.IRPCXService //rpcx服务接口 主要client->server
	agents    *AgentMgrComp     //客户端websocket连接管理
	wsservice *ginComp          //websocket服务 监听websocket连接
}

// GetType 获取模块服务类型
func (this *Gateway) GetType() core.M_Modules {
	return comm.ModuleGate
}

// NewOptions 模块自定义参数
func (this *Gateway) NewOptions() (options core.IModuleOptions) {
	return new(Options)
}

// Service 获取rpcx服务接口
func (this *Gateway) Service() base.IRPCXService {
	return this.service
}

// Init 模块初始化函数
func (this *Gateway) Init(service core.IService, module core.IModule, options core.IModuleOptions) (err error) {
	if err = this.ModuleBase.Init(service, module, options); err != nil {
		return
	}
	this.options = options.(*Options)
	this.service = service.(base.IRPCXService)
	return
}

// Start 模块启动函数 注册rpc服务接口提供用户相关的rpc接口服务
func (this *Gateway) Start() (err error) {
	_name2Func := map[string]any{
		// 向用户发送消息接口
		string(comm.Rpc_GatewayAgentSendMsg): this.agents.SendMsgToAgent,
		// 向多个用户对象发送消息接口
		string(comm.Rpc_GatewaySendBatchMsg): this.agents.SendMsgToAgents,
		// 向所有用户发送消息接口
		string(comm.Rpc_GatewaySendRadioMsg):      this.agents.SendMsgToAllAgent,
		string(comm.Rpc_GatewaySendBatchMsgByUid): this.agents.SendMsgToUsers,
		// 关闭用户socket连接接口
		string(comm.Rpc_GatewayAgentClose): this.agents.CloseAgent,
		// 关闭用户socket连接接口
		string(comm.Rpc_GatewayAgentUnBind): this.agents.CloseAgent,
	}
	for name, fn := range _name2Func {
		this.service.RegisterFunctionName(name, fn)
	}
	if err = this.ModuleBase.Start(); err != nil {
		return
	}
	return
}

// OnInstallComp 装备组件
func (this *Gateway) OnInstallComp() {
	this.ModuleBase.OnInstallComp()
	this.agents = this.RegisterComp(new(AgentMgrComp)).(*AgentMgrComp)
	this.wsservice = this.RegisterComp(new(ginComp)).(*ginComp)
}

// Connect 有新的连接对象进入
func (this *Gateway) Connect(a IAgent) {
	this.Debug("have new connect", log.Field{Key: "SessionId", Value: a.SessionId()})
	this.agents.Connect(a)
}

// DisConnect 有用户断开连接
func (this *Gateway) DisConnect(a IAgent) {
	this.Debug("have new disconnect", log.Field{Key: "SessionId", Value: a.SessionId()})
	this.agents.DisConnect(a)
}

// 登录通知
func (this *Gateway) LoginNotice(a IAgent) {
	this.agents.Logined(a)
}
