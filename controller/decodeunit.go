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

func (dc *DecodeUnit) SplitInstruction(){
	
	instructionRune := []rune(dc.instructionCode)

	// Get Opcode
	dc.opcode = string(instructionRune[0:6])

	// Instance map
	dc.instructionFormat = make(map[string]string)

	// Verify what's type of instruction
	if dc.instructionType == "Arithmetic"{
		dc.instructionFormat["reg1"] = string(instructionRune[6:11])
		dc.instructionFormat["reg2"] = string(instructionRune[11:16])
		dc.instructionFormat["regm"] = string(instructionRune[16:21])
	} else if dc.instructionType == "Comparison"{
		dc.instructionFormat["regn"] = string(instructionRune[6:11])
		dc.instructionFormat["regm"] = string(instructionRune[11:16])
	} else if dc.instructionType == "Bypass"{
		dc.instructionFormat["regn"] = string(instructionRune[6:11])
	} else{
		dc.instructionFormat["instruction"] = dc.instructionCode
	}
}
