package programcounter

import "testing"

var pc ProgramCounter 

func TestSetAddressMemmory(t *testing.T){
	want := 55
	pc.SetAddressMemmory(55)
	got := pc.GetDecAddressMemmory()

	if want != got{
		t.Errorf("SetAddressMemmory \n got: %v \n want %v \n", got, want)
	}
}

func TestToHexAddress(t *testing.T){
	want := "0x37"
	pc.SetAddressMemmory(55)
	pc.toHexAddress()
	got := pc.GetHexAddressMemmory()

	if want != got{
		t.Errorf("toHexAddress \n got: %v \n want %v \n", got, want)
	}
}

func TestToBinValue(t *testing.T){
	want := "111"
	pc.SetAddressMemmory(7)
	pc.toBinAddress()
	got := pc.GetBinAddressMemmory()

	if want != got{
		t.Errorf("toBinValue \n got: %v \n want %v \n", got, want)
	}
}