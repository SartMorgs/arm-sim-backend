package controller

import "strconv"

type ExecuteUnit struct{
	instruction string
	config map[string]int64
	functionInstructionMap map[[2]string]interface{}
}

func NewExecuteUnit() *ExecuteUnit{
	executeunit := new(ExecuteUnit)
	executeunit.config = make(map[string]int64)
	
	executeunit.functionInstructionMap = map[[2]string]interface{}{
		// Arithmetic
		{"ADDS", "1"}: executeunit.ConfigForAdds1,
		{"ADDS", "2"}: executeunit.ConfigForAdds2,
		{"SUBS", "1"}: executeunit.ConfigForSubs1,
		{"SUBS", "2"}: executeunit.ConfigForSubs2,
		//{"MULS", "1"}: executeunit.ConfigForMuls1,
		//{"MULS", "2"}: executeunit.ConfigForMuls2,
		{"ANDS", "1"}: executeunit.ConfigForAnds1,
		{"ANDS", "2"}: executeunit.ConfigForAnds2,
		//{"ORRS", "1"}: executeunit.ConfigForOrrs1,
		//{"ORRS", "2"}: executeunit.ConfigForOrrs2,
		//{"EORS", "1"}: executeunit.ConfigForEors1,
		//{"EORS", "2"}: executeunit.ConfigForEors2,
		//{"BICS", "1"}: executeunit.ConfigForBics1,
		//{"BICS", "2"}: executeunit.ConfigForBics2,
		//{"ASRS", "1"}: executeunit.ConfigForAsrs1,
		//{"ASRS", "2"}: executeunit.ConfigForAsrs2,
		//{"LSLS", "1"}: executeunit.ConfigForLsls1,
		//{"LSLS", "2"}: executeunit.ConfigForLsls2,
		//{"LSRS", "1"}: executeunit.ConfigForLsrs1,
		//{"LSRS", "2"}: executeunit.ConfigForLsrs2,
		//{"RORS", "1"}: executeunit.ConfigForRors1,
		//{"RORS", "2"}: executeunit.ConfigForRors2,

		// Comparison
		//{"CMN", "1"}: executeunit.ConfigForCmn1,
		//{"CMN", "2"}: executeunit.ConfigForCms2,
		//{"CMP", "1"}: executeunit.ConfigForCmp1,
		//{"CMP", "2"}: executeunit.ConfigForCmp2,

		// Bypass
		{"MOVS", "1"}: executeunit.ConfigForMovs1,
		{"MOVS", "2"}: executeunit.ConfigForMovs2,
		//{"BEQ", "1"}: executeunit.ConfigForBeq1,
		//{"BEQ", "2"}: executeunit.ConfigForBeq2,
		//{"BNE", "1"}: executeunit.ConfigForBne1,
		//{"BNE", "2"}: executeunit.ConfigForBne2,
		//{"BLT", "1"}: executeunit.ConfigForBlt1,
		//{"BLT", "2"}: executeunit.ConfigForBlt2,
		//{"BL", "1"}: executeunit.ConfigForBl1,
		//{"BL", "2"}: executeunit.ConfigForBl2,
		//{"BX", "1"}: executeunit.ConfigForBx1,
		//{"BX", "2"}: executeunit.ConfigForBx2,

		// Load and Store
		{"LDR", "1"}: executeunit.ConfigForLdr1,
		{"LDR", "2"}: executeunit.ConfigForLdr2,
		{"STR", "1"}: executeunit.ConfigForStr1,
		{"STR", "2"}: executeunit.ConfigForStr2,

		// Nothing
		//{"NOP", "1"}: executeunit.ConfigForNop,
	}

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