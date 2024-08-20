package user

import (
	"legoadmin/comm"
	"legoadmin/pb"
	"legoadmin/sys/db"
	"time"

	"github.com/liwei1dao/lego/sys/event"
	"github.com/liwei1dao/lego/sys/mysql"
	"github.com/liwei1dao/lego/utils/container/id"
)

// 登录
func (this *apiComp) Login(session comm.IUserContext, req *pb.UserLoginReq) (errdata *pb.ErrorData) {
	var (
		user *pb.DBUser
		err  error
	)
	if user, err = this.module.model.getmodelForAccount(req.Account); err != nil {
		this.module.Errorln(err)
		return
	}
	if err == mysql.ErrNoDocuments { //创建新的账号
		//如果是新玩家，创建一条基础的数据，页面会引导进入创角页面
		user = &pb.DBUser{
			Uid:         id.NewXId(),
			Account:     req.Account,
			Password:    req.Password,
			Createip:    session.GetCache().Ip,
			Ctime:       time.Now().Unix(),
			Lastlogints: time.Now().Unix(),
		}
		if err = db.MySql().Insert(comm.TableUser, &user); err != nil {
			errdata = &pb.ErrorData{
				Code:    pb.ErrorCode_DBError,
				Message: err.Error(),
			}
			return
		}
	} else {
		if user.Password != req.Password {
			errdata = &pb.ErrorData{
				Code:    pb.ErrorCode_PasswordErr,
				Message: pb.ErrorCode_PasswordErr.String(),
			}
			return
		}
	}

	session.SetMate(comm.Session_User, &user)
	session.SendMsg(string(this.module.GetType()), "login", &pb.UserLoginResp{
		User: user,
	})
	event.TriggerEvent(comm.EventUserLogin, session)
	return
}
