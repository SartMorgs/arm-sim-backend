package controller

import "testing"

var dcun DecodeUnit

func TestGetOpcode(t *testing.T){
	want := "111111"
	dcun.SetInstruction("1111110000000000000000000000000000")
	dcun.SplitInstruction()
	got := dcun.GetOpcode()

	if want != got{
		t.Errorf("GetOpcode \n got: %v \n want %v \n", got, want)
	}
}

func TestSplitInstruction(t *testing.T){
	// Instrução aritméticas e lógicas
	// Primeiro caso de teste
	want_opcode1 	:= "101110"
	want_reg1 		:= "01010"
	want_reg2 		:= "10000"
	want_regm		:= "10101"

	//	Segundo caso de teste
	want_opcode1 	:= "101110"
	want_regd 		:= "01010"
	want_regn 		:= "10000"

	// Instruções de comparação
	// Primeiro caso de teste
	want_opcode1 	:= "101110"
	want_regn 		:= "01010"
	want_regm 		:= "10000"

	//	Segundo caso de teste
	want_opcode1 	:= "101110"
	want_regn 		:= "01010"
	want_imediato	:= "10000110"

	// Instruções de controle e desvio
	// Primeiro caso de teste
	want_opcode1 	:= "101110"
	want_regn 		:= "01010"

	// Segundo caso de teste
	want_opcode1 	:= "101110"

}