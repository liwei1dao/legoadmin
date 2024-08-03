package user

import (
	"legoadmin/comm"
	"legoadmin/modules"
	"legoadmin/pb"

	"github.com/liwei1dao/lego/core"
	"github.com/liwei1dao/lego/sys/event"
)

var _ comm.IUser = (*User)(nil)

func NewModule() core.IModule {
	m := new(User)
	return m
}

type User struct {
	modules.ModuleBase
	service   comm.IService
	model     *userModel
	api       *apiComp
	configure *configureComp
}

func (this *User) GetType() core.M_Modules {
	return comm.ModuleUser
}
func (this *User) NewOptions() (options core.IModuleOptions) {
	return new(Options)
}

func (this *User) Init(service core.IService, module core.IModule, options core.IModuleOptions) (err error) {
	if err = this.ModuleBase.Init(service, module, options); err != nil {
		return
	}
	this.service = service.(comm.IService)
	return
}

func (this *User) Start() (err error) {
	if err = this.ModuleBase.Start(); err != nil {
		return
	}
	event.RegisterGO(comm.EventUserLogin, this.EventUserLogin)
	event.RegisterGO(comm.EventUserOffline, this.EventUserOffline)
	return
}

func (this *User) OnInstallComp() {
	this.ModuleBase.OnInstallComp()
	this.api = this.RegisterComp(new(apiComp)).(*apiComp)
	this.configure = this.RegisterComp(new(configureComp)).(*configureComp)
	this.model = this.RegisterComp(new(userModel)).(*userModel)
}

func (this *User) GetUser(uid string) (user *pb.DBUser, err error) {
	user, err = this.model.getmodel(uid)
	return
}

// Event------------------------------------------------------------------------------------------------------------
// 用户离线通知
func (this *User) EventUserLogin(session comm.IUserContext) {

}

// 用户离线通知
func (this *User) EventUserOffline(session comm.IUserContext) {

}
