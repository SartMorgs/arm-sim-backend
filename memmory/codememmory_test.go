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

func TestSetInputCode(t *testing.T){
	cmm := NewCodeMemmory()

	code_json, _ := json.MarshalIndent("{\r\n    \"main\": [\r\n        \"01010000000000000001101010000000\",\r\n        \"01010000001000000000100100000000\",\r\n        \"00000100010000000000100000000000\",\r\n        \"00001000011000000000100000000000\",\r\n        \"110101F1F00000000000000000000000\",\r\n        \"010111F0F00000000000000000000000\"\r\n    ],\r\n    \"func1\": [\r\n        \"01010000001000000000100100000000\",\r\n        \"00000100010000000000100000000000\"\r\n    ],\r\n    \"INT0_Handler\": [\r\n        \"00001000011000000000100000000000\",\r\n        \"00001000011000000000100000000000\"\r\n    ]\r\n}", "", "")
	code := string(code_json)

	want := code

	cmm.SetInputCode(code)

	got := cmm.inputCode

	if got != want{
		t.Errorf("SetInputCode \n got: %v \n want %v \n", got, want)
	}
}

func TestBuildFunctionMap(t *testing.T){
	cmm := NewCodeMemmory()

	code_json, _ := json.MarshalIndent("{\r\n    \"main\": [\r\n        \"01010000000000000001101010000000\",\r\n        \"01010000001000000000100100000000\",\r\n        \"00000100010000000000100000000000\",\r\n        \"00001000011000000000100000000000\",\r\n        \"110101F1F00000000000000000000000\",\r\n        \"010111F0F00000000000000000000000\"\r\n    ],\r\n    \"func1\": [\r\n        \"01010000001000000000100100000000\",\r\n        \"00000100010000000000100000000000\"\r\n    ],\r\n    \"INT0_Handler\": [\r\n        \"00001000011000000000100000000000\",\r\n        \"00001000011000000000100000000000\"\r\n    ]\r\n}", "", "")
	code := string(code_json)

	function_map_json, _ := json.MarshalIndent(`functions: {
    "main": ["0x0", "0x5"],
    "func1": ["0x6", "0x0x7"],
}
interruptions: {
}`, "", "")

	want := string(function_map_json)

	cmm.SetInputCode(code)
	cmm.BuildFunctionMap()

	got := cmm.functionMap 

	if got != want{
		t.Errorf("BuildFunctionMap \n got: %v \n want %v \n", got, want)
	}
}
