package test

import (
	"testing"

	"github.com/gwaylib/conf/ini"
)

func TestI18n(t *testing.T) {
	i18nDir := "./"
	cfg := ini.NewIniCache(i18nDir)
	// if the language file is not exist, read the default file
	msg_default := cfg.GetDefaultFile("", "app.default.en").String("error", "0")
	if msg_default != "zero" {
		t.Fatal(msg_default)
		return
	}
}
