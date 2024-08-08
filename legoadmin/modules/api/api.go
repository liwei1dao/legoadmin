package api

import (
	"legoadmin/modules"

	"github.com/liwei1dao/lego/base/cluster"
	"github.com/liwei1dao/lego/core"
)

type apiComp struct {
	modules.MCompHttpGate
	service cluster.IClusterService
	module  *Api
	options *Options
}

func (this *apiComp) Init(service core.IService, module core.IModule, comp core.IModuleComp, options core.IModuleOptions) (err error) {
	this.MCompHttpGate.Init(service, module, comp, options)
	this.service = service.(cluster.IClusterService)
	this.module = module.(*Api)
	this.options = options.(*Options)

	return
}

func (this *apiComp) Start() (err error) {
	err = this.MCompHttpGate.Start()
	return
}
