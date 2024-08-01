package modules

import (
	"errors"

	"github.com/liwei1dao/lego/core"
	"github.com/liwei1dao/lego/sys/log"

	"github.com/mitchellh/mapstructure"
)

type (
	IOptions interface {
		core.IModuleOptions
		GetDebug() bool
		GetLog() log.ILogger
	}
	Options struct {
		Debug bool //日志是否开启
		Log   log.ILogger
	}
)

func (this *Options) GetDebug() bool {
	return this.Debug
}

func (this *Options) GetLog() log.ILogger {
	return this.Log
}

func (this *Options) LoadConfig(settings map[string]interface{}) (err error) {
	this.Debug = false
	if settings != nil {
		err = mapstructure.Decode(settings, this)
	}

	if this.Log = log.NewTurnlog(this.Debug, log.Clone("", 4)); this.Log == nil {
		err = errors.New("log is nil")
	}
	return
}
