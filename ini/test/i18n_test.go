package test

import (
	"testing"

	"github.com/gwaylib/conf/ini"
)

func TestI18n(t *testing.T) {
	i18nDir := "./"
	cfg := ini.NewIniCache(i18nDir)
	// if the language file is not exist, load the default file
	val := cfg.GetDefaultFile("", "app.lang.en").String("error", "0")
	if val != "zero" {
		t.Fatal(val)
		return
	}

	val = cfg.GetDefaultFile("app.lang.zh_cn", "app.lang.en").String("error", "0")
	if val != "零" {
		t.Fatal(val)
		return
	}
}
