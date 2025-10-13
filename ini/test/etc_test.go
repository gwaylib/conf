package test

import (
	"testing"

	"github.com/gwaylib/conf/ini"
)

func TestEtc(t *testing.T) {
	cfg := ini.NewIni("./").GetFile("etc.cfg")
	sec := cfg.Section("test")
	if sec.Key("str").String() != "abc" {
		t.Fatal(sec.Key("str"))
	}
	if sec.Key("int").MustInt() != 1 {
		t.Fatal(sec.Key("int"))
	}
	if sec.Key("bool_true").MustBool() != true {
		t.Fatal(sec.Key("bool_true"))
	}
	if sec.Key("bool_false").MustBool() != false {
		t.Fatal(sec.Key("bool_false"))
	}
	if sec.Key("float").MustFloat64() != 3.20 {
		t.Fatal(sec.Key("float"))
	}
}
