package memmory

import "strconv"

type Memmory struct{
	decValue int
	hexValue string
	binValue string
	memmoryAddress string
}

func NewMemmory(memaddr string) *Memmory{
	memmory := new(Memmory)

	memmory.decValue = 0
	memmory.hexValue = "0x0"
	memmory.binValue = "00"
	memmory.memmoryAddress = memaddr 

	return memmory
}

func (m *Memmory) SetAddress(addr string){
	m.memmoryAddress = addr
}

func (m *Memmory) SetValue(value int){
	m.decValue = value
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

func (m *Memmory) toBinValue(){
	m.binValue = strconv.FormatInt(int64(m.decValue), 2)
}

func (m *Memmory) toHexValue() {
	m.hexValue = "0x" + strconv.FormatInt(int64(m.decValue), 16)
}