package controller

import "testing"

var dcun DecodeUnit

// Unit Tests
// Test all methods using mock
func TestGetOpcode(t *testing.T){
	want := "111111"
	dcun.opcode = "111111"
	got := dcun.GetOpcode()

	if want != got{
		t.Errorf("GetOpcode \n got: %v \n want %v \n", got, want)
	}
}
/*
func TestMapInstructionArithmetic1(t *testing.T){
	want := "Arithmetic1"
	dcun.instructionCode = "10111001010100001010100001011010"
	dcun.instructionMap = map[string]string{
		"101110": "Arithmetic1",
		"101111": "Arithmetic2",
		"100011": "Comparison1",
		"100111": "Comparison2",
		"000100": "Bypass1",
		"000101": "Bypass2",
		"101010": "LoadStore1",
		"111010": "LoadStore2",
		"000000": "Nop",
	}
	got := dcun.MapInstruction()

	if got != want{
		t.Errorf("MapInstruction for arithmetic type 1 \n got: %v \n want %v \n", got, want)
	}
}

func TestMapInstructionArithmetic2(t *testing.T){
	want := "Arithmetic2"
	dcun.instructionCode = "10111101010100001010100001011010"
	dcun.instructionMap = map[string]string{
		"101110": "Arithmetic1",
		"101111": "Arithmetic2",
		"100011": "Comparison1",
		"100111": "Comparison2",
		"000100": "Bypass1",
		"000101": "Bypass2",
		"101010": "LoadStore1",
		"111010": "LoadStore2",
		"000000": "Nop",
	}

	got := dcun.MapInstruction()

	if got != want{
		t.Errorf("MapInstruction for arithmetic type 2 \n got: %v \n want %v \n", got, want)
	}
}

func TestMapInstructionComparison1(t *testing.T){
	want := "Comparison1"
	dcun.instructionCode = "10001101010100001010100001011010"
	dcun.instructionMap = map[string]string{
		"101110": "Arithmetic1",
		"101111": "Arithmetic2",
		"100011": "Comparison1",
		"100111": "Comparison2",
		"000100": "Bypass1",
		"000101": "Bypass2",
		"101010": "LoadStore1",
		"111010": "LoadStore2",
		"000000": "Nop",
	}

	got := dcun.MapInstruction()

	if got != want{
		t.Errorf("MapInstruction for comparison type 1 \n got: %v \n want %v \n", got, want)
	}
}

func TestMapInstructionComparison2(t *testing.T){
	want := "Comparison2"
	dcun.instructionCode = "10011101010100001010100001011010"
	dcun.instructionMap = map[string]string{
		"101110": "Arithmetic1",
		"101111": "Arithmetic2",
		"100011": "Comparison1",
		"100111": "Comparison2",
		"000100": "Bypass1",
		"000101": "Bypass2",
		"101010": "LoadStore1",
		"111010": "LoadStore2",
		"000000": "Nop",
	}

	got := dcun.MapInstruction()

	if got != want{
		t.Errorf("MapInstruction for comparison type 2 \n got: %v \n want %v \n", got, want)
	}
}

func TestMapInstructionBypass1(t *testing.T){
	want := "Bypass1"
	dcun.instructionCode = "00010001010100001010100001011010"
	dcun.instructionMap = map[string]string{
		"101110": "Arithmetic1",
		"101111": "Arithmetic2",
		"100011": "Comparison1",
		"100111": "Comparison2",
		"000100": "Bypass1",
		"000101": "Bypass2",
		"101010": "LoadStore1",
		"111010": "LoadStore2",
		"000000": "Nop",
	}

	got := dcun.MapInstruction()

	if got != want{
		t.Errorf("MapInstruction for bypass and controll type 1 \n got: %v \n want %v \n", got, want)
	}
}

func TestMapInstructionBypass2(t *testing.T){
	want := "Bypass2"
	dcun.instructionCode = "00010101010100001010100001011010"
	dcun.instructionMap = map[string]string{
		"101110": "Arithmetic1",
		"101111": "Arithmetic2",
		"100011": "Comparison1",
		"100111": "Comparison2",
		"000100": "Bypass1",
		"000101": "Bypass2",
		"101010": "LoadStore1",
		"111010": "LoadStore2",
		"000000": "Nop",
	}

	got := dcun.MapInstruction()

	if got != want{
		t.Errorf("MapInstruction for bypass and controll type 2 \n got: %v \n want %v \n", got, want)
	}
}

func TestMapInstructionLoadStore1(t *testing.T){
	want := "LoadStore1"
	dcun.instructionCode = "10101001010100001010100001011010"
	dcun.instructionMap = map[string]string{
		"101110": "Arithmetic1",
		"101111": "Arithmetic2",
		"100011": "Comparison1",
		"100111": "Comparison2",
		"000100": "Bypass1",
		"000101": "Bypass2",
		"101010": "LoadStore1",
		"111010": "LoadStore2",
		"000000": "Nop",
	}

	got := dcun.MapInstruction()

	if got != want{
		t.Errorf("MapInstruction for load and store type 1 \n got: %v \n want %v \n", got, want)
	}
}

func TestMapInstructionLoadStore2(t *testing.T){
	want := "LoadStore2"
	dcun.instructionCode = "11101001010100001010100001011010"
	dcun.instructionMap = map[string]string{
		"101110": "Arithmetic1",
		"101111": "Arithmetic2",
		"100011": "Comparison1",
		"100111": "Comparison2",
		"000100": "Bypass1",
		"000101": "Bypass2",
		"101010": "LoadStore1",
		"111010": "LoadStore2",
		"000000": "Nop",
	}

	got := dcun.MapInstruction()

	if got != want{
		t.Errorf("MapInstruction for load and store type 2 \n got: %v \n want %v \n", got, want)
	}
}
/*
func TestMapInstructionNop(t *testing.T){
	want := "Nop"
	dcun.instructionCode = "00000001010100001010100001011010"
	dcun.instructionMap = map[string]string{
		"101110": "Arithmetic1",
		"101111": "Arithmetic2",
		"100011": "Comparison1",
		"100111": "Comparison2",
		"000100": "Bypass1",
		"000101": "Bypass2",
		"101010": "LoadStore1",
		"111010": "LoadStore2",
		"000000": "Nop",
	}

	got := dcun.MapInstruction()

	if got != want{
		t.Errorf("MapInstruction for nop \n got: %v \n want %v \n", got, want)
	}
}
*/

