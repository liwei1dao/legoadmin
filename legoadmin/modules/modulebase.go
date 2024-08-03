package modules

import (
	"fmt"
	"legoadmin/comm"
	"runtime"

	"github.com/liwei1dao/lego/core"
	"github.com/liwei1dao/lego/core/cbase"
	"github.com/liwei1dao/lego/sys/log"
)

/*
基础业务模块实现 封装一些通用的接口提供给业务模块使用
*/
type ModuleBase struct {
	cbase.ModuleBase
	service comm.IService
	options IOptions
}

func (this *ModuleBase) NewOptions() (options core.IModuleOptions) {
	return new(Options)
}

// 模块初始化接口
func (this *ModuleBase) Init(service core.IService, module core.IModule, options core.IModuleOptions) (err error) {
	this.service = service.(comm.IService)
	this.options = options.(IOptions)
	this.SetName("module." + string(module.GetType()))
	if err = this.ModuleBase.Init(service, module, options); err != nil {
		return
	}
	return
}

// 异步调用用户处理流
func (this *ModuleBase) AsynHandleSession(uctx comm.IUserContext, handle func(session comm.IUserContext)) {
	defer func() { //程序异常 收集异常信息传递给前端显示
		if r := recover(); r != nil {
			buf := make([]byte, 4096)
			l := runtime.Stack(buf, false)
			err := fmt.Errorf("%v: %s", r, buf[:l])
			log.Errorf("[AsynHandleSession] err:%s", err.Error())
		}
	}()
	handle(uctx)
	if uctx.IsOnline() {
		uctx.Push()
	}
	this.service.PutUserContext(uctx)
}

// 日志
func (this *ModuleBase) Enabled(lvl log.Loglevel) bool {
	return this.options.GetLog().Enabled(lvl)
}
func (this *ModuleBase) SetName(name string) {
	this.options.GetLog().SetName(name)
}

// 日志接口
func (this *ModuleBase) Debug(msg string, args ...log.Field) {
	this.options.GetLog().Debug(msg, args...)
}
func (this *ModuleBase) Info(msg string, args ...log.Field) {
	this.options.GetLog().Info(msg, args...)
}
func (this *ModuleBase) Print(msg string, args ...log.Field) {
	this.options.GetLog().Print(msg, args...)
}
func (this *ModuleBase) Warn(msg string, args ...log.Field) {
	this.options.GetLog().Warn(msg, args...)
}
func (this *ModuleBase) Error(msg string, args ...log.Field) {
	this.options.GetLog().Error(msg, args...)
}
func (this *ModuleBase) Panic(msg string, args ...log.Field) {
	this.options.GetLog().Panic(msg, args...)
}
func (this *ModuleBase) Fatal(msg string, args ...log.Field) {
	this.options.GetLog().Fatal(msg, args...)
}

func (this *ModuleBase) Debugf(format string, args ...interface{}) {
	this.options.GetLog().Debugf(format, args...)
}
func (this *ModuleBase) Infof(format string, args ...interface{}) {
	this.options.GetLog().Infof(format, args...)
}
func (this *ModuleBase) Printf(format string, args ...interface{}) {
	this.options.GetLog().Printf(format, args...)
}
func (this *ModuleBase) Warnf(format string, args ...interface{}) {
	this.options.GetLog().Warnf(format, args...)
}
func (this *ModuleBase) Errorf(format string, args ...interface{}) {
	this.options.GetLog().Errorf(format, args...)
}
func (this *ModuleBase) Fatalf(format string, args ...interface{}) {
	this.options.GetLog().Fatalf(format, args...)
}
func (this *ModuleBase) Panicf(format string, args ...interface{}) {
	this.options.GetLog().Panicf(format, args...)
}

func (this *ModuleBase) Debugln(args ...interface{}) {
	this.options.GetLog().Debugln(args...)
}
func (this *ModuleBase) Infoln(args ...interface{}) {
	this.options.GetLog().Infoln(args...)
}
func (this *ModuleBase) Println(args ...interface{}) {
	this.options.GetLog().Println(args...)
}
func (this *ModuleBase) Warnln(args ...interface{}) {
	this.options.GetLog().Warnln(args...)
}
func (this *ModuleBase) Errorln(args ...interface{}) {
	this.options.GetLog().Errorln(args...)
}
func (this *ModuleBase) Fatalln(args ...interface{}) {
	this.options.GetLog().Fatalln(args...)
}
func (this *ModuleBase) Panicln(args ...interface{}) {
	this.options.GetLog().Panicln(args...)
}
