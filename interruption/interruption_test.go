package interruption

import "testing"

var it Interruption

func TestTurnOnTrigger(t *testing.T){
	want := true
	it.TurnOnTrigger()
	got := it.trigger

	if got != want{
		t.Errorf("TurnOnTrigger \n got: %v \n want: %v \n", got, want)
	}
}

func TestTurnOffTrigger(t *testing.T){
	want := false
	it.TurnOffTrigger()
	got := it.trigger

	if got != want{
		t.Errorf("TurnOffTrigger \n got: %v \n want: %v \n", got, want)
	}
}

func TestGetTrigger(t *testing.T){
	want := false
	it.trigger = false
	got := it.GetTrigger()

	if got != want{
		t.Errorf("GetTriggerGetTrigger \n got: %v \n want: %v \n", got, want)
	}
}

func TestSetNumber(t *testing.T){
	want := 1
	it.SetNumber(1)
	got := it.number 

	if got != want{
		t.Errorf("SetNumber \n got: %v \n want: %v \n", got, want)
	}
}

func TestGetNumber(t *testing.T){
	want := 1
	it.number = 1
	got := it.GetNumber()

	if got != want{
		t.Errorf("GetNumber \n got: %v \n want: %v \n", got, want)
	}
}

func TestSetMemmoryAddress(t *testing.T){
	want := "0x55"
	it.SetMemmoryAddress("0x55")
	got := it.memmoryAddress

	if got != want{
		t.Errorf("SetMemmoryAddress \n got: %v \n want: %v \n", got, want)
	}
}

func TestGetMemmoryAddress(t *testing.T){
	want := "0x55"
	it.memmoryAddress = "0x55"
	got := it.GetMemmoryAddress()

	if got != want{
		t.Errorf("GetMemmoryAddress \n got: %v \n want: %v \n", got, want)
	}
}

func TestSetMemmoryAddressConstructor(t *testing.T){
	it := NewInterruption(false, 2)

	want := "0x100"
	got := it.GetMemmoryAddress()

	if got != want{
		t.Errorf("GetMemmoryAddressConstructor \n got: %v \n want: %v \n", got, want)
	}
}