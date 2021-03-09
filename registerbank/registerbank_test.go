package registerbank

import (
	"testing"
)

var rgbk RegisterBank 

func TestChangeRegister(t *testing.T){
	rgbk := NewRegisterBank()

	want := 77
	rgbk.ChangeRegister("R1", 77)
	got := rgbk.registerList["R1"].GetDecValue()

	if got != want{
		t.Errorf("ChangeRegister \n got: %v \n want %v \n", got, want)
	}
}