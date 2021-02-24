package registerbank

import "testing"

var reg Register

// Unit Tests
// Test all methods using mock
func TestSetValue(t *testing.T){
	want := 7
	reg.SetValue(7)
	got := reg.decValue

	if got != want{
		t.Errorf("SetValue \n got: %v \n want: %v \n", got, want)
	}
}

func TestSetName(t *testing.T){
	want := "NAME1"
	reg.SetRegisterName("NAME1")
	got := reg.registerName

	if got != want{
		t.Errorf("SetName \n got: %v \n want: %v \n", got, want)
	}
}

func TestSetRegisterFunction(t *testing.T){
	want := "Stack Pointer"
	reg.SetRegisterFunction("Stack Pointer")
	got := reg.registerFunction

	if got != want{
		t.Errorf("GetRegisterFunction \n got: %v \n want: %v \n", got, want)
	}
}

func TestGetValue(t *testing.T){
	want := 7
	reg.decValue = 7
	got := reg.GetDecValue()

	if got != want{
		t.Errorf("GetValue \n got: %v \n want: %v \n", got, want)
	}
}

func TestGetRegisterName(t *testing.T){
	want := "NAME1"
	reg.registerName = "NAME1"
	got := reg.GetRegisterName()

	if got != want{
		t.Errorf("GetRegisterName \n got %v \n want: %v \n", got, want)
	}
}

func TestGetRegisterFunction(t *testing.T){
	want := "Stack Pointer"
	reg.registerFunction = "Stack Pointer"
	got := reg.GetRegisterFunction()

	if got != want{
		t.Errorf("GetRegisterFunction \n got: %v \n want: %v \n", got, want)
	}
}

func TestGetBinValue(t *testing.T){
	want := "111"
	reg.binValue = "111"
	got := reg.GetBinValue()

	if got != want{
		t.Errorf("GetBinValue \n got: %v \n want: %v \n", got, want)
	}
}

func TestGetHexValue(t *testing.T){
	want := "0x37"
	reg.hexValue = "0x37"
	got := reg.GetHexValue()

	if got != want{
		t.Errorf("GetHexValue \n got: %v \n want %v \n", got, want)
	}
}

func TestToBinValue(t *testing.T){
	want := "111"
	reg.decValue = 7
	reg.toBinValue()
	got := reg.binValue

	if got != want{
		t.Errorf("toBinValue \n got: %v \n want: %v \n", got, want)
	}
}

func TestToHexValue(t *testing.T){
	want := "0x37"
	reg.decValue = 55
	reg.toHexValue()
	got := reg.hexValue

	if got != want{
		t.Errorf("toHexValue \n got: %v \n want: %v \n", got, want)
	}
}