package registerbank

import(
	"strconv"
	"encoding/json"
)

type RegisterBank struct{
	registerList map[string]*Register
}

func NewRegisterBank() *RegisterBank{
	registerbank := new(RegisterBank)
	registerbank.registerList = make(map[string]*Register)

	var registerCode string
	for count := 0; count < 17; count++{
		registerCode = "R" + strconv.Itoa(count)
		registerbank.registerList[registerCode] = NewRegister(registerCode, "")
	}

	return registerbank
}

// With register refered on argument, set new value
func (rbank *RegisterBank) ChangeRegister(regcode string, regvalue int){
	rbank.registerList[regcode].SetValue(regvalue)
}

func (rbank *RegisterBank) ResetRegisterBank(){
	var registerCode string
	for count := 0; count < 17; count++{
		registerCode = "R" + strconv.Itoa(count)
		rbank.registerList[registerCode].decValue = 0
	}
}

func (rbank *RegisterBank) GetRegisterBank() map[string]*Register{
	return rbank.registerList
}

func (rbank *RegisterBank) GetRegisterBankJson() string{
	register_bank := make(map[string]string)

	var registerCode string
	for count := 0; count < 17; count++{
		registerCode = "R" + strconv.Itoa(count)
		str_register_bank := map[string]string{
			"decimal_value": strconv.FormatInt(int64(rbank.registerList[registerCode].GetDecValue()), 10),
			"hexadecimal_value": rbank.registerList[registerCode].GetHexValue(),
			"binary_value": rbank.registerList[registerCode].GetBinValue(),
			"register_name": rbank.registerList[registerCode].GetRegisterName(),
			"register_function": rbank.registerList[registerCode].GetRegisterFunction()}

		jrbank, _ := json.MarshalIndent(str_register_bank, "", "")

		register_bank[registerCode] = string(jrbank)
	}

	jregister_bank, _ := json.MarshalIndent(register_bank, "", "")

	return string(jregister_bank)
}