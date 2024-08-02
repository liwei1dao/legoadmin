package db

import (
	"context"
	"crypto/tls"

	"github.com/liwei1dao/lego/sys/log"
	"github.com/liwei1dao/lego/sys/mysql"

	"github.com/redis/go-redis/v9"
)

func newSys(options *Options) (sys *DB, err error) {
	sys = &DB{
		options: options,
	}
	err = sys.init()
	return
}

type DB struct {
	options    *Options
	admindb    mysql.ISys
	adminredis redis.UniversalClient
	gamedb     mysql.ISys
	gameredis  redis.UniversalClient
}

func (this *DB) init() (err error) {
	rconf := &redis.UniversalOptions{
		Addrs:    this.options.AdminRedisAddr,
		Password: this.options.AdminRedisPassword, // 如果有密码
		DB:       this.options.AdminRedisDB,
	}
	if this.options.AdminRedisTLS {
		rconf.TLSConfig = &tls.Config{}
	}
	// 使用集群模式
	this.adminredis = redis.NewUniversalClient(rconf)

	_, err = this.adminredis.Ping(context.Background()).Result()
	if err != nil {
		this.options.Log.Error(err.Error(), log.Field{Key: "options", Value: this.options})
		return
	}
	if this.admindb, err = mysql.NewSys(
		mysql.SetMySQLDsn(this.options.AdminMysqlDNS),
	); err != nil {
		this.options.Log.Error(err.Error(), log.Field{Key: "options", Value: this.options})
		return
	}
	if this.options.GameMysqlDNS != "" {
		if this.gamedb, err = mysql.NewSys(
			mysql.SetMySQLDsn(this.options.GameMysqlDNS),
		); err != nil {
			this.options.Log.Error(err.Error(), log.Field{Key: "options", Value: this.options})
			return
		}
	}
	if this.options.GameRedisAddr != nil && len(this.options.GameRedisAddr) > 0 {
		gconf := &redis.UniversalOptions{
			Addrs:    this.options.AdminRedisAddr,
			Password: this.options.AdminRedisPassword, // 如果有密码
			DB:       this.options.AdminRedisDB,
		}
		if this.options.GameRedisTLS {
			gconf.TLSConfig = &tls.Config{}
		}
		// 使用集群模式
		this.gameredis = redis.NewUniversalClient(gconf)

		_, err = this.gameredis.Ping(context.Background()).Result()
		if err != nil {
			this.options.Log.Error(err.Error(), log.Field{Key: "options", Value: this.options})
			return
		}
	}
	return
}

func (this *DB) AdminDB() mysql.ISys {
	return this.admindb
}
func (this *DB) GameDB() mysql.ISys {
	return this.gamedb
}
func (this *DB) AdminRedis() redis.UniversalClient {
	return this.adminredis
}
func (this *DB) GameRedis() redis.UniversalClient {
	return this.gameredis
}
