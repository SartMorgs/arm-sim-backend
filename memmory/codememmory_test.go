package memmory

import(
	"testing"
	"strconv"
	"encoding/json"
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

func TestGetCodeMemmoryJson(t *testing.T){
	cmm := NewCodeMemmory()
	test := NewCodeMemmory()

	code_memmory := make(map[int]string)

	var addressMemmoryCode string
	for count := 0; count < (12 * 1024); count++{
		addressMemmoryCode = "0x" + strconv.FormatInt(int64(count), 16)
		str_memmory := map[string]string{
			"decimal_value": strconv.FormatInt(int64(test.GetRomList()[addressMemmoryCode].GetDecValue()), 10),
			"hexadecimal_value": test.GetRomList()[addressMemmoryCode].GetHexValue(),
			"binary_value": test.GetRomList()[addressMemmoryCode].GetBinValue(),
			"full_binary": test.GetRomList()[addressMemmoryCode].GetFullBinValue(),
			"config_type": test.GetRomList()[addressMemmoryCode].GetConfigType(),
			"memmory_address": test.GetRomList()[addressMemmoryCode].GetAddress(),
			"alias_field": test.GetRomList()[addressMemmoryCode].GetAliasField()}

		jmem, _ := json.Marshal(str_memmory)

		code_memmory[count] = string(jmem)
	}

	jcode_memmory, _ := json.Marshal(code_memmory)
	want := string(jcode_memmory)
	got := cmm.GetCodeMemmoryJson()

	if got != want{
		t.Errorf("GetCodeMemmoryJson \n got: %v \n want %v \n", got, want)
	}
}
