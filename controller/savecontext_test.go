package controller

import "testing"

var sv SaveStatus

func TestSetAddressRom(t *testing.T){
	want := "0x55"
	sv.SetAddressRom("0x55")
	got := sv.lastAddressRom

	if got != want{
		t.Errorf("SetAddressRom \n got: %v \n want: %v \n", got, want)
	}
}