package appconf

import (
	"fmt"

	appconf "github.com/eshu0/appconfig/pkg"
)

//AppConfigHelper This struct is a helper forstoring config
type AppConfigHelper struct {
	Config   *appconf.AppConfig `json:"-"`
	FilePath string             `json:"-"`
}

//NewAppConfigHelper creates new server config
func NewAppConfigHelper() *AppConfigHelper {
	return NewAppConfigHelperWithDefault(nil)
}

//NewAppConfigHelperWithDefaults creates new apphelper with default function
func NewAppConfigHelperWithDefaults(DefaultFunc func()) *AppConfigHelper {
	conf := NewAppConfig()
	dc := &AppConfigHelper{}
	Config, ok := conf.(*appconf.AppConfig)
	if ok {
		dc.Config = Config
		if DefaultFunc != nil {
			dc.Config.SetDefaultFunc(DefaultFunc)
			dc.Config.SetDefaults()
		}
		return dc
	}
	return nil

}

//Save saves config to disk
func (ach *AppConfigHelper) Save() error {

	if ach.Config == nil {
		return fmt.Errorf("Config was nil")
	}

	if len(ach.FilePath) <= 0 {
		return fmt.Errorf("Config File Path was not set could not save")
	}

	if err := ach.Config.Save(ach.FilePath); err != nil {
		return err
	}
	return nil
}

//Load loads server config from disk
func (ach *AppConfigHelper) Load() error {

	if ach.Config == nil {
		return fmt.Errorf("Config was nil")
	}

	if len(ach.FilePath) <= 0 {
		return fmt.Errorf("Config File Path was not set could not load")
	}

	newconfig, err := ach.Config.Load(ach.FilePath)
	if err != nil {
		return err
	}

	if newconfig == nil {
		return fmt.Errorf("Loading resulted with a nil")
	}

	ccat, ok := newconfig.(*appconf.AppConfig)
	if ok {
		ach.Config = ccat
		return nil
	}
	return fmt.Errorf("Cast failed")
}
