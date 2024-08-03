package modules

import (
	"context"
	"legoadmin/comm"
	"legoadmin/pb"
	"reflect"

	"github.com/gogo/protobuf/proto"
	"github.com/liwei1dao/lego/core"
	"github.com/liwei1dao/lego/sys/log"
)

var (
	typeOfSession   = reflect.TypeOf((*comm.IUserContext)(nil)).Elem()
	typeOfMessage   = reflect.TypeOf((*proto.Message)(nil)).Elem()
	typeOfErrorCode = reflect.TypeOf((*pb.ErrorCode)(nil)).Elem()
	typeOfErrorData = reflect.TypeOf((*pb.ErrorData)(nil))
	typeOfError     = reflect.TypeOf((*error)(nil)).Elem()
	typeOfContext   = reflect.TypeOf((*context.Context)(nil)).Elem()
)

type (
	//业务模块基类接口 定义所有业务模块都可以使用的接口
	IModuleBase interface {
		core.IModule
		log.ILogger
	}
	IMCompModel interface {
		core.IModuleComp
		GetModelName() string
		ReadKey(key string) string
		ReadListKey(uid string, id string) string
	}
)
