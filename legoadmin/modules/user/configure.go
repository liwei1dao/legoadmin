package user

import (
	"legoadmin/modules"

	"github.com/liwei1dao/lego/core"
)

type configureComp struct {
	modules.MCompConfigure
}

func (this *configureComp) Init(service core.IService, module core.IModule, comp core.IModuleComp, options core.IModuleOptions) (err error) {
	this.MCompConfigure.Init(service, module, comp, options)

	return
}

func (this *configureComp) Start() (err error) {
	err = this.MCompConfigure.Start()

	return
}
