package controller

import (
	"testing"
)

var exun ExecuteUnit

func TestSetExecuteInstruction(t *testing.T){
	want := "10111001010100001010100001011010"
	exun.SetInstruction("10111001010100001010100001011010")
	got := exun.instruction 

	if want != got{
		t.Errorf("SetInstruction \n got: %v want %v \n", got, want)
	}
}