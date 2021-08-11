package controller

type DecodeUnit struct {
	InstructionMap

	instructionCode   string
	instructionType1  string
	instructionType2  string
	instructionName   string
	instructionFormat map[string]string
	opcode            string
	gapAddress        bool
}

func NewDecodeUnit() *DecodeUnit {
	decodeunit := new(DecodeUnit)

	decodeunit.instructionMap = map[string][3]string{
		// Arithmetic
		"000001": {"ADDS", "Arithmetic", "1"},
		"100001": {"ADDS", "Arithmetic", "2"},
		"000010": {"SUBS", "Arithmetic", "1"},
		"100010": {"SUBS", "Arithmetic", "2"},
		"000011": {"MULS", "Arithmetic", "1"},
		"100011": {"MULS", "Arithmetic", "2"},
		"000100": {"ANDS", "Arithmetic", "1"},
		"100100": {"ANDS", "Arithmetic", "2"},
		"000101": {"ORRS", "Arithmetic", "1"},
		"100101": {"ORRS", "Arithmetic", "2"},
		"000110": {"EORS", "Arithmetic", "1"},
		"100110": {"EORS", "Arithmetic", "2"},
		"000111": {"BICS", "Arithmetic", "1"},
		"100111": {"BICS", "Arithmetic", "2"},
		"001000": {"ASRS", "Arithmetic", "1"},
		"101000": {"ASRS", "Arithmetic", "2"},
		"001001": {"LSLS", "Arithmetic", "1"},
		"101001": {"LSLS", "Arithmetic", "2"},
		"001010": {"LSRS", "Arithmetic", "1"},
		"101010": {"LSRS", "Arithmetic", "2"},
		"001011": {"RORS", "Arithmetic", "1"},
		"101011": {"RORS", "Arithmetic", "2"},

		// Comparison
		"001100": {"CMN", "Comparison", "1"},
		"101100": {"CMN", "Comparison", "2"},
		"001101": {"CMP", "Comparison", "1"},
		"101101": {"CMP", "Comparison", "2"},

		// Bypass
		"001110": {"MOVS", "Bypass", "1"},
		"101110": {"MOVS", "Bypass", "2"},
		"001111": {"BEQ", "Bypass", "1"},
		"101111": {"BEQ", "Bypass", "2"},
		"010000": {"BNE", "Bypass", "1"},
		"110000": {"BNE", "Bypass", "2"},
		"010001": {"BLT", "Bypass", "1"},
		"110001": {"BLT", "Bypass", "2"},
		"010010": {"BL", "Bypass", "1"},
		"110010": {"BL", "Bypass", "2"},
		"010011": {"BX", "Bypass", "1"},
		//"110011": {"BX", "Bypass", "2"}, PERGUNTAR PARA PROFESSOR
		"011011": {"B", "Bypass", "1"},
		"111011": {"B", "Bypass", "2"},

		// Load and Store
		"010100": {"LDR", "Load", "1"},
		"110100": {"LDR", "Load", "2"},
		"010101": {"STR", "Store", "1"},
		"110101": {"STR", "Store", "2"},

		// Nothing
		"000000": {"NOP", "Nothing", "1"},
	}

	decodeunit.gapAddress = false
	decodeunit.instructionCode = "00000000000000000000000000000000"
	decodeunit.instructionType1 = ""
	decodeunit.instructionType2 = ""
	decodeunit.instructionFormat = make(map[string]string)

	return decodeunit
}

func (it *DecodeUnit) SetInstruction(inst string) {
	it.instructionCode = inst
}

func (dc *DecodeUnit) GetOpcode() string {
	return dc.opcode
}

func (dc *DecodeUnit) getOpcode(inst string) string {
	instructionRune := []rune(dc.instructionCode)
	opcode := string(instructionRune[0:6])
	return opcode
}

func (dc *DecodeUnit) GetGapAddressFlag() bool {
	return dc.gapAddress
}

func (dc *DecodeUnit) MapInstruction() {
	dc.opcode = dc.getOpcode(dc.instructionCode)

	dc.instructionName = dc.instructionMap[dc.opcode][0]
	dc.instructionType1 = dc.instructionMap[dc.opcode][1]
	dc.instructionType2 = dc.instructionMap[dc.opcode][2]
}

func (dc *DecodeUnit) SplitInstruction() {

	instructionRune := []rune(dc.instructionCode)

	dc.MapInstruction()

	// Instance map
	dc.instructionFormat = make(map[string]string)

	// Verify what's type of instruction
	if dc.instructionType1 == "Arithmetic" {
		dc.gapAddress = false
		if dc.instructionType2 == "1" {
			dc.instructionFormat["rd"] = string(instructionRune[6:11])
			dc.instructionFormat["rn"] = string(instructionRune[11:16])
			dc.instructionFormat["rm"] = string(instructionRune[16:21])
		} else {
			dc.instructionFormat["rd"] = string(instructionRune[6:11])
			dc.instructionFormat["rn"] = string(instructionRune[11:16])
			dc.instructionFormat["data"] = string(instructionRune[16:32])
		}
	} else if dc.instructionType1 == "Comparison" {
		dc.gapAddress = false
		if dc.instructionType2 == "1" {
			dc.instructionFormat["rn"] = string(instructionRune[6:11])
			dc.instructionFormat["rm"] = string(instructionRune[11:16])
		} else {
			dc.instructionFormat["rn"] = string(instructionRune[6:11])
			dc.instructionFormat["imediato"] = string(instructionRune[11:19])
		}
	} else if dc.instructionType1 == "Bypass" {
		dc.gapAddress = true
		if dc.instructionType2 == "1" {
			dc.instructionFormat["rn"] = string(instructionRune[6:11])
		} else {
			dc.instructionFormat["label"] = string(instructionRune[6:32])
		}
	} else if dc.instructionType1 == "Load" || dc.instructionType1 == "Store" {
		dc.gapAddress = false
		if dc.instructionType2 == "1" {
			dc.instructionFormat["rd"] = string(instructionRune[6:11])
			dc.instructionFormat["rn"] = string(instructionRune[11:16])
		} else {
			dc.instructionFormat["rd"] = string(instructionRune[6:11])
			dc.instructionFormat["address"] = string(instructionRune[11:25])
		}
	} else if dc.instructionType1 == "Nop" {
		dc.gapAddress = false
		dc.instructionFormat["rest"] = string(instructionRune[6:32])
	} else {
		dc.gapAddress = false
		dc.instructionFormat["rest"] = dc.instructionCode
	}
}

func (dc *DecodeUnit) GetInstruction() string {
	return dc.instructionCode
}

func (dc *DecodeUnit) GetInstructionFormat() map[string]string {
	return dc.instructionFormat
}

func (dc *DecodeUnit) GetInstructionName() string {
	return dc.instructionName
}

func (dc *DecodeUnit) GetInstructionType2() string {
	return dc.instructionType2
}
