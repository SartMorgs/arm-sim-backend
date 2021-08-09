package program

import (
	"strconv"
	//"fmt"

	"arm/controller"
	"arm/interruption"
	"arm/io"
	"arm/memmory"
	"arm/programcounter"
	"arm/registerbank"
	"arm/ula"
)

type Program struct {
	rom           *memmory.CodeMemmory
	ram           *memmory.DataMemmory
	deviceMemmory *memmory.DeviceMemmory

	registerBank *registerbank.RegisterBank

	ula *ula.Ula

	pc *programcounter.ProgramCounter

	controller *controller.Controller

	interruptionBank *interruption.InterruptionBank

	deviceList map[string]*io.Device

	behaviorInstruction map[[2]string]func()
}

func NewProgram(rom *memmory.CodeMemmory) *Program {
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
		{"ADDS", "2"}: program.ExecuteAdds2,
		{"SUBS", "1"}: program.ExecuteSubs1,
		{"SUBS", "2"}: program.ExecuteSubs2,
		{"MULS", "1"}: program.ExecuteMuls1,
		{"MULS", "2"}: program.ExecuteMuls2,
		{"ANDS", "1"}: program.ExecuteAnds1,
		{"ANDS", "2"}: program.ExecuteAnds2,
		{"ORRS", "1"}: program.ExecuteOrrs1,
		{"ORRS", "2"}: program.ExecuteOrrs2,
		{"EORS", "1"}: program.ExecuteEors1,
		{"EORS", "2"}: program.ExecuteEors2,
		{"BICS", "1"}: program.ExecuteBics1,
		{"BICS", "2"}: program.ExecuteBics2,
		{"ASRS", "1"}: program.ExecuteAsrs1,
		{"ASRS", "2"}: program.ExecuteAsrs2,
		{"LSLS", "1"}: program.ExecuteLsls1,
		{"LSLS", "2"}: program.ExecuteLsls2,
		{"LSRS", "1"}: program.ExecuteLsrs1,
		{"LSRS", "2"}: program.ExecuteLsrs2,
		{"RORS", "1"}: program.ExecuteRors1,
		{"RORS", "2"}: program.ExecuteRors2,

		// Comparison
		//{"CMN", "1"}: program.ExecuteCmn1,
		//{"CMN", "2"}: program.ExecuteCms2,
		//{"CMP", "1"}: program.ExecuteCmp1,
		//{"CMP", "2"}: program.ExecuteCmp2,

		// Bypass
		//{"MOVS", "1"}: program.ExecuteMovs1,
		//{"MOVS", "2"}: program.ExecuteMovs2,
		//{"BEQ", "1"}: program.ExecuteBeq1,
		//{"BEQ", "2"}: program.ExecuteBeq2,
		//{"BNE", "1"}: program.ExecuteBne1,
		//{"BNE", "2"}: program.ExecuteBne2,
		//{"BLT", "1"}: program.ExecuteBlt1,
		//{"BLT", "2"}: program.ExecuteBlt2,
		//{"BL", "1"}: program.ExecuteBl1,
		//{"BL", "2"}: program.ExecuteBl2,
		//{"BX", "1"}: program.ExecuteBx1,
		//{"BX", "2"}: program.ExecuteBx2,
		//{"B", "1"}: program.ExecuteB1,
		//{"B", "2"}: program.ExecuteB2,

		// Load and Store
		{"LDR", "1"}: program.ExecuteLdr1,
		{"LDR", "2"}: program.ExecuteLdr2,
		{"STR", "1"}: program.ExecuteStr1,
		{"STR", "2"}: program.ExecuteStr2,

		// Nothing
		{"NOP", "1"}: program.ExecuteNop,
	}

	return program
}

func (p *Program) initialConfigurations() {
	p.registerBank.GetRegisterBank()["R13"].SetRegisterName("SP")
	p.registerBank.GetRegisterBank()["R14"].SetRegisterName("LR")
	p.registerBank.GetRegisterBank()["R15"].SetRegisterName("PC")
	p.registerBank.GetRegisterBank()["R16"].SetRegisterName("PSR")
}

