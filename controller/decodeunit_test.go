package controller

import (
	"testing"
)

var dcun DecodeUnit

// Unit Tests
// Test all methods using mock
func TestGetOpcode(t *testing.T){
	want := "111111"
	dcun.opcode = "111111"
	got := dcun.GetOpcode()

	if want != got{
		t.Errorf("GetOpcode \n got: %v \n want %v \n", got, want)
	}
}

func TestGetNameInstruction(t *testing.T){
	want := "ORRS"
	dcun.instructionName = "ORRS"
	got := dcun.GetInstructionName()

	if want != got{
		t.Errorf("GetInstructionName \n got: %v \n want %v \n", got, want)
	}
}

func TestGapAddressFlag(t *testing.T){
	dcun := NewDecodeUnit()

	want := true
	dcun.instructionCode = "01000010101001001110100001011010"
	dcun.SplitInstruction()
	got := dcun.gapAddress

	if want != got{
		t.Errorf("GapAddressFlag \n got: %v \n want %v \n", got, want)
	}
}

func TestGetGapAddressFlag(t *testing.T){
	want := false
	dcun.gapAddress = false
	got := dcun.GetGapAddressFlag()

	if want != got{
		t.Errorf("GetGapAddressFlag \n got: %v \n want %v \n", got, want)
	}
}