func TestSplitInstructionArithmetic1(t *testing.T){
	// Arithmetic and Logical Instructions
	// First test case
	want_opcode	:= "101110"
	want_rd 	:= "01010"
	want_rn 	:= "10000"
	want_rm		:= "10101"

	dcun.instructionCode = "10111001010100001010100001011010"
	dcun.instructionType = "Arithmetic1"
	dcun.SplitInstruction()

	got_opcode 	:= dcun.opcode
	got_rd 		:= dcun.instructionFormat["rd"]
	got_rn 		:= dcun.instructionFormat["rn"]
	got_rm 		:= dcun.instructionFormat["rm"]

	if 	got_opcode != want_opcode || got_rd != want_rd || 
		got_rn != want_rn || got_rm != want_rm{
			t.Errorf("SplitInstruction for arithmetic instruction type 1 \n got_opcode: %v \n want_opcode %v \n got_rd: %v \n want_rd %v \n got_rn: %v \n want_rn %v \n got_rm: %v \n want_rm %v", 
					got_opcode, want_opcode, got_rd, want_rd, got_rn, want_rn, got_rm, want_rm)
	}
}

func TestSplitInstructionArithmetic2(t *testing.T){
	// Arithmetic and Logical Instructions
	// Second test case
	want_opcode	:= "101111"
	want_rd 	:= "01010"
	want_rn		:= "10101"

	dcun.instructionCode = "10111101010101011010100001011010"
	dcun.instructionType = "Arithmetic2"
	dcun.SplitInstruction()

	got_opcode 	:= dcun.opcode
	got_rd 		:= dcun.instructionFormat["rd"]
	got_rn 		:= dcun.instructionFormat["rn"]

	if 	got_opcode != want_opcode || got_rd != want_rd || got_rn != want_rn{
			t.Errorf("SplitInstruction for arithmetic instruction type 2 \n got_opcode: %v \n want_opcode %v \n got_rd: %v \n want_rd %v \n got_rn: %v \n want_rn %v", 
					got_opcode, want_opcode, got_rd, want_rd, got_rn, want_rn)
	}
}


func TestSplitInstruictionComparison1(t *testing.T){
	// Comparison Instructions
	// First test case
	want_opcode	:= "100011"
	want_rn 	:= "10000"
	want_rm		:= "10101"

	dcun.instructionCode = "10001110000101011010100001011010"
	dcun.instructionType = "Comparison1"
	dcun.SplitInstruction()

	got_opcode 	:= dcun.opcode
	got_rn 		:= dcun.instructionFormat["rn"]
	got_rm 		:= dcun.instructionFormat["rm"]

	if 	got_opcode != want_opcode || got_rn != want_rn || got_rm != want_rm{
			t.Errorf("SplitInstruction for comparison instruction type 1 \n got_opcode: %v \n want_opcode %v \n got_rn: %v \n want_rn %v \n got_rm: %v \n want_rm %v", 
					got_opcode, want_opcode, got_rn, want_rn, got_rm, want_rm)
	}
}

