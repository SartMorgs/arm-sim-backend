package controller

import "testing"

var imp InstructionMap

// Unit Tests
// Test all methods using mock

func SetInstructionMap(){
	// Instruction map
	imp.instructionMap = map[string]string{
		// Arithmetic
		"000001": "ADDS",
		"000010": "SUBS",
		"000011": "MULS",
		"000100": "ANDS",
		"000101": "ORRS",
		"000110": "EORS",
		"000111": "BICS",
		"001000": "ASRS",
		"001001": "LSLS",
		"001010": "LSRS",
		"001011": "RORS",

		// Comparison
		"001100": "CMN",
		"001101": "CMP",

		// Bypass
		"001110": "MOVS",
		"001111": "BEQ",
		"010000": "BNE",
		"010001": "BLT",
		"010010": "BL",
		"010011": "BX",

		// Load and Store
		"010100": "LDR",
		"010101": "STR",

		// Nothing
		"111111": "NOP",
	}
}

// Instruction Adds
func TestMapInstructionAdds(t *testing.T){
	SetInstructionMap()

	want := "ADDS"
	got := imp.InstructionMap("00000101010100001010100001011010")

	if got != want{
		t.Errorf("MapInstructionAdds \n got: %v \n want %v \n", got, want)
	}
}

func TestMapInstructionSubs(t *testing.T){
	SetInstructionMap()

	want := "SUBS"
	got := imp.InstructionMap("00001001010100001010100001011010")

	if got != want{
		t.Errorf("MapInstructionSubs \n got: %v \n want %v \n", got, want)
	}
}

func TestMapInstructionMuls(t *testing.T){
	SetInstructionMap()

	want := "MULS"
	got := imp.InstructionMap("00001101010100001010100001011010")

	if got != want{
		t.Errorf("MapInstructionMuls \n got: %v \n want %v \n", got, want)
	}
}

func TestMapInstructionAnds(t *testing.T){
	SetInstructionMap()

	want := "ANDS"
	got := imp.InstructionMap("00010001010100001010100001011010")

	if got != want{
		t.Errorf("MapInstructionAnds \n got: %v \n want %v \n", got, want)
	}
}

func TestMapInstructionOrrs(t *testing.T){
	SetInstructionMap()

	want := "ORRS"
	got := imp.InstructionMap("00010101010100001010100001011010")

	if got != want{
		t.Errorf("MapInstructionOrrs \n got: %v \n want %v \n", got, want)
	}
}

func TestMapInstructionEors(t *testing.T){
	SetInstructionMap()

	want := "EORS"
	got := imp.InstructionMap("00011001010100001010100001011010")

	if got != want{
		t.Errorf("MapInstructionEors \n got: %v \n want %v \n", got, want)
	}
}

func TestMapInstructionBics(t *testing.T){
	SetInstructionMap()

	want := "BICS"
	got := imp.InstructionMap("00011101010100001010100001011010")

	if got != want{
		t.Errorf("MapInstructionBics \n got: %v \n want %v \n", got, want)
	}
}

func TestMapInstructionAsrs(t *testing.T){
	SetInstructionMap()

	want := "ASRS"
	got := imp.InstructionMap("00100001010100001010100001011010")

	if got != want{
		t.Errorf("MapInstructionAsrs \n got: %v \n want %v \n", got, want)
	}
}

func TestMapInstructionLsls(t *testing.T){
	SetInstructionMap()

	want := "LSLS"
	got := imp.InstructionMap("00100101010100001010100001011010")

	if got != want{
		t.Errorf("MapInstructionLsls \n got: %v \n want %v \n", got, want)
	}
}

func TestMapInstructionLsrs(t *testing.T){
	SetInstructionMap()

	want := "LSRS"
	got := imp.InstructionMap("00101001010100001010100001011010")

	if got != want{
		t.Errorf("MapInstructionLsrs \n got: %v \n want %v \n", got, want)
	}
}

func TestMapInstructionRors(t *testing.T){
	SetInstructionMap()

	want := "RORS"
	got := imp.InstructionMap("00101101010100001010100001011010")

	if got != want{
		t.Errorf("MapInstructionRors \n got: %v \n want %v \n", got, want)
	}
}

func TestMapInstructionCmn(t *testing.T){
	SetInstructionMap()

	want := "CMN"
	got := imp.InstructionMap("00110001010100001010100001011010")

	if got != want{
		t.Errorf("MapInstructionCmn \n got: %v \n want %v \n", got, want)
	}
}

func TestMapInstructionCmp(t *testing.T){
	SetInstructionMap()

	want := "CMP"
	got := imp.InstructionMap("00110101010100001010100001011010")

	if got != want{
		t.Errorf("MapInstructionCmp \n got: %v \n want %v \n", got, want)
	}
}

func TestMapInstructionMovs(t *testing.T){
	SetInstructionMap()

	want := "MOVS"
	got := imp.InstructionMap("00111001010100001010100001011010")

	if got != want{
		t.Errorf("MapInstructionMovs \n got: %v \n want %v \n", got, want)
	}
}

func TestMapInstructionBeq(t *testing.T){
	SetInstructionMap()

	want := "BEQ"
	got := imp.InstructionMap("00111101010100001010100001011010")

	if got != want{
		t.Errorf("MapInstructionBeq \n got: %v \n want %v \n", got, want)
	}
}

func TestMapInstructionBne(t *testing.T){
	SetInstructionMap()

	want := "BNE"
	got := imp.InstructionMap("01000001010100001010100001011010")

	if got != want{
		t.Errorf("MapInstructionBne \n got: %v \n want %v \n", got, want)
	}
}

func TestMapInstructionBlt(t *testing.T){
	SetInstructionMap()

	want := "BLT"
	got := imp.InstructionMap("01000101010100001010100001011010")

	if got != want{
		t.Errorf("MapInstructionBlt \n got: %v \n want %v \n", got, want)
	}
}

func TestMapInstructionBl(t *testing.T){
	SetInstructionMap()

	want := "BL"
	got := imp.InstructionMap("01001001010100001010100001011010")

	if got != want{
		t.Errorf("MapInstructionBl \n got: %v \n want %v \n", got, want)
	}
}

func TestMapInstructionBx(t *testing.T){
	SetInstructionMap()

	want := "BX"
	got := imp.InstructionMap("01001101010100001010100001011010")

	if got != want{
		t.Errorf("MapInstructionBx \n got: %v \n want %v \n", got, want)
	}
}

func TestMapInstructionLdr(t *testing.T){
	SetInstructionMap()

	want := "LDR"
	got := imp.InstructionMap("01010001010100001010100001011010")

	if got != want{
		t.Errorf("MapInstructionLdr \n got: %v \n want %v \n", got, want)
	}
}

func TestMapInstructionStr(t *testing.T){
	SetInstructionMap()

	want := "STR"
	got := imp.InstructionMap("01010101010100001010100001011010")

	if got != want{
		t.Errorf("MapInstructionStr \n got: %v \n want %v \n", got, want)
	}
}

func TestMapInstructionNop(t *testing.T){
	SetInstructionMap()

	want := "NOP"
	got := imp.InstructionMap("11111101010100001010100001011010")

	if got != want{
		t.Errorf("MapInstructionNop \n got: %v \n want %v \n", got, want)
	}
}