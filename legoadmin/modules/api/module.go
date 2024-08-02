/*
管理系统
*/
package api

import (
	"legoadmin/comm"
	"legoadmin/modules"

	"github.com/liwei1dao/lego/core"
)

func NewModule() core.IModule {
	m := new(Api)
	return m
}

type Api struct {
	modules.ModuleBase
	api     *apiComp
	model   *modelComp
	options *Options
}

func (this *Api) GetType() core.M_Modules {
	return comm.ModuleApi
}
func (this *Api) NewOptions() (options core.IModuleOptions) {
	return new(Options)
}
func (this *Api) Init(service core.IService, module core.IModule, options core.IModuleOptions) (err error) {
	this.options = options.(*Options)
	if err = this.ModuleBase.Init(service, module, options); err != nil {
		return
	}
	return
}

func (this *Api) Start() (err error) {
	if err = this.ModuleBase.Start(); err != nil {
		return
	}

	return
}

func (this *Api) OnInstallComp() {
	this.ModuleBase.OnInstallComp()
	this.api = this.RegisterComp(new(apiComp)).(*apiComp)
	this.model = this.RegisterComp(new(modelComp)).(*modelComp)
}
