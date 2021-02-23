package controller


type DecodeUnit struct{
	instruction string
	opcode string
}

func (dc * DecodeUnit) SetInstruction(inst string){
	dc.instruction = inst
}

func (dc *DecodeUnit) GetOpcode() string{
	return dc.opcode
}

func (dc *DecodeUnit) SplitInstruction(){
	// Aqui falta
	instructionRune := []rune(dc.instruction)

	dc.opcode = string(instructionRune[0:6])
}