func TestSplitInstruictionComparison2(t *testing.T){
	// Comparison Instructions
	// Second test case
	want_opcode		:= "100111"
	want_rn			:= "10101"
	want_imediato	:= "00100111"

	dcun.instructionCode = "10011110101001001110100001011010"
	dcun.instructionType = "Comparison2"
	dcun.SplitInstruction()

	got_opcode 		:= dcun.opcode
	got_rn 			:= dcun.instructionFormat["rn"]
	got_imediato 	:= dcun.instructionFormat["imediato"]

	if 	got_opcode != want_opcode || got_rn != want_rn || got_imediato != want_imediato{
			t.Errorf("SplitInstruction for comparison instruction type 2 \n got_opcode: %v \n want_opcode %v \n got_rn: %v \n want_rn %v \n got_imediato: %v \n want_imediato %v", 
					got_opcode, want_opcode, got_rn, want_rn, got_imediato, want_imediato)
	}
}

func TestSplitInstructionBypass1(t *testing.T){
	// Bypass and Controll Instructions
	// First test case
	want_opcode	:= "000100"
	want_rn		:= "10101"

	dcun.instructionCode = "00010010101001001110100001011010"
	dcun.instructionType = "Bypass1"
	dcun.SplitInstruction()

	got_opcode 	:= dcun.opcode
	got_rn 		:= dcun.instructionFormat["rn"]

	if 	got_opcode != want_opcode || got_rn != want_rn{
			t.Errorf("SplitInstruction for bypass and controll instructions type 1 \n got_opcode: %v \n want_opcode %v \n got_rn: %v \n want_rn %v", 
					got_opcode, want_opcode, got_rn, want_rn)
	}
}

func TestSplitInstructionBypass2(t *testing.T){
	// Bypass and Controll Instructions
	// First test case
	want_opcode	:= "000101"
	want_label	:= "00010101001110100001011010"

	dcun.instructionCode = "00010100010101001110100001011010"
	dcun.instructionType = "Bypass2"
	dcun.SplitInstruction()

	got_opcode 		:= dcun.opcode
	got_label		:= dcun.instructionFormat["label"]

	if 	got_opcode != want_opcode || got_label != want_label{
			t.Errorf("SplitInstruction for bypass and controll instructions type 2 \n got_opcode: %v \n want_opcode %v \n got_label: %v \n want_label %v", 
					got_opcode, want_opcode, got_label, want_label)
	}
}

func TestSplitInstructionLoadStore1(t *testing.T){
	// Load and Store Instruction
	// First test case
	want_opcode	:= "101010"
	want_rd		:= "10111"
	want_rn		:= "10001"

	dcun.instructionCode = "10101010111100011110100001011010"
	dcun.instructionType = "LoadStore1"
	dcun.SplitInstruction()

	got_opcode	:= dcun.opcode
	got_rd		:= dcun.instructionFormat["rd"]
	got_rn		:= dcun.instructionFormat["rn"]

	if 	got_opcode != want_opcode || got_rd != want_rd || got_rn != want_rn{
			t.Errorf("SplitInstruction for load and store instructions type 1 \n got_opcode: %v \n want_opcode %v \n got_rd: %v \n want_rd %v \n got_rn: %v \n want_rn %v", 
					got_opcode, want_opcode, got_rd, want_rd, got_rn, want_rn)
	}
}

func TestSplitInstructionLoadStore2(t *testing.T){
	// Load and Store Instruction
	// Second test case
	want_opcode		:= "111010"
	want_rn			:= "10111"
	want_address	:= "11110100001011"

	dcun.instructionCode = "11101010111111101000010111011010"
	dcun.instructionType = "LoadStore2"
	dcun.SplitInstruction()

	got_opcode	:= dcun.opcode
	got_rn		:= dcun.instructionFormat["rn"]
	got_address	:= dcun.instructionFormat["address"]

	if 	got_opcode != want_opcode || got_rn != want_rn || got_address != want_address{
			t.Errorf("SplitInstruction for load and store instructions type 2 \n got_opcode: %v \n want_opcode %v \n got_rn: %v \n want_rn %v \n got_address: %v \n want_address %v", 
					got_opcode, want_opcode, got_rn, want_rn, got_address, want_address)
	}
}

func TestSplitInstructionNop(t *testing.T){
	// Nop Instruction
	// Unic test case
	want_opcode	:= "000000"
	want_rest	:= ""

	dcun.instructionCode = "00000010111111101000010111011010"
	dcun.instructionType = "Nop"
	dcun.SplitInstruction()

	got_opcode	:= dcun.opcode
	got_rest	:= ""

	if 	got_opcode != want_opcode || got_rest != want_rest{
			t.Errorf("SplitInstruction for nop instruction \n got_opcode: %v \n want_opcode %v \n got_rest: %v \n want_rest %v", 
					got_opcode, want_opcode, got_rest, want_rest)
	}
}