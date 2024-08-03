package services

import (
	"context"
	"legoadmin/comm"
	"legoadmin/pb"

	"github.com/liwei1dao/lego/base/cluster"
	"github.com/liwei1dao/lego/core"
	"github.com/liwei1dao/lego/sys/pools"
)

// 基础服务对象
type ServiceBase struct {
	cluster.ClusterService
}

func (this *ServiceBase) Init(service core.IService) (err error) {
	if err = this.ClusterService.Init(service); err != nil {
		return
	}
	pools.Add(comm.Pool_UserSession, func() comm.IUserContext { return comm.NewUserContext() })
	return
}

// 初始化相关系统
func (this *ServiceBase) InitSys() {
	this.ClusterService.InitSys()
}

func (this *ServiceBase) GetUserContext(ctx context.Context, cache *pb.UserCacheData) (session comm.IUserContext) {
	session = pools.Get[comm.IUserContext](comm.Pool_UserSession)
	session.SetSession(ctx, this, cache)
	return
}

func (this *ServiceBase) PutUserContext(ctx comm.IUserContext) {
	pools.Put[comm.IUserContext](comm.Pool_UserSession, ctx)
}
