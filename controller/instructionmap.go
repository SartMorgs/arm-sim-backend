package controller

type InstructionMap struct{	
	instructionMap map[string]string
}

func (im *InstructionMap) InstructionMap(inst string) string{
	opcode := im.getOpcode(inst)

	return im.instructionMap[opcode]
}

func (im *InstructionMap) getOpcode(inst string) string{
	instructionRune := []rune(inst)
	return string(instructionRune[0:6])
}