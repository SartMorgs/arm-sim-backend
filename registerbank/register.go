package registerbank

import "strconv"

type Register struct {
	value int
	binValue string
	name string
}

func (r *Register) GetValue() int{
	return r.value
}

func (r *Register) SetValue(val int){
	r.value = val
}

func (r *Register) GetBinValue() string{
	return r.binValue
}

func (r *Register) SetBinValue(bin_value string) {
	r.binValue = bin_value
}

func (r *Register) SetName(nm string){
	r.name = nm
}

func (r *Register) GetName() string{
	return r.name
}

func (r *Register) buildBinValue() {
	var value_iteraction int = r.value
	var mod_value int = 0
	var binVal string

	for ok:= true; ok; ok = value_iteraction > 0 {
		mod_value = value_iteraction % 2
		binVal = binVal + strconv.Itoa(mod_value)
		value_iteraction = value_iteraction / 2
	}

	r.SetBinValue(binVal)
}