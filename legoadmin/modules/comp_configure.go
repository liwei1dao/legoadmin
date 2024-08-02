package modules

import (
	"fmt"
	"legoadmin/sys/configure"
	cfg "legoadmin/sys/configure/structs"
	"sync"

	"github.com/liwei1dao/lego/core"
	"github.com/liwei1dao/lego/core/cbase"
)

// /配置管理基础组件
type MCompConfigure struct {
	cbase.ModuleCompBase
	module   core.IModule
	lock     sync.RWMutex
	playerlv map[int32]map[int32]*cfg.GamePlayerLvData
}

// 组件初始化接口
func (this *MCompConfigure) Init(service core.IService, module core.IModule, comp core.IModuleComp, options core.IModuleOptions) (err error) {
	err = this.ModuleCompBase.Init(service, module, comp, options)
	this.module = module
	this.playerlv = make(map[int32]map[int32]*cfg.GamePlayerLvData)

	return
}

func (this *MCompConfigure) LoadConfigure(name string, fn interface{}) (err error) {
	return configure.RegisterConfigure(name, fn, nil)
}

func (this *MCompConfigure) UpdateConfigure(name string, fn interface{}, update func()) (err error) {
	return configure.RegisterConfigure(name, fn, update)
}

// 读取配置数据
func (this *MCompConfigure) GetConfigure(name string) (v interface{}, err error) {
	return configure.GetConfigure(name)
}

// 创建配置表错误对象
func (this *MCompConfigure) NewNotFoundConfErr(filename string, id interface{}) error {
	return fmt.Errorf("NotFoundConf Err module:%s ,file:%s,id:%v", this.module.GetType(), filename, id)
}
