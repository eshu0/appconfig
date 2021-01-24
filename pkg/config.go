package appconf

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	appconfint "github.com/eshu0/appconfig/pkg/interfaces"
)

//DefaultFilePath is the default path for the configuration file
const DefaultFilePath = "./config.json"

//DefaultVersion the default version of the config
const DefaultVersion = "0.0.0.1"

//AppConfig This struct is the configuration for the REST server
type AppConfig struct {
	Version string `json:"version,omitempty"`

	//Inherit from Interface
	appconfint.IAppConfig `json:"-"`

	//Data for storage
	Data interface{} `json:"data,omitempty"`

	DefaultFunction func(Config appconfint.IAppConfig) `json:"-"`
}

//NewAppConfig creates a new configuation with default settings
func NewAppConfig() appconfint.IAppConfig {
	Config := &AppConfig{}
	Config.Version = DefaultVersion
	Config.Data = nil
	return Config
}

//GetData gets the data fr that keyed item
func (Config *AppConfig) GetData() interface{} {
	return Config.Data
}

//SetData sets the data for this keyed item
func (Config *AppConfig) SetData(data interface{}) {
	Config.Data = data
}

//SetDefaults sets the defaults based off the function
func (Config *AppConfig) SetDefaults() {
	if Config.DefaultFunction != nil {
		Config.DefaultFunction(Config)
	}
}

//SetDefaultFunc sets the defaults this is a noop
func (Config *AppConfig) SetDefaultFunc(f func(Config appconfint.IAppConfig)) {
	Config.DefaultFunction = f
}

//Save This saves the configuration from a file path
func (Config *AppConfig) Save(ConfigFilePath string) error {
	bytes, err1 := json.MarshalIndent(Config, "", "\t") //json.Marshal(p)
	if err1 != nil {
		return err1
	}

	err2 := ioutil.WriteFile(ConfigFilePath, bytes, 0644)
	if err2 != nil {
		return err2
	}

	return nil

}

//Load This loads the configuration from a file path
func (Config *AppConfig) Load(ConfigFilePath string) (appconfint.IAppConfig, error) {
	ok, err := Config.checkFileExists(ConfigFilePath)
	if ok {
		bytes, err1 := ioutil.ReadFile(ConfigFilePath) //ReadAll(jsonFile)
		if err1 != nil {
			return nil, fmt.Errorf("Reading '%s' failed with %s ", ConfigFilePath, err1.Error())
		}

		appconfig := NewAppConfig()

		err2 := json.Unmarshal(bytes, appconfig)

		if err2 != nil {
			return nil, fmt.Errorf("Loading %s failed with %s ", ConfigFilePath, err2.Error())
		}

		fmt.Printf("appconfig %v\n", appconfig)

		//Log.LogDebugf("LoadFile()", "Read Port %s ", rserverconfig.Port)
		//rs.Log.LogDebugf("LoadFile()", "Port in config %s ", rs.Config.Port)
		return appconfig, nil
	}

	if err != nil {
		return nil, fmt.Errorf("'%s' was not found to load with error: %s", ConfigFilePath, err.Error())
	}

	return nil, fmt.Errorf("'%s' was not found to load", ConfigFilePath)
}

func (Config *AppConfig) checkFileExists(filename string) (bool, error) {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false, err
	}
	return !info.IsDir(), nil
}
