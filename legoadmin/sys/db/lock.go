package db

import (
	"context"
	"errors"
	"time"

	"github.com/redis/go-redis/v9"
)

/*
Redis Scard 命令返回集合中元素的数量
*/
func (this *DB) NewGameRedisMutex(key string, opt ...RMutexOption) (result *RedisMutex, err error) {
	opts := newRMutexOptions(opt...)
	result = &RedisMutex{
		sys:    this.gameredis,
		key:    key,
		expiry: opts.expiry,
		delay:  opts.delay,
	}
	return
}

type RedisMutex struct {
	sys    redis.UniversalClient
	key    string
	expiry time.Duration //过期时间 单位秒
	delay  time.Duration
}

// 此接口未阻塞接口
func (this *RedisMutex) Lock() (err error) {
	wait := make(chan error)
	go func() {
		start := time.Now()
		for time.Now().Sub(start) <= this.expiry {
			if _, err := this.sys.SetEx(context.Background(), this.key, 1, this.expiry).Result(); err == nil {
				wait <- nil
				return
			} else if err == nil {
				time.Sleep(this.delay)
			} else {
				wait <- err
				return
			}
		}
		wait <- errors.New("time out")
	}()
	err = <-wait
	return
}

func (this *RedisMutex) Unlock() {
	this.sys.Del(context.Background(), this.key)
}
