package configure

import "time"

/*
系统 服务配置中心

*/

type (
	//配置文件信息
	FileInfo struct {
		Name    string    // 文件名
		Size    int64     // 文件大小
		ModTime time.Time // 文件修改时间
	}
	ISys interface {
		ConfigurePath() string
		Start() (err error)
		Stop() (err error)
		RegisterConfigure(name string, fn interface{}, callback func()) (err error) //注册配置
		GetConfigure(name string) (v interface{}, err error)                        //获取配置
	}
)

var defsys ISys

func OnInit(config map[string]interface{}, option ...Option) (err error) {
	var options Options
	if options, err = newOptions(config, option...); err != nil {
		return
	}
	defsys, err = newSys(options)
	return
}

func NewSys(option ...Option) (sys ISys, err error) {
	var options Options
	if options, err = newOptionsByOption(option...); err != nil {
		return
	}
	defsys, err = newSys(options)
	return
}
func ConfigurePath() string {
	return defsys.ConfigurePath()
}

func Start() (err error) {
	return defsys.Start()
}
func Stop() (err error) {
	return defsys.Stop()
}

func RegisterConfigure(name string, fn interface{}, callback func()) (err error) {
	return defsys.RegisterConfigure(name, fn, callback)
}

func GetConfigure(name string) (v interface{}, err error) {
	return defsys.GetConfigure(name)
}
