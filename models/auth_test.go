package models

import (
	"github.com/BaiMeow/HduHelpLogin/conf"
	"testing"
)

func TestAddAuth(t *testing.T) {
	conf.Init()
	Init()
	id, err := AddAuth("1234", "4567")
	if err != nil {
		t.Error(err)
		return
	}
	if id == 0 {
		t.Error("existed username")
	}
	t.Log(id)
}

func TestCheckAuth(t *testing.T) {
	conf.Init()
	Init()
	id, err := CheckAuth("123", "456")
	if err != nil {
		t.Error(err)
		return
	}
	if id != 0 {
		t.Log("pass")
	} else {
		t.Log("fail")
	}
}
