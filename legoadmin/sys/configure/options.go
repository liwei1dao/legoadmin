package configure

import (
	"github.com/liwei1dao/lego/sys/log"
	"github.com/liwei1dao/lego/utils/mapstructure"
)

type Option func(*Options)
type Options struct {
	ConfigurePath string //配置中心路径
	TimestampFile string //时间戳配置文件路径
	CheckInterval int    //配置文件更新检查间隔时间 单位秒
	Debug         bool   //日志是否开启
	Log           log.ILogger
}

func SetConfigurePath(v string) Option {
	return func(o *Options) {
		o.ConfigurePath = v
	}
}
func SetTimestampFile(v string) Option {
	return func(o *Options) {
		o.TimestampFile = v
	}
}
func SetCheckInterval(v int) Option {
	return func(o *Options) {
		o.CheckInterval = v
	}
}

func SetDebug(v bool) Option {
	return func(o *Options) {
		o.Debug = v
	}
}

func SetLog(v log.ILogger) Option {
	return func(o *Options) {
		o.Log = v
	}
}

func newOptions(config map[string]interface{}, opts ...Option) (Options, error) {
	options := Options{
		CheckInterval: 60,
	}
	if config != nil {
		mapstructure.Decode(config, &options)
	}
	for _, o := range opts {
		o(&options)
	}
	if options.Log == nil {
		options.Log = log.NewTurnlog(options.Debug, log.Clone("sys.configure", 3))
	}
	return options, nil
}

func newOptionsByOption(opts ...Option) (Options, error) {
	options := Options{
		CheckInterval: 60,
	}
	for _, o := range opts {
		o(&options)
	}
	if options.Log == nil {
		options.Log = log.NewTurnlog(options.Debug, log.Clone("sys.configure", 3))
	}
	return options, nil
}

func SetConfigPath(v string) Option {
	return func(o *Options) {
		o.ConfigurePath = v
	}
}
