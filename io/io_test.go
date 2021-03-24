package io

import "testing"

var dev Device

func TestSetName(t *testing.T){
	want := "count-1"
	dev.SetName("count-1")
	got := dev.name

	if got != want{
		t.Errorf("SetName \n got: %v \n want: %v \n", got, want)
	}
}

func TestGetName(t *testing.T){
	want := "count-1"
	dev.name = "count-1"
	got := dev.GetName()

	if got != want{
		t.Errorf("GetName \n got: %v \n want: %v \n", got, want)
	}
}

func TestSetAddress(t *testing.T){
	want := "0x55"
	dev.SetAddress("0x55")
	got := dev.address 

	if got != want{
		t.Errorf("SetAddress \n got: %v \n want: %v \n", got, want)
	}
}

func TestGetAddress(t *testing.T){
	want := "0x55"
	dev.address = "0x55"
	got := dev.GetAddress()

	if got != want{
		t.Errorf("GetAddress \n got: %v \n want: %v \n", got, want)
	}
}

func TestSetConfig(t *testing.T){
	want := "10111001010100001010100001011010"
	dev.SetConfig("10111001010100001010100001011010")
	got := dev.config

	if got != want{
		t.Errorf("SetConfig \n got: %v \n want: %v \n", got, want)
	}
}

func TestGetConfig(t *testing.T){
	want := "10111001010100001010100001011010"
	dev.config = "10111001010100001010100001011010"
	got := dev.GetConfig()

	if got != want{
		t.Errorf("GetConfig \n got: %v \n want: %v \n", got, want)
	}
}

func TestSetType(t *testing.T){
	want := "In"
	dev.SetType("In")
	got := dev.dType

	if got != want{
		t.Errorf("SetType \n got: %v \n want: %v \n", got, want)
	}
}

func TestGetType(t *testing.T){
	want := "Out"
	dev.dType = "Out"
	got := dev.GetType()

	if got != want{
		t.Errorf("GetType \n got: %v \n want: %v \n", got, want)
	}
}