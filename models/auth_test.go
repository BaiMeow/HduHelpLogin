package models

import (
	"testing"
)

func TestAddAuth(t *testing.T) {
	Init("data.db")
	err := AddAuth("1234", "4567")
	if err != nil {
		t.Error(err)
	}
}

func TestCheckAuth(t *testing.T) {
	Init("data.db")
	ok, err := CheckAuth("123", "456")
	if ok {
		t.Log("pass")
	}
	if err != nil {
		t.Error(err)
	}
}
