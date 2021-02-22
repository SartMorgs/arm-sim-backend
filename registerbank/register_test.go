package registerbank

import "testing"

var reg Register

func TestSetValue(t *testing.T){
	want := 7
	reg.SetValue(7)
	got := reg.GetDecValue()

	if got != want{
		t.Errorf("SetValue \n got: %v \n want: %v \n", got, want)
	}
}

func TestToBinValue(t *testing.T){
	want := "111"
	reg.SetValue(7)
	reg.toBinValue()
	got := reg.GetBinValue()

	if got != want{
		t.Errorf("toBinValue \n got: %v \n want: %v \n", got, want)
	}
}

func TestSetName(t *testing.T){
	want := "NAME1"
	reg.SetRegisterName("NAME1")
	got := reg.GetRegisterName()

	if got != want{
		t.Errorf("SetName \n got: %v \n want: %v \n", got, want)
	}
}

func TestToHexValue(t *testing.T){
	want := "0x37"
	reg.SetValue(55)
	reg.toHexValue()
	got := reg.GetHexValue()

	if got != want{
		t.Errorf("toHexValue \n got: %v \n want: %v \n", got, want)
	}
}

func TestSetRegisterFunction(t *testing.T){
	want := "Stack Pointer"
	reg.SetRegisterFunction("Stack Pointer")
	got := reg.GetRegisterFunction()

	if got != want{
		t.Errorf("GetRegisterFunction \n got: %v \n want: %v \n", got, want)
	}
}