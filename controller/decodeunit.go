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

func NewDecodeUnit() *DecodeUnit{
	decodeunit := new(DecodeUnit)
	decodeunit.instructionCode = ""
	decodeunit.instructionType = ""
	decodeunit.instructionFormat = make(map[string]string)

	return decodeunit
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
	if dc.instructionType == "Arithmetic1"{
		dc.instructionFormat["rd"] = string(instructionRune[6:11])
		dc.instructionFormat["rn"] = string(instructionRune[11:16])
		dc.instructionFormat["rm"] = string(instructionRune[16:21])
	} else if dc.instructionType == "Arithmetic2"{
		dc.instructionFormat["rd"] = string(instructionRune[6:11])
		dc.instructionFormat["rn"] = string(instructionRune[11:16])
	} else if dc.instructionType == "Comparison1"{
		dc.instructionFormat["rn"] = string(instructionRune[6:11])
		dc.instructionFormat["rm"] = string(instructionRune[11:16])
	} else if dc.instructionType == "Comparison2"{
		dc.instructionFormat["rn"] = string(instructionRune[6:11])
		dc.instructionFormat["imediato"] = string(instructionRune[11:19])
	} else if dc.instructionType == "Bypass1"{
		dc.instructionFormat["rn"] = string(instructionRune[6:11])
	} else if dc.instructionType == "Bypass2"{
		dc.instructionFormat["label"] = string(instructionRune[6:32])
	} else if dc.instructionType == "LoadStore1"{
		dc.instructionFormat["rd"] = string(instructionRune[6:11])
		dc.instructionFormat["rn"] = string(instructionRune[11:16])
	} else if dc.instructionType == "LoadStore2"{
		dc.instructionFormat["rn"] = string(instructionRune[6:11])
		dc.instructionFormat["address"] = string(instructionRune[11:25])
	} else if dc.instructionType == "Nop"{
		dc.instructionFormat["rest"] = ""
	} else{
		dc.instructionFormat["instruction"] = dc.instructionCode
	}
}
