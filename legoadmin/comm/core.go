package comm

import (
	"context"
	"legoadmin/pb"
	"reflect"

	"github.com/golang-jwt/jwt"
	"github.com/liwei1dao/lego/base"
	"github.com/liwei1dao/lego/core"
	"google.golang.org/protobuf/proto"
)

type IService interface {
	base.IRPCXService
	GetUserSession(cache *pb.UserSessionData) (session IUserSession)
	PutUserSession(session IUserSession)
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

// 用户会话
type IUserSession interface {
	SetSession(ctx context.Context, service IService, cache *pb.UserSessionData)
	GetCache() *pb.UserSessionData
	GetUserId() string
	IsOnline() bool
	UnBind() (err error)
	SendMsg(mainType, subType string, msg proto.Message) (err error)
	Polls() []*pb.UserMessage
	Push() (err error)       //警告 api传递过来的会话禁用此接口
	SyncPush() (err error)   //警告 api传递过来的会话禁用此接口 同步
	SetOffline(offline bool) //设置离线状态
	GetOffline() bool        //设置离线状态
	Close() (err error)
	Reset()
	SetMate(name string, value interface{})
	GetMate(name string) (ok bool, value interface{})
	Clone() (session IUserSession) //克隆
}

// Claims struct to define JWT claims
type TokenClaims struct {
	Account  string      `json:"account"`
	Identity pb.Identity `json:"identity"`
	jwt.StandardClaims
}
