package memmory

import (
	"encoding/json"
	"reflect"
	"strconv"
	"testing"
)

var cmm CodeMemmory

// Unit Tests
// Test all methods using mock
func TestAddAliasMemmoryField(t *testing.T) {
	cmm := NewCodeMemmory()

	want := "Area-A"
	cmm.AddAliasMemmoryField("0x55", "Area-A")
	got := cmm.romList["0x55"].GetAliasField()

	if got != want {
		t.Errorf("AddAliasMemmoryField \n got: %v \n want %v \n", got, want)
	}
}

func TestAddInstructionField(t *testing.T) {
	cmm := NewCodeMemmory()

	want := "10001101010100001010100001011010"
	cmm.AddInstructionField("0x55", "10001101010100001010100001011010")
	got := cmm.romList["0x55"].GetFullBinValue()

	if got != want {
		t.Errorf("AddInstructionField \n got: %v \n want %v \n", got, want)
	}
}

func TestGetInterruptionAlias(t *testing.T) {
	cmm := NewCodeMemmory()

	want := "Int3"
	got := cmm.romList["0x2f8"].GetAliasField()

	if got != want {
		t.Errorf("GetInterruptionAlias \n got: %v \n want %v \n", got, want)
	}
}

func TestResetCodeMemmory(t *testing.T) {
	cmm := NewCodeMemmory()

	for count := 0; count < (12 * 1024); count++ {
		cmm.AddInstructionField("0x"+strconv.FormatInt(int64(count), 16), "00000000000000000000000000000000")
	}

	cmm.ResetCodeMemmory()
	want := "00000000000000000000000000000000"
	got := cmm.romList["0x55"].GetFullBinValue()

	if got != want {
		t.Errorf("ResetCodeMemmory \n got: %v \n want %v \n", got, want)
	}
}

func TestSetInputCode(t *testing.T) {
	cmm := NewCodeMemmory()

	json_input := `{
		"main": [
			"01010000000000000001101010000000",
			"01010000001000000000100100000000",
			"00000100010000000000100000000000",
			"00001000011000000000100000000000",
			"110101F1F00000000000000000000000",
			"010111F0F00000000000000000000000"
		],
		"func1": [
			"01010000001000000000100100000000",
			"00000100010000000000100000000000"
		],
		"INT0_Handler": [
			"00001000011000000000100000000000",
			"00001000011000000000100000000000"
		]
	}`
	code_json := make(map[string]interface{})
	json.Unmarshal([]byte(json_input), &code_json)

	want := code_json

	cmm.SetInputCode(code_json)

	got := cmm.inputCode

	if !reflect.DeepEqual(want, got) {
		t.Errorf("SetInputCode \n got: %v \n want %v \n", got, want)
	}
}

func TestBuildFunctionMap(t *testing.T) {
	cmm := NewCodeMemmory()

	json_input := `{
		"main": [
			"01010000000000000001101010000000",
			"01010000001000000000100100000000",
			"00000100010000000000100000000000",
			"00001000011000000000100000000000",
			"110101F1F00000000000000000000000",
			"010111F0F00000000000000000000000"
		],
		"func1": [
			"01010000001000000000100100000000",
			"00000100010000000000100000000000"
		],
		"INT0_Handler": [
			"00001000011000000000100000000000",
			"00001000011000000000100000000000"
		]
	}`
	code_json := make(map[string]interface{})
	json.Unmarshal([]byte(json_input), &code_json)

	function_map_json := `{
		"functions": {
			"main": {
				"start": "0x0", 
				"end": "0x5"
			},
			"func1": {
				"start": "0x6",
				"end": "0x7"
			}
		},
		"interruptions": {
			"INT0_Handler": {
				"start": "0x0", 
				"end": "0x1"
			}
		}
	}`
	json_output := make(map[string]map[string]map[string]string)
	json.Unmarshal([]byte(function_map_json), &json_output)

	want := json_output

	cmm.SetInputCode(code_json)
	cmm.BuildFunctionMap()

	got := cmm.functionMap

	if !reflect.DeepEqual(want, got) {
		t.Errorf("BuildFunctionMap \n got: %v \n want %v \n", got, want)
	}
}

