package handlers

import "testing"

func TestConfigInit(t *testing.T) {
	home := ".relayer"
	err := ConfigInit(home)
	if err != nil {

		t.Fatal("init config error")
	}
}
