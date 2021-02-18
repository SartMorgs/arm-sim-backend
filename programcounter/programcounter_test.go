package programcounter

import "testing"

var pc ProgramCounter 

func TestSetAddressMemmory(t *testing.T){
	want := 0x37

	pc.SetAddressMemmory(55)
	pc.decAddressToHexAddress()

	got := pc.GetHexAddressMemmory()

	if want != got{
		t.Errorf("SetAddressMesmmory \n got: %v \n want %v \n", got, want)
	}
}