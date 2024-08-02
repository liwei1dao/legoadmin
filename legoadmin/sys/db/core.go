package db

import (
	"sync"
	"time"

	"github.com/liwei1dao/lego/sys/mysql"
	"github.com/liwei1dao/lego/sys/redis"
)

/*
系统描述:数据模型系统,redis和mgo 存储策略的集成方案
*/

type (
	M    map[string]interface{}
	ISys interface {
		MySql() mysql.ISys
		Redis() redis.ISys
	}

	//过期数据
	ModelDataExpired struct {
		key     string              //主key
		mu      sync.RWMutex        //安全锁
		keys    map[string]struct{} //数据集合
		expired time.Time           //过期时间
	}
)

var (
	defsys ISys
)

func OnInit(config map[string]interface{}, opt ...Option) (err error) {
	var option *Options
	if option, err = newOptions(config, opt...); err != nil {
		return
	}
	defsys, err = newSys(option)
	return
}
func NewSys(opt ...Option) (sys ISys, err error) {
	var option *Options
	if option, err = newOptionsByOption(opt...); err != nil {
		return
	}
	sys, err = newSys(option)
	return
}

func MySql() mysql.ISys {
	return defsys.MySql()
}
func Redis() redis.ISys {
	return defsys.Redis()
}
