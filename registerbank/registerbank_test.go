package registerbank

import (
	"testing"
	"strconv"
)

var rgbk RegisterBank 

func TestChangeRegister(t *testing.T){
	rgbk := NewRegisterBank()

	want := 77
	rgbk.ChangeRegister("R1", 77)
	got := rgbk.registerList["R1"].decValue

	if got != want{
		t.Errorf("ChangeRegister \n got: %v \n want %v \n", got, want)
	}
}

func TestResetRegisterBank(t *testing.T){
	rgbk := NewRegisterBank()

	want := [17]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	rgbk.ResetRegisterBank()
	got := rgbk.registerList

	var registerCode string
	for count := 0; count < 17; count++{
		registerCode = "R" + strconv.Itoa(count)
		if got[registerCode].decValue != want[count]{
			t.Errorf("ResetRegisterBank \n got: %v \n want %v \n", got, want)
			break
		}
	}
}