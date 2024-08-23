package comm

import (
	"context"
	"legoadmin/pb"
	"reflect"

	"github.com/golang-jwt/jwt"
	"github.com/liwei1dao/lego/base/cluster"
	"github.com/liwei1dao/lego/core"
	"google.golang.org/protobuf/proto"
)

type IService interface {
	cluster.IClusterService
	GetHttpContext(ctx context.Context) (context IHttpContext)
	PutHttpContext(ctx IHttpContext)
	GetUserContext(ctx context.Context, cache *pb.UserCacheData) (session IUserContext)
	PutUserContext(ctx IUserContext)
}

// 服务网关组件接口定义
type ISC_GateRouteComp interface {
	core.IServiceComp
	Rpc_GatewayRoute(ctx context.Context, args *pb.Rpc_GatewayRouteReq, reply *pb.Rpc_GatewayRouteResp) error
	RegisterRoute(methodName string, comp reflect.Value, msg reflect.Type, handle reflect.Method)
}

// 服务网关组件接口定义
type ISC_HttpRouteComp interface {
	core.IServiceComp
	Rpc_GatewayHttpRoute(ctx context.Context, args *pb.Rpc_GatewayHttpRouteReq, reply *pb.Rpc_GatewayHttpRouteResp) error
	RegisterRoute(methodName string, comp reflect.Value, msg reflect.Type, handle reflect.Method)
}

type IHttpContext interface {
	context.Context
	SetMate(name string, value interface{})
	GetMate(name string) (ok bool, value interface{})
}

// 用户会话
type IUserContext interface {
	context.Context
	SetSession(ctx context.Context, service IService, cache *pb.UserCacheData)
	GetCache() *pb.UserCacheData
	GetUserId() string
	IsOnline() bool
	UnBind() (err error)
	SendMsg(mainType, subType string, msg proto.Message) (err error)
	Polls() []*pb.UserMessage //获取缓存消息
	Push() (err error)        //警告 api传递过来的会话禁用此接口
	SyncPush() (err error)    //警告 api传递过来的会话禁用此接口 同步
	Close() (err error)
	Reset()
	SetMate(name string, value interface{})
	GetMate(name string) (ok bool, value interface{})
	Clone(ctx context.Context) (session IUserContext) //克隆
}

type HttpResult struct {
	Code    pb.ErrorCode `json:"code"`
	Message string       `json:"message"`
	Data    interface{}  `json:"data"`
}

// Claims struct to define JWT claims
type TokenClaims struct {
	Account  string      `json:"account"`
	Identity pb.Identity `json:"identity"`
	jwt.StandardClaims
}
