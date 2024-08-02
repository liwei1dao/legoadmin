package db

import (
	"sync"
	"time"

	"github.com/liwei1dao/lego/sys/mysql"

	"github.com/redis/go-redis/v9"
)

/*
系统描述:数据模型系统,redis和mgo 存储策略的集成方案
*/

type (
	M    map[string]interface{}
	ISys interface {
		AdminDB() mysql.ISys
		AdminRedis() redis.UniversalClient
		GameDB() mysql.ISys
		GameRedis() redis.UniversalClient
		NewGameRedisMutex(key string, opt ...RMutexOption) (result *RedisMutex, err error)
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

func AdminDB() mysql.ISys {
	return defsys.AdminDB()
}
func AdminRedis() redis.UniversalClient {
	return defsys.AdminRedis()
}

func GameDB() mysql.ISys {
	return defsys.GameDB()
}
func GameRedis() redis.UniversalClient {
	return defsys.GameRedis()
}

/*Lock*/
func NewGameRedisMutex(key string, opt ...RMutexOption) (result *RedisMutex, err error) {
	return defsys.NewGameRedisMutex(key, opt...)
}
