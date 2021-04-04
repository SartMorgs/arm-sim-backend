package controller

import "testing"

var imp InstructionMap

// Unit Tests
// Test all methods using mock

func SetInstructionMap(){
	// Instruction map
	imp.instructionMap = map[string][3]string{
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
		"110011": {"BX", "Bypass", "2"},

		// Load and Store
		"010100": {"LDR", "Load", "1"},
		"110100": {"LDR", "Load", "2"},
		"010101": {"STR", "Store", "1"},
		"110101": {"STR", "Store", "2"},

		// Nothing
		"111111": {"NOP", "Nothing", "1"},
	}
}

func TestGetTypeInstructionAdds(t *testing.T){
	SetInstructionMap()

	want := "Arithmetic"
	got := imp.GetTypeInstruction("00000101010100001010100001011010")

	if got != want{
		t.Errorf("GetTypeInstructionAdds \n got: %v \n want %v \n", got, want)
	}
}

func TestGetTypeInstructionSubs(t *testing.T){
	SetInstructionMap()

	want := "Arithmetic"
	got := imp.GetTypeInstruction("00001001010100001010100001011010")

	if got != want{
		t.Errorf("GetTypeInstructionSubs \n got: %v \n want %v \n", got, want)
	}
}

func TestGetTypeInstructionMuls(t *testing.T){
	SetInstructionMap()

	want := "Arithmetic"
	got := imp.GetTypeInstruction("00001101010100001010100001011010")

	if got != want{
		t.Errorf("GetTypeInstructionMuls \n got: %v \n want %v \n", got, want)
	}
}

func TestGetTypeInstructionAnds(t *testing.T){
	SetInstructionMap()

	want := "Arithmetic"
	got := imp.GetTypeInstruction("00010001010100001010100001011010")

	if got != want{
		t.Errorf("GetTypeInstructionAnds \n got: %v \n want %v \n", got, want)
	}
}

func TestGetTypeInstructionOrrs(t *testing.T){
	SetInstructionMap()

	want := "Arithmetic"
	got := imp.GetTypeInstruction("00010101010100001010100001011010")

	if got != want{
		t.Errorf("GetTypeInstructionOrrs \n got: %v \n want %v \n", got, want)
	}
}

func TestGetTypeInstructionEors(t *testing.T){
	SetInstructionMap()

	want := "Arithmetic"
	got := imp.GetTypeInstruction("00011001010100001010100001011010")

	if got != want{
		t.Errorf("GetTypeInstructionEors \n got: %v \n want %v \n", got, want)
	}
}

func TestGetTypeInstructionBics(t *testing.T){
	SetInstructionMap()

	want := "Arithmetic"
	got := imp.GetTypeInstruction("00011101010100001010100001011010")

	if got != want{
		t.Errorf("GetTypeInstructionBics \n got: %v \n want %v \n", got, want)
	}
}

func TestGetTypeInstructionAsrs(t *testing.T){
	SetInstructionMap()

	want := "Arithmetic"
	got := imp.GetTypeInstruction("00100001010100001010100001011010")

	if got != want{
		t.Errorf("GetTypeInstructionAsrs \n got: %v \n want %v \n", got, want)
	}
}

func TestGetTypeInstructionLsls(t *testing.T){
	SetInstructionMap()

	want := "Arithmetic"
	got := imp.GetTypeInstruction("00100101010100001010100001011010")

	if got != want{
		t.Errorf("GetTypeInstructionLsls \n got: %v \n want %v \n", got, want)
	}
}

func TestGetTypeInstructionLsrs(t *testing.T){
	SetInstructionMap()

	want := "Arithmetic"
	got := imp.GetTypeInstruction("00101001010100001010100001011010")

	if got != want{
		t.Errorf("GetTypeInstructionLsrs \n got: %v \n want %v \n", got, want)
	}
}

func TestGetTypeInstructionRors(t *testing.T){
	SetInstructionMap()

	want := "Arithmetic"
	got := imp.GetTypeInstruction("00101101010100001010100001011010")

	if got != want{
		t.Errorf("GetTypeInstructionRors \n got: %v \n want %v \n", got, want)
	}
}

func TestGetTypeInstructionCmn(t *testing.T){
	SetInstructionMap()

	want := "Comparison"
	got := imp.GetTypeInstruction("00110001010100001010100001011010")

	if got != want{
		t.Errorf("GetTypeInstructioCmn \n got: %v \n want %v \n", got, want)
	}
}

func TestGetTypeInstructionCmp(t *testing.T){
	SetInstructionMap()

	want := "Comparison"
	got := imp.GetTypeInstruction("00110101010100001010100001011010")

	if got != want{
		t.Errorf("GetTypeInstructioCmp \n got: %v \n want %v \n", got, want)
	}
}

func TestGetTypeInstructionMovs(t *testing.T){
	SetInstructionMap()

	want := "Bypass"
	got := imp.GetTypeInstruction("00111001010100001010100001011010")

	if got != want{
		t.Errorf("GetTypeInstructioMovs \n got: %v \n want %v \n", got, want)
	}
}

func TestGetTypeInstructionBeq(t *testing.T){
	SetInstructionMap()

	want := "Bypass"
	got := imp.GetTypeInstruction("00111101010100001010100001011010")

	if got != want{
		t.Errorf("GetTypeInstructioBeq \n got: %v \n want %v \n", got, want)
	}
}

func TestGetTypeInstructionBne(t *testing.T){
	SetInstructionMap()

	want := "Bypass"
	got := imp.GetTypeInstruction("01000001010100001010100001011010")

	if got != want{
		t.Errorf("GetTypeInstructioBne \n got: %v \n want %v \n", got, want)
	}
}

func TestGetTypeInstructionBlt(t *testing.T){
	SetInstructionMap()

	want := "Bypass"
	got := imp.GetTypeInstruction("01000101010100001010100001011010")

	if got != want{
		t.Errorf("GetTypeInstructioBlt \n got: %v \n want %v \n", got, want)
	}
}

func TestGetTypeInstructionBl(t *testing.T){
	SetInstructionMap()

	want := "Bypass"
	got := imp.GetTypeInstruction("01001001010100001010100001011010")

	if got != want{
		t.Errorf("GetTypeInstructioBl \n got: %v \n want %v \n", got, want)
	}
}

func TestGetTypeInstructionBx(t *testing.T){
	SetInstructionMap()

	want := "Bypass"
	got := imp.GetTypeInstruction("01001101010100001010100001011010")

	if got != want{
		t.Errorf("GetTypeInstructioBx \n got: %v \n want %v \n", got, want)
	}
}

func TestGetTypeInstructionLdr(t *testing.T){
	SetInstructionMap()

	want := "Load"
	got := imp.GetTypeInstruction("01010001010100001010100001011010")

	if got != want{
		t.Errorf("GetTypeInstructioLdr \n got: %v \n want %v \n", got, want)
	}
}

func TestGetTypeInstructionStr(t *testing.T){
	SetInstructionMap()

	want := "Store"
	got := imp.GetTypeInstruction("01010101010100001010100001011010")

	if got != want{
		t.Errorf("GetTypeInstructioStr \n got: %v \n want %v \n", got, want)
	}
}

func TestGetTypeInstructionNop(t *testing.T){
	SetInstructionMap()

	want := "Nothing"
	got := imp.GetTypeInstruction("11111101010100001010100001011010")

	if got != want{
		t.Errorf("GetTypeInstructioNop \n got: %v \n want %v \n", got, want)
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