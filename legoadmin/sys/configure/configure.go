package configure

import (
	"fmt"
	"io"
	"io/fs"
	"io/ioutil"

	"github.com/liwei1dao/lego"
	"github.com/liwei1dao/lego/sys/log"
	"github.com/liwei1dao/lego/utils/codec"
	"github.com/liwei1dao/lego/utils/codec/json"

	"os"
	"path"
	"reflect"
	"sync"
	"time"
)

var typeOfIn = reflect.TypeOf(([]map[string]interface{})(nil)).Elem()
var typeOfError = reflect.TypeOf((*error)(nil)).Elem()

type configurehandle struct {
	configureType reflect.Type
	fn            reflect.Value
	events        []func()
}

func newSys(options Options) (sys *Configure, err error) {
	sys = &Configure{
		options:          options,
		closeSignal:      make(chan struct{}),
		configurehandles: make(map[string]*configurehandle),
		configure:        make(map[string]interface{}),
		fileinfos:        make(map[string]*FileInfo),
	}
	return
}

type Configure struct {
	options          Options
	closeSignal      chan struct{}
	hlock            sync.RWMutex
	configurehandles map[string]*configurehandle
	clock            sync.RWMutex
	configure        map[string]interface{}
	fileinfos        map[string]*FileInfo
	offsettime       time.Duration //偏移时间
}

func (this *Configure) ConfigurePath() string {
	return this.options.ConfigurePath
}

func (this *Configure) Start() (err error) {
	timer := time.NewTicker(time.Second * time.Duration(this.options.CheckInterval))
	go func() {
	locp:
		for {
			select {
			case <-this.closeSignal:
				break locp
			case <-timer.C:
				this.checkConfigure()
			}
		}
		timer.Stop()
	}()
	this.readoffsettime()
	return
}
func (this *Configure) Stop() (err error) {
	this.closeSignal <- struct{}{}
	return
}

func (this *Configure) Update() {
	this.checkConfigure()
	this.readoffsettime()
}

// 加载配置文件
func (this *Configure) RegisterConfigure(name string, fn interface{}, callback func()) (err error) {
	this.hlock.RLock()
	handle, ok := this.configurehandles[name]
	this.hlock.RUnlock()
	if ok {
		// err = fmt.Errorf("重复 注册配置【%s】", name)
		if callback != nil {
			handle.events = append(handle.events, callback)
			callback()
		}
		return
	}
	fnvalue := reflect.ValueOf(fn)
	if fnvalue.Type().NumIn() != 1 {
		err = fmt.Errorf("LoadConfigure fn 类型错误! 只接受fn( _buf []map[string]interface{})(v,error) 函数参数")
		return
	}
	inType := fnvalue.Type().In(0)
	if inType.Elem() != typeOfIn {
		err = fmt.Errorf("LoadConfigure fn 类型错误! 只接受fn( _buf []map[string]interface{})(v,error) 函数参数")
		return
	}
	if fnvalue.Type().NumOut() != 2 {
		err = fmt.Errorf("LoadConfigure fn 类型错误! 只接受fn( _buf []map[string]interface{})(v,error) 函数参数")
		return
	}
	dataType := fnvalue.Type().Out(0)
	if dataType.Kind() != reflect.Ptr {
		err = fmt.Errorf("LoadConfigure fn 类型错误! 只接受fn( _buf []map[string]interface{})(v,error) 函数参数")
		return
	}
	if returnType := fnvalue.Type().Out(1); returnType != typeOfError {
		err = fmt.Errorf("LoadConfigure fn 类型错误! 只接受fn( _buf []map[string]interface{})(v,error) 函数参数")
		return
	}
	handle = &configurehandle{
		configureType: dataType,
		fn:            fnvalue,
		events:        []func(){callback},
	}
	if err = this.loaderConfigure(name, handle); err != nil {
		log.Errorf("loaderConfigure name:%s err:%v", name, err)
		return
	}
	this.hlock.Lock()
	this.configurehandles[name] = handle
	this.hlock.Unlock()
	if callback != nil {
		callback()
	}

	return
}