func TestGetInputCode(t *testing.T) {
	cmm := NewCodeMemmory()

	json_input := `{
		"main": [
			"01010000000000000001101010000000",
			"01010000001000000000100100000000",
			"00000100010000000000100000000000",
			"00001000011000000000100000000000",
			"110101F1F00000000000000000000000",
			"010111F0F00000000000000000000000"
		],
		"func1": [
			"01010000001000000000100100000000",
			"00000100010000000000100000000000"
		],
		"INT0_Handler": [
			"00001000011000000000100000000000",
			"00001000011000000000100000000000"
		]
	}`
	code_json := make(map[string]interface{})
	json.Unmarshal([]byte(json_input), &code_json)

	want := code_json
	cmm.inputCode = code_json
	got := cmm.GetInputCode()

	if !reflect.DeepEqual(want, got) {
		t.Errorf("GetInputCode \n got: %v \n want %v \n", got, want)
	}
}

func TestFixFunctionAddressIntoMainCode(t *testing.T) {
	cmm := NewCodeMemmory()

	json_input := `{
		"main": [
			"01010000000000000001101010000000",
			"01010000001000000000100100000000",
			"00000100010000000000100000000000",
			"00001000011000000000100000000000",
			"110101F1F00000000000000000000000",
			"010111F0F00000000000000000000000"
		],
		"func1": [
			"01010000001000000000100100000000",
			"00000100010000000000100000000000"
		],
		"INT0_Handler": [
			"00001000011000000000100000000000",
			"00001000011000000000100000000000"
		]
	}`
	code_json := make(map[string]interface{})
	json.Unmarshal([]byte(json_input), &code_json)

	json_want := `{
		"main": [
			"01010000000000000001101010000000",
			"01010000001000000000100100000000",
			"00000100010000000000100000000000",
			"00001000011000000000100000000000",
			"11010100000000000000000000000110",
			"01011100000000000000000000000000"
		],
		"func1": [
			"01010000001000000000100100000000",
			"00000100010000000000100000000000"
		],
		"INT0_Handler": [
			"00001000011000000000100000000000",
			"00001000011000000000100000000000"
		]
	}`
	want_json := make(map[string]interface{})
	json.Unmarshal([]byte(json_want), &want_json)

	want := want_json
	cmm.inputCode = code_json
	cmm.BuildFunctionMap()
	cmm.FixFunctionAddressIntoMainCode()
	got := cmm.inputCode

	if !reflect.DeepEqual(want, got) {
		t.Errorf("FixFunctionAddressIntoMainCode \n got: %v \n want %v \n", got, want)
	}
}

