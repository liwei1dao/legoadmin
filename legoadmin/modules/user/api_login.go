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
func (this *apiComp) Login(session comm.IUserSession, req *pb.UserLoginReq) (errdata *pb.ErrorData) {
	var (
		user pb.DBUser
		err  error
	)
	if err = db.MySql().FindOne(comm.TableUser, &user, db.M{"binduid": req.Account}); err != nil && err != mysql.ErrNoDocuments {
		errdata = &pb.ErrorData{
			Code:    pb.ErrorCode_DBError,
			Message: err.Error(),
		}
		this.module.Errorln(err)
		return
	}
	if err == mysql.ErrNoDocuments { //创建新的账号
		//如果是新玩家，创建一条基础的数据，页面会引导进入创角页面
		user = pb.DBUser{
			Uid:      id.NewXId(),
			Binduid:  req.Account,
			Createip: session.GetCache().Ip,
			Ctime:    time.Now().Unix(),
			Balance:  0,
		}
		if err = db.MySql().Insert(comm.TableUser, &user); err != nil {
			errdata = &pb.ErrorData{
				Code:    pb.ErrorCode_DBError,
				Message: err.Error(),
			}
			this.module.Errorln(err)
			return
		}
	}
	session.SetMate(comm.Session_User, &user)
	session.SendMsg(string(this.module.GetType()), "login", &pb.UserLoginResp{
		Uid:           user.Uid,
		Agentid:       user.Agentid,
		Playeraccount: user.Binduid,
		Playername:    user.Playname,
	})
	event.TriggerEvent(comm.EventUserLogin, session)
	return
}
