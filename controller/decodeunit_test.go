package controller

import "testing"

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