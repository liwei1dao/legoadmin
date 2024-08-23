package main

import (
	"flag"
	"fmt"
	"legoadmin/modules/api"
	"legoadmin/services"
	"legoadmin/sys/db"

	"github.com/liwei1dao/lego"
	"github.com/liwei1dao/lego/base/cluster"
	"github.com/liwei1dao/lego/core"
	"github.com/liwei1dao/lego/sys/log"
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
		services.NewHttpRouteComp(),
	)
	lego.Run(s, //运行模块
		api.NewModule(),
	)

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
	//存储系统
	if err := db.OnInit(this.GetSettings().Sys["db"]); err != nil {
		panic(fmt.Sprintf("init sys.db err: %s", err.Error()))
	} else {
		log.Infof("init sys.db success!")
	}
}
