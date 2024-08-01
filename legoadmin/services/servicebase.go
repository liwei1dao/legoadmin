package services

import (
	"github.com/liwei1dao/lego/base/cluster"
)

// 基础服务对象
type ServiceBase struct {
	cluster.ClusterService
}

// 初始化相关系统
func (this *ServiceBase) InitSys() {
	this.ClusterService.InitSys()
}
