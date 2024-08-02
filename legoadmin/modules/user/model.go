package user

import (
	"legoadmin/comm"
	"legoadmin/pb"
	"legoadmin/sys/db"

	"github.com/liwei1dao/lego/core"
	"github.com/liwei1dao/lego/core/cbase"
)

// 代理模型
type userModel struct {
	cbase.ModuleCompBase
	module *User
}

// 组件初始化接口
func (this *userModel) Init(service core.IService, module core.IModule, comp core.IModuleComp, opt core.IModuleOptions) (err error) {
	this.ModuleCompBase.Init(service, module, comp, opt)
	this.module = module.(*User)
	return
}
func (this *userModel) Start() (err error) {
	err = this.ModuleCompBase.Start()
	return
}

// 获取代理数据
func (this *userModel) getmodel(uid string) (model *pb.DBUser, err error) {
	model = &pb.DBUser{}
	err = db.MySql().FindOne(comm.TableUser, model, db.M{"uid": uid})
	return
}
