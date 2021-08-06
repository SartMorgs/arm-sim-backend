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
	want_rd			:= "10111"
	want_address	:= "11110100001011"

	cn.decodeUnit.instructionCode = "11010110111111101000010111011010"
	cn.decodeUnit.SplitInstruction()

	got_opcode	:= cn.decodeUnit.opcode
	got_rd		:= cn.decodeUnit.instructionFormat["rd"]
	got_address	:= cn.decodeUnit.instructionFormat["address"]

	if 	got_opcode != want_opcode || got_rd != want_rd || got_address != want_address{
			t.Errorf("SplitInstruction for load and store instructions type 2 \n got_opcode: %v \n want_opcode %v \n got_rd: %v \n want_rd %v \n got_address: %v \n want_address %v", 
					got_opcode, want_opcode, got_rd, want_rd, got_address, want_address)
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

func TestGetInstructionFormat(t *testing.T){
	cn := NewController()
	want := map[string]string{
		"rd": "10111",
		"address": "11110100001011",
	}

	cn.decodeUnit.instructionCode = "11010110111111101000010111011010"
	cn.decodeUnit.SplitInstruction()
	got := cn.decodeUnit.GetInstructionFormat()

	if !reflect.DeepEqual(want, got){
		t.Errorf("GetInstructionFormat \n got: %v \n want %v \n", got, want)
	}
}

func TestGetInstructionName(t *testing.T){
	cn := NewController()
	want := "ORRS"

	cn.decodeUnit.instructionCode = "00010101010100001010100001011010"
	cn.decodeUnit.SplitInstruction()
	got := cn.decodeUnit.GetInstructionName()

	if got != want{
		t.Errorf("GetInstructionName after split \n got: %v \n want %v \n", got, want)
	}
}

// ---------------------------------------------------------------------------------------
// Tests about execute unit
// ---------------------------------------------------------------------------------------

func TestConfigForLdr1(t *testing.T){
	cn := NewController()

	want := map[string]int64{
		"target_register": 2,
		"source_register1": 5,
	}
	cn.executeUnit.SetTargetRegister("00010")
	cn.executeUnit.SetSourceRegister1("00101")
	cn.executeUnit.ConfigForLdr1()
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
	cn.executeUnit.SetTargetRegister("00010")
	cn.executeUnit.SetAddressInstruction("00000001000110")
	cn.executeUnit.ConfigForLdr2()
	got := cn.executeUnit.config

	if !reflect.DeepEqual(want, got){
		t.Errorf("ConfigForLdr type 2 \n got: %v \n want %v \n", got, want)
	}
}

func TestConfigForStr1(t *testing.T){
	cn := NewController()

	want := map[string]int64{
		"target_register": 2,
		"source_register1": 5,
	}
	cn.executeUnit.SetTargetRegister("00010")
	cn.executeUnit.SetSourceRegister1("00101")
	cn.executeUnit.ConfigForStr1()
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
	cn.executeUnit.SetTargetRegister("00010")
	cn.executeUnit.SetAddressInstruction("00000001000110")
	cn.executeUnit.ConfigForStr2()
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
	cn.executeUnit.SetTargetRegister("00010")
	cn.executeUnit.SetSourceRegister1("00010")
	cn.executeUnit.SetSourceRegister2("00100")
	cn.executeUnit.ConfigForAdds1()
	got := cn.executeUnit.config

	if !reflect.DeepEqual(want, got){
		t.Errorf("ConfigForAdds type 1 \n got: %v \n want %v \n", got, want)
	}
}

func TestConfigForAdds2(t *testing.T){
	cn := NewController()

	want := map[string]int64{
		"target_register": 2,
		"source_register1": 2,
		"value": 70,
	}
	cn.executeUnit.SetTargetRegister("00010")
	cn.executeUnit.SetSourceRegister1("00010")
	cn.executeUnit.SetValueInstruction("0000000001000110")
	cn.executeUnit.ConfigForAdds2()
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
	cn.executeUnit.SetTargetRegister("00010")
	cn.executeUnit.SetSourceRegister1("00010")
	cn.executeUnit.SetSourceRegister2("00100")
	cn.executeUnit.ConfigForSubs1()
	got := cn.executeUnit.config

	if !reflect.DeepEqual(want, got){
		t.Errorf("ConfigForSubs type 1 \n got: %v \n want %v \n", got, want)
	}
}

func TestConfigForSubs2(t *testing.T){
	cn := NewController()

	want := map[string]int64{
		"target_register": 2,
		"source_register1": 2,
		"value": 70,
	}
	cn.executeUnit.SetTargetRegister("00010")
	cn.executeUnit.SetSourceRegister1("00010")
	cn.executeUnit.SetValueInstruction("0000000001000110")
	cn.executeUnit.ConfigForSubs2()
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
	cn.executeUnit.SetTargetRegister("00010")
	cn.executeUnit.SetSourceRegister1("00010")
	cn.executeUnit.SetSourceRegister2("00100")
	cn.executeUnit.ConfigForAnds1()
	got := cn.executeUnit.config

	if !reflect.DeepEqual(want, got){
		t.Errorf("ConfigForAnds type 1 \n got: %v \n want %v \n", got, want)
	}
}

func TestConfigForAnds2(t *testing.T){
	cn := NewController()

	want := map[string]int64{
		"target_register": 2,
		"source_register1": 2,
		"value": 70,
	}
	cn.executeUnit.SetTargetRegister("00010")
	cn.executeUnit.SetSourceRegister1("00010")
	cn.executeUnit.SetValueInstruction("0000000001000110")
	cn.executeUnit.ConfigForAnds2()
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
	cn.executeUnit.SetTargetRegister("00010")
	cn.executeUnit.SetValueRegister("00101")
	cn.executeUnit.ConfigForMovs1()
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
	cn.executeUnit.SetTargetRegister("00010")
	cn.executeUnit.SetValueInstruction("00000001000110")
	cn.executeUnit.ConfigForMovs2()
	got := cn.executeUnit.config

	if !reflect.DeepEqual(want, got){
		t.Errorf("ConfigForMovs type 2 \n got: %v \n want %v \n", got, want)
	}
}

func TestConfigForMuls1(t *testing.T){
	cn := NewController()

	want := map[string]int64{
		"target_register": 2,
		"source_register1": 2,
		"source_register2": 4,
	}
	cn.executeUnit.SetTargetRegister("00010")
	cn.executeUnit.SetSourceRegister1("00010")
	cn.executeUnit.SetSourceRegister2("00100")
	cn.executeUnit.ConfigForMuls1()

	got := cn.executeUnit.config

	if !reflect.DeepEqual(want, got){
		t.Errorf("ConfigForMuls type 1 \n got: %v \n want %v \n", got, want)
	}
}

func TestConfigForMuls2(t *testing.T){
	cn := NewController()

	want := map[string]int64{
		"target_register": 2,
		"source_register1": 2,
		"value": 70,
	}
	cn.executeUnit.SetTargetRegister("00010")
	cn.executeUnit.SetSourceRegister1("00010")
	cn.executeUnit.SetValueInstruction("0000000001000110")
	cn.executeUnit.ConfigForMuls2()
	got := cn.executeUnit.config

	if !reflect.DeepEqual(want, got){
		t.Errorf("ConfigForMuls type 2 \n got: %v \n want %v \n", got, want)
	}
}

func TestConfigForOrrs1(t *testing.T){
	cn := NewController()

	want := map[string]int64{
		"target_register": 2,
		"source_register1": 2,
		"source_register2": 4,
	}
	cn.executeUnit.SetTargetRegister("00010")
	cn.executeUnit.SetSourceRegister1("00010")
	cn.executeUnit.SetSourceRegister2("00100")
	cn.executeUnit.ConfigForOrrs1()

	got := cn.executeUnit.config

	if !reflect.DeepEqual(want, got){
		t.Errorf("ConfigForOrrs type 1 \n got: %v \n want %v \n", got, want)
	}
}

func TestConfigForOrrs2(t *testing.T){
	cn := NewController()

	want := map[string]int64{
		"target_register": 2,
		"source_register1": 2,
		"value": 70,
	}
	cn.executeUnit.SetTargetRegister("00010")
	cn.executeUnit.SetSourceRegister1("00010")
	cn.executeUnit.SetValueInstruction("0000000001000110")
	cn.executeUnit.ConfigForOrrs2()
	got := cn.executeUnit.config

	if !reflect.DeepEqual(want, got){
		t.Errorf("ConfigForOrrs type 2 \n got: %v \n want %v \n", got, want)
	}
}

func TestConfigForEors1(t *testing.T){
	cn := NewController()

	want := map[string]int64{
		"target_register": 2,
		"source_register1": 2,
		"source_register2": 4,
	}
	cn.executeUnit.SetTargetRegister("00010")
	cn.executeUnit.SetSourceRegister1("00010")
	cn.executeUnit.SetSourceRegister2("00100")
	cn.executeUnit.ConfigForEors1()

	got := cn.executeUnit.config

	if !reflect.DeepEqual(want, got){
		t.Errorf("ConfigForEors type 1 \n got: %v \n want %v \n", got, want)
	}
}

func TestConfigForEors2(t *testing.T){
	cn := NewController()

	want := map[string]int64{
		"target_register": 2,
		"source_register1": 2,
		"value": 70,
	}
	cn.executeUnit.SetTargetRegister("00010")
	cn.executeUnit.SetSourceRegister1("00010")
	cn.executeUnit.SetValueInstruction("0000000001000110")
	cn.executeUnit.ConfigForEors2()
	got := cn.executeUnit.config

	if !reflect.DeepEqual(want, got){
		t.Errorf("ConfigForEors type 2 \n got: %v \n want %v \n", got, want)
	}
}

func TestConfigForBics1(t *testing.T){
	cn := NewController()

	want := map[string]int64{
		"target_register": 2,
		"source_register1": 2,
		"source_register2": 4,
	}
	cn.executeUnit.SetTargetRegister("00010")
	cn.executeUnit.SetSourceRegister1("00010")
	cn.executeUnit.SetSourceRegister2("00100")
	cn.executeUnit.ConfigForBics1()

	got := cn.executeUnit.config

	if !reflect.DeepEqual(want, got){
		t.Errorf("ConfigForBics type 1 \n got: %v \n want %v \n", got, want)
	}
}

func TestConfigForBics2(t *testing.T){
	cn := NewController()

	want := map[string]int64{
		"target_register": 2,
		"source_register1": 2,
		"value": 70,
	}
	cn.executeUnit.SetTargetRegister("00010")
	cn.executeUnit.SetSourceRegister1("00010")
	cn.executeUnit.SetValueInstruction("0000000001000110")
	cn.executeUnit.ConfigForBics2()
	got := cn.executeUnit.config

	if !reflect.DeepEqual(want, got){
		t.Errorf("ConfigForBics type 2 \n got: %v \n want %v \n", got, want)
	}
}

func TestConfigForAsrs1(t *testing.T){
	cn := NewController()

	want := map[string]int64{
		"target_register": 2,
		"source_register1": 2,
		"source_register2": 4,
	}
	cn.executeUnit.SetTargetRegister("00010")
	cn.executeUnit.SetSourceRegister1("00010")
	cn.executeUnit.SetSourceRegister2("00100")
	cn.executeUnit.ConfigForAsrs1()

	got := cn.executeUnit.config

	if !reflect.DeepEqual(want, got){
		t.Errorf("ConfigForAsrs type 1 \n got: %v \n want %v \n", got, want)
	}
}

func TestConfigForAsrs2(t *testing.T){
	cn := NewController()

	want := map[string]int64{
		"target_register": 2,
		"source_register1": 2,
		"value": 70,
	}
	cn.executeUnit.SetTargetRegister("00010")
	cn.executeUnit.SetSourceRegister1("00010")
	cn.executeUnit.SetValueInstruction("0000000001000110")
	cn.executeUnit.ConfigForAsrs2()
	got := cn.executeUnit.config

	if !reflect.DeepEqual(want, got){
		t.Errorf("ConfigForAsrs type 2 \n got: %v \n want %v \n", got, want)
	}
}

func TestConfigForLsls1(t *testing.T){
	cn := NewController()

	want := map[string]int64{
		"target_register": 2,
		"source_register1": 2,
		"source_register2": 4,
	}
	cn.executeUnit.SetTargetRegister("00010")
	cn.executeUnit.SetSourceRegister1("00010")
	cn.executeUnit.SetSourceRegister2("00100")
	cn.executeUnit.ConfigForLsls1()

	got := cn.executeUnit.config

	if !reflect.DeepEqual(want, got){
		t.Errorf("ConfigForLsls type 1 \n got: %v \n want %v \n", got, want)
	}
}

func TestConfigForLsls2(t *testing.T){
	cn := NewController()

	want := map[string]int64{
		"target_register": 2,
		"source_register1": 2,
		"value": 70,
	}
	cn.executeUnit.SetTargetRegister("00010")
	cn.executeUnit.SetSourceRegister1("00010")
	cn.executeUnit.SetValueInstruction("0000000001000110")
	cn.executeUnit.ConfigForLsls2()
	got := cn.executeUnit.config

	if !reflect.DeepEqual(want, got){
		t.Errorf("ConfigForLsls type 2 \n got: %v \n want %v \n", got, want)
	}
}

func TestConfigForLsrs1(t *testing.T){
	cn := NewController()

	want := map[string]int64{
		"target_register": 2,
		"source_register1": 2,
		"source_register2": 4,
	}
	cn.executeUnit.SetTargetRegister("00010")
	cn.executeUnit.SetSourceRegister1("00010")
	cn.executeUnit.SetSourceRegister2("00100")
	cn.executeUnit.ConfigForLsrs1()

	got := cn.executeUnit.config

	if !reflect.DeepEqual(want, got){
		t.Errorf("ConfigForLsrs type 1 \n got: %v \n want %v \n", got, want)
	}
}

func TestConfigForLsrs2(t *testing.T){
	cn := NewController()

	want := map[string]int64{
		"target_register": 2,
		"source_register1": 2,
		"value": 70,
	}
	cn.executeUnit.SetTargetRegister("00010")
	cn.executeUnit.SetSourceRegister1("00010")
	cn.executeUnit.SetValueInstruction("0000000001000110")
	cn.executeUnit.ConfigForLsrs2()
	got := cn.executeUnit.config

	if !reflect.DeepEqual(want, got){
		t.Errorf("ConfigForLsrs type 2 \n got: %v \n want %v \n", got, want)
	}
}

func TestConfigForRors1(t *testing.T){
	cn := NewController()

	want := map[string]int64{
		"target_register": 2,
		"source_register1": 2,
		"source_register2": 4,
	}
	cn.executeUnit.SetTargetRegister("00010")
	cn.executeUnit.SetSourceRegister1("00010")
	cn.executeUnit.SetSourceRegister2("00100")
	cn.executeUnit.ConfigForRors1()

	got := cn.executeUnit.config

	if !reflect.DeepEqual(want, got){
		t.Errorf("ConfigForRors type 1 \n got: %v \n want %v \n", got, want)
	}
}

func TestConfigForRors2(t *testing.T){
	cn := NewController()

	want := map[string]int64{
		"target_register": 2,
		"source_register1": 2,
		"value": 70,
	}
	cn.executeUnit.SetTargetRegister("00010")
	cn.executeUnit.SetSourceRegister1("00010")
	cn.executeUnit.SetValueInstruction("0000000001000110")
	cn.executeUnit.ConfigForRors2()
	got := cn.executeUnit.config

	if !reflect.DeepEqual(want, got){
		t.Errorf("ConfigForRors type 2 \n got: %v \n want %v \n", got, want)
	}
}

func TestConfigForCmp1(t *testing.T){
	cn := NewController()

	want := map[string]int64{
		"source_register1": 2,
		"source_register2": 2,
	}

	cn.executeUnit.SetSourceRegister1("00010")
	cn.executeUnit.SetSourceRegister2("00010")
	cn.executeUnit.ConfigForCmp1()
	got := cn.executeUnit.config

	if !reflect.DeepEqual(want, got){
		t.Errorf("ConfigForCmp1 type 1 \n got: %v \n want %v \n", got, want)
	}
}

func TestConfigForCmp2(t *testing.T){
	cn := NewController()

	want := map[string]int64{
		"source_register1": 2,
		"value": 2,
	}
	cn.executeUnit.SetSourceRegister1("00010")
	cn.executeUnit.SetValueInstruction("0000000000000010")
	cn.executeUnit.ConfigForCmp2()
	got := cn.executeUnit.config

	if !reflect.DeepEqual(want, got){
		t.Errorf("ConfigForCmp type 2 \n got: %v \n want %v \n", got, want)
	}
}

func TestConfigForBeq1(t *testing.T){
	cn := NewController()

	want := map[string]int64{
		"source_register1": 2,
	}

	cn.executeUnit.SetSourceRegister1("00010")
	cn.executeUnit.ConfigForBeq1()
	got := cn.executeUnit.config

	if !reflect.DeepEqual(want, got){
		t.Errorf("ConfigForBeq type 1 \n got: %v \n want %v \n", got, want)
	}
}

func TestConfigForBeq2(t *testing.T){
	cn := NewController()

	want := map[string]int64{
		"value": 2,
	}

	cn.executeUnit.SetValueInstruction("0000000000000010")
	cn.executeUnit.ConfigForBeq2()
	got := cn.executeUnit.config

	if !reflect.DeepEqual(want, got){
		t.Errorf("ConfigForBeq type 2 \n got: %v \n want %v \n", got, want)
	}
}

func TestConfigForBne1(t *testing.T){
	cn := NewController()

	want := map[string]int64{
		"source_register1": 2,
	}

	cn.executeUnit.SetSourceRegister1("00010")
	cn.executeUnit.ConfigForBne1()
	got := cn.executeUnit.config

	if !reflect.DeepEqual(want, got){
		t.Errorf("ConfigForBne type 1 \n got: %v \n want %v \n", got, want)
	}
}

func TestConfigForBne2(t *testing.T){
	cn := NewController()

	want := map[string]int64{
		"value": 2,
	}

	cn.executeUnit.SetValueInstruction("0000000000000010")
	cn.executeUnit.ConfigForBne2()
	got := cn.executeUnit.config

	if !reflect.DeepEqual(want, got){
		t.Errorf("ConfigForBne type 2 \n got: %v \n want %v \n", got, want)
	}
}

func TestConfigForBlt1(t *testing.T){
	cn := NewController()

	want := map[string]int64{
		"source_register1": 2,
	}

	cn.executeUnit.SetSourceRegister1("00010")
	cn.executeUnit.ConfigForBlt1()
	got := cn.executeUnit.config

	if !reflect.DeepEqual(want, got){
		t.Errorf("ConfigForBlt type 1 \n got: %v \n want %v \n", got, want)
	}
}

func TestConfigForBlt2(t *testing.T){
	cn := NewController()

	want := map[string]int64{
		"value": 2,
	}

	cn.executeUnit.SetValueInstruction("0000000000000010")
	cn.executeUnit.ConfigForBlt2()
	got := cn.executeUnit.config

	if !reflect.DeepEqual(want, got){
		t.Errorf("ConfigForBlt type 2 \n got: %v \n want %v \n", got, want)
	}
}

func TestConfigForBl1(t *testing.T){
	cn := NewController()

	want := map[string]int64{
		"source_register1": 2,
	}

	cn.executeUnit.SetSourceRegister1("00010")
	cn.executeUnit.ConfigForBl1()
	got := cn.executeUnit.config

	if !reflect.DeepEqual(want, got){
		t.Errorf("ConfigForBl type 1 \n got: %v \n want %v \n", got, want)
	}
}

func TestConfigForBl2(t *testing.T){
	cn := NewController()

	want := map[string]int64{
		"value": 2,
	}

	cn.executeUnit.SetValueInstruction("0000000000000010")
	cn.executeUnit.ConfigForBl2()
	got := cn.executeUnit.config

	if !reflect.DeepEqual(want, got){
		t.Errorf("ConfigForBl type 2 \n got: %v \n want %v \n", got, want)
	}
}

func TestConfigForBx1(t *testing.T){
	cn := NewController()

	want := map[string]int64{
		"source_register1": 2,
	}

	cn.executeUnit.SetSourceRegister1("00010")
	cn.executeUnit.ConfigForBx1()
	got := cn.executeUnit.config

	if !reflect.DeepEqual(want, got){
		t.Errorf("ConfigForBx type 1 \n got: %v \n want %v \n", got, want)
	}
}

func TestConfigForB1(t *testing.T){
	cn := NewController()

	want := map[string]int64{
		"source_register1": 2,
	}

	cn.executeUnit.SetSourceRegister1("00010")
	cn.executeUnit.ConfigForB1()
	got := cn.executeUnit.config

	if !reflect.DeepEqual(want, got){
		t.Errorf("ConfigForB type 1 \n got: %v \n want %v \n", got, want)
	}
}

func TestConfigForB2(t *testing.T){
	cn := NewController()

	want := map[string]int64{
		"value": 2,
	}

	cn.executeUnit.SetValueInstruction("0000000000000010")
	cn.executeUnit.ConfigForB2()
	got := cn.executeUnit.config

	if !reflect.DeepEqual(want, got){
		t.Errorf("ConfigForB type 2 \n got: %v \n want %v \n", got, want)
	}
}

// -------------------------------------------------------------------------------

func TestSetFunctionForConfigInstruction(t *testing.T){
	cn := NewController()

	want := map[string]int64{
		"target_register": 2,
		"source_register1": 2,
		"value": 70,
	}
	cn.executeUnit.SetTargetRegister("00010")
	cn.executeUnit.SetSourceRegister1("00010")
	cn.executeUnit.SetValueInstruction("0000000001000110")
	cn.executeUnit.ConfigInstruction("ANDS", "2")
	got := cn.executeUnit.config
	
	if !reflect.DeepEqual(want, got){
		t.Errorf("ConfigForMovs type 2 \n got: %v \n want %v \n", got, want)
	}
}

func TestSetFunctionForConfigFormat(t *testing.T){
	cn := NewController()

	want := int64(2)
	cn.GetExecuteUnit().SetFunctionForConfigFormat("rd", "00010")
	got := cn.GetExecuteUnit().targetRegisterDec

	if got != want{
		t.Errorf("SetFunctionForConfigFormat \n got: %v \n want %v \n", got, want)
	}
}
