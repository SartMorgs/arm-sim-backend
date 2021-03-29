package controller

import (
	"testing"
	"reflect"
)

var exun ExecuteUnit

func TestConfigForLdi(t *testing.T){
	exun := NewExecuteUnit()

	want := map[string]int64{
		"register": 2,
		"value": 70,
	}
	exun.ConfigForLdi("00010", "00000001000110")
	got := exun.config

	if !reflect.DeepEqual(want, got){
		t.Errorf("ConfigForLdi \n got: %v \n want %v \n", got, want)
	}
}

func TestConfigForLdr(t *testing.T){
	exun := NewExecuteUnit()

	want := map[string]int64{
		"register_rt": 2,
		"register_rn": 4, 
	}
	exun.ConfigForLdr("00010", "00100")
	got := exun.config

	if !reflect.DeepEqual(want, got){
		t.Errorf("ConfigForLdr \n got: %v \n want %v \n", got, want)
	}
}

func TestConfigForStr(t *testing.T){
	exun := NewExecuteUnit()

	want := map[string]int64{
		"register_rt": 2,
		"register_rn": 4,
		"ram_address": 100, 
	}
	exun.ConfigForStr("00010", "00100", "1100100")
	got := exun.config

	if !reflect.DeepEqual(want, got){
		t.Errorf("ConfigForStr \n got: %v \n want %v \n", got, want)
	}
}

func TestConfigForAdds(t *testing.T){
	exun := NewExecuteUnit()

	want := map[string]int64{
		"register_rd": 2,
		"register_rn": 2,
		"register_rm": 4,
		"immediate": 0, 
	}
	exun.ConfigForAdds("00010", "00010", "00100", "0")
	got := exun.config

	if !reflect.DeepEqual(want, got){
		t.Errorf("ConfigForAdds \n got: %v \n want %v \n", got, want)
	}
}

func TestConfigForSubs(t *testing.T){
	exun := NewExecuteUnit()

	want := map[string]int64{
		"register_rd": 2,
		"register_rn": 2,
		"register_rm": 4,
		"immediate": 0, 
	}
	exun.ConfigForSubs("00010", "00010", "00100", "0")
	got := exun.config

	if !reflect.DeepEqual(want, got){
		t.Errorf("ConfigForSubs \n got: %v \n want %v \n", got, want)
	}
}