package appconf

import (
	"testing"
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
