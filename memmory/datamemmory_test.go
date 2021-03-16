package memmory

import "testing"

var dmm DataMemmory

// Unit Tests
// Test all methods using mock
func TestChangeDataMemmoryField(t *testing.T){
	dmm := NewDataMemmory()

	want := 77
	dmm.ChangeField("0x55", 77)
	got := dmm.memmoryList["0x55"].GetDecValue()

	if got != want{
		t.Errorf("ChangeDataMemmoryField \n got: %v \n want %v \n", got, want)
	}
}

