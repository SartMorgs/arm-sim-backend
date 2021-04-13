package controller

import (
	"testing"
	"reflect"
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

func TestGetInstructionFormat(t *testing.T){
	want := map[string]string{
		"rn": "10111",
		"address": "11110100001011",
	}

	cn.decodeUnit.instructionCode = "11010110111111101000010111011010"
	cn.decodeUnit.SplitInstruction()
	got := cn.decodeUnit.GetInstructionFormat()

	if !reflect.DeepEqual(want, got){
		t.Errorf("GetInstructionFormat \n got: %v \n want %v \n", got, want)
	}
}