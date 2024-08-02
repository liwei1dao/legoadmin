package api

import (
	"legoadmin/comm"
	"legoadmin/pb"
	"legoadmin/sys/db"

	"github.com/liwei1dao/lego/core"
	"github.com/liwei1dao/lego/core/cbase"
)

// 代理模型
type modelComp struct {
	cbase.ModuleCompBase
	module *Api
}

// 组件初始化接口
func (this *modelComp) Init(service core.IService, module core.IModule, comp core.IModuleComp, opt core.IModuleOptions) (err error) {
	this.ModuleCompBase.Init(service, module, comp, opt)
	this.module = module.(*Api)

	if err = db.MySql().CreateTable(comm.TableWebUser, &pb.DBApiUser{}); err != nil {
		this.module.Errorln(err)
		return
	}
	model := &pb.DBApiUser{
		Account:  this.module.options.AdninAccount,
		Password: this.module.options.AdninPassword,
		Identity: pb.Identity_Admin,
	}
	// 执行分页查询
	db.MySql().Insert(comm.TableWebUser, model)
	return
}

// 获取用户列表
func (this *modelComp) getusers() (models []*pb.DBApiUser, err error) {

	models = make([]*pb.DBApiUser, 0)
	// 执行分页查询
	err = db.MySql().Find(comm.TableWebUser, &models, db.M{})
	if err != nil {
		this.module.Errorln(err)
		return
	}
	return
}

// 查找账号
func (this *modelComp) findByAccount(account string) (model *pb.DBApiUser, err error) {
	model = &pb.DBApiUser{}
	err = db.MySql().FindOne(comm.TableWebUser, model, db.M{"account": account})
	return
}
