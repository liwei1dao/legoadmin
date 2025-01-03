package modules

import (
	"fmt"
	"legoadmin/comm"
	"log"
	"reflect"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/liwei1dao/lego/base/cluster"
	"github.com/liwei1dao/lego/core"
	"github.com/liwei1dao/lego/core/cbase"
)

/*
模块 网关组件 接收处理用户传递消息
*/
type MCompHttpGate struct {
	cbase.ModuleCompBase
	service cluster.IClusterService //rpc服务对象
	module  core.IModule            //当前业务模块
	comp    core.IModuleComp        //网关组件自己
	scomp   comm.ISC_HttpRouteComp
}

// 组件初始化接口
func (this *MCompHttpGate) Init(service core.IService, module core.IModule, comp core.IModuleComp, options core.IModuleOptions) (err error) {
	this.ModuleCompBase.Init(service, module, comp, options)
	this.service = service.(cluster.IClusterService)
	this.module = module
	this.comp = comp
	return
}

// 组件启动接口，启动时将自己接收用户消息的处理函数注册到services/comp_gateroute.go 对象中
func (this *MCompHttpGate) Start() (err error) {
	if err = this.ModuleCompBase.Start(); err != nil {
		return
	}
	var comp core.IServiceComp
	//注册远程路由
	if comp, err = this.service.GetComp(comm.SC_ServiceHttpRouteComp); err != nil {
		return
	}
	this.scomp = comp.(comm.ISC_HttpRouteComp)
	this.suitableMethods()
	return
}

// 反射注册相关接口道services/comp_gateroute.go 对象中
func (this *MCompHttpGate) suitableMethods() {
	typ := reflect.TypeOf(this.comp)
	for m := 0; m < typ.NumMethod(); m++ {
		method := typ.Method(m)
		mname := method.Name
		if mname == "Start" ||
			mname == "Init" ||
			mname == "Destroy" ||
			strings.HasSuffix(mname, "Check") {
			continue
		}
		this.reflectionRouteHandle(typ, method)
	}
}

// 反射注册路由处理函数
func (this *MCompHttpGate) reflectionRouteHandle(typ reflect.Type, method reflect.Method) (ret bool) {
	mtype := method.Type
	mname := method.Name
	if method.PkgPath != "" {
		log.Panicf("反射注册用户处理函数错误 [%s-%s] Api接口格式错误", this.module.GetType(), mname)
		return
	}
	if mtype.NumIn() != 3 {
		log.Panicf("反射注册用户处理函数错误 [%s-%s] Api接口格式错误", this.module.GetType(), mname)
		return
	}
	contextType := mtype.In(1)
	if !contextType.Implements(typeOfHttpContext) {
		log.Panicf("反射注册用户处理函数错误 [%s-%s] Api接口格式错误", this.module.GetType(), mname)
		return
	}
	agrType := mtype.In(2)
	if !agrType.Implements(typeOfMessage) {
		log.Panicf("反射注册用户处理函数错误 [%s-%s] Api接口格式错误", this.module.GetType(), mname)
		return
	}
	if mtype.NumOut() != 2 {
		log.Panicf("反射注册用户处理函数错误 [%s-%s] Api接口格式错误", this.module.GetType(), mname)
		return
	}
	returnDataType := mtype.Out(0)
	if !returnDataType.Implements(typeOfMessage) {
		log.Panicf("反射注册用户处理函数错误 [%s-%s] Api接口格式错误", this.module.GetType(), mname)
		return
	}

	returnErrDataType := mtype.Out(1)
	if returnErrDataType != typeOfErrorData {
		log.Panicf("反射注册用户处理函数错误 [%s-%s] Api接口格式错误", this.module.GetType(), mname)
		return
	}

	//注册路由函数
	this.scomp.RegisterRoute(fmt.Sprintf("%s.%s", this.module.GetType(), strings.ToLower(mname)), reflect.ValueOf(this.comp), agrType, method)

	return true
}

func (this *MCompHttpGate) isExportedOrBuiltinType(t reflect.Type) bool {
	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return this.isExported(t.Name()) || t.PkgPath() == ""
}

func (this *MCompHttpGate) isExported(name string) bool {
	rune, _ := utf8.DecodeRuneInString(name)
	return unicode.IsUpper(rune)
}