// 读取配置文件
func (this *Configure) GetConfigure(name string) (v interface{}, err error) {
	this.clock.RLock()
	v, ok := this.configure[name]
	this.clock.RUnlock()
	if !ok {
		err = fmt.Errorf("no LoadConfigure:%s", name)
	}
	return
}

// 加载配置文件
func (this *Configure) loaderConfigure(name string, handle *configurehandle) (err error) {
	var (
		fliepath     string
		fileInfo     fs.FileInfo
		file         *os.File
		bytes        []byte
		data         []map[string]interface{}
		returnValues []reflect.Value
	)

	fliepath = path.Join(this.options.ConfigurePath, name)
	if fileInfo, err = os.Stat(fliepath); err != nil {
		err = fmt.Errorf("no found file:%s", fliepath)
		return
	}

	if file, err = os.Open(fliepath); err != nil {
		err = fmt.Errorf("no found file:%s", fliepath)
		return
	}
	defer file.Close()
	if bytes, err = io.ReadAll(file); err != nil {
		err = fmt.Errorf("read file:%s err:%v", fliepath, err)
		return
	}
	if err = json.Unmarshal(bytes, &data); err != nil {
		err = fmt.Errorf("read file:%s json.Unmarshal err:%v", fliepath, err)
		return
	}
	returnValues = handle.fn.Call([]reflect.Value{reflect.ValueOf(data)})
	errInter := returnValues[1].Interface()
	if errInter != nil {
		err = fmt.Errorf("read file:%s load.fn err:%v", fliepath, errInter)
		return
	}
	this.clock.Lock()
	this.configure[name] = returnValues[0].Interface()
	this.fileinfos[fileInfo.Name()] = &FileInfo{Name: name, Size: fileInfo.Size(), ModTime: fileInfo.ModTime()}
	this.clock.Unlock()
	return
}

// 检查配置文件是否有更新
func (this *Configure) checkConfigure() {
	// log.Debug("Check Configure Update")
	if dir, err := os.ReadDir(this.options.ConfigurePath); err != nil {
		log.Errorf("[Configure Sys] checkConfigure err:%v", err)
	} else {
		for _, fi := range dir {
			if !fi.IsDir() { //不处理目录代码
				this.clock.RLock()
				v, ok := this.fileinfos[fi.Name()]
				this.clock.RUnlock()
				f, _ := fi.Info()
				if ok && f.ModTime().After(v.ModTime) {
					this.hlock.RLock()
					handle := this.configurehandles[v.Name]
					this.hlock.RUnlock()
					if err = this.loaderConfigure(v.Name, handle); err != nil {
						log.Errorln(err)
						return
					}
					log.Debug("UpDate Configure", log.Field{Key: "table", Value: v.Name})
					v.ModTime = f.ModTime() //重置配置文件修改时间
					for _, v := range handle.events {
						if v != nil {
							go func(f func()) {
								defer lego.Recover("configure")
								f()
							}(v)

						}
					}
				}
			}
		}
	}
	this.readoffsettime()
}

// 读取偏移时间
func (this *Configure) readoffsettime() {
	var (
		file       *os.File
		data       []byte
		offtimeStr string
		err        error
	)
	if file, err = os.Open(this.options.TimestampFile); err != nil {
		// log.Errorf("[Configure Sys] readoffsettime err:%v", err)
		return
	}
	defer file.Close()
	if data, err = ioutil.ReadAll(file); err != nil {
		// log.Errorf("[Configure Sys] readoffsettime err:%v", err)
		return
	}
	offtimeStr = codec.BytesToString(data)
	this.offsettime = time.Second * time.Duration(codec.StringToFloat64(offtimeStr))
}

// 写入开服时间偏移
func (this *Configure) writeoffsettime(offset time.Duration) {
	var (
		file *os.File
		err  error
	)
	if file, err = os.OpenFile(this.options.TimestampFile, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666); err != nil {
		// log.Errorf("[Configure Sys] writeoffsettime err:%v", err)
		return
	}
	defer file.Close()
	file.WriteString(fmt.Sprintf("%f", offset.Seconds()))
}
