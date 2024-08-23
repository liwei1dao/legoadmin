package services

import (
	"legoadmin/comm"
	"reflect"

	"github.com/liwei1dao/lego/utils/mapstructure"
)

// 组件参数
type (
	//路由组件选项
	RouteCompOptions struct {
		MaxTime int32 //路由执行任务 超时警告
	}
)

var (
	httpResultTyoe reflect.Type = reflect.TypeOf(&comm.HttpResult{})
)

func (this *RouteCompOptions) LoadConfig(settings map[string]interface{}) (err error) {
	if settings != nil {
		err = mapstructure.Decode(settings, this)
	}
	return
}