func (p *Program) setPSRRegister() {

	register_form := make(map[bool]string)
	register_form = map[bool]string{
		true:  "1",
		false: "0"}

	register_string := register_form[p.ula.GetNegativeResultFlag()] +
		register_form[p.ula.GetZeroResultFlag()] +
		register_form[p.ula.GetCarryResultFlag()] +
		register_form[p.ula.GetOverflowResultFlag()] +
		"0000000000000000000000000000"

	register_int, _ := strconv.ParseInt(register_string, 2, 64)
	p.registerBank.ChangeRegister("R16", int(register_int))
}

func (p *Program) GetCodeMemmory() *memmory.CodeMemmory {
	return p.rom
}

func (p *Program) GetDataMemmory() *memmory.DataMemmory {
	return p.ram
}

func (p *Program) GetDeviceMemmory() *memmory.DeviceMemmory {
	return p.deviceMemmory
}

func (p *Program) GetRegisterBank() *registerbank.RegisterBank {
	return p.registerBank
}

func (p *Program) GetUla() *ula.Ula {
	return p.ula
}

func (p *Program) GetProgramCounter() *programcounter.ProgramCounter {
	return p.pc
}

func (p *Program) GetController() *controller.Controller {
	return p.controller
}

func (p *Program) GetInterruptionBank() *interruption.InterruptionBank {
	return p.interruptionBank
}

func (p *Program) GetDevice(name string) *io.Device {
	return p.deviceList[name]
}

func (p *Program) ExecuteLdr1() {
	target_register := p.controller.GetExecuteUnit().GetTargetRegisterDec()
	source_register1 := p.controller.GetExecuteUnit().GetSourceRegister1Dec()
	target_register_r := "R" + strconv.Itoa(int(target_register))
	source_register1_r := "R" + strconv.Itoa(int(source_register1))
	memmory_addr := p.registerBank.GetRegisterBank()[source_register1_r].GetHexValue()
	value := p.ram.GetDataMemmoryList()[memmory_addr].GetDecValue()
	p.registerBank.ChangeRegister(target_register_r, value)
}

func (p *Program) ExecuteLdr2() {
	target_register := p.controller.GetExecuteUnit().GetTargetRegisterDec()
	address := p.controller.GetExecuteUnit().GetAddressInstructionDec()
	target_register_r := "R" + strconv.Itoa(int(target_register))
	p.registerBank.ChangeRegister(target_register_r, int(address))
}

