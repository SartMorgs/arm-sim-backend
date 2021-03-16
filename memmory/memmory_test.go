package memmory

import "testing"

var mem Memmory 

func TestSetAddress(t *testing.T){
	want := "0x55"
	mem.SetAddress("0x55")
	got := mem.memmoryAddress

	if got != want{
		t.Errorf("SetAddress \n got: %v \n want %v \n", got, want)
	}
}

func TestGetAddress(t *testing.T){
	want := "0x55"
	mem.memmoryAddress = "0x55"
	got := mem.GetAddress()

	if got != want{
		t.Errorf("GetAddress \n got: %v \n want %v \n", got, want)
	}
}

func TestSetValue(t *testing.T){
	want := 77
	mem.SetValue(77)
	got := mem.decValue

	if got != want{
		t.Errorf("SetValue \n got: %v \n want %v \n", got, want)
	}
}

func TestGetDecValue(t *testing.T){
	want := 77
	mem.decValue = 77
	got := mem.GetDecValue()

	if got != want{
		t.Errorf("GetValue \n got: %v \n want %v \n", got, want)
	}
}

func TestGetBinValue(t *testing.T){
	want := "111"
	mem.binValue = "111"
	got := mem.GetBinValue()

	if got != want{
		t.Errorf("GetBinValue \n got: %v \n want: %v \n", got, want)
	}
}

func TestGetHexValue(t *testing.T){
	want := "0x37"
	mem.hexValue = "0x37"
	got := mem.GetHexValue()

	if got != want{
		t.Errorf("GetHexValue \n got: %v \n want %v \n", got, want)
	}
}

func TestToBinValue(t *testing.T){
	want := "111"
	mem.decValue = 7
	mem.toBinValue()
	got := mem.binValue

	if got != want{
		t.Errorf("toBinValue \n got: %v \n want: %v \n", got, want)
	}
}

func TestToHexValue(t *testing.T){
	want := "0x37"
	mem.decValue = 55
	mem.toHexValue()
	got := mem.hexValue

	if got != want{
		t.Errorf("toHexValue \n got: %v \n want: %v \n", got, want)
	}
}