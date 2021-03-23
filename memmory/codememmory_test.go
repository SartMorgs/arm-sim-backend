package memmory

import(
	"testing"
	"strconv"
)

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

func TestAddInstructionField(t *testing.T){
	cmm := NewCodeMemmory()

	want := "10001101010100001010100001011010"
	cmm.AddInstructionField("0x55", "10001101010100001010100001011010")
	got:= cmm.romList["0x55"].GetFullBinValue()

	if got != want{
		t.Errorf("AddInstructionField \n got: %v \n want %v \n", got, want)
	}
}

func TestGetInterruptionAlias(t *testing.T){
	cmm := NewCodeMemmory()

	want := "Int3"
	got := cmm.romList["0x2f8"].GetAliasField()

	if got != want{
		t.Errorf("GetInterruptionAlias \n got: %v \n want %v \n", got, want)
	}
}

func TestResetCodeMemmory(t *testing.T){
	cmm := NewCodeMemmory()

	for count := 0; count < (12 * 1024); count++{
		cmm.AddInstructionField("0x" + strconv.FormatInt(int64(count), 16), "00000000000000000000000000000000")
	}

	cmm.ResetCodeMemmory()
	want := "00000000000000000000000000000000"
	got := cmm.romList["0x55"].GetFullBinValue()

	if got != want{
		t.Errorf("ResetCodeMemmory \n got: %v \n want %v \n", got, want)
	}
}
