package controller

import "testing"

var fdun FindUnit 

func TestSetInstruction(t *testing.T){
	want := "000001"
	fdun.SetInstruction("000001")
	got := fdun.GetInstruction()

	if want != got{
		t.Errorf("SetInstruction \n got: %v want %v \n", got, want)
	}
}

func TestSetProgramArea(t *testing.T){
	want := "func-a"
	fdun.SetProgramArea("func-a")
	got := fdun.GetProgramArea()

	if want != got{
		t.Errorf("SetProgramArea \n got: %v \n want %v \n", got, want)
	}
}

func TestSetMemmoryAddress(t *testing.T){
	want := 77
	fdun.SetMemmoryAddress(77)
	got := fdun.GetMemmoryAddress()

	if want != got{
		t.Errorf("SetMemmoryAddress \n got: %v \n want %v \n", got, want)
	}
}