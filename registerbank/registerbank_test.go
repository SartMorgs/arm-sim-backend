package registerbank

import "testing"

var rgbk RegisterBank 

func TestChangeRegister(t *testing.T){
	want := 77
	rgbk.ChangeRegister("R01", 77)
	got := rgbk.GetDecRegister()

	if got != want{
		t.Errorf("ChangeRegister \n got: %v \n want %v \n", got, want)
	}
}