package memmory

import (
	"strconv"
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