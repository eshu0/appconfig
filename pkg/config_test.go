package appconf

import (
	"testing"
)

// TestSave saves the config
func TestSave(t *testing.T) {
	conf := NewAppConfig()

	err := conf.Save(DefaultFilePath)
	if err != nil {
		t.Fatalf(`Save(%s) = %v should not error`, DefaultFilePath, err)
	}
}