func (p *Program) ExecuteAdds1() {
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

func (p *Program) ExecuteAdds2() {
	target_register := p.controller.GetExecuteUnit().GetTargetRegisterDec()
	source_register1 := p.controller.GetExecuteUnit().GetSourceRegister1Dec()
	data := p.controller.GetExecuteUnit().GetValueInstructionDec()
	target_register_r := "R" + strconv.Itoa(int(target_register))
	source_register1_r := "R" + strconv.Itoa(int(source_register1))
	value_reg1 := p.registerBank.GetRegisterBank()[source_register1_r].GetDecValue()
	p.ula.SetValue1(value_reg1)
	p.ula.SetValue2(int(data))
	result := p.ula.Add()
	p.registerBank.ChangeRegister(target_register_r, result)
}

func (p *Program) ExecuteSubs1() {
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

func (p *Program) ExecuteSubs2() {
	target_register := p.controller.GetExecuteUnit().GetTargetRegisterDec()
	source_register1 := p.controller.GetExecuteUnit().GetSourceRegister1Dec()
	data := p.controller.GetExecuteUnit().GetValueInstructionDec()
	target_register_r := "R" + strconv.Itoa(int(target_register))
	source_register1_r := "R" + strconv.Itoa(int(source_register1))
	value_reg1 := p.registerBank.GetRegisterBank()[source_register1_r].GetDecValue()
	p.ula.SetValue1(value_reg1)
	p.ula.SetValue2(int(data))
	result := p.ula.Sub()
	p.registerBank.ChangeRegister(target_register_r, result)
}

func (p *Program) ExecuteMuls1() {
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
	result := p.ula.Mult()
	p.registerBank.ChangeRegister(target_register_r, result)
}

func (p *Program) ExecuteMuls2() {
	target_register := p.controller.GetExecuteUnit().GetTargetRegisterDec()
	source_register1 := p.controller.GetExecuteUnit().GetSourceRegister1Dec()
	data := p.controller.GetExecuteUnit().GetValueInstructionDec()
	target_register_r := "R" + strconv.Itoa(int(target_register))
	source_register1_r := "R" + strconv.Itoa(int(source_register1))
	value_reg1 := p.registerBank.GetRegisterBank()[source_register1_r].GetDecValue()
	p.ula.SetValue1(value_reg1)
	p.ula.SetValue2(int(data))
	result := p.ula.Mult()
	p.registerBank.ChangeRegister(target_register_r, result)
}

func (p *Program) ExecuteAnds1() {
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
	result := p.ula.And()
	p.registerBank.ChangeRegister(target_register_r, result)
}

func (p *Program) ExecuteAnds2() {
	target_register := p.controller.GetExecuteUnit().GetTargetRegisterDec()
	source_register1 := p.controller.GetExecuteUnit().GetSourceRegister1Dec()
	data := p.controller.GetExecuteUnit().GetValueInstructionDec()
	target_register_r := "R" + strconv.Itoa(int(target_register))
	source_register1_r := "R" + strconv.Itoa(int(source_register1))
	value_reg1 := p.registerBank.GetRegisterBank()[source_register1_r].GetDecValue()
	p.ula.SetValue1(value_reg1)
	p.ula.SetValue2(int(data))
	result := p.ula.And()
	p.registerBank.ChangeRegister(target_register_r, result)
}

func (p *Program) ExecuteOrrs1() {
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
	result := p.ula.Or()
	p.registerBank.ChangeRegister(target_register_r, result)
}

func (p *Program) ExecuteOrrs2() {
	target_register := p.controller.GetExecuteUnit().GetTargetRegisterDec()
	source_register1 := p.controller.GetExecuteUnit().GetSourceRegister1Dec()
	data := p.controller.GetExecuteUnit().GetValueInstructionDec()
	target_register_r := "R" + strconv.Itoa(int(target_register))
	source_register1_r := "R" + strconv.Itoa(int(source_register1))
	value_reg1 := p.registerBank.GetRegisterBank()[source_register1_r].GetDecValue()
	p.ula.SetValue1(value_reg1)
	p.ula.SetValue2(int(data))
	result := p.ula.Or()
	p.registerBank.ChangeRegister(target_register_r, result)
}

func (p *Program) ExecuteEors1() {
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
	result := p.ula.Eor()
	p.registerBank.ChangeRegister(target_register_r, result)
}

func (p *Program) ExecuteEors2() {
	target_register := p.controller.GetExecuteUnit().GetTargetRegisterDec()
	source_register1 := p.controller.GetExecuteUnit().GetSourceRegister1Dec()
	data := p.controller.GetExecuteUnit().GetValueInstructionDec()
	target_register_r := "R" + strconv.Itoa(int(target_register))
	source_register1_r := "R" + strconv.Itoa(int(source_register1))
	value_reg1 := p.registerBank.GetRegisterBank()[source_register1_r].GetDecValue()
	p.ula.SetValue1(value_reg1)
	p.ula.SetValue2(int(data))
	result := p.ula.Eor()
	p.registerBank.ChangeRegister(target_register_r, result)
}

func (p *Program) ExecuteBics1() {
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
	result := p.ula.Bic()
	p.registerBank.ChangeRegister(target_register_r, result)
}

func (p *Program) ExecuteBics2() {
	target_register := p.controller.GetExecuteUnit().GetTargetRegisterDec()
	source_register1 := p.controller.GetExecuteUnit().GetSourceRegister1Dec()
	data := p.controller.GetExecuteUnit().GetValueInstructionDec()
	target_register_r := "R" + strconv.Itoa(int(target_register))
	source_register1_r := "R" + strconv.Itoa(int(source_register1))
	value_reg1 := p.registerBank.GetRegisterBank()[source_register1_r].GetDecValue()
	p.ula.SetValue1(value_reg1)
	p.ula.SetValue2(int(data))
	result := p.ula.Bic()
	p.registerBank.ChangeRegister(target_register_r, result)
}

func (p *Program) ExecuteAsrs1() {
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
	result := p.ula.Asr()
	p.registerBank.ChangeRegister(target_register_r, result)
}

func (p *Program) ExecuteAsrs2() {
	target_register := p.controller.GetExecuteUnit().GetTargetRegisterDec()
	source_register1 := p.controller.GetExecuteUnit().GetSourceRegister1Dec()
	data := p.controller.GetExecuteUnit().GetValueInstructionDec()
	target_register_r := "R" + strconv.Itoa(int(target_register))
	source_register1_r := "R" + strconv.Itoa(int(source_register1))
	value_reg1 := p.registerBank.GetRegisterBank()[source_register1_r].GetDecValue()
	p.ula.SetValue1(value_reg1)
	p.ula.SetValue2(int(data))
	result := p.ula.Asr()
	p.registerBank.ChangeRegister(target_register_r, result)
}

func (p *Program) ExecuteLsls1() {
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
	result := p.ula.Lsl()
	p.registerBank.ChangeRegister(target_register_r, result)
}

func (p *Program) ExecuteLsls2() {
	target_register := p.controller.GetExecuteUnit().GetTargetRegisterDec()
	source_register1 := p.controller.GetExecuteUnit().GetSourceRegister1Dec()
	data := p.controller.GetExecuteUnit().GetValueInstructionDec()
	target_register_r := "R" + strconv.Itoa(int(target_register))
	source_register1_r := "R" + strconv.Itoa(int(source_register1))
	value_reg1 := p.registerBank.GetRegisterBank()[source_register1_r].GetDecValue()
	p.ula.SetValue1(value_reg1)
	p.ula.SetValue2(int(data))
	result := p.ula.Lsl()
	p.registerBank.ChangeRegister(target_register_r, result)
}

func (p *Program) ExecuteLsrs1() {
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
	result := p.ula.Lsr()
	p.registerBank.ChangeRegister(target_register_r, result)
}

func (p *Program) ExecuteLsrs2() {
	target_register := p.controller.GetExecuteUnit().GetTargetRegisterDec()
	source_register1 := p.controller.GetExecuteUnit().GetSourceRegister1Dec()
	data := p.controller.GetExecuteUnit().GetValueInstructionDec()
	target_register_r := "R" + strconv.Itoa(int(target_register))
	source_register1_r := "R" + strconv.Itoa(int(source_register1))
	value_reg1 := p.registerBank.GetRegisterBank()[source_register1_r].GetDecValue()
	p.ula.SetValue1(value_reg1)
	p.ula.SetValue2(int(data))
	result := p.ula.Lsr()
	p.registerBank.ChangeRegister(target_register_r, result)
}

func (p *Program) ExecuteRors1() {
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
	result := p.ula.Ror()
	p.registerBank.ChangeRegister(target_register_r, result)
}

func (p *Program) ExecuteRors2() {
	target_register := p.controller.GetExecuteUnit().GetTargetRegisterDec()
	source_register1 := p.controller.GetExecuteUnit().GetSourceRegister1Dec()
	data := p.controller.GetExecuteUnit().GetValueInstructionDec()
	target_register_r := "R" + strconv.Itoa(int(target_register))
	source_register1_r := "R" + strconv.Itoa(int(source_register1))
	value_reg1 := p.registerBank.GetRegisterBank()[source_register1_r].GetDecValue()
	p.ula.SetValue1(value_reg1)
	p.ula.SetValue2(int(data))
	result := p.ula.Ror()
	p.registerBank.ChangeRegister(target_register_r, result)
}

func (p *Program) ExecuteStr1() {
	target_register := p.controller.GetExecuteUnit().GetTargetRegisterDec()
	source_register1 := p.controller.GetExecuteUnit().GetSourceRegister1Dec()
	target_register_r := "R" + strconv.Itoa(int(target_register))
	source_register1_r := "R" + strconv.Itoa(int(source_register1))
	memmory_addr := p.registerBank.GetRegisterBank()[source_register1_r].GetHexValue()
	value := p.registerBank.GetRegisterBank()[target_register_r].GetDecValue()
	p.ram.ChangeField(memmory_addr, value)
}

func (p *Program) ExecuteStr2() {
	target_register := p.controller.GetExecuteUnit().GetTargetRegisterDec()
	address := p.controller.GetExecuteUnit().GetAddressInstructionDec()
	target_register_r := "R" + strconv.Itoa(int(target_register))
	value_register := p.registerBank.GetRegisterBank()[target_register_r].GetDecValue()
	address_hex := "0x" + strconv.FormatInt(int64(address), 16)
	p.ram.ChangeField(address_hex, value_register)
}

func (p *Program) ExecuteFunctionForInstruction(inst string, type_inst string) {
	x := [2]string{inst, type_inst}
	p.behaviorInstruction[x]()
}

func (p *Program) ExecuteNop() {

}
