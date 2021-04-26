package memmory

import (
	"strconv"
	"encoding/json"
)

type CodeMemmory struct{
	romList map[string]*Memmory
}

func NewCodeMemmory() *CodeMemmory{
	codememmory := new(CodeMemmory)
	codememmory.romList = make(map[string]*Memmory)

	var addressMemmoryCode string

	// Interruption Area
	for countInt := 1; countInt < 17; countInt++{
		for count := ((countInt - 1) * 256); count < (countInt * 256); count++{
			addressMemmoryCode = "0x" + strconv.FormatInt(int64(count), 16)
			codememmory.romList[addressMemmoryCode] = NewMemmory(addressMemmoryCode)
			codememmory.romList[addressMemmoryCode].SetAliasField("Int" + strconv.Itoa(countInt))
		}
	}

	for count := (4 * 1024); count < (12 * 1024); count++{
		addressMemmoryCode = "0x" + strconv.FormatInt(int64(count), 16)
		codememmory.romList[addressMemmoryCode] = NewMemmory(addressMemmoryCode)
	}

	return codememmory
}

func (cdmem *CodeMemmory) AddAliasMemmoryField(addr string, alias string){
	cdmem.romList[addr].SetAliasField(alias)
}

func (cdmem *CodeMemmory) AddInstructionField(addr string, inst string){
	cdmem.romList[addr].SetFullBinValue(inst)
}

func (cdmem *CodeMemmory) ResetCodeMemmory(){
	var addressMemmoryCode string
	for count := 0; count < (12 * 1024); count++{
		addressMemmoryCode = "0x" + strconv.FormatInt(int64(count), 16)
		cdmem.romList[addressMemmoryCode].SetFullBinValue("00000000000000000000000000000000")
	}
}

func (cdmem *CodeMemmory) GetRomList() map[string]*Memmory{
	return cdmem.romList
}

func (cdmem *CodeMemmory) GetCodeMemmoryJson() string{
	code_memmory := make(map[int]string)

	var addressMemmoryCode string
	for count := 0; count < (12 * 1024); count++{
		addressMemmoryCode = "0x" + strconv.FormatInt(int64(count), 16)
		str_memmory := map[string]string{
			"decimal_value": strconv.FormatInt(int64(cdmem.romList[addressMemmoryCode].GetDecValue()), 10),
			"hexadecimal_value": cdmem.romList[addressMemmoryCode].GetHexValue(),
			"binary_value": cdmem.romList[addressMemmoryCode].GetBinValue(),
			"full_binary": cdmem.romList[addressMemmoryCode].GetFullBinValue(),
			"config_type": cdmem.romList[addressMemmoryCode].GetConfigType(),
			"memmory_address": cdmem.romList[addressMemmoryCode].GetAddress(),
			"alias_field": cdmem.romList[addressMemmoryCode].GetAliasField()}

		jmem, _ := json.MarshalIndent(str_memmory, "", "")

		code_memmory[count] = string(jmem)
	}

	jcode_memmory, _ := json.MarshalIndent(code_memmory, "", "")

	return string(jcode_memmory)
}