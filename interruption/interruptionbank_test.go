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

func TestResetInterruptionBank(t *testing.T){
	itbk := NewInterruptionBank()

	want := [17]bool{false, false, false, false, false, false, false, false, false, false, false, false, false, false, false}
	itbk.ResetInterruptionBank()
	got := itbk.interruptionList

	for count := 1; count < 17; count++{
		if got[count].trigger != want[count]{
			t.Errorf("ResetInterruptionBank \n got: %v \n want %v \n", got, want)
			break
		}
	}
}