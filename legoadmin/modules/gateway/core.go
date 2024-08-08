package gateway

import (
	"legoadmin/modules"
	"legoadmin/pb"
	"net/http"
	"reflect"

	"github.com/liwei1dao/lego/base/cluster"
	"github.com/liwei1dao/lego/sys/gin/engine"
)

const (
	Msg_GatewayHeartbeat   string = "gateway/heartbeat"   //网关心跳协议
	Msg_GatewayErrornotify string = "gateway/errornotify" //网关错误推送
	Msg_UserLogin          string = "user/login"          //用户登录协议
)

var (
	gatewayReqTyoe, gatewayRespTyoe reflect.Type = reflect.TypeOf(&pb.Rpc_GatewayRouteReq{}), reflect.TypeOf(&pb.Rpc_GatewayRouteResp{})
	httpReqTyoe, httpRespTyoe       reflect.Type = reflect.TypeOf(&pb.Rpc_GatewayHttpRouteReq{}), reflect.TypeOf(&pb.Rpc_GatewayHttpRouteResp{})
)

type (
	// IAgent 用户代理对象接口定义
	IAgent interface {
		ServicePath() string
		SessionId() string
		UserId() string
		IP() string
		GetSessionData() *pb.UserCacheData
		UnBuild()
		WriteMsgs(msgs ...*pb.UserMessage) (err error)
		WriteBytes(data []byte) (err error)
		HandleMessage(msg *pb.UserMessage) (err error)
		Close() //主动关闭接口
	}
	// IGateway 网关模块 接口定义
	IGateway interface {
		modules.IModuleBase
		Service() cluster.IClusterService
		Connect(a IAgent)
		DisConnect(a IAgent)
		LoginNotice(a IAgent)
	}

	HttpResult struct {
		Code    pb.ErrorCode `json:"code"`
		Message string       `json:"message"`
		Data    interface{}  `json:"data"`
	}
)

// 设置跨域
func cors() engine.HandlerFunc {
	return func(c *engine.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		if origin != "" {
			c.Header("Access-Control-Allow-Origin", "*") // 可将将 * 替换为指定的域名
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "true")
		}
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}
