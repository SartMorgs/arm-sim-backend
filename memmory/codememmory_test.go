package memmory

import "testing"

var cmm CodeMemmory

// Unit Tests
// Test all methods using mock
func TestAddAliasMemmoryField(t *testing.T){
	cmm := NewCodeMemmory()

	want := "Area-A"
	cmm.AddAliasMemmoryField("0x55", "Area-A")
	got := cmm.romList["0x55"].GetAliasField()

	if got != want{
		t.Errorf("AddAliasMemmoryField \n got: %v \n want %v \n", got, want)
	}
}