func TestBuildCodeMemmory(t *testing.T) {
	cmm := NewCodeMemmory()

	var test_program []string = []string{
		"01010000000000000001101010000000",
		"01010000001000000000100100000000",
		"00000100010000000000100000000000",
		"00001000011000000000100000000000",
		"11010100000000000000000000000110",
		"01011100000000000000000000000000"}
	var test_function []string = []string{
		"01010000001000000000100100000000",
		"00000100010000000000100000000000"}
	var test_interruption []string = []string{
		"00001000011000000000100000000000",
		"00001000011000000000100000000000"}

	json_input := `{
		"main": [
			"01010000000000000001101010000000",
			"01010000001000000000100100000000",
			"00000100010000000000100000000000",
			"00001000011000000000100000000000",
			"110101F1F00000000000000000000000",
			"010111F0F00000000000000000000000"
		],
		"func1": [
			"01010000001000000000100100000000",
			"00000100010000000000100000000000"
		],
		"INT0_Handler": [
			"00001000011000000000100000000000",
			"00001000011000000000100000000000"
		]
	}`
	code_json := make(map[string]interface{})
	json.Unmarshal([]byte(json_input), &code_json)

	cmm.SetInputCode(code_json)
	cmm.BuildFunctionMap()
	cmm.FixFunctionAddressIntoMainCode()
	cmm.AddCodeIntoMemmory()

	cmm_want := NewCodeMemmory()

	var addressMemmoryCode string
	for iterator := 0; iterator < len(test_interruption); iterator++ {
		addressMemmoryCode = "0x" + strconv.FormatInt(int64(iterator), 16)
		cmm_want.AddInstructionField(addressMemmoryCode, test_interruption[iterator])
	}

	start_address_main := (4 * 1024)
	for iterator := start_address_main; iterator < (start_address_main + len(test_program)); iterator++ {
		addressMemmoryCode = "0x" + strconv.FormatInt(int64(iterator), 16)
		cmm_want.AddInstructionField(addressMemmoryCode, test_program[iterator-start_address_main])
	}

	start_address_function := start_address_main + len(test_program)
	for iterator := start_address_function; iterator < (start_address_function + len(test_function)); iterator++ {
		addressMemmoryCode = "0x" + strconv.FormatInt(int64(iterator), 16)
		cmm_want.AddInstructionField(addressMemmoryCode, test_function[iterator-start_address_function])
	}

	want := cmm.GetCodeMemmoryJson()
	got := cmm_want.GetCodeMemmoryJson()

	if got != want {
		t.Errorf("BuildCodeMemmory \n got: \n want \n")
	}
}

func TestSetDirectiveCode(t *testing.T) {
	cmm := NewCodeMemmory()

	json_directive := `{
		"1": [
			"addr1",
			"EQU",
			"0x10"
		],
		"2": [
			"AREA",
			"main",
			"CODE",
			"READONLY"
		]
	}`
	directive_json := make(map[string]interface{})
	json.Unmarshal([]byte(json_directive), &directive_json)

	want := directive_json

	cmm.SetDirectiveCode(directive_json)

	got := cmm.directiveCode

	if !reflect.DeepEqual(want, got) {
		t.Errorf("SetDirectiveCode \n got: %v \n want %v \n", got, want)
	}
}

func TestGetDirectiveCode(t *testing.T) {
	cmm := NewCodeMemmory()

	json_directive := `{
		"1": [
			"addr1",
			"EQU",
			"0x10"
		],
		"2": [
			"AREA",
			"main",
			"CODE",
			"READONLY"
		]
	}`
	directive_json := make(map[string]interface{})
	json.Unmarshal([]byte(json_directive), &directive_json)

	want := directive_json

	cmm.directiveCode = directive_json

	got := cmm.GetDirectiveCode()

	if !reflect.DeepEqual(want, got) {
		t.Errorf("SetDirectiveCode \n got: %v \n want %v \n", got, want)
	}
}

func TestAddDirectivesAliasIntoMemmory(t *testing.T) {
	cmm := NewCodeMemmory()

	json_directive := `{
		"1": [
			"addr1",
			"EQU",
			"0x10"
		],
		"2": [
			"AREA",
			"main",
			"CODE",
			"READONLY"
		]
	}`
	directive_json := make(map[string]interface{})
	json.Unmarshal([]byte(json_directive), &directive_json)
	//print(directive_json["1"].([]interface{})[2].(string))
	//print(len(directive_json))

	cmm.directiveCode = directive_json
	cmm.AddDirectivesAliasIntoMemmory()

	cmm_want := NewCodeMemmory()
	cmm_want.AddAliasMemmoryField("0x10", "addr1")

	got := cmm.GetCodeMemmoryJson()
	want := cmm_want.GetCodeMemmoryJson()

	if got != want {
		t.Errorf("AddDirectivesAliasIntoMemmory \n got: \n want \n")
	}
}
