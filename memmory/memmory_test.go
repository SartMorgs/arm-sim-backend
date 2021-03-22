package memmory

import "testing"

var mem Memmory 

func TestSetAddress(t *testing.T){
	want := "0x55"
	mem.SetAddress("0x55")
	got := mem.memmoryAddress

	if got != want{
		t.Errorf("SetAddress \n got: %v \n want %v \n", got, want)
	}
}

func TestGetAddress(t *testing.T){
	want := "0x55"
	mem.memmoryAddress = "0x55"
	got := mem.GetAddress()

	if got != want{
		t.Errorf("GetAddress \n got: %v \n want %v \n", got, want)
	}
}

func TestSetConfigType(t *testing.T){
	want := "L"
	mem.SetConfigType("L")
	got := mem.configType

	if got != want{
		t.Errorf("SetConfigType \n got: %v \n want %v \n", got, want)
	}
}

func TestSetValue(t *testing.T){
	want := 77
	mem.SetValue(77)
	got := mem.decValue

	if got != want{
		t.Errorf("SetValue \n got: %v \n want %v \n", got, want)
	}
}

func TestGetDecValue(t *testing.T){
	want := 77
	mem.decValue = 77
	got := mem.GetDecValue()

	if got != want{
		t.Errorf("GetValue \n got: %v \n want %v \n", got, want)
	}
}

func TestGetBinValue(t *testing.T){
	want := "111"
	mem.binValue = "111"
	got := mem.GetBinValue()

	if got != want{
		t.Errorf("GetBinValue \n got: %v \n want: %v \n", got, want)
	}
}

func TestGetHexValue(t *testing.T){
	want := "0x37"
	mem.hexValue = "0x37"
	got := mem.GetHexValue()

	if got != want{
		t.Errorf("GetHexValue \n got: %v \n want %v \n", got, want)
	}
}

func TestGetConfigType(t *testing.T){
	want := "L"
	mem.configType = "L"
	got := mem.GetConfigType()

	if got != want{
		t.Errorf("GetConfigType\n got: %v \n want %v \n", got, want)
	}
}

func TestToBinValue(t *testing.T){
	want := "111"
	mem.decValue = 7
	mem.toBinValue()
	got := mem.binValue

	if got != want{
		t.Errorf("toBinValue \n got: %v \n want: %v \n", got, want)
	}
}

func TestToHexValue(t *testing.T){
	want := "0x37"
	mem.decValue = 55
	mem.toHexValue()
	got := mem.hexValue

	if got != want{
		t.Errorf("toHexValue \n got: %v \n want: %v \n", got, want)
	}
}

func TestAppendBinValue(t *testing.T){
	want := "00000000000001010010001101000001"
	mem.binValue = "1010010001101000001"
	mem.appendBinValue()
	got := mem.fullBinValue

	if got != want{
		t.Errorf("appendBinValue \n got: %v \n want: %v \n", got, want)
	}
}

func TestGetSplitBin(t *testing.T){
	want := [4]string{"00000000", "00000101", "00100011", "01000001"}
	mem.splitedBinValue = [4]string{"00000000", "00000101", "00100011", "01000001"}
	got := mem.GetSplitBin()

	if got != want{
		t.Errorf("GetSplitBin \n got: %v \n want: %v \n", got, want)
	}
}

func TestSplitBin(t *testing.T){
	want := [4]string{"00000000", "00000101", "00100011", "01000001"}
	mem.fullBinValue = "00000000000001010010001101000001"
	mem.configType = "L"
	mem.splitBin()
	got := mem.splitedBinValue

	if got != want{
		t.Errorf("GetSplitBin \n got: %v \n want: %v \n", got, want)
	}
}

func TestSetAliasField(t *testing.T){
	want := "Area-A"
	mem.SetAliasField("Area-A")
	got := mem.aliasField

	if got != want{
		t.Errorf("SetAliasField \n got: %v \n want: %v \n", got, want)
	}
}

func TestGetAliasField(t *testing.T){
	want := "Area-A"
	mem.aliasField = "Area-A"
	got := mem.GetAliasField()

	if got != want{
		t.Errorf("GetAliasField \n got: %v \n want: %v \n", got, want)
	}
}

func TestSetFullBinValue(t *testing.T){
	want := "10001101010100001010100001011010"
	mem.SetFullBinValue("10001101010100001010100001011010")
	got := mem.fullBinValue

	if got != want{
		t.Errorf("SetFullBinValue \n got: %v \n want: %v \n", got, want)
	}
}

func TestGetFullBinValue(t *testing.T){
	want := "10001101010100001010100001011010"
	mem.fullBinValue = "10001101010100001010100001011010"
	got := mem.GetFullBinValue()

	if got != want{
		t.Errorf("GetFullBinValue \n got: %v \n want: %v \n", got, want)
	}
}