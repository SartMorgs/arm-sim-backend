package controller

type FindUnit struct{
	memmoryAddress int
	instruction string
	programArea string
}

func (fu *FindUnit) SetInstruction(inst string){
	fu.instruction = inst
}

func (fu *FindUnit) SetProgramArea(prg_area string){
	fu.programArea = prg_area
}

func (fu *FindUnit) SetMemmoryAddress(mem_addr int){
	fu.memmoryAddress = mem_addr
}

func (fu *FindUnit) GetInstruction() string{
	return fu.instruction
}

func (fu *FindUnit) GetProgramArea() string{
	return fu.programArea
}

func (fu *FindUnit) GetMemmoryAddress() int{
	return fu.memmoryAddress
}