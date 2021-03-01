package controller

import "testing"

var ftun FetchUnit 

// Unit Tests
// Test all methods using mock
func TestSetInstruction(t *testing.T){
	want := "000001"
	ftun.SetInstruction("000001")
	got := ftun.instruction

	if want != got{
		t.Errorf("SetInstruction \n got: %v want %v \n", got, want)
	}
}

func TestSetProgramArea(t *testing.T){
	want := "func-a"
	ftun.SetProgramArea("func-a")
	got := ftun.programArea

	if want != got{
		t.Errorf("SetProgramArea \n got: %v \n want %v \n", got, want)
	}
}

func TestGetInstruction(t *testing.T){
	want := "000001"
	ftun.instruction = "000001"
	got := ftun.GetInstruction()

	if got != want{
		t.Errorf("GetInstruction \n got: %v \n want %v \n", got, want)
	}
}

func TestGetProgramArea(t *testing.T){
	want := "func-a"
	ftun.programArea = "func-a"
	got := ftun.GetProgramArea()

	if got != want{
		t.Errorf("GetProgramArea \n got %v \n want %v \n", got, want)
	}
}
