package controller

import (
	"testing"
	"reflect"
)

var cn Controller

func TestChangeInstructionFetch(t *testing.T){
	cn := NewController()

	want := "10111001010100001010100001011010"
	cn.ChangeInstructionFetch("10111001010100001010100001011010")
	got := cn.fetchUnit.instruction 

	if got != want{
		t.Errorf("ChangeInstructionFetch \n got: %v \n want %v \n", got, want)
	}
}

func TestChangeInstructionDecode(t *testing.T){
	cn := NewController()

	want := "10111001010100001010100001011010"
	cn.ChangeInstructionDecode("10111001010100001010100001011010")
	got := cn.decodeUnit.instructionCode 

	if got != want{
		t.Errorf("ChangeInstructionDecode \n got: %v \n want %v \n", got, want)
	}
}

func TestChangeInstructionExecute(t *testing.T){
	cn := NewController()

	want := "10111001010100001010100001011010"
	cn.ChangeInstructionExecute("10111001010100001010100001011010")
	got := cn.executeUnit.instruction

	if got != want{
		t.Errorf("ChangeInstructionExecute \n got: %v \n want %v \n", got, want)
	}
}

// ---------------------------------------------------------------------------------------
// Tests about decode unit
// ---------------------------------------------------------------------------------------
func TestSplitInstructionArithmetic1(t *testing.T){
	cn := NewController()
	// Arithmetic and Logical Instructions
	// First test case
	want_opcode	:= "000011"
	want_rd 	:= "01010"
	want_rn 	:= "10000"
	want_rm		:= "10101"

	cn.decodeUnit.instructionCode = "00001101010100001010100001011010"
	cn.decodeUnit.SplitInstruction()

	got_opcode 	:= cn.decodeUnit.opcode
	got_rd 		:= cn.decodeUnit.instructionFormat["rd"]
	got_rn 		:= cn.decodeUnit.instructionFormat["rn"]
	got_rm 		:= cn.decodeUnit.instructionFormat["rm"]

	if 	got_opcode != want_opcode || got_rd != want_rd || 
		got_rn != want_rn || got_rm != want_rm{
			t.Errorf("SplitInstruction for arithmetic instruction type 1 \n got_opcode: %v \n want_opcode %v \n got_rd: %v \n want_rd %v \n got_rn: %v \n want_rn %v \n got_rm: %v \n want_rm %v", 
					got_opcode, want_opcode, got_rd, want_rd, got_rn, want_rn, got_rm, want_rm)
	}
}

func TestSplitInstructionArithmetic2(t *testing.T){
	cn := NewController()
	// Arithmetic and Logical Instructions
	// Second test case
	want_opcode	:= "100100"
	want_rd 	:= "01010"
	want_rn		:= "10101"

	cn.decodeUnit.instructionCode = "10010001010101011010100001011010"
	cn.decodeUnit.SplitInstruction()

	got_opcode 	:= cn.decodeUnit.opcode
	got_rd 		:= cn.decodeUnit.instructionFormat["rd"]
	got_rn 		:= cn.decodeUnit.instructionFormat["rn"]

	if 	got_opcode != want_opcode || got_rd != want_rd || got_rn != want_rn{
			t.Errorf("SplitInstruction for arithmetic instruction type 2 \n got_opcode: %v \n want_opcode %v \n got_rd: %v \n want_rd %v \n got_rn: %v \n want_rn %v", 
					got_opcode, want_opcode, got_rd, want_rd, got_rn, want_rn)
	}
}


func TestSplitInstruictionComparison1(t *testing.T){
	cn := NewController()
	// Comparison Instructions
	// First test case
	want_opcode	:= "001100"
	want_rn 	:= "10000"
	want_rm		:= "10101"

	cn.decodeUnit.instructionCode = "00110010000101011010100001011010"
	cn.decodeUnit.SplitInstruction()

	got_opcode 	:= cn.decodeUnit.opcode
	got_rn 		:= cn.decodeUnit.instructionFormat["rn"]
	got_rm 		:= cn.decodeUnit.instructionFormat["rm"]

	if 	got_opcode != want_opcode || got_rn != want_rn || got_rm != want_rm{
			t.Errorf("SplitInstruction for comparison instruction type 1 \n got_opcode: %v \n want_opcode %v \n got_rn: %v \n want_rn %v \n got_rm: %v \n want_rm %v", 
					got_opcode, want_opcode, got_rn, want_rn, got_rm, want_rm)
	}
}

func TestSplitInstruictionComparison2(t *testing.T){
	cn := NewController()
	// Comparison Instructions
	// Second test case
	want_opcode		:= "101101"
	want_rn			:= "10101"
	want_imediato	:= "00100111"

	cn.decodeUnit.instructionCode = "10110110101001001110100001011010"
	cn.decodeUnit.SplitInstruction()

	got_opcode 		:= cn.decodeUnit.opcode
	got_rn 			:= cn.decodeUnit.instructionFormat["rn"]
	got_imediato 	:= cn.decodeUnit.instructionFormat["imediato"]

	if 	got_opcode != want_opcode || got_rn != want_rn || got_imediato != want_imediato{
			t.Errorf("SplitInstruction for comparison instruction type 2 \n got_opcode: %v \n want_opcode %v \n got_rn: %v \n want_rn %v \n got_imediato: %v \n want_imediato %v", 
					got_opcode, want_opcode, got_rn, want_rn, got_imediato, want_imediato)
	}
}

