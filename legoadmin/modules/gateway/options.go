package gateway

import (
	"legoadmin/modules"

	"github.com/liwei1dao/lego/utils/mapstructure"
)

/*
网关模块 参数定义
*/

type (
	Options struct {
		modules.Options
		GinDebug   bool   //web引擎日志开关
		ListenPort int    //websocket 监听端口
		ApiKey     string //后台加密Key
	}
)

// LoadConfig 配置文件序列化为Options
func (this *Options) LoadConfig(settings map[string]interface{}) (err error) {
	if settings != nil {
		if err = this.Options.LoadConfig(settings); err != nil {
			return
		}
		if err = mapstructure.Decode(settings, this); err != nil {
			return
		}
	}
	return
}
