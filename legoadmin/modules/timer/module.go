/*
定时任务
*/
package timer

import (
	"legoadmin/comm"
	"legoadmin/modules"

	"github.com/liwei1dao/lego/base/cluster"
	"github.com/liwei1dao/lego/core"
)

func NewModule() core.IModule {
	m := new(Timer)
	return m
}

type Timer struct {
	modules.ModuleBase
	service cluster.IClusterService
	options *Options
}

func (this *Timer) GetType() core.M_Modules {
	return comm.ModuleTimer
}
func (this *Timer) NewOptions() (options core.IModuleOptions) {
	return new(Options)
}
func (this *Timer) Init(service core.IService, module core.IModule, options core.IModuleOptions) (err error) {
	this.service = service.(cluster.IClusterService)
	this.options = options.(*Options)
	if err = this.ModuleBase.Init(service, module, options); err != nil {
		return
	}
	return
}

func (this *Timer) Start() (err error) {
	if err = this.ModuleBase.Start(); err != nil {
		return
	}
	return
}

func (this *Timer) OnInstallComp() {
	this.ModuleBase.OnInstallComp()

}
