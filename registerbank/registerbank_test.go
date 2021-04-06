package registerbank

import (
	"testing"
	"strconv"
	"encoding/json"
)

var rgbk RegisterBank 

func TestChangeRegister(t *testing.T){
	rgbk := NewRegisterBank()

	want := 77
	rgbk.ChangeRegister("R1", 77)
	got := rgbk.registerList["R1"].decValue

	if got != want{
		t.Errorf("ChangeRegister \n got: %v \n want %v \n", got, want)
	}
}

func TestResetRegisterBank(t *testing.T){
	rgbk := NewRegisterBank()

	want := [17]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	rgbk.ResetRegisterBank()
	got := rgbk.registerList

	var registerCode string
	for count := 0; count < 17; count++{
		registerCode = "R" + strconv.Itoa(count)
		if got[registerCode].decValue != want[count]{
			t.Errorf("ResetRegisterBank \n got: %v \n want %v \n", got, want)
			break
		}
	}
}

func TestGetRegisterBankJson(t *testing.T){
	rgbk := NewRegisterBank()
	test := NewRegisterBank()

	register_bank := make(map[string]string)

	var registerCode string
	for count := 0; count < 17; count++{
		registerCode = "R" + strconv.Itoa(count)
		str_register_bank := map[string]string{
			"decimal_value": strconv.FormatInt(int64(test.GetRegisterBank()[registerCode].GetDecValue()), 10),
			"hexadecimal_value": test.GetRegisterBank()[registerCode].GetHexValue(),
			"binary_value": test.GetRegisterBank()[registerCode].GetBinValue(),
			"register_name": test.GetRegisterBank()[registerCode].GetRegisterName(),
			"register_function": test.GetRegisterBank()[registerCode].GetRegisterFunction()}

		jrbank, _ := json.Marshal(str_register_bank)

		register_bank[registerCode] = string(jrbank)
	}

	jregister_bank, _ := json.Marshal(register_bank)
	want := string(jregister_bank)
	got := rgbk.GetRegisterBankJson()
	
	if got != want{
		t.Errorf("GetRegisterBankJson \n got: %v \n want %v \n", got, want)
	}
}