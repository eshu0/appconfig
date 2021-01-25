package appconf

import (
	"fmt"
	"testing"

	appconfint "github.com/eshu0/appconfig/pkg/interfaces"
)

//DummyConfigController This struct is the configuration for the REST server
type DummyConfigController struct {
	Helper *AppConfigHelper `json:"-"`
	cache  *ConfigData      `json:"-"`
}

//ComplexData some data
type ComplexData struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

//ConfigData the data to be stored
type ConfigData struct {
	Port              string        `json:"port,omitempty"`
	Handlers          []ComplexData `json:"handlers,omitempty"`
	DefaultHandlers   []ComplexData `json:"defaulthandlers,omitempty"`
	TemplateFilepath  string        `json:"templatefilepath,omitempty"`
	TemplateFileTypes []string      `json:"templatefiletypes,omitempty"`
	CacheTemplates    bool          `json:"cachetemplates,omitempty"`
}

//NewDummyConfigController creates new server config
func NewDummyConfigController(filepath string) *DummyConfigController {
	dc := &DummyConfigController{}
	helper := NewAppConfigHelperWithDefault(filepath, dc.SetServerDefaultConfig)

	if helper != nil {
		dc.Helper = helper
		// we call this after the helper has been set!
		dc.Helper.LoadedConfig.SetDefaults()
	}

	return dc

}

//Load Loads the config from disk
func (rsc *DummyConfigController) Load(t *testing.T) error {
	fmt.Printf("conf before load %v\n", rsc.Helper.LoadedConfig)

	// load the data
	if err := rsc.Helper.Load(); err != nil {
		return err
	}

	fmt.Printf("conf after load %v\n", rsc.Helper.LoadedConfig)

	// reset the cache
	rsc.cache = nil

	// this rebuilds the cache
	rsc.GetConfigData(t)

	fmt.Printf("conf after data config %v\n", rsc.Helper.LoadedConfig)

	return nil
}

//SetServerDefaultConfig ets the defult items
func (rsc *DummyConfigController) SetServerDefaultConfig(Config appconfint.IAppConfig) {

	Data := &ConfigData{}
	Data.DefaultHandlers = []ComplexData{}
	Data.Handlers = []ComplexData{}
	Data.Port = "7777"
	Data.TemplateFileTypes = []string{".tmpl", ".html"}
	Data.CacheTemplates = false

	rsc.SetConfigData(Data)
}

//GetConfigData returns the config data from the store
func (rsc *DummyConfigController) GetConfigData(t *testing.T) *ConfigData {
	if rsc.cache == nil {
		fmt.Println("cache is nil")
		data := rsc.Helper.LoadedConfig.GetData()
		fmt.Printf("data %v\n", data)
		Config, ok := data.(*ConfigData) //(map[string]*ConfigData)
		if ok {
			fmt.Printf("cast ok %v\n", Config)
			rsc.cache = Config
			return Config
		}

		Config1, ok1 := data.(ConfigData) //(map[string]*ConfigData)
		if ok1 {
			fmt.Printf("cast1 ok %v\n", Config1)
			rsc.cache = &Config1
			return &Config1
		}

		t.Fatalf(`GetConfigData %v should cast ok`, data)
		/*
			Config2, ok2 := data.((map[string]*ConfigData))
			if ok2 {
				fmt.Printf("cast2 ok %v\n", Config2)
			}

			Config3, ok3 := data.((map[string]interface{}))
			if ok3 {
				fmt.Printf("cast3 ok %v\n", Config3)
				for key, element := range Config3 {
					fmt.Println("Key:", key, "=>", "Element:", element)
				}
				Config4, ok4 := Config3["Data"].(*ConfigData)
				if ok4 {
					fmt.Printf("cast4 ok %v\n", Config4)
					rsc.cache = Config4
					return Config4
				}
			}
		*/
		fmt.Printf("cast failed %v\n", Config)
		return nil

	}
	return rsc.cache

}

//SetConfigData sets the config data to the store
func (rsc *DummyConfigController) SetConfigData(data *ConfigData) {

	// reset the cache
	rsc.cache = nil
	fmt.Printf("before %v\n", rsc.Helper.LoadedConfig)

	// set the data ietm
	rsc.Helper.LoadedConfig.SetData(data)

	fmt.Printf("after %v\n", rsc.Helper.LoadedConfig)

}

//HasTemplate returns if a teplate path has been set
func (rsc *DummyConfigController) HasTemplate(t *testing.T) bool {
	d := rsc.GetConfigData(t)
	if d == nil {
		panic("HasTemplate - config data was nil")
	}

	return &(d.TemplateFilepath) != nil && len(d.TemplateFilepath) > 0
}

