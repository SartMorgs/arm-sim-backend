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

func TestConfigForAnds(t *testing.T){
	exun := NewExecuteUnit()

	want := map[string]int64{
		"register_rd": 2,
		"register_rn": 2,
		"register_rm": 4,
	}
	exun.ConfigForAnds("00010", "00010", "00100")
	got := exun.config

	if !reflect.DeepEqual(want, got){
		t.Errorf("ConfigForAnds \n got: %v \n want %v \n", got, want)
	}
}

func TestConfigForLdm(t* testing.T){
	exun := NewExecuteUnit()

	want := map[string]int64{
		"register_rn": 2,
		"register_1": 4,
		"register_2": 5,
		"register_3": 6,
		"register_4": 7,
	}

	var register_list = []string{
		"00100",
		"00101",
		"00110",
		"00111"}

	exun.ConfigForLdm("00010", register_list)
	got := exun.config

	if !reflect.DeepEqual(want, got){
		t.Errorf("ConfigForLdm \n got: %v \n want %v \n", got, want)
	}
}