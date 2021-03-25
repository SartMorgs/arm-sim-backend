package controller

import (
	"testing"
	"reflect"
)

var exun ExecuteUnit

func TestConfigForLdi(t *testing.T){
	want := map[string]int{
		"register": 2,
		"value": 70,
	}
	exun.ConfigForLdi("00010", "00000001000110")
	got := exun.config

	if reflect.DeepEqual(want, got){
		t.Errorf("ConfigForLdi \n got: %v \n want %v \n", got, want)
	}
}