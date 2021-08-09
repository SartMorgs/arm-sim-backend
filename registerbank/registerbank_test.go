package registerbank

import (
	"encoding/json"
	"strconv"
	"testing"
)

var rgbk RegisterBank

func TestChangeRegister(t *testing.T) {
	rgbk := NewRegisterBank()

	want := 77
	rgbk.ChangeRegister("R1", 77)
	got := rgbk.registerList["R1"].decValue

	if got != want {
		t.Errorf("ChangeRegister \n got: %v \n want %v \n", got, want)
	}
}

func TestResetRegisterBank(t *testing.T) {
	rgbk := NewRegisterBank()

	want := [17]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	rgbk.ResetRegisterBank()
	got := rgbk.registerList

	var registerCode string
	for count := 0; count < 17; count++ {
		registerCode = "R" + strconv.Itoa(count)
		if got[registerCode].decValue != want[count] {
			t.Errorf("ResetRegisterBank \n got: %v \n want %v \n", got, want)
			break
		}
	}
}

func TestSetSPRegisterAlias(t *testing.T) {
	rgbk := NewRegisterBank()

	want := "SP"
	rgbk.GetRegisterBank()["R13"].SetRegisterName("SP")
	got := rgbk.GetRegisterBank()["R13"].GetRegisterName()

	if got != want {
		t.Errorf("SetSPRegisterAlias \n got: %v \n want %v \n", got, want)
	}
}

func TestSetLRRegisterAlias(t *testing.T) {
	rgbk := NewRegisterBank()

	want := "LR"
	rgbk.GetRegisterBank()["R14"].SetRegisterName("LR")
	got := rgbk.GetRegisterBank()["R14"].GetRegisterName()

	if got != want {
		t.Errorf("SetLRRegisterAlias \n got: %v \n want %v \n", got, want)
	}
}

func TestSetPCRegisterAlias(t *testing.T) {
	rgbk := NewRegisterBank()

	want := "PC"
	rgbk.GetRegisterBank()["R15"].SetRegisterName("PC")
	got := rgbk.GetRegisterBank()["R15"].GetRegisterName()

	if got != want {
		t.Errorf("SetPCRegisterAlias \n got: %v \n want %v \n", got, want)
	}
}

func TestSetPSRRegisterAlias(t *testing.T) {
	rgbk := NewRegisterBank()

	want := "PSR"
	rgbk.GetRegisterBank()["R16"].SetRegisterName("PSR")
	got := rgbk.GetRegisterBank()["R16"].GetRegisterName()

	if got != want {
		t.Errorf("SetPSRRegisterAlias \n got: %v \n want %v \n", got, want)
	}
}

func TestGetRegisterBankJson(t *testing.T) {
	rgbk := NewRegisterBank()
	test := NewRegisterBank()

	register_bank := make(map[string]string)

	var registerCode string
	for count := 0; count < 17; count++ {
		registerCode = "R" + strconv.Itoa(count)
		str_register_bank := map[string]string{
			"decimal_value":     strconv.FormatInt(int64(test.GetRegisterBank()[registerCode].GetDecValue()), 10),
			"hexadecimal_value": test.GetRegisterBank()[registerCode].GetHexValue(),
			"binary_value":      test.GetRegisterBank()[registerCode].GetBinValue(),
			"register_name":     test.GetRegisterBank()[registerCode].GetRegisterName(),
			"register_function": test.GetRegisterBank()[registerCode].GetRegisterFunction()}

		jrbank, _ := json.MarshalIndent(str_register_bank, "", "")

		register_bank[registerCode] = string(jrbank)
	}

	jregister_bank, _ := json.MarshalIndent(register_bank, "", "")
	want := string(jregister_bank)
	got := rgbk.GetRegisterBankJson()

	if got != want {
		t.Errorf("GetRegisterBankJson \n got: %v \n want %v \n", got, want)
	}
}