func TestSplitInstructionBypass1(t *testing.T){
	cn := NewController()
	// Bypass and Controll Instructions
	// First test case
	want_opcode	:= "010000"
	want_rn		:= "10101"

	cn.decodeUnit.instructionCode = "01000010101001001110100001011010"
	cn.decodeUnit.SplitInstruction()

	got_opcode 	:= cn.decodeUnit.opcode
	got_rn 		:= cn.decodeUnit.instructionFormat["rn"]

	if 	got_opcode != want_opcode || got_rn != want_rn{
			t.Errorf("SplitInstruction for bypass and controll instructions type 1 \n got_opcode: %v \n want_opcode %v \n got_rn: %v \n want_rn %v", 
					got_opcode, want_opcode, got_rn, want_rn)
	}
}

func TestSplitInstructionBypass2(t *testing.T){
	cn := NewController()
	// Bypass and Controll Instructions
	// First test case
	want_opcode	:= "110001"
	want_label	:= "00010101001110100001011010"

	cn.decodeUnit.instructionCode = "11000100010101001110100001011010"
	cn.decodeUnit.SplitInstruction()

	got_opcode 		:= cn.decodeUnit.opcode
	got_label		:= cn.decodeUnit.instructionFormat["label"]

	if 	got_opcode != want_opcode || got_label != want_label{
			t.Errorf("SplitInstruction for bypass and controll instructions type 2 \n got_opcode: %v \n want_opcode %v \n got_label: %v \n want_label %v", 
					got_opcode, want_opcode, got_label, want_label)
	}
}

func TestSplitInstructionLoadStore1(t *testing.T){
	cn := NewController()
	// Load and Store Instruction
	// First test case
	want_opcode	:= "010100"
	want_rd		:= "10111"
	want_rn		:= "10001"

	cn.decodeUnit.instructionCode = "01010010111100011110100001011010"
	cn.decodeUnit.SplitInstruction()

	got_opcode	:= cn.decodeUnit.opcode
	got_rd		:= cn.decodeUnit.instructionFormat["rd"]
	got_rn		:= cn.decodeUnit.instructionFormat["rn"]

	if 	got_opcode != want_opcode || got_rd != want_rd || got_rn != want_rn{
			t.Errorf("SplitInstruction for load and store instructions type 1 \n got_opcode: %v \n want_opcode %v \n got_rd: %v \n want_rd %v \n got_rn: %v \n want_rn %v", 
					got_opcode, want_opcode, got_rd, want_rd, got_rn, want_rn)
	}
}

func TestSplitInstructionLoadStore2(t *testing.T){
	cn := NewController()
	// Load and Store Instruction
	// Second test case
	want_opcode		:= "110101"
	want_rn			:= "10111"
	want_address	:= "11110100001011"

	cn.decodeUnit.instructionCode = "11010110111111101000010111011010"
	cn.decodeUnit.SplitInstruction()

	got_opcode	:= cn.decodeUnit.opcode
	got_rn		:= cn.decodeUnit.instructionFormat["rn"]
	got_address	:= cn.decodeUnit.instructionFormat["address"]

	if 	got_opcode != want_opcode || got_rn != want_rn || got_address != want_address{
			t.Errorf("SplitInstruction for load and store instructions type 2 \n got_opcode: %v \n want_opcode %v \n got_rn: %v \n want_rn %v \n got_address: %v \n want_address %v", 
					got_opcode, want_opcode, got_rn, want_rn, got_address, want_address)
	}
}

func TestSplitInstructionNop(t *testing.T){
	cn := NewController()
	// Nop Instruction
	// Unic test case
	want_opcode	:= "111111"
	want_rest	:= ""

	cn.decodeUnit.instructionCode = "11111111111110111111101000010111011010"
	cn.decodeUnit.instructionType1 = "Nop"
	cn.decodeUnit.SplitInstruction()

	got_opcode	:= cn.decodeUnit.opcode
	got_rest	:= ""

	if 	got_opcode != want_opcode || got_rest != want_rest{
			t.Errorf("SplitInstruction for nop instruction \n got_opcode: %v \n want_opcode %v \n got_rest: %v \n want_rest %v", 
					got_opcode, want_opcode, got_rest, want_rest)
	}
}

// ---------------------------------------------------------------------------------------
// Tests about execute unit
// ---------------------------------------------------------------------------------------
func TestConfigForLdr1(t *testing.T){
	cn := NewController()

	want := map[string]int64{
		"target_register": 2,
		"address_register": 5,
	}
	cn.executeUnit.ConfigForLdr1("00010", "00101")
	got := cn.executeUnit.config

	if !reflect.DeepEqual(want, got){
		t.Errorf("ConfigForLdr type 1 \n got: %v \n want %v \n", got, want)
	}
}

