package user

import (
	"legoadmin/modules"

	"github.com/liwei1dao/lego/utils/mapstructure"
)

type (
	IOptions interface {
		modules.IOptions
	}
	Options struct {
		modules.Options
	}
)

func (this *Options) LoadConfig(settings map[string]interface{}) (err error) {
	if settings != nil {
		if err = this.Options.LoadConfig(settings); err != nil {
			return
		}
		if err = mapstructure.Decode(settings, this); err != nil {
			return
		}
	}
	return
}
