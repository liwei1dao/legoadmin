package gateway

import (
	"fmt"
	"legoadmin/comm"
	"legoadmin/pb"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
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
	this.gin.POST("/api/*path", this.api)
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
		body   []byte
		params string
		claims *comm.TokenClaims
		args   *pb.Rpc_GatewayHttpRouteReq  = pools.GetForType(httpReqTyoe).(*pb.Rpc_GatewayHttpRouteReq)
		reply  *pb.Rpc_GatewayHttpRouteResp = pools.GetForType(httpRespTyoe).(*pb.Rpc_GatewayHttpRouteResp)
		ok     bool
		err    error
	)

	defer func() {
		pools.PutForType(httpReqTyoe, args)
		pools.PutForType(httpRespTyoe, reply)
	}()

	fullPath := c.Request.URL.Path
	// 移除开头的 '/'
	if len(fullPath) > 0 && fullPath[0] == '/' {
		fullPath = fullPath[1:]
	}

	//非登录 验证token
	if fullPath != "api/login" {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusOK, &comm.HttpResult{
				Code:    pb.ErrorCode_TokenInvalid,
				Message: pb.ErrorCode_TokenInvalid.String(),
			})
			return
		}
		token, err := jwt.ParseWithClaims(tokenString, &comm.TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(this.options.ApiKey), nil
		})
		if err != nil || !token.Valid {
			c.JSON(http.StatusOK, &comm.HttpResult{
				Code:    pb.ErrorCode_TokenInvalid,
				Message: "Invalid token",
			})
			return
		}

		claims, ok = token.Claims.(*comm.TokenClaims)
		if !ok {
			c.JSON(http.StatusOK, &comm.HttpResult{
				Code:    pb.ErrorCode_TokenInvalid,
				Message: "Invalid token claims",
			})
			return
		}
		args.Meta = map[string]string{
			comm.HttpContext_UserId:   claims.Account,
			comm.HttpContext_Identity: fmt.Sprintf("%d", claims.Identity),
		}
	}

	if body, err = c.GetRawData(); err != nil {
		c.JSON(http.StatusOK, &comm.HttpResult{
			Code:    pb.ErrorCode_ReqParameterError,
			Message: err.Error(),
		})
		this.module.Errorln(err)
		return
	}
	args.MsgName = fullPath
	args.Message = body
	stime := time.Now()
	if err = this.module.Service().RpcCall(c, comm.Service_Api, string(comm.Rpc_GatewayHttpRoute), args, reply); err != nil {
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
	c.RenderForBytes(http.StatusOK, "application/json", reply.Body)
}
