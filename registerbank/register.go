package registerbank

import "strconv"

type Register struct {
	Value int
	BinValue string
	Name string
}

func (r *Register) GetValue() int{
	return r.Value
}

func (r *Register) SetValue(value int){
	r.Value = value
}

func (r *Register) GetBinValue() string{
	return r.BinValue
}

func (r *Register) SetBinValue(bin_value string) {
	r.BinValue = bin_value
}

func (r *Register) SetName(name string){
	r.Name = name
}

func (r *Register) GetName() string{
	return r.Name
}

func (r *Register) buildBinValue() {
	var value_iteraction int = r.Value
	var mod_value int = 0
	var binValue string

	for ok:= true; ok; ok = value_iteraction > 0 {
		mod_value = value_iteraction % 2
		binValue = binValue + strconv.Itoa(mod_value)
		value_iteraction = value_iteraction / 2
	}
}