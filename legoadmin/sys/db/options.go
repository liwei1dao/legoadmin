package db

import (
	"errors"
	"time"

	"github.com/liwei1dao/lego/sys/log"
	"github.com/liwei1dao/lego/utils/mapstructure"
)

type Option func(*Options)
type Options struct {
	AdminMysqlDNS      string   //平台MySql 数据库
	AdminRedisAddr     []string //平台Rdis 的集群地址
	AdminRedisPassword string   //平台Rdis 密码
	AdminRedisTLS      bool     //平台Rdis 是否开启tls
	AdminRedisDB       int      //平台Rdis DB
	GameMysqlDNS       string   //游戏MySql 游戏数据库
	GameRedisAddr      []string //游戏Rdis 的集群地址
	GameRedisPassword  string   //游戏Rdis 的密码
	GameRedisTLS       bool     //游戏Rdis 是否开启tls
	GameRedisDB        int      //游戏Rdis 数据库位置
	Debug              bool     //日志是否开启
	Log                log.ILogger
}

func SetDebug(v bool) Option {
	return func(o *Options) {
		o.Debug = v
	}
}

func SetLog(v log.ILogger) Option {
	return func(o *Options) {
		o.Log = v
	}
}
func newOptions(config map[string]interface{}, opts ...Option) (options *Options, err error) {
	options = &Options{}
	if config != nil {
		if err = mapstructure.Decode(config, options); err != nil {
			return
		}
	}
	for _, o := range opts {
		o(options)
	}
	if options.Log = log.NewTurnlog(options.Debug, log.Clone("sys.db", 3)); options.Log == nil {
		err = errors.New("log is nil")
	}
	return
}

func newOptionsByOption(opts ...Option) (options *Options, err error) {
	options = &Options{}
	for _, o := range opts {
		o(options)
	}
	if options.Log = log.NewTurnlog(options.Debug, log.Clone("sys.db", 3)); options.Log == nil {
		err = errors.New("log is nil")
	}
	return
}

type DBOption func(*DBOptions)
type DBOptions struct {
	IsSyncRedis bool //是否写redis
}

// 设置是否写mgor日志
func SetIsSyncRedis(v bool) DBOption {
	return func(o *DBOptions) {
		o.IsSyncRedis = v
	}
}

// 更具 Option 序列化 系统参数对象
func NewDBOption(opts ...DBOption) DBOptions {
	options := DBOptions{
		IsSyncRedis: true,
	}
	for _, o := range opts {
		o(&options)
	}
	return options
}

type RMutexOption func(*RMutexOptions)
type RMutexOptions struct {
	expiry time.Duration
	delay  time.Duration
}

func SetExpiry(v time.Duration) RMutexOption {
	return func(o *RMutexOptions) {
		o.expiry = v
	}
}
func Setdelay(v time.Duration) RMutexOption {
	return func(o *RMutexOptions) {
		o.delay = v
	}
}

func newRMutexOptions(opts ...RMutexOption) RMutexOptions {
	opt := RMutexOptions{
		expiry: 5,
		delay:  time.Millisecond * 50,
	}
	for _, o := range opts {
		o(&opt)
	}
	return opt
}
