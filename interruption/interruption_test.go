package interruption

import "testing"

var int Interruption

func TestTurnOnTrigger(t *testing.T){
	want := true
	int.TurnOnTrigger()
	got := int.trigger

	if got != want{
		t.Errorf("TurnOnTrigger \n got: %v \n want: %v \n", got, want)
	}
}

func TestTurnOffTrigger(t *testing.T){
	want := false
	int.TurnOffTrigger()
	got := int.trigger

	if got != want{
		t.Errorf("TurnOffTrigger \n got: %v \n want: %v \n", got, want)
	}
}

func TestSaveSystemStatus(t *testing.T){
	want := "00000000000001010010001101000001"
	int.SaveSystemStatus("00000000000001010010001101000001")
	got := int.systemSaved

	if got != want{
		t.Errorf("SaveSystemStatus \n got: %v \n want: %v \n", got, want)
	}
}