package registerbank

import "testing"

var reg Register

func TestSetValue(t *testing.T){
	want := 7
	reg.SetValue(7)
	got := reg.GetValue()

	if got != want{
		t.Errorf("SetValue \n got: %v \n want: \n%v", got, want)
	}
}

func TestbuildBinValue(t *testing.T){
	want := "0111"
	reg.SetValue(7)
	reg.buildBinValue()
	got := reg.GetBinValue()

	if got != want{
		t.Errorf("builBinValue \n got: %v \n want: \n%v", got, want)
	}
}

func TestSetName(t *testing.T){
	want := "NAME1"
	reg.SetName("NAME1")
	got := reg.GetName()

	if got != want{
		t.Errorf("SetName \n got: %v \n want: \n%v", got, want)
	}
}