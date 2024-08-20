package main

import (
	"flag"
	"legoadmin/services"

	"github.com/liwei1dao/lego"
	"github.com/liwei1dao/lego/base/cluster"
	"github.com/liwei1dao/lego/core"
)

/*
服务类型:后台服务
*/
var (
	conf = flag.String("conf", "./conf/api.yaml", "获取需要启动的服务配置文件") //启动服务的Id
)

/*服务启动的入口函数*/
func main() {
	flag.Parse()
	s := NewService(
		cluster.SetConfPath(*conf),
		cluster.SetVersion("1.0.0.0"),
	)
	s.OnInstallComp( //装备组件

	)
	lego.Run(s) //运行模块

}

func NewService(ops ...cluster.Option) core.IService {
	s := new(Service)
	s.Configure(ops...)
	return s
}

// worker 的服务对象定义
type Service struct {
	services.ServiceBase
}

// 初始化worker需要的一些系统工具
func (this *Service) InitSys() {
	this.ServiceBase.InitSys()
}
