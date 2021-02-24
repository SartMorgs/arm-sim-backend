package programcounter

import "testing"

var pc ProgramCounter 

// Unit Tests
// Test all methods using mock
func TestSetAddressMemmory(t *testing.T){
	want := 55
	pc.SetAddressMemmory(55)
	got := pc.decAddressMemmory

	if want != got{
		t.Errorf("SetAddressMemmory \n got: %v \n want %v \n", got, want)
	}
}

func TestGetDecAddressMemmory(t *testing.T){
	want := 55
	pc.decAddressMemmory = 55
	got := pc.GetDecAddressMemmory()

	if got != want {
		t.Errorf("GetDecAddressMemmory \n got: %v \n want %v \n", got, want)
	}
}

func TestGetBinAddressMemmory(t *testing.T){
	want := "111"
	pc.binAddressMemmory = "111"
	got := pc.GetBinAddressMemmory()

	if got != want{
		t.Errorf("GetBinAddressMemmory \n got: %v \n want %v \n", got, want)
	}
}

func TestGetHexAddressMemmory(t *testing.T){
	want := "0x37"
	pc.hexAddressMemmory = "0x37"
	got := pc.GetHexAddressMemmory()

	if got != want{
		t.Errorf("GetHexAddressMemmory \n got: %v \n want %v \n", got, want)
	}
}

func TestToHexAddress(t *testing.T){
	want := "0x37"
	pc.decAddressMemmory = 55
	pc.toHexAddress()
	got := pc.hexAddressMemmory

	if want != got{
		t.Errorf("toHexAddress \n got: %v \n want %v \n", got, want)
	}
}

func TestToBinValue(t *testing.T){
	want := "111"
	pc.decAddressMemmory = 7
	pc.toBinAddress()
	got := pc.binAddressMemmory

	if want != got{
		t.Errorf("toBinValue \n got: %v \n want %v \n", got, want)
	}
}

func TestGePointerPC(t *testing.T)(){
	want := 75
	pc.currentPointer = 0

	for i := 0; i < 75; i++{
		pc.nextPointer()
	}

	got := pc.currentPointer

	if want != got{
		t.Errorf("GetPointerPC \n got: %v \n want %v \n", got, want)
	}
}
