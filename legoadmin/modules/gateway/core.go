package gateway

import (
	"legoadmin/modules"
	"legoadmin/pb"
	"net/http"
	"sync"

	"github.com/liwei1dao/lego/base"

	"github.com/liwei1dao/lego/sys/gin/engine"
)

type (
	// IAgent 用户代理对象接口定义
	IAgent interface {
		SessionId() string
		IP() string
		AgentId() string
		GameId() string
		UserId() string
		ServicePath() string
		GetSessionData() *pb.UserSessionData
		WriteMsgs(msgs ...*pb.UserMessage) (err error)
		WriteBytes(data []byte) (err error)
		HandleMessage(msg *pb.UserMessage) (err error)
		Close() //主动关闭接口
	}
	// IGateway 网关模块 接口定义
	IGateway interface {
		modules.IModuleBase
		Service() base.IRPCXService
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

var msgPool = &sync.Pool{
	New: func() interface{} {
		return &pb.AgentMessage{
			UserSession: &pb.UserSessionData{},
		}
	},
}

func getmsg() *pb.AgentMessage {
	req := msgPool.Get().(*pb.AgentMessage)
	return req
}

func putmsg(r *pb.AgentMessage) {
	msgPool.Put(r)
}

var msgreplyPool = &sync.Pool{
	New: func() interface{} {
		return &pb.RPCMessageReply{}
	},
}

func getmsgreply() *pb.RPCMessageReply {
	reply := msgreplyPool.Get().(*pb.RPCMessageReply)
	return reply
}

func putmsgreply(r *pb.RPCMessageReply) {
	msgreplyPool.Put(r)
}

var httpreplyPool = &sync.Pool{
	New: func() interface{} {
		return &pb.RPCHttpMessageReply{}
	},
}

func gethttpreply() *pb.RPCHttpMessageReply {
	reply := httpreplyPool.Get().(*pb.RPCHttpMessageReply)
	return reply
}

func puthttpreply(r *pb.RPCHttpMessageReply) {
	httpreplyPool.Put(r)
}

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
