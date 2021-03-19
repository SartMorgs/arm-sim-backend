package memmory

import "strconv"

type CodeMemmory struct{
	romList map[string]*Memmory
}

func NewCodeMemmory() *CodeMemmory{
	codememmory := new(CodeMemmory)
	codememmory.romList = make(map[string]*Memmory)

	var addressMemmoryCode string
	for count := 0; count < 1000; count++{
		addressMemmoryCode = "0x" + strconv.FormatInt(int64(count), 16)
		codememmory.romList[addressMemmoryCode] = NewMemmory(addressMemmoryCode)
	}

	return codememmory
}

func (cdmem *CodeMemmory) AddAliasMemmoryField(addr string, alias string){
	cdmem.romList[addr].SetAliasField(alias)
}