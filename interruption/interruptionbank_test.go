package interruption

import "testing"

var itbk InterruptionBank

func TestTriggerInterruption(t *testing.T){
	itbk := NewInterruptionBank()

	want := true
	itbk.TriggerInterruption(5)
	got := itbk.interruptionList[5].GetTrigger()

	if got != want{
		t.Errorf("TriggerInterruption \n got: %v \n want %v \n", got, want)
	}
}

func TestTurnOffInterruption(t *testing.T){
	itbk := NewInterruptionBank()

	want := false
	itbk.TurnOffInterruption(5)
	got := itbk.interruptionList[5].GetTrigger()

	if got != want{
		t.Errorf("TurnOffInterruption \n got: %v \n want %v \n", got, want)
	}
}