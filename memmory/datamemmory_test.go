package memmory

import(
	"testing"
	"strconv"
)

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

func TestResetDataMemmory(t *testing.T){
	dmm := NewDataMemmory()

	for count := 0; count < (4 * 1024); count++{
		dmm.ChangeField("0x" + strconv.FormatInt(int64(count), 16), count)
	}

	dmm.ResetDataMemmory()
	want := 0
	got := dmm.memmoryList["0x55"].GetDecValue()

	if got != want{
		t.Errorf("ResetDataMemmory \n got: %v \n want %v \n", got, want)
	}
}