func TestConfigForLdr2(t *testing.T){
	cn := NewController()

	want := map[string]int64{
		"target_register": 2,
		"address": 70,
	}
	cn.executeUnit.ConfigForLdr2("00010", "00000001000110")
	got := cn.executeUnit.config

	if !reflect.DeepEqual(want, got){
		t.Errorf("ConfigForLdr type 2 \n got: %v \n want %v \n", got, want)
	}
}

func TestConfigForStr1(t *testing.T){
	cn := NewController()

	want := map[string]int64{
		"target_register": 2,
		"address_register": 5,
	}
	cn.executeUnit.ConfigForStr1("00010", "00101")
	got := cn.executeUnit.config

	if !reflect.DeepEqual(want, got){
		t.Errorf("ConfigForStr type 1 \n got: %v \n want %v \n", got, want)
	}
}

func TestConfigForStr2(t *testing.T){
	cn := NewController()

	want := map[string]int64{
		"target_register": 2,
		"address": 70,
	}
	cn.executeUnit.ConfigForStr2("00010", "00000001000110")
	got := cn.executeUnit.config

	if !reflect.DeepEqual(want, got){
		t.Errorf("ConfigForStr type 2 \n got: %v \n want %v \n", got, want)
	}
}

func TestConfigForAdds1(t *testing.T){
	cn := NewController()

	want := map[string]int64{
		"target_register": 2,
		"source_register1": 2,
		"source_register2": 4,
	}
	cn.executeUnit.ConfigForAdds1("00010", "00010", "00100")
	got := cn.executeUnit.config

	if !reflect.DeepEqual(want, got){
		t.Errorf("ConfigForAdds type 1 \n got: %v \n want %v \n", got, want)
	}
}

func TestConfigForAdds2(t *testing.T){
	cn := NewController()

	want := map[string]int64{
		"target_register": 2,
		"source_register": 2,
		"value": 70,
	}
	cn.executeUnit.ConfigForAdds2("00010", "00010", "0000000001000110")
	got := cn.executeUnit.config

	if !reflect.DeepEqual(want, got){
		t.Errorf("ConfigForAdds type 2 \n got: %v \n want %v \n", got, want)
	}
}

func TestConfigForSubs1(t *testing.T){
	cn := NewController()

	want := map[string]int64{
		"target_register": 2,
		"source_register1": 2,
		"source_register2": 4, 
	}
	cn.executeUnit.ConfigForSubs1("00010", "00010", "00100")
	got := cn.executeUnit.config

	if !reflect.DeepEqual(want, got){
		t.Errorf("ConfigForSubs type 1 \n got: %v \n want %v \n", got, want)
	}
}

func TestConfigForSubs2(t *testing.T){
	cn := NewController()

	want := map[string]int64{
		"target_register": 2,
		"source_register": 2,
		"value": 70,
	}
	cn.executeUnit.ConfigForSubs2("00010", "00010", "0000000001000110")
	got := cn.executeUnit.config

	if !reflect.DeepEqual(want, got){
		t.Errorf("ConfigForSubs type 2 \n got: %v \n want %v \n", got, want)
	}
}

func TestConfigForAnds1(t *testing.T){
	cn := NewController()

	want := map[string]int64{
		"target_register": 2,
		"source_register1": 2,
		"source_register2": 4,
	}
	cn.executeUnit.ConfigForAnds1("00010", "00010", "00100")
	got := cn.executeUnit.config

	if !reflect.DeepEqual(want, got){
		t.Errorf("ConfigForAnds type 1 \n got: %v \n want %v \n", got, want)
	}
}

func TestConfigForAnds2(t *testing.T){
	cn := NewController()

	want := map[string]int64{
		"target_register": 2,
		"source_register": 2,
		"value": 70,
	}
	cn.executeUnit.ConfigForAnds2("00010", "00010", "0000000001000110")
	got := cn.executeUnit.config

	if !reflect.DeepEqual(want, got){
		t.Errorf("ConfigForAnds type 2 \n got: %v \n want %v \n", got, want)
	}
}

func TestConfigForMovs1(t *testing.T){
	cn := NewController()

	want := map[string]int64{
		"target_register": 2,
		"value_register": 5,
	}

	cn.executeUnit.ConfigForMovs1("00010", "00101")
	got := cn.executeUnit.config

	if !reflect.DeepEqual(want, got){
		t.Errorf("ConfigFormovs type 1 \n got: %v \n want %v \n", got, want)
	}
}

func TestConfigForMovs2(t *testing.T){
	cn := NewController()

	want := map[string]int64{
		"target_register": 2,
		"value": 70,
	}
	cn.executeUnit.ConfigForMovs2("00010", "00000001000110")
	got := cn.executeUnit.config

	if !reflect.DeepEqual(want, got){
		t.Errorf("ConfigForMovs type 2 \n got: %v \n want %v \n", got, want)
	}
}