package appconf

import (
	"testing"

	appconfint "github.com/eshu0/appconfig/pkg/interfaces"
)

// TestSave saves the config
func TestSave(t *testing.T) {
	conf := NewAppConfig()
	conf.SetData("Monkey")
	err := conf.Save(DefaultFilePath)
	if err != nil {
		t.Fatalf(`Save(%s) = %v should not error`, DefaultFilePath, err)
	}
}

func TestLoad(t *testing.T) {

	// old config
	conf := NewAppConfig()
	conf.SetData("Monkey1")

	err := conf.Save(DefaultFilePath)
	if err != nil {
		t.Fatalf(`Save(%s) = %v should not error`, DefaultFilePath, err)
	}

	newconf, err1 := conf.Load(DefaultFilePath)
	if err1 != nil {
		t.Fatalf(`Load(%s) = %v should not error`, DefaultFilePath, err1)
	}

	monkey1 := newconf.GetData()
	Monkeystring, ok := monkey1.(string)

	if !ok && Monkeystring != "Monkey1" {
		t.Fatalf(`"Monkeystring != Monkey1 instead was %s`, Monkeystring)
	}

	newconf.SetData("Monkey2")
	updatederr := newconf.Save(DefaultFilePath)
	if updatederr != nil {
		t.Fatalf(`Save(%s) = %v should not error`, DefaultFilePath, updatederr)
	}
}

func TestUpdate(t *testing.T) {

	// set config
	conf := NewAppConfig()
	conf.SetData("Monkey1")

	// save it
	err := conf.Save(DefaultFilePath)
	if err != nil {
		t.Fatalf(`Save(%s) = %v should not error`, DefaultFilePath, err)
	}

	// load it
	newconf, err1 := conf.Load(DefaultFilePath)
	if err1 != nil {
		t.Fatalf(`Load(%s) = %v should not error`, DefaultFilePath, err1)
	}

	// check new one was loaded
	monkey1 := newconf.GetData()
	Monkeystring, ok := monkey1.(string)

	if !ok && Monkeystring != "Monkey1" {
		t.Fatalf(`"Monkeystring != Monkey1 instead was %s`, Monkeystring)
	}

	// udate it
	newconf.SetData("Monkey2")
	updatederr := newconf.Save(DefaultFilePath)
	if updatederr != nil {
		t.Fatalf(`Save(%s) = %v should not error`, DefaultFilePath, updatederr)
	}

	// load it again
	newconf, err1 = conf.Load(DefaultFilePath)
	if err1 != nil {
		t.Fatalf(`Load(%s) = %v should not error`, DefaultFilePath, err1)
	}

	monkey2 := newconf.GetData()
	Monkeystring, ok = monkey2.(string)

	if !ok && Monkeystring != "Monkey2" {
		t.Fatalf(`"Monkeystring != Monkey2 instead was %s`, Monkeystring)
	}
}

type DummyConfig struct {
	Parent *AppConfig
}

type DummyData struct {
	DumbProperty   string   `json:"dumdum,omitempty"`
	DumbProperties []string `json:"dumanddumer,omitempty"`
}

func NewDummyConfig() *DummyConfig {
	conf := NewAppConfig()
	dc := &DummyConfig{}
	Config, ok := conf.(*AppConfig)
	if ok {
		dc.Parent = Config
		dc.Parent.SetDefaultFunc(SetDummyDefaults)
		dc.Parent.SetDefaults()
		return dc
	}

	return nil

}

//SetDummyDefaults sets the defaults this is a noop
func SetDummyDefaults(Config appconfint.IAppConfig) {
	dd := DummyData{}
	dd.DumbProperty = "Monkey3"
	dd.DumbProperties = []string{"dummy", "summy", "lummy"}
	Config.SetData(dd)
}

func TestSaveDummy(t *testing.T) {
	conf := NewDummyConfig()
	if conf != nil {
		err := conf.Parent.Save(DefaultFilePath)
		if err != nil {
			t.Fatalf(`Save(%s) = %v should not error`, DefaultFilePath, err)
		}
	} else {
		t.Fatal(`NewDummyConfig is nil `)
	}
}

func TestLoadDummy(t *testing.T) {
	conf := NewDummyConfig()
	if conf != nil {
		err := conf.Parent.Save(DefaultFilePath)
		if err != nil {
			t.Fatalf(`Save(%s) = %v should not error`, DefaultFilePath, err)
		}

		newconfig, err := conf.Parent.Load(DefaultFilePath)
		if err != nil || newconfig == nil {
			t.Fatalf(`Load(%s) = %v should not error`, DefaultFilePath, err)
			return
		}
		ccat, ok := newconfig.(*AppConfig)

		if ok {
			ccat.SetData("Monkey2")
			err := conf.Parent.Save("./afterwards.json")
			if err != nil {
				t.Fatalf(`Save(%s) = %v should not error`, DefaultFilePath, err)
			}
			return
		}

		t.Fatal("LoadConfig Cast failed")

	} else {
		t.Fatal(`NewDummyConfig is nil `)
	}
}
