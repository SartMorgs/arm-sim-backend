package program

import(
	"strconv"

	"arm/memmory"
	"arm/registerbank"
	"arm/ula"
	"arm/programcounter"
	"arm/controller"
	"arm/interruption"
	"arm/io"
)

type Program struct{
	rom *memmory.CodeMemmory
	ram *memmory.DataMemmory
	deviceMemmory *memmory.DeviceMemmory

	registerBank *registerbank.RegisterBank

	ula *ula.Ula

	pc *programcounter.ProgramCounter

	controller *controller.Controller

	interruptionBank *interruption.InterruptionBank

	deviceList map[string]*io.Device

	behaviorInstruction map[[2]string]func()
}

func NewProgram(rom *memmory.CodeMemmory) *Program{
	program := new(Program)
	program.rom = rom
	program.ram = memmory.NewDataMemmory()
	program.deviceMemmory = memmory.NewDeviceMemmory()

	program.registerBank = registerbank.NewRegisterBank()

	program.ula = ula.NewUla()

	program.pc = programcounter.NewProgramCounter()

	program.controller = controller.NewController()

	program.interruptionBank = interruption.NewInterruptionBank()

	program.deviceList = make(map[string]*io.Device)
	program.deviceList["standard"] = io.NewDevice("standard", "0x0", "00000000000000000000000000000000")

	
	program.behaviorInstruction = map[[2]string]func(){
		// Arithmetic
		{"ADDS", "1"}: program.ExecuteAdds1,
		//{"ADDS", "2"}: program.ExecuteAdds2,
		{"SUBS", "1"}: program.ExecuteSubs1,
		//{"SUBS", "2"}: program.ExecuteSubs2,
		//{"MULS", "1"}: executeunit.ConfigForMuls1,
		//{"MULS", "2"}: executeunit.ConfigForMuls2,
		//{"ANDS", "1"}: program.ExecuteAnds1,
		//{"ANDS", "2"}: program.ExecuteAnds2,
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
		//{"MOVS", "1"}: program.ExecuteMovs1,
		//{"MOVS", "2"}: program.ExecuteMovs2,
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
		//{"LDR", "1"}: program.ExecuteLdr1,
		{"LDR", "2"}: program.ExecuteLdr2,
		//{"STR", "1"}: program.ExecuteStr1,
		{"STR", "2"}: program.ExecuteStr2,

		// Nothing
		//{"NOP", "1"}: executeunit.ConfigForNop,
	}
	

	return program
}

func (p *Program) GetCodeMemmory() *memmory.CodeMemmory{
	return p.rom
}

func (p *Program) GetDataMemmory() *memmory.DataMemmory{
	return p.ram
}

func (p *Program) GetDeviceMemmory() *memmory.DeviceMemmory{
	return p.deviceMemmory
}

func (p *Program) GetRegisterBank() *registerbank.RegisterBank{
	return p.registerBank
}

func (p *Program) GetUla() *ula.Ula{
	return p.ula
}

func (p *Program) GetProgramCounter() *programcounter.ProgramCounter{
	return p.pc
}

func (p *Program) GetController() *controller.Controller{
	return p.controller
}

func (p *Program) GetInterruptionBank() *interruption.InterruptionBank{
	return p.interruptionBank
}

func (p *Program) GetDevice(name string) *io.Device{
	return p.deviceList[name]
}

func (p *Program) ExecuteLdr2(){
	target_register := p.controller.GetExecuteUnit().GetTargetRegisterDec()
	address := p.controller.GetExecuteUnit().GetAddressInstructionDec()
	target_register_r := "R" + strconv.Itoa(int(target_register))
	p.registerBank.ChangeRegister(target_register_r, int(address))
}

func (p *Program) ExecuteAdds1(){
	target_register := p.controller.GetExecuteUnit().GetTargetRegisterDec()
	source_register1 := p.controller.GetExecuteUnit().GetSourceRegister1Dec()
	source_register2 := p.controller.GetExecuteUnit().GetSourceRegister2Dec()
	target_register_r := "R" + strconv.Itoa(int(target_register))
	source_register1_r := "R" + strconv.Itoa(int(source_register1))
	source_register2_r := "R" + strconv.Itoa(int(source_register2))
	value_reg1 := p.registerBank.GetRegisterBank()[source_register1_r].GetDecValue()
	value_reg2 := p.registerBank.GetRegisterBank()[source_register2_r].GetDecValue()
	p.ula.SetValue1(value_reg1)
	p.ula.SetValue2(value_reg2)
	result := p.ula.Add()
	p.registerBank.ChangeRegister(target_register_r, result)
}

func (p *Program) ExecuteSubs1(){
	target_register := p.controller.GetExecuteUnit().GetTargetRegisterDec()
	source_register1 := p.controller.GetExecuteUnit().GetSourceRegister1Dec()
	source_register2 := p.controller.GetExecuteUnit().GetSourceRegister2Dec()
	target_register_r := "R" + strconv.Itoa(int(target_register))
	source_register1_r := "R" + strconv.Itoa(int(source_register1))
	source_register2_r := "R" + strconv.Itoa(int(source_register2))
	value_reg1 := p.registerBank.GetRegisterBank()[source_register1_r].GetDecValue()
	value_reg2 := p.registerBank.GetRegisterBank()[source_register2_r].GetDecValue()
	p.ula.SetValue1(value_reg1)
	p.ula.SetValue2(value_reg2)
	result := p.ula.Sub()
	p.registerBank.ChangeRegister(target_register_r, result)
}

func (p *Program) ExecuteStr2(){
	target_register := p.controller.GetExecuteUnit().GetTargetRegisterDec()
	address := p.controller.GetExecuteUnit().GetAddressInstructionDec()
	target_register_r := "R" + strconv.Itoa(int(target_register))
	value_register := p.registerBank.GetRegisterBank()[target_register_r].GetDecValue()
	address_hex := "0x" + strconv.FormatInt(int64(address), 16)
	p.ram.ChangeField(address_hex, value_register)
}

func (p *Program) ExecuteFunctionForInstruction(inst string, type_inst string){
	x := [2]string{inst, type_inst}
	p.behaviorInstruction[x]()
}