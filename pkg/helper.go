package appconf

import (
	"fmt"

	appconfint "github.com/eshu0/appconfig/pkg/interfaces"
)

//AppConfigHelper This struct is a helper forstoring config
type AppConfigHelper struct {
	LoadedConfig *AppConfig `json:"-"`
	FilePath     string     `json:"-"`
}

//NewAppConfigHelper creates new server config
func NewAppConfigHelper(filepath string) *AppConfigHelper {
	return NewAppConfigHelperWithDefault(filepath, nil)
}

//NewAppConfigHelperWithDefault creates new apphelper with default function
func NewAppConfigHelperWithDefault(filepath string, DefaultFunc func(Config appconfint.IAppConfig)) *AppConfigHelper {
	// has a conifg file been provided?
	if len(filepath) == 0 || filepath == "" {
		// load this first
		filepath = DefaultFilePath
	}

	conf := NewAppConfig()
	dc := &AppConfigHelper{}
	dc.FilePath = filepath
	Config, ok := conf.(*AppConfig)
	if ok {
		dc.LoadedConfig = Config
		if DefaultFunc != nil {
			dc.LoadedConfig.SetDefaultFunc(DefaultFunc)
		}
		return dc
	}
	return nil

}

//Save saves config to disk
func (ach *AppConfigHelper) Save() error {

	if ach.LoadedConfig == nil {
		return fmt.Errorf("Loaded Config was nil")
	}

	if len(ach.FilePath) <= 0 {
		return fmt.Errorf("Config File Path was not set could not save")
	}

	if err := ach.LoadedConfig.Save(ach.FilePath); err != nil {
		return err
	}
	return nil
}

//Load loads server config from disk
func (ach *AppConfigHelper) Load() error {

	if ach.LoadedConfig == nil {
		return fmt.Errorf("Loaded Config was nil")
	}

	if len(ach.FilePath) <= 0 {
		return fmt.Errorf("Config File Path was not set could not load")
	}

	newconfig, err := ach.LoadedConfig.Load(ach.FilePath)
	if err != nil {
		return err
	}

	if newconfig == nil {
		return fmt.Errorf("Loading resulted with a nil")
	}

	ccat, ok := newconfig.(*AppConfig)
	if ok {
		ach.LoadedConfig = ccat
		return nil
	}
	return fmt.Errorf("AppConfig cast failed")
}