//GetTemplatePath returns the template path
func (rsc *DummyConfigController) GetTemplatePath(t *testing.T) string {
	d := rsc.GetConfigData(t)
	if d == nil {
		panic("GetTemplatePath - config data was nil")
	}
	return d.TemplateFilepath
}

//GetCacheTemplates returns the cached template paths
func (rsc *DummyConfigController) GetCacheTemplates(t *testing.T) bool {
	d := rsc.GetConfigData(t)
	if d == nil {
		panic("GetCacheTemplates - config data was nil")
	}
	return d.CacheTemplates
}

//GetTemplateFileTypes returns the file types for the templates, such as .tmpl, .html
func (rsc *DummyConfigController) GetTemplateFileTypes(t *testing.T) []string {
	d := rsc.GetConfigData(t)
	if d == nil {
		panic("GetTemplateFileTypes - config data was nil")
	}
	return d.TemplateFileTypes
}

//GetHandlersLen this gets length handlers
func (rsc *DummyConfigController) GetHandlersLen(t *testing.T) int {
	d := rsc.GetConfigData(t)
	if d == nil {
		panic("GetHandlersLen - config data was nil")
	}
	return len(d.Handlers)
}

//GetHandlers this gets the handlers from the config
func (rsc *DummyConfigController) GetHandlers(t *testing.T) []ComplexData {
	d := rsc.GetConfigData(t)
	if d == nil {
		panic("GetHandlers - config data was nil")
	}
	return d.Handlers
}

//GetDefaultHandlers this gets the default handlers
func (rsc *DummyConfigController) GetDefaultHandlers(t *testing.T) []ComplexData {
	d := rsc.GetConfigData(t)
	if d == nil {
		panic("GetDefaultHandlers - config data was nil")
	}
	return d.DefaultHandlers
}

//GetDefaultHandlersLen this gets length default handlers
func (rsc *DummyConfigController) GetDefaultHandlersLen(t *testing.T) int {
	d := rsc.GetConfigData(t)
	if d == nil {
		panic("GetDefaultHandlersLen - config data was nil")
	}
	return len(d.DefaultHandlers)
}

//GetAddress this gets the server address
func (rsc *DummyConfigController) GetAddress(t *testing.T) string {
	d := rsc.GetConfigData(t)
	if d == nil {
		panic("GetAddress - config data was nil")
	}
	return ":" + d.Port
}

//AddDefaultHandler this adds a default handler to the configuration
func (rsc *DummyConfigController) AddDefaultHandler(t *testing.T, Handler ComplexData) {
	d := rsc.GetConfigData(t)
	if d != nil {
		handlers := d.DefaultHandlers
		handlers = append(handlers, Handler)
		d.DefaultHandlers = handlers
		rsc.SetConfigData(d)
	} else {
		panic("AddHandler - config data was nil")
	}
}

//AddHandler this adds a handler to the configuration
func (rsc *DummyConfigController) AddHandler(t *testing.T, Handler ComplexData) {
	d := rsc.GetConfigData(t)
	if d != nil {
		handlers := d.Handlers
		handlers = append(handlers, Handler)
		d.Handlers = handlers
		rsc.SetConfigData(d)
	} else {
		panic("AddHandler - config data was nil")
	}
}

//SaveConfig saves server config to disk
func TestSaveConfig(t *testing.T) {
	conf := NewDummyConfigController("")

	if err := conf.Helper.Save(); err != nil {
		t.Fatalf(`Save(%s) = %v should not error`, conf.Helper.FilePath, err)
		return
	}
}

//LoadConfig loads server config from disk
func TestLoadConfig(t *testing.T) {

	conf := NewDummyConfigController("")

	if err := conf.Load(t); err != nil {
		t.Fatalf(`Load(%s) = %v should not error`, conf.Helper.FilePath, err)
		return
	}
	CheckDetails(t, conf)
}

func CheckDetails(t *testing.T, Config *DummyConfigController) {
	Config.GetAddress(t)
	/*
		rs.LogInfof("PrintDetails", "Address: %s", rs.Config.GetAddress())
		rs.LogInfof("PrintDetails", "Template Filepath: %s", rs.Config.GetTemplatePath())
		rs.LogInfof("PrintDetails", "Template FileTypes: %s", strings.Join(rs.Config.GetTemplateFileTypes(), ","))
		rs.LogInfof("PrintDetails", "Cache Templates: %t", rs.Config.GetCacheTemplates())
		rs.LogInfo("PrintDetails", "Handlers: ")
		for _, handl := range rs.Config.GetHandlers() {
			rs.LogInfof("PrintDetails", "Handler: %s", handl.MethodName)
		}

		rs.LogInfo("PrintDetails", "DefaultHandlers: ")

		for _, handl := range rs.Config.GetDefaultHandlers() {
			rs.LogInfof("PrintDetails", "Default Handler: %s", handl.MethodName)
		}
	*/
}
