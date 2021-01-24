package appconfint

// IAppConfig  Application configuration interface
type IAppConfig interface {
	Save(ConfigFilePath string) error
	Load(ConfigFilePath string) (IAppConfig, error)
	SetDefaults()
	SetDefaultFunc(f func(Config IAppConfig))
	GetData() interface{}
	SetData(data interface{})
}
