package user

import (
	"legoadmin/comm"
	"legoadmin/modules"

	"github.com/liwei1dao/lego/core"
)

type apiComp struct {
	modules.MCompGate
	service comm.IService
	module  *User
	options *Options
}

func (this *apiComp) Init(service core.IService, module core.IModule, comp core.IModuleComp, options core.IModuleOptions) (err error) {
	this.MCompGate.Init(service, module, comp, options)
	this.service = service.(comm.IService)
	this.module = module.(*User)
	this.options = options.(*Options)
	return
}

func (this *apiComp) Start() (err error) {
	err = this.MCompGate.Start()
	return
}
