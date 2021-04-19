package controller

import "strconv"

type ConfigForInstruction struct{
	targetRegisterBin string
	targetRegisterDec int64

	addressRegisterBin string
	addressRegisterDec int64

	addressInstructionBin string
	addressInstructionDec int64

	sourceRegister1Bin string 
	sourceRegister1Dec int64

	sourceRegister2Bin string 
	sourceRegister2Dec int64

	valueInstructionBin string
	valueInstructionDec int64

	valueRegisterBin string
	valueRegisterDec int64
}

type ExecuteUnit struct{
	instruction string
	ConfigForInstruction
	config map[string]int64
	functionInstructionMap map[[2]string]func()
	instructionFormatMap map[string]func(string)
}

func NewExecuteUnit() *ExecuteUnit{
	executeunit := new(ExecuteUnit)
	executeunit.config = make(map[string]int64)
	
	executeunit.functionInstructionMap = map[[2]string]func(){
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

	executeunit.instructionFormatMap = map[string]func(string){
		"rd": executeunit.SetTargetRegister,
		"rn": executeunit.SetSourceRegister1,
		"rm": executeunit.SetSourceRegister2,
		"data": executeunit.SetValueInstruction,
		"imediato": executeunit.SetAddressInstruction,
		"label": executeunit.SetValueInstruction,
		"address": executeunit.SetAddressInstruction,
	}

	return executeunit
}

func (eu *ExecuteUnit) SetInstruction(inst string){
	eu.instruction = inst
}

func (eu *ExecuteUnit) SetTargetRegister(target_reg string){
	eu.targetRegisterBin = target_reg
	eu.targetRegisterDec, _ = strconv.ParseInt(eu.targetRegisterBin, 2, 64)
}

func (eu *ExecuteUnit) SetAddressInstruction(addr string){
	eu.addressInstructionBin = addr
	eu.addressInstructionDec, _ = strconv.ParseInt(eu.addressInstructionBin, 2, 64)
}

func (eu *ExecuteUnit) SetSourceRegister1(source_reg1 string){
	eu.sourceRegister1Bin = source_reg1
	eu.sourceRegister1Dec, _ = strconv.ParseInt(eu.sourceRegister1Bin, 2, 64)
}

func (eu *ExecuteUnit) SetSourceRegister2(source_reg2 string){
	eu.sourceRegister2Bin = source_reg2
	eu.sourceRegister2Dec, _ = strconv.ParseInt(eu.sourceRegister2Bin, 2, 64)
}

func (eu *ExecuteUnit) SetValueInstruction(value string){
	eu.valueInstructionBin = value
	eu.valueInstructionDec, _ = strconv.ParseInt(eu.valueInstructionBin, 2, 64)
}

func (eu *ExecuteUnit) SetValueRegister(value_reg string){
	eu.valueRegisterBin = value_reg
	eu.valueRegisterDec, _ = strconv.ParseInt(eu.valueRegisterBin, 2, 64)
}

func (eu *ExecuteUnit) GetTargetRegisterDec()int64{
	return eu.targetRegisterDec
}

func (eu *ExecuteUnit) GetAddressInstructionDec()int64{
	return eu.addressInstructionDec
}

func (eu *ExecuteUnit) GetSourceRegister1Dec()int64{
	return eu.sourceRegister1Dec
}

func (eu *ExecuteUnit) GetSourceRegister2Dec()int64{
	return eu.sourceRegister2Dec
}

func (eu *ExecuteUnit) GetValueInstructionDec()int64{
	return eu.valueInstructionDec
}

func (eu *ExecuteUnit) SetValueRegisterDec()int64{
	return eu.valueRegisterDec
}

func (eu *ExecuteUnit) GetInstruction() string{
	return eu.instruction
}

func (eu *ExecuteUnit) ConfigForLdr1(){
	eu.config["target_register"] = eu.targetRegisterDec
	eu.config["source_register1"] = eu.sourceRegister1Dec
}


func (eu *ExecuteUnit) ConfigForLdr2(){
	eu.config["target_register"] = eu.targetRegisterDec
	eu.config["address"] = eu.addressInstructionDec
}

func (eu *ExecuteUnit) ConfigForStr1(){
	eu.config["target_register"] = eu.targetRegisterDec
	eu.config["source_register1"] = eu.sourceRegister1Dec
}

func (eu *ExecuteUnit) ConfigForStr2(){
	eu.config["target_register"] = eu.targetRegisterDec
	eu.config["address"] = eu.addressInstructionDec
}

func (eu *ExecuteUnit) ConfigForAdds1(){
	eu.config["target_register"] = eu.targetRegisterDec
	eu.config["source_register1"] = eu.sourceRegister1Dec
	eu.config["source_register2"] = eu.sourceRegister2Dec
}

func (eu *ExecuteUnit) ConfigForAdds2(){
	eu.config["target_register"] = eu.targetRegisterDec
	eu.config["source_register1"] = eu.sourceRegister1Dec
	eu.config["value"] = eu.valueInstructionDec
}

func (eu *ExecuteUnit) ConfigForSubs1(){
	eu.config["target_register"] = eu.targetRegisterDec
	eu.config["source_register1"] = eu.sourceRegister1Dec
	eu.config["source_register2"] = eu.sourceRegister2Dec
}

func (eu *ExecuteUnit) ConfigForSubs2(){
	eu.config["target_register"] = eu.targetRegisterDec
	eu.config["source_register1"] = eu.sourceRegister1Dec
	eu.config["value"] =  eu.valueInstructionDec
}

func (eu *ExecuteUnit) ConfigForAnds1(){
	eu.config["target_register"] = eu.targetRegisterDec
	eu.config["source_register1"] = eu.sourceRegister1Dec
	eu.config["source_register2"] = eu.sourceRegister2Dec
}

func (eu *ExecuteUnit) ConfigForAnds2(){
	eu.config["target_register"] = eu.targetRegisterDec
	eu.config["source_register1"] = eu.sourceRegister1Dec
	eu.config["value"] = eu.valueInstructionDec
}

func (eu *ExecuteUnit) ConfigForMovs1(){
	eu.config["target_register"] = eu.targetRegisterDec
	eu.config["value_register"] = eu.valueRegisterDec
}

func (eu *ExecuteUnit) ConfigForMovs2(){
	eu.config["target_register"] = eu.targetRegisterDec
	eu.config["value"] = eu.valueInstructionDec
}

func (eu *ExecuteUnit) ConfigInstruction(inst string, type_inst string){
	x := [2]string{inst, type_inst}
	eu.functionInstructionMap[x]()
}

func (eu *ExecuteUnit) SetFunctionForConfigFormat(format string, value string){
	eu.instructionFormatMap[format](value)
}