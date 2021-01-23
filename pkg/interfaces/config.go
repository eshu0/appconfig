package appconfint

// IAppConfig  Application configuration interface
type IAppConfig interface {
	Save(ConfigFilePath string) error
	Load(ConfigFilePath string) (IAppConfig, error)
	SetDefaults()
	GetItem(key string) interface{}
	SetItem(key string, data interface{})
}
