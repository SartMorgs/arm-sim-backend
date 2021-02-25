package controller

import "arm/programcounter"

type FindUnit struct{
	programcounter.ProgramCounter
	instruction string
	programArea string
}

func (fu *FindUnit) SetInstruction(inst string){
	fu.instruction = inst
}

func (fu *FindUnit) SetProgramArea(prg_area string){
	fu.programArea = prg_area
}

func (fu *FindUnit) GetInstruction() string{
	return fu.instruction
}

func (fu *FindUnit) GetProgramArea() string{
	return fu.programArea
}