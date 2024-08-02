package api

import (
	"context"
	"legoadmin/comm"
	"legoadmin/pb"
	"time"

	"github.com/liwei1dao/lego/sys/log"

	"github.com/golang-jwt/jwt"
)

// 创建代理
func (this *apiComp) Login(ctx context.Context, req *pb.ApiLoginReq) (resp *pb.ApiLoginResp, errdata *pb.ErrorData) {
	var (
		model       *pb.DBApiUser
		tokenString string
		err         error
	)

	if model, err = this.module.model.findByAccount(req.Account); err != nil {
		errdata = &pb.ErrorData{
			Code:    pb.ErrorCode_ReqParameterError,
			Message: pb.ErrorCode_ReqParameterError.String(),
		}
		this.module.Errorln(err)
		return
	}
	if model.Password != req.Password {
		errdata = &pb.ErrorData{
			Code:    pb.ErrorCode_ReqParameterError,
			Message: pb.ErrorCode_ReqParameterError.String(),
		}
		return
	}

	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &comm.TokenClaims{
		Account:  req.Account,
		Identity: model.Identity,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	if tokenString, err = token.SignedString([]byte(this.options.TokenKey)); err != nil {
		errdata = &pb.ErrorData{
			Code:    pb.ErrorCode_ReqParameterError,
			Message: pb.ErrorCode_ReqParameterError.String(),
		}
		this.module.Errorln(err)
		return
	}
	this.module.Debug("Login", log.Field{Key: "token", Value: tokenString})
	resp = &pb.ApiLoginResp{
		Account:  req.Account,
		Identity: model.Identity,
		Token:    tokenString,
	}
	return
}
