package register_bank

import "strconv"

type Register struct {
	Value int
	BinValue string
}

func (r Register) getValue() int{
	return r.Value
}

func (r Register) setValue(value int){
	r.Value = value
}

func (r Register) getBinValue() string{
	return r.BinValue
}

func (r Register) setBinValue(bin_value string) {
	r.BinValue = bin_value
}

func (r Register) buildBinValue() {
	var value_iteraction int = r.Value
	var mod_value int = 0
	var binValue string

	for ok:= true; ok; ok = value_iteraction > 0 {
		mod_value = value_iteraction % 2
		binValue = binValue + strconv.Itoa(mod_value)
		value_iteraction = value_iteraction / 2
	}
}