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

func TestMapInstructionArithmetic(t *testing. T){
	want := "Arithmetic"
	dcun.instruction = "10111001010100001010100001011010"
	dcun.instructionMap = map[string]string{
		"101110": "Arithmetic",
		"100011": "Comparison",
		"000100": "Bypass",
	}
	got := dcun.MapInstruction()

	if got != want{
		t.Errorf("MapInstruction for group1 \n got: %v \n want %v \n", got, want)
	}
}

/*
func TestSplitInstruction(t *testing.T){
	// Instrução aritméticas e lógicas
	// Primeiro caso de teste
	want_opcode1 		:= "101110"
	want_reg1 			:= "01010"
	want_reg2 			:= "10000"
	want_regm			:= "10101"
	dcun.instruction 	= "10111001010100001010100001011010"

	//	Segundo caso de teste
	want_opcode1 		= "101110"
	want_regd 			:= "01010"
	want_regn 			:= "10000"
	dcun.instruction 	= "10111010101000000101000000000001"

	// Instruções de comparação
	// Primeiro caso de teste
	want_opcode1 		= "101110"
	want_regn 			= "01010"
	want_regm 			= "10000"
	dcun.instruction 	= "10111001010100000000100111100000"

	//	Segundo caso de teste
	want_opcode1 		= "101110"
	want_regn 			= "01010"
	want_imediato		:= "10000110"
	dcun.instruction 	= "10111001010100001100110001100000"

	// Instruções de controle e desvio
	// Primeiro caso de teste
	want_opcode1 		= "101110"
	want_regn 			= "01010"
	dcun.instruction 	= "10111001010000000000000111100000"

	// Segundo caso de teste
	want_opcode1 		= "101110"
	dcun.instruction 	= "10111000000000010000000000000001"

}
*/