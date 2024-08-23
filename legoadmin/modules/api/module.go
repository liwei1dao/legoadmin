/*
管理系统
*/
package api

import (
	"legoadmin/comm"
	"legoadmin/modules"

	"github.com/liwei1dao/lego/core"
	"github.com/liwei1dao/lego/sys/httppool"
)

func NewModule() core.IModule {
	m := new(Api)
	return m
}

type Api struct {
	modules.ModuleBase
	pool    httppool.ISys
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
	if this.pool, err = httppool.NewSys(
		httppool.SetMaxIdleConns(5),
		httppool.SetMaxIdleConnsPerHost(1),
		httppool.SetIdleConnTimeout(15),
		httppool.SetRequestTimeout(3),
	); err != nil {
		this.Errorln(err)
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
