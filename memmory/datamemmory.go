package memmory

import (
	"strconv"
	"encoding/json"
)

type DataMemmory struct{
	memmoryList map[string]*Memmory
}

func NewDataMemmory() *DataMemmory{
	datamemmory := new(DataMemmory)
	datamemmory.memmoryList = make(map[string]*Memmory)

	var addressMemmoryCode string
	for count := 0; count < (4 * 1024); count++{
		addressMemmoryCode = "0x" + strconv.FormatInt(int64(count), 16)
		datamemmory.memmoryList[addressMemmoryCode] = NewMemmory(addressMemmoryCode)
	}

	return datamemmory
}

func (dtmem *DataMemmory) ChangeField(addr string, value int){
	dtmem.memmoryList[addr].SetValue(value)
}

func (dtmem *DataMemmory) ResetDataMemmory(){
	var addressMemmoryCode string
	for count := 0; count < (4 * 1024); count++{
		addressMemmoryCode = "0x" + strconv.FormatInt(int64(count), 16)
		dtmem.memmoryList[addressMemmoryCode].SetValue(0)
	}
}

func (dtmem *DataMemmory) GetDataMemmoryList() map[string]*Memmory{
	return dtmem.memmoryList
}

func (dtmem *DataMemmory) GetDataMemmoryJson() string{
	data_memmory := make(map[int]string)

	var addressMemmoryCode string
	for count := 0; count < (4 * 1024); count++{
		addressMemmoryCode = "0x" + strconv.FormatInt(int64(count), 16)
		str_memmory := map[string]string{
			"decimal_value": strconv.FormatInt(int64(dtmem.memmoryList[addressMemmoryCode].GetDecValue()), 10),
			"hexadecimal_value": dtmem.memmoryList[addressMemmoryCode].GetHexValue(),
			"binary_value": dtmem.memmoryList[addressMemmoryCode].GetBinValue(),
			"full_binary": dtmem.memmoryList[addressMemmoryCode].GetFullBinValue(),
			"config_type": dtmem.memmoryList[addressMemmoryCode].GetConfigType(),
			"memmory_address": dtmem.memmoryList[addressMemmoryCode].GetAddress(),
			"alias_field": dtmem.memmoryList[addressMemmoryCode].GetAliasField()}

		jmem, _ := json.MarshalIndent(str_memmory, "", "")

		data_memmory[count] = string(jmem)
	}

	jdata_memmory, _ := json.MarshalIndent(data_memmory, "", "")

	return string(jdata_memmory)
}