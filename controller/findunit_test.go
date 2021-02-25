package controller

import "testing"

var fdun FindUnit 

// Unit Tests
// Test all methods using mock
func TestSetInstruction(t *testing.T){
	want := "000001"
	fdun.SetInstruction("000001")
	got := fdun.instruction

	if want != got{
		t.Errorf("SetInstruction \n got: %v want %v \n", got, want)
	}
}

func TestSetProgramArea(t *testing.T){
	want := "func-a"
	fdun.SetProgramArea("func-a")
	got := fdun.programArea

	if want != got{
		t.Errorf("SetProgramArea \n got: %v \n want %v \n", got, want)
	}
}

func TestGetInstruction(t *testing.T){
	want := "000001"
	fdun.instruction = "000001"
	got := fdun.GetInstruction()

	if got != want{
		t.Errorf("GetInstruction \n got: %v \n want %v \n", got, want)
	}
}

func TestGetProgramArea(t *testing.T){
	want := "func-a"
	fdun.programArea = "func-a"
	got := fdun.GetProgramArea()

	if got != want{
		t.Errorf("GetProgramArea \n got %v \n want %v \n", got, want)
	}
}
