package appconf

import (
	"testing"

	appconfint "github.com/eshu0/appconfig/pkg/interfaces"
)

// TestSave saves the config
func TestSave(t *testing.T) {
	conf := NewAppConfig()
	conf.SetItem("Banana", "Monkey")
	err := conf.Save(DefaultFilePath)
	if err != nil {
		t.Fatalf(`Save(%s) = %v should not error`, DefaultFilePath, err)
	}
}

func TestLoad(t *testing.T) {

	// old config
	conf := NewAppConfig()
	conf.SetItem("Banana1", "Monkey1")

	err := conf.Save(DefaultFilePath)
	if err != nil {
		t.Fatalf(`Save(%s) = %v should not error`, DefaultFilePath, err)
	}

	newconf, err1 := conf.Load(DefaultFilePath)
	if err1 != nil {
		t.Fatalf(`Load(%s) = %v should not error`, DefaultFilePath, err1)
	}

	monkey1 := newconf.GetItem("Banana1")
	Monkeystring, ok := monkey1.(string)

	if !ok && Monkeystring != "Monkey1" {
		t.Fatalf(`"Monkeystring != Monkey1 instead was %s`, Monkeystring)
	}

	newconf.SetItem("Banana1", "Monkey2")
	updatederr := newconf.Save(DefaultFilePath)
	if updatederr != nil {
		t.Fatalf(`Save(%s) = %v should not error`, DefaultFilePath, updatederr)
	}
}

func TestUpdate(t *testing.T) {

	// set config
	conf := NewAppConfig()
	conf.SetItem("Banana1", "Monkey1")

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
	monkey1 := newconf.GetItem("Banana1")
	Monkeystring, ok := monkey1.(string)

	if !ok && Monkeystring != "Monkey1" {
		t.Fatalf(`"Monkeystring != Monkey1 instead was %s`, Monkeystring)
	}

	// udate it
	newconf.SetItem("Banana1", "Monkey2")
	updatederr := newconf.Save(DefaultFilePath)
	if updatederr != nil {
		t.Fatalf(`Save(%s) = %v should not error`, DefaultFilePath, updatederr)
	}

	// load it again
	newconf, err1 = conf.Load(DefaultFilePath)
	if err1 != nil {
		t.Fatalf(`Load(%s) = %v should not error`, DefaultFilePath, err1)
	}

	monkey2 := newconf.GetItem("Banana1")
	Monkeystring, ok = monkey2.(string)

	if !ok && Monkeystring != "Monkey2" {
		t.Fatalf(`"Monkeystring != Monkey2 instead was %s`, Monkeystring)
	}
}

type DummyConfig struct {
	Parent *AppConfig
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
	Config.SetItem("DummyItem", "Monkey3")
	Config.SetItem("Anotherdummy", []string{"dummy", "summy", "lummy"})
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
			ccat.SetItem("Banana1", "Monkey2")

			return
		}

		t.Fatal("LoadConfig Cast failed")

	} else {
		t.Fatal(`NewDummyConfig is nil `)
	}
}
