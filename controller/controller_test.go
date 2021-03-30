package controller

import "testing"

var cn Controller

func TestChangeInstructionFetch(t *testing.T){
	cn := NewController()

	want := "10111001010100001010100001011010"
	cn.ChangeInstructionFetch("10111001010100001010100001011010")
	got := cn.fetchUnit.instruction 

	if got != want{
		t.Errorf("ChangeInstructionFetch \n got: %v \n want %v \n", got, want)
	}
}

func TestChangeInstructionDecode(t *testing.T){
	cn := NewController()

	want := "10111001010100001010100001011010"
	cn.ChangeInstructionDecode("10111001010100001010100001011010")
	got := cn.decodeUnit.instructionCode 

	if got != want{
		t.Errorf("ChangeInstructionDecode \n got: %v \n want %v \n", got, want)
	}
}