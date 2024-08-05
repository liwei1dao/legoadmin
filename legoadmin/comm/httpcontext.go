package comm

import (
	"context"
	"sync"
)

func NewHttpContext() IHttpContext {
	return &HttpContext{}
}

// 用户会话
type HttpContext struct {
	context.Context
	lock sync.RWMutex
	mate map[string]interface{}
}

// 重置
func (this *HttpContext) SetSession(ctx context.Context) {
	this.Context = ctx
	this.mate = make(map[string]interface{})
}

// 重置
func (this *HttpContext) Reset() {
	this.mate = make(map[string]interface{})
}

// 写入元数据
func (this *HttpContext) SetMate(name string, value interface{}) {
	this.lock.Lock()
	this.mate[name] = value
	this.lock.Unlock()
}

// 写入元数据
func (this *HttpContext) GetMate(name string) (ok bool, value interface{}) {
	this.lock.RLock()
	value, ok = this.mate[name]
	this.lock.RUnlock()
	return
}
