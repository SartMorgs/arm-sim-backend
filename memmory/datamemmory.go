package memmory

import "strconv"

type DataMemmory struct{
	memmoryList map[string]*Memmory
}

func NewDataMemmory() *DataMemmory{
	datamemmory := new(DataMemmory)
	datamemmory.memmoryList = make(map[string]*Memmory)

	var addressMemmoryCode string
	for count := 0; count < 1000; count++{
		addressMemmoryCode = "0x" + strconv.FormatInt(int64(count), 16)
		datamemmory.memmoryList[addressMemmoryCode] = NewMemmory(addressMemmoryCode)
	}

	return datamemmory
}

func (dtmem *DataMemmory) ChangeField(addr string, value int){
	dtmem.memmoryList[addr].SetValue(value)
}