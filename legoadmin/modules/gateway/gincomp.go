package gateway

import (
	"legoadmin/comm"
	"legoadmin/pb"
	"net/http"
	"time"

	"github.com/liwei1dao/lego/core"
	"github.com/liwei1dao/lego/core/cbase"
	"github.com/liwei1dao/lego/sys/gin"
	"github.com/liwei1dao/lego/sys/gin/engine"
	"github.com/liwei1dao/lego/sys/log"
	"github.com/liwei1dao/lego/sys/pools"

	"github.com/gorilla/websocket"
)

type ginComp struct {
	cbase.ModuleCompBase
	options *Options
	module  *Gateway // 网关
	gin     gin.ISys // gin服务接口
}

// Init websocket服务初始化
func (this *ginComp) Init(service core.IService, module core.IModule, comp core.IModuleComp, options core.IModuleOptions) (err error) {
	err = this.ModuleCompBase.Init(service, module, comp, options)
	this.options = options.(*Options)
	this.module = module.(*Gateway)
	this.gin, err = gin.NewSys(
		gin.SetListenPort(this.options.ListenPort),
		gin.SetDebug(this.options.GinDebug),
	)
	//游戏业务逻辑处理
	this.gin.GET("/ws", this.ws)
	//api 业务请求
	this.gin.POST("/web/:param1/:param2", this.api)
	return
}

// ws 升级websocket连接处理本次请求
func (this *ginComp) ws(c *engine.Context) {
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

// 后台接口转发
func (this *ginComp) api(c *engine.Context) {
	var (
		param1, param2 string
		body           []byte
		params         string
		args           *pb.Rpc_GatewayHttpRouteReq  = pools.GetForType(httpReqTyoe).(*pb.Rpc_GatewayHttpRouteReq)
		reply          *pb.Rpc_GatewayHttpRouteResp = pools.GetForType(httpRespTyoe).(*pb.Rpc_GatewayHttpRouteResp)
		err            error
	)

	defer func() {
		pools.PutForType(httpReqTyoe, args)
		pools.PutForType(httpRespTyoe, reply)
	}()

	param1 = c.Param("param1")
	param2 = c.Param("param2")

	if body, err = c.GetRawData(); err != nil {
		c.JSON(http.StatusOK, &comm.HttpResult{
			Code:    pb.ErrorCode_ReqParameterError,
			Message: err.Error(),
		})
		this.module.Errorln(err)
		return
	}
	args.MsgName = param2
	args.Message = body
	stime := time.Now()
	if err = this.module.Service().RpcCall(c, param1, string(comm.Rpc_GatewayHttpRoute), args, reply); err != nil {
		this.module.Error("[API]",
			log.Field{Key: "req", Value: params},
			log.Field{Key: "err", Value: err.Error()},
		)
		return
	}
	this.module.Debug("[API]",
		log.Field{Key: "t", Value: time.Since(stime).Milliseconds()},
		log.Field{Key: "req", Value: params},
		log.Field{Key: "reply", Value: reply.String()},
	)
	c.RenderForBytes(http.StatusOK, reply.ContentType, reply.Body)
}
