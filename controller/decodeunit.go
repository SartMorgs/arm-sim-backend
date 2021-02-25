package controller

type Instruction struct{
	instructionCode string
	instructionType string
	instructionFormat map[string]string
}

type DecodeUnit struct{
	Instruction
	opcode string
	instructionMap map[string]string
}

func (it *Instruction) SetInstruction(inst string){
	it.instructionCode = inst
}

func (dc *DecodeUnit) GetOpcode() string{
	return dc.opcode
}

func (dc *DecodeUnit) MapInstruction() string{
	instructionRune := []rune(dc.instructionCode)
	dc.opcode = string(instructionRune[0:6])

	return dc.instructionMap[dc.opcode]
}

/*
func (dc *DecodeUnit) SplitInstruction(){
	
	instructionRune := []rune(dc.instructionCode)

	// Get Opcode
	dc.opcode = string(instructionRune[0:6])

	if dc.instructionType == "Arithmetic"{
		dc.instructionFormat["reg1"] = string(instructionRune[6:5])
		dc.instructionFormat["reg2"] = string(instructionRune[11:5])
		dc.instructionFormat["regm"] = string
	} else if dc.instructionType == "Comparison"{

	} else if dc.instructionType == "Bypass"{

	} else{

	}
}
*/