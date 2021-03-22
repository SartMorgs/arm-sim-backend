package memmory

import "strconv"

type Memmory struct{
	decValue int
	hexValue string
	binValue string
	fullBinValue string
	splitedBinValue [4]string
	configType string
	memmoryAddress string
	aliasField string
}

func NewMemmory(memaddr string) *Memmory{
	memmory := new(Memmory)

	memmory.decValue = 0
	memmory.hexValue = "0x0"
	memmory.binValue = "00"
	memmory.fullBinValue = "00000000000000000000000000000000"
	memmory.configType = "D"
	memmory.memmoryAddress = memaddr 
	memmory.aliasField = ""

	return memmory
}

func (m *Memmory) SetAddress(addr string){
	m.memmoryAddress = addr
}

func (m *Memmory) SetValue(value int){
	m.decValue = value
}

func (m *Memmory) SetConfigType(tp string){
	m.configType = tp
}

func (m* Memmory) SetAliasField(alias string){
	m.aliasField = alias
}

func (m* Memmory) SetFullBinValue(value string){
	m.fullBinValue = value
}

func (m *Memmory) GetAddress() string{
	return m.memmoryAddress
}

func (m *Memmory) GetDecValue() int{
	return m.decValue
}

func (m *Memmory) GetBinValue() string{
	return m.binValue
}

func (m *Memmory) GetHexValue() string{
	return m.hexValue
}

func (m *Memmory) GetConfigType() string{
	return m.configType
}

func (m *Memmory) GetSplitBin() [4]string{
	return m.splitedBinValue
}

func (m *Memmory) GetAliasField() string{
	return m.aliasField
}

func (m *Memmory) GetFullBinValue() string{
	return m.fullBinValue
}

func (m *Memmory) toBinValue(){
	m.binValue = strconv.FormatInt(int64(m.decValue), 2)
}

func (m *Memmory) appendBinValue(){
	max_value := 32

	appendBin := m.binValue

	for count := (max_value - len(appendBin)); count > 0; count--{
		appendBin = "0" + appendBin
	} 

	m.fullBinValue = appendBin
}

func (m *Memmory) toHexValue() {
	m.hexValue = "0x" + strconv.FormatInt(int64(m.decValue), 16)
}

func (m *Memmory) splitBin(){

	bit1 := m.fullBinValue[0:8]
	bit2 := m.fullBinValue[8:16]
	bit3 := m.fullBinValue[16:24]
	bit4 := m.fullBinValue[24:32]

	if m.configType == "L"{
		m.splitedBinValue = [4]string{bit1, bit2, bit3, bit4}
	} else if m.configType == "B"{
		m.splitedBinValue = [4]string{bit4, bit3, bit2, bit1}
	} else{
		m.splitedBinValue = [4]string{"00000000", "00000000", "00000000", "00000000"}
	}
}