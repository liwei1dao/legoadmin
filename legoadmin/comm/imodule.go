package comm

import (
	"legoadmin/pb"
)

type (
	//玩家用户模块
	IUser interface {
		GetUser(uid string) (user *pb.DBUser, err error)
	}
)
