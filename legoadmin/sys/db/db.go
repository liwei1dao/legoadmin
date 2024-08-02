package db

import (
	"github.com/liwei1dao/lego/sys/mysql"
	"github.com/liwei1dao/lego/sys/redis"
)

func newSys(options *Options) (sys *DB, err error) {
	sys = &DB{
		options: options,
	}
	err = sys.init()
	return
}

type DB struct {
	options *Options
	mysql   mysql.ISys
	redis   redis.ISys
}

func (this *DB) init() (err error) {
	if this.redis, err = redis.NewSys(
		redis.SetRedisAddr(this.options.RedisAddr),
		redis.SetRedisPassword(this.options.RedisPassword),
		redis.SetRedisTLS(this.options.RedisTLS),
		redis.SetRedisDB(this.options.RedisDB),
	); err != nil {
		this.options.Log.Errorln(err)
		return
	}
	if this.mysql, err = mysql.NewSys(
		mysql.SetMySQLDsn(this.options.MysqlDNS),
	); err != nil {
		this.options.Log.Errorln(err)
		return
	}
	return
}

func (this *DB) MySql() mysql.ISys {
	return this.mysql
}

func (this *DB) Redis() redis.ISys {
	return this.redis
}
