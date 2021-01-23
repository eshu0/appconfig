package appconf

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	appconfint "github.com/eshu0/appconfig/pkg/interfaces"
)

//DefaultFilePath is the default path for the server config
const DefaultFilePath = "./config.json"

//AppConfig This struct is the configuration for the REST server
type AppConfig struct {
	//Inherit from Interface
	appconfint.IAppConfig

	//Items for storage
	Items map[string]interface{} `json:"Items,omitempty"`
}

//NewAppConfig creates a new configuation with default settings
func NewAppConfig() appconfint.IAppConfig {
	Config := AppConfig{}
	Config.Items = make(map[string]interface{})
	return &Config
}

//GetItem gets the data fr that keyed item
func (rsc *AppConfig) GetItem(key string) interface{} {
	return Config.Items[key]
}

//SetItem sets the data for this keyed item
func (rsc *AppConfig) SetItem(key string, data interface{}) {
	Config.Items[key] = data
}

//SetDefaults sets the defaults this is a noop
func (rsc *AppConfig) SetDefaults() {

}

//Save This saves the configuration from a file path
func (rsc *AppConfig) Save(ConfigFilePath string) error {
	bytes, err1 := json.MarshalIndent(rsc, "", "\t") //json.Marshal(p)
	if err1 != nil {
		//Log.LogErrorf("SaveToFile()", "Marshal json for %s failed with %s ", ConfigFilePath, err1.Error())
		return err1
	}

	err2 := ioutil.WriteFile(ConfigFilePath, bytes, 0644)
	if err2 != nil {
		//Log.LogErrorf("SaveToFile()", "Saving %s failed with %s ", ConfigFilePath, err2.Error())
		return err2
	}

	return nil

}

//Load This loads the configuration from a file path
func (rsc *AppConfig) Load(ConfigFilePath string) (appconfint.IAppConfig, error) {
	ok, err := rsc.checkFileExists(ConfigFilePath)
	if ok {
		bytes, err1 := ioutil.ReadFile(ConfigFilePath) //ReadAll(jsonFile)
		if err1 != nil {
			return nil, fmt.Errorf("Reading '%s' failed with %s ", ConfigFilePath, err1.Error())
		}

		rserverconfig := RServerConfig{}

		err2 := json.Unmarshal(bytes, &rserverconfig)

		if err2 != nil {
			return nil, fmt.Errorf("Loading %s failed with %s ", ConfigFilePath, err2.Error())
		}

		//Log.LogDebugf("LoadFile()", "Read Port %s ", rserverconfig.Port)
		//rs.Log.LogDebugf("LoadFile()", "Port in config %s ", rs.Config.Port)
		return &rserverconfig, nil
	}

	if err != nil {
		return nil, fmt.Errorf("'%s' was not found to load with error: %s", ConfigFilePath, err.Error())
	}

	return nil, fmt.Errorf("'%s' was not found to load", ConfigFilePath)
}

func (rsc *RServerConfig) checkFileExists(filename string) (bool, error) {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false, err
	}
	return !info.IsDir(), nil
}
