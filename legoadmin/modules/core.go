package modules

import (
	"github.com/liwei1dao/lego/core"
	"github.com/liwei1dao/lego/sys/log"
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
