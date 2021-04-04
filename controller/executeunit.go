package controller

import "strconv"

type ExecuteUnit struct{
	instruction string
	config map[string]int64
}

func NewExecuteUnit() *ExecuteUnit{
	executeunit := new(ExecuteUnit)
	executeunit.config = make(map[string]int64)

	return executeunit
}

func (eu *ExecuteUnit) SetInstruction(inst string){
	eu.instruction = inst
}

func (eu *ExecuteUnit) GetInstruction() string{
	return eu.instruction
}

func (eu *ExecuteUnit) ConfigForLdr1(target_reg string, addr_reg string){
	x, _ := strconv.ParseInt(target_reg, 2, 64)
	eu.config["target_register"] = x

	y, _ := strconv.ParseInt(addr_reg, 2, 64)
	eu.config["address_register"] = y
}


func (eu *ExecuteUnit) ConfigForLdr2(target_reg string, addr string){
	x, _ := strconv.ParseInt(target_reg, 2, 64)
	eu.config["target_register"] = x

	y, _ := strconv.ParseInt(addr, 2, 64)
	eu.config["address"] = y
}

func (eu *ExecuteUnit) ConfigForStr1(target_reg string, addr_reg string){
	x, _ := strconv.ParseInt(target_reg, 2, 64)
	eu.config["target_register"] = x

	y, _ := strconv.ParseInt(addr_reg, 2, 64)
	eu.config["address_register"] = y
}

func (eu *ExecuteUnit) ConfigForStr2(target_reg string, addr string){
	x, _ := strconv.ParseInt(target_reg, 2, 64)
	eu.config["target_register"] = x

	y, _ := strconv.ParseInt(addr, 2, 64)
	eu.config["address"] = y
}

func (eu *ExecuteUnit) ConfigForAdds1(target_reg string, source_reg1 string, source_reg2 string){
	x, _ := strconv.ParseInt(target_reg, 2, 64)
	eu.config["target_register"] = x

	y, _ := strconv.ParseInt(source_reg1, 2, 64)
	eu.config["source_register1"] = y

	z, _ := strconv.ParseInt(source_reg2, 2, 64)
	eu.config["source_register2"] = z
}

func (eu *ExecuteUnit) ConfigForAdds2(target_reg string, source_reg string, value string){
	x, _ := strconv.ParseInt(target_reg, 2, 64)
	eu.config["target_register"] = x

	y, _ := strconv.ParseInt(source_reg, 2, 64)
	eu.config["source_register"] = y

	z, _ := strconv.ParseInt(value, 2, 64)
	eu.config["value"] = z
}

func (eu *ExecuteUnit) ConfigForSubs1(target_reg string, source_reg1 string, source_reg2 string){
	x, _ := strconv.ParseInt(target_reg, 2, 64)
	eu.config["target_register"] = x

	y, _ := strconv.ParseInt(source_reg1, 2, 64)
	eu.config["source_register1"] = y

	z, _ := strconv.ParseInt(source_reg2, 2, 64)
	eu.config["source_register2"] = z
}

func (eu *ExecuteUnit) ConfigForSubs2(target_reg string, source_reg string, value string){
	x, _ := strconv.ParseInt(target_reg, 2, 64)
	eu.config["target_register"] = x

	y, _ := strconv.ParseInt(source_reg, 2, 64)
	eu.config["source_register"] = y

	z, _ := strconv.ParseInt(value, 2, 64)
	eu.config["value"] = z
}

func (eu *ExecuteUnit) ConfigForAnds1(target_reg string, source_reg1 string, source_reg2 string){
	x, _ := strconv.ParseInt(target_reg, 2, 64)
	eu.config["target_register"] = x

	y, _ := strconv.ParseInt(source_reg1, 2, 64)
	eu.config["source_register1"] = y

	z, _ := strconv.ParseInt(source_reg2, 2, 64)
	eu.config["source_register2"] = z
}

func (eu *ExecuteUnit) ConfigForAnds2(target_reg string, source_reg string, value string){
	x, _ := strconv.ParseInt(target_reg, 2, 64)
	eu.config["target_register"] = x

	y, _ := strconv.ParseInt(source_reg, 2, 64)
	eu.config["source_register"] = y

	z, _ := strconv.ParseInt(value, 2, 64)
	eu.config["value"] = z
}

func (eu *ExecuteUnit) ConfigForMovs1(target_reg string, value_reg string){
	x, _ := strconv.ParseInt(target_reg, 2, 64)
	eu.config["target_register"] = x

	y, _ := strconv.ParseInt(value_reg, 2, 64)
	eu.config["value_register"] = y
}

func (eu *ExecuteUnit) ConfigForMovs2(target_reg string, value string){
	x, _ := strconv.ParseInt(target_reg, 2, 64)
	eu.config["target_register"] = x

	y, _ := strconv.ParseInt(value, 2, 64)
	eu.config["value"] = y
}