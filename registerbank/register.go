package registerbank

import "strconv"

type Register struct {
	decValue int
	hexValue string
	binValue string
	registerName string
	registerFunction string
}

func NewRegister(rgnm string, rgfc string) *Register{
	register := new(Register)

	register.decValue = 0
	register.hexValue = "0x0"
	register.binValue = "00"
	register.registerName = rgnm
	register.registerFunction = rgfc

	return register
}

func (r *Register) SetValue(value int){
	r.decValue = value
	r.toBinValue()
	r.toHexValue()
}

func (r *Register) SetRegisterName(name string){
	r.registerName = name
}

func (r *Register) SetRegisterFunction(function string){
	r.registerFunction = function
}

func (r *Register) GetDecValue() int{
	return r.decValue
}

func (r *Register) GetBinValue() string{
	return r.binValue
}

func (r *Register) GetHexValue() string{
	return r.hexValue
}

func (r *Register) GetRegisterName() string{
	return r.registerName
}

func (r *Register) GetRegisterFunction() string{
	return r.registerFunction
}

func (r *Register) toBinValue() {
	r.binValue = strconv.FormatInt(int64(r.decValue), 2)
}

func (r *Register) toHexValue(){
	r.hexValue = "0x" + strconv.FormatInt(int64(r.decValue), 16)
}
