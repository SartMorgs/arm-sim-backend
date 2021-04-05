package program

import(
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