package gateway

import (
	"net/http"

	"github.com/liwei1dao/lego/core"
	"github.com/liwei1dao/lego/core/cbase"
	"github.com/liwei1dao/lego/sys/gin"
	"github.com/liwei1dao/lego/sys/gin/engine"
	"github.com/liwei1dao/lego/sys/log"

	"github.com/gorilla/websocket"
)

type WSServiceComp struct {
	cbase.ModuleCompBase
	options *Options
	module  *Gateway // 网关
	gin     gin.ISys // gin服务接口
}

// Init websocket服务初始化
func (this *WSServiceComp) Init(service core.IService, module core.IModule, comp core.IModuleComp, options core.IModuleOptions) (err error) {
	err = this.ModuleCompBase.Init(service, module, comp, options)
	this.options = options.(*Options)
	this.module = module.(*Gateway)
	this.gin, err = gin.NewSys(
		gin.SetListenPort(this.options.ListenPort),
		gin.SetDebug(this.options.GinDebug),
	)
	//游戏业务逻辑处理
	this.gin.GET("/ws", this.ws)
	return
}

// ws 升级websocket连接处理本次请求
func (this *WSServiceComp) ws(c *engine.Context) {
	upGrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	if wsConn, err := upGrader.Upgrade(c.Writer, c.Request, nil); err != nil {
		log.Errorf("accept faile client:%s err:%v", c.RemoteIP(), err)
		return
	} else {
		agent := newAgent(this.module, wsConn)
		this.module.Connect(agent)
	}
}
