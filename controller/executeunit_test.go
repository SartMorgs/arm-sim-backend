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

func TestSetTargetRegister(t *testing.T){
	want := "00010"
	exun.SetTargetRegister("00010")
	got := exun.targetRegisterBin

	if want != got{
		t.Errorf("SetTargetRegister \n got: %v want %v \n", got, want)
	}
}

func TestSetAddressInstruction(t *testing.T){
	want := "00000001000110"
	exun.SetAddressInstruction("00000001000110")
	got := exun.addressInstructionBin

	if want != got{
		t.Errorf("SetAddressInstruction \n got: %v want %v \n", got, want)
	}
}

func TestSetSourceRegister1(t *testing.T){
	want := "00010"
	exun.SetSourceRegister1("00010")
	got := exun.sourceRegister1Bin

	if want != got{
		t.Errorf("SetSourceRegister1 \n got: %v want %v \n", got, want)
	}
}

func TestSetSourceRegister2(t *testing.T){
	want := "00010"
	exun.SetSourceRegister2("00010")
	got := exun.sourceRegister2Bin

	if want != got{
		t.Errorf("SetSourceRegister2 \n got: %v want %v \n", got, want)
	}
}

func TestSetValueInstruction(t *testing.T){
	want := "0000000001000110"
	exun.SetValueInstruction("0000000001000110")
	got := exun.valueInstructionBin

	if want != got{
		t.Errorf("SetValueInstruction \n got: %v want %v \n", got, want)
	}
}

func TestSetValueRegister(t *testing.T){
	want := "00010"
	exun.SetValueRegister("00010")
	got := exun.valueRegisterBin

	if want != got{
		t.Errorf("SetValueRegister \n got: %v want %v \n", got, want)
	}
}