package registerbank

import "strconv"

type RegisterBank struct{
	registerList map[string]Register
}

func NewRegisterBank() *RegisterBank{
	registerbank := new(RegisterBank)
	registerbank.registerList = make(map[string]Register)

	var registerCode string
	for count := 0; count < 17; count++{
		registerCode = "R" + strconv.Itoa(count)
		registerbank.registerList[registerCode] = *NewRegister(registerCode, "")
	}

	return registerbank
}

func (rbank *RegisterBank) ChangeRegister(regcode string, regvalue int){
	
}