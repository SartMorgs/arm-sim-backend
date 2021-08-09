package program

import (
	"arm/memmory"
	"arm/registerbank"

	//"arm/ula"
	//"arm/programcounter"
	//"arm/controller"
	//"arm/interruption"
	//"arm/io"

	"testing"
	//"reflect"
	"encoding/json"
	"strconv"
)

var pr Program
var test_program []string = []string{
	"11010000101000000010001100000000", // LDR 5, #70
	"11010000001000000000010000000000", // LDR 1, #8
	"00000100100001010000100000000000", // ADDS 4, 5, 1
	"00001000111001010000100000000000", // SUBS 7, 5, 1
	"11010100100000000000000010000000", // STR 4, #1
	"11010100111000000000000100000000"} // STR 7, #2
var rom_teste *memmory.CodeMemmory

func TestInitialConfigurations(t *testing.T) {
	rom_teste := memmory.NewCodeMemmory()
	var addressMemmoryCode string
	for count := (4 * 1024); count < ((4 * 1024) + len(test_program)); count++ {
		addressMemmoryCode = "0x" + strconv.FormatInt(int64(count), 16)
		rom_teste.AddInstructionField(addressMemmoryCode, test_program[count-(4*1024)])
	}
	pr := NewProgram(rom_teste)

	regbk := registerbank.NewRegisterBank()
	regbk.GetRegisterBank()["R13"].SetRegisterName("SP")
	regbk.GetRegisterBank()["R14"].SetRegisterName("LR")
	regbk.GetRegisterBank()["R15"].SetRegisterName("PC")
	regbk.GetRegisterBank()["R16"].SetRegisterName("PSR")

	want := regbk.GetRegisterBankJson()

	pr.initialConfigurations()

	got := pr.registerBank.GetRegisterBankJson()

	if got != want {
		t.Errorf("InitialConfigurations \n got: %v \n want %v \n", got, want)
	}
}

func TestSetPSRRegister(t *testing.T) {
	rom_teste := memmory.NewCodeMemmory()
	var addressMemmoryCode string
	for count := (4 * 1024); count < ((4 * 1024) + len(test_program)); count++ {
		addressMemmoryCode = "0x" + strconv.FormatInt(int64(count), 16)
		rom_teste.AddInstructionField(addressMemmoryCode, test_program[count-(4*1024)])
	}
	pr := NewProgram(rom_teste)

	regbk := registerbank.NewRegisterBank()
	regbk.ChangeRegister("R16", 1073741824)

	want := regbk.GetRegisterBankJson()

	pr.ula.SetValue1(36)
	pr.ula.SetValue2(36)
	pr.ula.Cmp()
	pr.ula.AllResultFlag()
	pr.setPSRRegister()

	got := pr.registerBank.GetRegisterBankJson()

	if got != want {
		t.Errorf("SetPSRRegister \n got: %v \n want %v \n", got, want)
	}
}

//-----------------------------------------------------------------------------------
// Program execution
//-----------------------------------------------------------------------------------
func TestExecuteFunctionForInstruction(t *testing.T) {
	rom_teste := memmory.NewCodeMemmory()
	var addressMemmoryCode string
	for count := (4 * 1024); count < ((4 * 1024) + len(test_program)); count++ {
		addressMemmoryCode = "0x" + strconv.FormatInt(int64(count), 16)
		rom_teste.AddInstructionField(addressMemmoryCode, test_program[count-(4*1024)])
	}
	pr := NewProgram(rom_teste)

	// Set for execute instruction
	pr.controller.GetExecuteUnit().SetTargetRegister("00010")
	pr.controller.GetExecuteUnit().SetAddressInstruction("1000110")

	// LDR
	want := 70
	pr.ExecuteFunctionForInstruction("LDR", "2")
	got := pr.GetRegisterBank().GetRegisterBank()["R2"].GetDecValue()

	if got != want {
		t.Errorf("ExecuteFunctionForInstruction \n got: %v \n want %v \n", got, want)
	}
}

func TestExecuteAdds1(t *testing.T) {
	rom_teste := memmory.NewCodeMemmory()
	var addressMemmoryCode string
	for count := (4 * 1024); count < ((4 * 1024) + len(test_program)); count++ {
		addressMemmoryCode = "0x" + strconv.FormatInt(int64(count), 16)
		rom_teste.AddInstructionField(addressMemmoryCode, test_program[count-(4*1024)])
	}
	pr := NewProgram(rom_teste)

	pr.registerBank.ChangeRegister("R5", 70)
	pr.registerBank.ChangeRegister("R1", 8)
	pr.controller.GetExecuteUnit().SetTargetRegister("00100")
	pr.controller.GetExecuteUnit().SetSourceRegister1("00101")
	pr.controller.GetExecuteUnit().SetSourceRegister2("00001")

	// ADDS 4, 5, 1
	want := 78
	pr.ExecuteAdds1()
	got := pr.GetRegisterBank().GetRegisterBank()["R4"].GetDecValue()

	if got != want {
		t.Errorf("ExecuteAdds1 \n got: %v \n want %v \n", got, want)
	}
}

func TestExecuteAdds2(t *testing.T) {
	rom_teste := memmory.NewCodeMemmory()
	var addressMemmoryCode string
	for count := (4 * 1024); count < ((4 * 1024) + len(test_program)); count++ {
		addressMemmoryCode = "0x" + strconv.FormatInt(int64(count), 16)
		rom_teste.AddInstructionField(addressMemmoryCode, test_program[count-(4*1024)])
	}
	pr := NewProgram(rom_teste)

	pr.registerBank.ChangeRegister("R1", 8)
	pr.controller.GetExecuteUnit().SetTargetRegister("00100")
	pr.controller.GetExecuteUnit().SetSourceRegister1("00001")
	pr.controller.GetExecuteUnit().SetValueInstruction("0000000001000110")

	// ADDS 4, 1, 70
	want := 78
	pr.ExecuteAdds2()
	got := pr.GetRegisterBank().GetRegisterBank()["R4"].GetDecValue()

	if got != want {
		t.Errorf("ExecuteAdds2 \n got: %v \n want %v \n", got, want)
	}
}

func TestExecuteSubs1(t *testing.T) {
	rom_teste := memmory.NewCodeMemmory()
	var addressMemmoryCode string
	for count := (4 * 1024); count < ((4 * 1024) + len(test_program)); count++ {
		addressMemmoryCode = "0x" + strconv.FormatInt(int64(count), 16)
		rom_teste.AddInstructionField(addressMemmoryCode, test_program[count-(4*1024)])
	}
	pr := NewProgram(rom_teste)

	pr.registerBank.ChangeRegister("R5", 70)
	pr.registerBank.ChangeRegister("R1", 8)
	pr.controller.GetExecuteUnit().SetTargetRegister("00110")
	pr.controller.GetExecuteUnit().SetSourceRegister1("00101")
	pr.controller.GetExecuteUnit().SetSourceRegister2("00001")

	// SUBS 6, 5, 1
	want := 62
	pr.ExecuteSubs1()
	got := pr.GetRegisterBank().GetRegisterBank()["R6"].GetDecValue()

	if got != want {
		t.Errorf("ExecuteSubs1 \n got: %v \n want %v \n", got, want)
	}
}

func TestExecuteSubs2(t *testing.T) {
	rom_teste := memmory.NewCodeMemmory()
	var addressMemmoryCode string
	for count := (4 * 1024); count < ((4 * 1024) + len(test_program)); count++ {
		addressMemmoryCode = "0x" + strconv.FormatInt(int64(count), 16)
		rom_teste.AddInstructionField(addressMemmoryCode, test_program[count-(4*1024)])
	}
	pr := NewProgram(rom_teste)

	pr.registerBank.ChangeRegister("R5", 70)
	pr.controller.GetExecuteUnit().SetTargetRegister("00110")
	pr.controller.GetExecuteUnit().SetSourceRegister1("00101")
	pr.controller.GetExecuteUnit().SetValueInstruction("0000000000001000")

	// SUBS 6, 5, 70
	want := 62
	pr.ExecuteSubs2()
	got := pr.GetRegisterBank().GetRegisterBank()["R6"].GetDecValue()

	if got != want {
		t.Errorf("ExecuteSubs2 \n got: %v \n want %v \n", got, want)
	}
}

func TestExecuteMuls1(t *testing.T) {
	rom_teste := memmory.NewCodeMemmory()
	var addressMemmoryCode string
	for count := (4 * 1024); count < ((4 * 1024) + len(test_program)); count++ {
		addressMemmoryCode = "0x" + strconv.FormatInt(int64(count), 16)
		rom_teste.AddInstructionField(addressMemmoryCode, test_program[count-(4*1024)])
	}
	pr := NewProgram(rom_teste)

	pr.registerBank.ChangeRegister("R5", 2)
	pr.registerBank.ChangeRegister("R1", 8)
	pr.controller.GetExecuteUnit().SetTargetRegister("00110")
	pr.controller.GetExecuteUnit().SetSourceRegister1("00101")
	pr.controller.GetExecuteUnit().SetSourceRegister2("00001")

	// MULS 6, 5, 1
	want := 16
	pr.ExecuteMuls1()
	got := pr.GetRegisterBank().GetRegisterBank()["R6"].GetDecValue()

	if got != want {
		t.Errorf("ExecuteMuls1 \n got: %v \n want %v \n", got, want)
	}
}

func TestExecuteMuls2(t *testing.T) {
	rom_teste := memmory.NewCodeMemmory()
	var addressMemmoryCode string
	for count := (4 * 1024); count < ((4 * 1024) + len(test_program)); count++ {
		addressMemmoryCode = "0x" + strconv.FormatInt(int64(count), 16)
		rom_teste.AddInstructionField(addressMemmoryCode, test_program[count-(4*1024)])
	}
	pr := NewProgram(rom_teste)

	pr.registerBank.ChangeRegister("R5", 2)
	pr.controller.GetExecuteUnit().SetTargetRegister("00110")
	pr.controller.GetExecuteUnit().SetSourceRegister1("00101")
	pr.controller.GetExecuteUnit().SetValueInstruction("0000000000001000")

	// MULS 6, 5, 70
	want := 16
	pr.ExecuteMuls2()
	got := pr.GetRegisterBank().GetRegisterBank()["R6"].GetDecValue()

	if got != want {
		t.Errorf("ExecuteMuls2 \n got: %v \n want %v \n", got, want)
	}
}

func TestExecuteAnds1(t *testing.T) {
	rom_teste := memmory.NewCodeMemmory()
	var addressMemmoryCode string
	for count := (4 * 1024); count < ((4 * 1024) + len(test_program)); count++ {
		addressMemmoryCode = "0x" + strconv.FormatInt(int64(count), 16)
		rom_teste.AddInstructionField(addressMemmoryCode, test_program[count-(4*1024)])
	}
	pr := NewProgram(rom_teste)

	pr.registerBank.ChangeRegister("R5", 2)
	pr.registerBank.ChangeRegister("R1", 3)
	pr.controller.GetExecuteUnit().SetTargetRegister("00110")
	pr.controller.GetExecuteUnit().SetSourceRegister1("00101")
	pr.controller.GetExecuteUnit().SetSourceRegister2("00001")

	// ANDS 6, 5, 1
	want := 2
	pr.ExecuteAnds1()
	got := pr.GetRegisterBank().GetRegisterBank()["R6"].GetDecValue()

	if got != want {
		t.Errorf("ExecuteAnds1 \n got: %v \n want %v \n", got, want)
	}
}

func TestExecuteAnds2(t *testing.T) {
	rom_teste := memmory.NewCodeMemmory()
	var addressMemmoryCode string
	for count := (4 * 1024); count < ((4 * 1024) + len(test_program)); count++ {
		addressMemmoryCode = "0x" + strconv.FormatInt(int64(count), 16)
		rom_teste.AddInstructionField(addressMemmoryCode, test_program[count-(4*1024)])
	}
	pr := NewProgram(rom_teste)

	pr.registerBank.ChangeRegister("R5", 2)
	pr.controller.GetExecuteUnit().SetTargetRegister("00110")
	pr.controller.GetExecuteUnit().SetSourceRegister1("00101")
	pr.controller.GetExecuteUnit().SetValueInstruction("0000000000000011")

	// ANDS 6, 5, 70
	want := 2
	pr.ExecuteAnds2()
	got := pr.GetRegisterBank().GetRegisterBank()["R6"].GetDecValue()

	if got != want {
		t.Errorf("ExecuteAnds2 \n got: %v \n want %v \n", got, want)
	}
}

func TestExecuteOrrs1(t *testing.T) {
	rom_teste := memmory.NewCodeMemmory()
	var addressMemmoryCode string
	for count := (4 * 1024); count < ((4 * 1024) + len(test_program)); count++ {
		addressMemmoryCode = "0x" + strconv.FormatInt(int64(count), 16)
		rom_teste.AddInstructionField(addressMemmoryCode, test_program[count-(4*1024)])
	}
	pr := NewProgram(rom_teste)

	pr.registerBank.ChangeRegister("R5", 2)
	pr.registerBank.ChangeRegister("R1", 3)
	pr.controller.GetExecuteUnit().SetTargetRegister("00110")
	pr.controller.GetExecuteUnit().SetSourceRegister1("00101")
	pr.controller.GetExecuteUnit().SetSourceRegister2("00001")

	// ORRS 6, 5, 1
	want := 3
	pr.ExecuteOrrs1()
	got := pr.GetRegisterBank().GetRegisterBank()["R6"].GetDecValue()

	if got != want {
		t.Errorf("ExecuteOrrs1 \n got: %v \n want %v \n", got, want)
	}
}

func TestExecuteOrrs2(t *testing.T) {
	rom_teste := memmory.NewCodeMemmory()
	var addressMemmoryCode string
	for count := (4 * 1024); count < ((4 * 1024) + len(test_program)); count++ {
		addressMemmoryCode = "0x" + strconv.FormatInt(int64(count), 16)
		rom_teste.AddInstructionField(addressMemmoryCode, test_program[count-(4*1024)])
	}
	pr := NewProgram(rom_teste)

	pr.registerBank.ChangeRegister("R5", 2)
	pr.controller.GetExecuteUnit().SetTargetRegister("00110")
	pr.controller.GetExecuteUnit().SetSourceRegister1("00101")
	pr.controller.GetExecuteUnit().SetValueInstruction("0000000000000011")

	// ORRS 6, 5, 70
	want := 3
	pr.ExecuteOrrs2()
	got := pr.GetRegisterBank().GetRegisterBank()["R6"].GetDecValue()

	if got != want {
		t.Errorf("ExecuteOrrs2 \n got: %v \n want %v \n", got, want)
	}
}

func TestExecuteEors1(t *testing.T) {
	rom_teste := memmory.NewCodeMemmory()
	var addressMemmoryCode string
	for count := (4 * 1024); count < ((4 * 1024) + len(test_program)); count++ {
		addressMemmoryCode = "0x" + strconv.FormatInt(int64(count), 16)
		rom_teste.AddInstructionField(addressMemmoryCode, test_program[count-(4*1024)])
	}
	pr := NewProgram(rom_teste)

	pr.registerBank.ChangeRegister("R5", 2)
	pr.registerBank.ChangeRegister("R1", 3)
	pr.controller.GetExecuteUnit().SetTargetRegister("00110")
	pr.controller.GetExecuteUnit().SetSourceRegister1("00101")
	pr.controller.GetExecuteUnit().SetSourceRegister2("00001")

	// EORS 6, 5, 1
	want := 1
	pr.ExecuteEors1()
	got := pr.GetRegisterBank().GetRegisterBank()["R6"].GetDecValue()

	if got != want {
		t.Errorf("ExecuteEors1 \n got: %v \n want %v \n", got, want)
	}
}

func TestExecuteEors2(t *testing.T) {
	rom_teste := memmory.NewCodeMemmory()
	var addressMemmoryCode string
	for count := (4 * 1024); count < ((4 * 1024) + len(test_program)); count++ {
		addressMemmoryCode = "0x" + strconv.FormatInt(int64(count), 16)
		rom_teste.AddInstructionField(addressMemmoryCode, test_program[count-(4*1024)])
	}
	pr := NewProgram(rom_teste)

	pr.registerBank.ChangeRegister("R5", 2)
	pr.controller.GetExecuteUnit().SetTargetRegister("00110")
	pr.controller.GetExecuteUnit().SetSourceRegister1("00101")
	pr.controller.GetExecuteUnit().SetValueInstruction("0000000000000011")

	// EORS 6, 5, 70
	want := 1
	pr.ExecuteEors2()
	got := pr.GetRegisterBank().GetRegisterBank()["R6"].GetDecValue()

	if got != want {
		t.Errorf("ExecuteEors2 \n got: %v \n want %v \n", got, want)
	}
}

func TestExecuteBics1(t *testing.T) {
	rom_teste := memmory.NewCodeMemmory()
	var addressMemmoryCode string
	for count := (4 * 1024); count < ((4 * 1024) + len(test_program)); count++ {
		addressMemmoryCode = "0x" + strconv.FormatInt(int64(count), 16)
		rom_teste.AddInstructionField(addressMemmoryCode, test_program[count-(4*1024)])
	}
	pr := NewProgram(rom_teste)

	pr.registerBank.ChangeRegister("R5", 2)
	pr.registerBank.ChangeRegister("R1", 3)
	pr.controller.GetExecuteUnit().SetTargetRegister("00110")
	pr.controller.GetExecuteUnit().SetSourceRegister1("00101")
	pr.controller.GetExecuteUnit().SetSourceRegister2("00001")

	// BICS 6, 5, 1
	want := 0
	pr.ExecuteBics1()
	got := pr.GetRegisterBank().GetRegisterBank()["R6"].GetDecValue()

	if got != want {
		t.Errorf("ExecuteBics1 \n got: %v \n want %v \n", got, want)
	}
}

func TestExecuteBics2(t *testing.T) {
	rom_teste := memmory.NewCodeMemmory()
	var addressMemmoryCode string
	for count := (4 * 1024); count < ((4 * 1024) + len(test_program)); count++ {
		addressMemmoryCode = "0x" + strconv.FormatInt(int64(count), 16)
		rom_teste.AddInstructionField(addressMemmoryCode, test_program[count-(4*1024)])
	}
	pr := NewProgram(rom_teste)

	pr.registerBank.ChangeRegister("R5", 2)
	pr.controller.GetExecuteUnit().SetTargetRegister("00110")
	pr.controller.GetExecuteUnit().SetSourceRegister1("00101")
	pr.controller.GetExecuteUnit().SetValueInstruction("0000000000000011")

	// BICS 6, 5, 70
	want := 0
	pr.ExecuteBics2()
	got := pr.GetRegisterBank().GetRegisterBank()["R6"].GetDecValue()

	if got != want {
		t.Errorf("ExecuteBics2 \n got: %v \n want %v \n", got, want)
	}
}

func TestExecuteAsrs1(t *testing.T) {
	rom_teste := memmory.NewCodeMemmory()
	var addressMemmoryCode string
	for count := (4 * 1024); count < ((4 * 1024) + len(test_program)); count++ {
		addressMemmoryCode = "0x" + strconv.FormatInt(int64(count), 16)
		rom_teste.AddInstructionField(addressMemmoryCode, test_program[count-(4*1024)])
	}
	pr := NewProgram(rom_teste)

	pr.registerBank.ChangeRegister("R5", 2)
	pr.registerBank.ChangeRegister("R1", 3)
	pr.controller.GetExecuteUnit().SetTargetRegister("00110")
	pr.controller.GetExecuteUnit().SetSourceRegister1("00101")
	pr.controller.GetExecuteUnit().SetSourceRegister2("00001")

	// ASRS 6, 5, 1
	want := 16
	pr.ExecuteAsrs1()
	got := pr.GetRegisterBank().GetRegisterBank()["R6"].GetDecValue()

	if got != want {
		t.Errorf("ExecuteAsrs1 \n got: %v \n want %v \n", got, want)
	}
}

func TestExecuteAsrs2(t *testing.T) {
	rom_teste := memmory.NewCodeMemmory()
	var addressMemmoryCode string
	for count := (4 * 1024); count < ((4 * 1024) + len(test_program)); count++ {
		addressMemmoryCode = "0x" + strconv.FormatInt(int64(count), 16)
		rom_teste.AddInstructionField(addressMemmoryCode, test_program[count-(4*1024)])
	}
	pr := NewProgram(rom_teste)

	pr.registerBank.ChangeRegister("R5", 2)
	pr.controller.GetExecuteUnit().SetTargetRegister("00110")
	pr.controller.GetExecuteUnit().SetSourceRegister1("00101")
	pr.controller.GetExecuteUnit().SetValueInstruction("0000000000000011")

	// ASRS 6, 5, 70
	want := 16
	pr.ExecuteAsrs2()
	got := pr.GetRegisterBank().GetRegisterBank()["R6"].GetDecValue()

	if got != want {
		t.Errorf("ExecuteAsrs2 \n got: %v \n want %v \n", got, want)
	}
}

func TestExecuteLsls1(t *testing.T) {
	rom_teste := memmory.NewCodeMemmory()
	var addressMemmoryCode string
	for count := (4 * 1024); count < ((4 * 1024) + len(test_program)); count++ {
		addressMemmoryCode = "0x" + strconv.FormatInt(int64(count), 16)
		rom_teste.AddInstructionField(addressMemmoryCode, test_program[count-(4*1024)])
	}
	pr := NewProgram(rom_teste)

	pr.registerBank.ChangeRegister("R5", 2)
	pr.registerBank.ChangeRegister("R1", 3)
	pr.controller.GetExecuteUnit().SetTargetRegister("00110")
	pr.controller.GetExecuteUnit().SetSourceRegister1("00101")
	pr.controller.GetExecuteUnit().SetSourceRegister2("00001")

	// LSLS 6, 5, 1
	want := 16
	pr.ExecuteLsls1()
	got := pr.GetRegisterBank().GetRegisterBank()["R6"].GetDecValue()

	if got != want {
		t.Errorf("ExecuteLsls1 \n got: %v \n want %v \n", got, want)
	}
}

func TestExecuteLsls2(t *testing.T) {
	rom_teste := memmory.NewCodeMemmory()
	var addressMemmoryCode string
	for count := (4 * 1024); count < ((4 * 1024) + len(test_program)); count++ {
		addressMemmoryCode = "0x" + strconv.FormatInt(int64(count), 16)
		rom_teste.AddInstructionField(addressMemmoryCode, test_program[count-(4*1024)])
	}
	pr := NewProgram(rom_teste)

	pr.registerBank.ChangeRegister("R5", 2)
	pr.controller.GetExecuteUnit().SetTargetRegister("00110")
	pr.controller.GetExecuteUnit().SetSourceRegister1("00101")
	pr.controller.GetExecuteUnit().SetValueInstruction("0000000000000011")

	// LSLS 6, 5, 70
	want := 16
	pr.ExecuteLsls2()
	got := pr.GetRegisterBank().GetRegisterBank()["R6"].GetDecValue()

	if got != want {
		t.Errorf("ExecuteLsls2 \n got: %v \n want %v \n", got, want)
	}
}

func TestExecuteLsrs1(t *testing.T) {
	rom_teste := memmory.NewCodeMemmory()
	var addressMemmoryCode string
	for count := (4 * 1024); count < ((4 * 1024) + len(test_program)); count++ {
		addressMemmoryCode = "0x" + strconv.FormatInt(int64(count), 16)
		rom_teste.AddInstructionField(addressMemmoryCode, test_program[count-(4*1024)])
	}
	pr := NewProgram(rom_teste)

	pr.registerBank.ChangeRegister("R5", 36)
	pr.registerBank.ChangeRegister("R1", 2)
	pr.controller.GetExecuteUnit().SetTargetRegister("00110")
	pr.controller.GetExecuteUnit().SetSourceRegister1("00101")
	pr.controller.GetExecuteUnit().SetSourceRegister2("00001")

	// LSRS 6, 5, 1
	want := 9
	pr.ExecuteLsrs1()
	got := pr.GetRegisterBank().GetRegisterBank()["R6"].GetDecValue()

	if got != want {
		t.Errorf("ExecuteLsrs1 \n got: %v \n want %v \n", got, want)
	}
}

func TestExecuteLsrs2(t *testing.T) {
	rom_teste := memmory.NewCodeMemmory()
	var addressMemmoryCode string
	for count := (4 * 1024); count < ((4 * 1024) + len(test_program)); count++ {
		addressMemmoryCode = "0x" + strconv.FormatInt(int64(count), 16)
		rom_teste.AddInstructionField(addressMemmoryCode, test_program[count-(4*1024)])
	}
	pr := NewProgram(rom_teste)

	pr.registerBank.ChangeRegister("R5", 36)
	pr.controller.GetExecuteUnit().SetTargetRegister("00110")
	pr.controller.GetExecuteUnit().SetSourceRegister1("00101")
	pr.controller.GetExecuteUnit().SetValueInstruction("0000000000000010")

	// LSRS 6, 5, 70
	want := 9
	pr.ExecuteLsrs2()
	got := pr.GetRegisterBank().GetRegisterBank()["R6"].GetDecValue()

	if got != want {
		t.Errorf("ExecuteLsrs2 \n got: %v \n want %v \n", got, want)
	}
}

func TestExecuteRors1(t *testing.T) {
	rom_teste := memmory.NewCodeMemmory()
	var addressMemmoryCode string
	for count := (4 * 1024); count < ((4 * 1024) + len(test_program)); count++ {
		addressMemmoryCode = "0x" + strconv.FormatInt(int64(count), 16)
		rom_teste.AddInstructionField(addressMemmoryCode, test_program[count-(4*1024)])
	}
	pr := NewProgram(rom_teste)

	pr.registerBank.ChangeRegister("R5", 36)
	pr.registerBank.ChangeRegister("R1", 2)
	pr.controller.GetExecuteUnit().SetTargetRegister("00110")
	pr.controller.GetExecuteUnit().SetSourceRegister1("00101")
	pr.controller.GetExecuteUnit().SetSourceRegister2("00001")

	// RORS 6, 5, 1
	want := 9
	pr.ExecuteRors1()
	got := pr.GetRegisterBank().GetRegisterBank()["R6"].GetDecValue()

	if got != want {
		t.Errorf("ExecuteRors1 \n got: %v \n want %v \n", got, want)
	}
}

func TestExecuteRors2(t *testing.T) {
	rom_teste := memmory.NewCodeMemmory()
	var addressMemmoryCode string
	for count := (4 * 1024); count < ((4 * 1024) + len(test_program)); count++ {
		addressMemmoryCode = "0x" + strconv.FormatInt(int64(count), 16)
		rom_teste.AddInstructionField(addressMemmoryCode, test_program[count-(4*1024)])
	}
	pr := NewProgram(rom_teste)

	pr.registerBank.ChangeRegister("R5", 36)
	pr.controller.GetExecuteUnit().SetTargetRegister("00110")
	pr.controller.GetExecuteUnit().SetSourceRegister1("00101")
	pr.controller.GetExecuteUnit().SetValueInstruction("0000000000000010")

	// RORS 6, 5, 70
	want := 9
	pr.ExecuteRors2()
	got := pr.GetRegisterBank().GetRegisterBank()["R6"].GetDecValue()

	if got != want {
		t.Errorf("ExecuteRors2 \n got: %v \n want %v \n", got, want)
	}
}

func TestExecuteLdr1(t *testing.T) {
	rom_teste := memmory.NewCodeMemmory()
	var addressMemmoryCode string
	for count := (4 * 1024); count < ((4 * 1024) + len(test_program)); count++ {
		addressMemmoryCode = "0x" + strconv.FormatInt(int64(count), 16)
		rom_teste.AddInstructionField(addressMemmoryCode, test_program[count-(4*1024)])
	}

	pr := NewProgram(rom_teste)

	pr.ram.ChangeField("0x1", 70)
	pr.registerBank.ChangeRegister("R1", 1)

	pr.controller.GetExecuteUnit().SetTargetRegister("00000")
	pr.controller.GetExecuteUnit().SetSourceRegister1("00001")

	// LDR R0, R1
	want := 70
	pr.ExecuteLdr1()
	got := pr.GetRegisterBank().GetRegisterBank()["R0"].GetDecValue()

	if got != want {
		t.Errorf("ExecuteLdr1 \n got: %v \n want %v \n", got, want)
	}
}

func TestExecuteLdr2(t *testing.T) {
	rom_teste := memmory.NewCodeMemmory()
	var addressMemmoryCode string
	for count := (4 * 1024); count < ((4 * 1024) + len(test_program)); count++ {
		addressMemmoryCode = "0x" + strconv.FormatInt(int64(count), 16)
		rom_teste.AddInstructionField(addressMemmoryCode, test_program[count-(4*1024)])
	}
	pr := NewProgram(rom_teste)

	pr.controller.GetExecuteUnit().SetTargetRegister("00101")
	pr.controller.GetExecuteUnit().SetAddressInstruction("1000110")

	// LDR R5, #70
	want := 70
	pr.ExecuteLdr2()
	got := pr.GetRegisterBank().GetRegisterBank()["R5"].GetDecValue()

	if got != want {
		t.Errorf("ExecuteLdr2 \n got: %v \n want %v \n", got, want)
	}
}

func TestExecuteStr1(t *testing.T) {
	rom_teste := memmory.NewCodeMemmory()
	var addressMemmoryCode string
	for count := (4 * 1024); count < ((4 * 1024) + len(test_program)); count++ {
		addressMemmoryCode = "0x" + strconv.FormatInt(int64(count), 16)
		rom_teste.AddInstructionField(addressMemmoryCode, test_program[count-(4*1024)])
	}

	pr := NewProgram(rom_teste)

	pr.registerBank.ChangeRegister("R0", 70)
	pr.registerBank.ChangeRegister("R1", 1)
	pr.controller.GetExecuteUnit().SetTargetRegister("00000")
	pr.controller.GetExecuteUnit().SetSourceRegister1("00001")

	// STR R0, R1
	want := 70
	pr.ExecuteStr1()
	got := pr.ram.GetDataMemmoryList()["0x1"].GetDecValue()

	if got != want {
		t.Errorf("ExecuteStr1 \n got: %v \n want %v \n", got, want)
	}
}

func TestExecuteStr2(t *testing.T) {
	rom_teste := memmory.NewCodeMemmory()
	var addressMemmoryCode string
	for count := (4 * 1024); count < ((4 * 1024) + len(test_program)); count++ {
		addressMemmoryCode = "0x" + strconv.FormatInt(int64(count), 16)
		rom_teste.AddInstructionField(addressMemmoryCode, test_program[count-(4*1024)])
	}
	pr := NewProgram(rom_teste)

	pr.registerBank.ChangeRegister("R4", 50)
	pr.controller.GetExecuteUnit().SetTargetRegister("00100")
	pr.controller.GetExecuteUnit().SetAddressInstruction("001")

	// STR 4, #1
	want := 50
	pr.ExecuteStr2()
	got := pr.GetDataMemmory().GetDataMemmoryList()["0x1"].GetDecValue()

	if got != want {
		t.Errorf("ExecuteStr2 \n got: %v \n want %v \n", got, want)
	}
}

//-----------------------------------------------------------------------------------
// Json generates
//-----------------------------------------------------------------------------------

func TestBuildCodeMemmoryJson(t *testing.T) {
	rom_teste := memmory.NewCodeMemmory()
	var addressMemmoryCode string
	for count := (4 * 1024); count < ((4 * 1024) + len(test_program)); count++ {
		addressMemmoryCode = "0x" + strconv.FormatInt(int64(count), 16)
		rom_teste.AddInstructionField(addressMemmoryCode, test_program[count-(4*1024)])
	}
	pr := NewProgram(rom_teste)

	code_memmory := make(map[int]string)

	for count := 0; count < (12 * 1024); count++ {
		addressMemmoryCode = "0x" + strconv.FormatInt(int64(count), 16)
		str_memmory := map[string]string{
			"decimal_value":     strconv.FormatInt(int64(rom_teste.GetRomList()[addressMemmoryCode].GetDecValue()), 10),
			"hexadecimal_value": rom_teste.GetRomList()[addressMemmoryCode].GetHexValue(),
			"binary_value":      rom_teste.GetRomList()[addressMemmoryCode].GetBinValue(),
			"full_binary":       rom_teste.GetRomList()[addressMemmoryCode].GetFullBinValue(),
			"config_type":       rom_teste.GetRomList()[addressMemmoryCode].GetConfigType(),
			"memmory_address":   rom_teste.GetRomList()[addressMemmoryCode].GetAddress(),
			"alias_field":       rom_teste.GetRomList()[addressMemmoryCode].GetAliasField()}

		jmem, _ := json.MarshalIndent(str_memmory, "", "")

		code_memmory[count] = string(jmem)
	}

	jcode_memmory, _ := json.MarshalIndent(code_memmory, "", "")
	want := string(jcode_memmory)
	got := pr.GetCodeMemmory().GetCodeMemmoryJson()

	if got != want {
		t.Errorf("GetCodeMemmoryJson to Program \n got: %v \n want %v \n", got, want)
	}
}

func TestBuildDataMemmoryJson(t *testing.T) {
	rom_teste := memmory.NewCodeMemmory()
	var addressMemmoryCode string
	for count := (4 * 1024); count < ((4 * 1024) + len(test_program)); count++ {
		addressMemmoryCode = "0x" + strconv.FormatInt(int64(count), 16)
		rom_teste.AddInstructionField(addressMemmoryCode, test_program[count-(4*1024)])
	}
	pr := NewProgram(rom_teste)

	data_memmory_test := memmory.NewDataMemmory()
	data_memmory := make(map[int]string)

	for count := 0; count < (4 * 1024); count++ {
		addressMemmoryCode = "0x" + strconv.FormatInt(int64(count), 16)
		str_memmory := map[string]string{
			"decimal_value":     strconv.FormatInt(int64(data_memmory_test.GetDataMemmoryList()[addressMemmoryCode].GetDecValue()), 10),
			"hexadecimal_value": data_memmory_test.GetDataMemmoryList()[addressMemmoryCode].GetHexValue(),
			"binary_value":      data_memmory_test.GetDataMemmoryList()[addressMemmoryCode].GetBinValue(),
			"full_binary":       data_memmory_test.GetDataMemmoryList()[addressMemmoryCode].GetFullBinValue(),
			"config_type":       data_memmory_test.GetDataMemmoryList()[addressMemmoryCode].GetConfigType(),
			"memmory_address":   data_memmory_test.GetDataMemmoryList()[addressMemmoryCode].GetAddress(),
			"alias_field":       data_memmory_test.GetDataMemmoryList()[addressMemmoryCode].GetAliasField()}

		jmem, _ := json.MarshalIndent(str_memmory, "", "")

		data_memmory[count] = string(jmem)
	}

	jdata_memmory, _ := json.MarshalIndent(data_memmory, "", "")
	want := string(jdata_memmory)
	got := pr.GetDataMemmory().GetDataMemmoryJson()

	if got != want {
		t.Errorf("GetDataMemmoryJson to Program \n got: %v \n want %v \n", got, want)
	}
}

func TestBuildDeviceMemmoryJson(t *testing.T) {
	rom_teste := memmory.NewCodeMemmory()
	var addressMemmoryCode string
	for count := (4 * 1024); count < ((4 * 1024) + len(test_program)); count++ {
		addressMemmoryCode = "0x" + strconv.FormatInt(int64(count), 16)
		rom_teste.AddInstructionField(addressMemmoryCode, test_program[count-(4*1024)])
	}
	pr := NewProgram(rom_teste)

	device_memmory_test := memmory.NewDeviceMemmory()
	device_memmory := make(map[int]string)

	for count := 0; count < (4 * 1024); count++ {
		addressMemmoryCode = "0x" + strconv.FormatInt(int64(count), 16)
		str_memmory := map[string]string{
			"decimal_value":     strconv.FormatInt(int64(device_memmory_test.GetDeviceMemmoryList()[addressMemmoryCode].GetDecValue()), 10),
			"hexadecimal_value": device_memmory_test.GetDeviceMemmoryList()[addressMemmoryCode].GetHexValue(),
			"binary_value":      device_memmory_test.GetDeviceMemmoryList()[addressMemmoryCode].GetBinValue(),
			"full_binary":       device_memmory_test.GetDeviceMemmoryList()[addressMemmoryCode].GetFullBinValue(),
			"config_type":       device_memmory_test.GetDeviceMemmoryList()[addressMemmoryCode].GetConfigType(),
			"memmory_address":   device_memmory_test.GetDeviceMemmoryList()[addressMemmoryCode].GetAddress(),
			"alias_field":       device_memmory_test.GetDeviceMemmoryList()[addressMemmoryCode].GetAliasField()}

		jmem, _ := json.MarshalIndent(str_memmory, "", "")

		device_memmory[count] = string(jmem)
	}

	jdevice_memmory, _ := json.MarshalIndent(device_memmory, "", "")
	want := string(jdevice_memmory)
	got := pr.GetDataMemmory().GetDataMemmoryJson()

	if got != want {
		t.Errorf("GetDeviceMemmoryJson to Program \n got: %v \n want %v \n", got, want)
	}
}

func TestBuildRegisterBankJson(t *testing.T) {
	rom_teste := memmory.NewCodeMemmory()
	var addressMemmoryCode string
	for count := (4 * 1024); count < ((4 * 1024) + len(test_program)); count++ {
		addressMemmoryCode = "0x" + strconv.FormatInt(int64(count), 16)
		rom_teste.AddInstructionField(addressMemmoryCode, test_program[count-(4*1024)])
	}
	pr := NewProgram(rom_teste)

	register_bank_test := registerbank.NewRegisterBank()
	register_bank := make(map[string]string)

	var registerCode string
	for count := 0; count < 17; count++ {
		registerCode = "R" + strconv.Itoa(count)
		str_register_bank := map[string]string{
			"decimal_value":     strconv.FormatInt(int64(register_bank_test.GetRegisterBank()[registerCode].GetDecValue()), 10),
			"hexadecimal_value": register_bank_test.GetRegisterBank()[registerCode].GetHexValue(),
			"binary_value":      register_bank_test.GetRegisterBank()[registerCode].GetBinValue(),
			"register_name":     register_bank_test.GetRegisterBank()[registerCode].GetRegisterName(),
			"register_function": register_bank_test.GetRegisterBank()[registerCode].GetRegisterFunction()}

		jrbank, _ := json.MarshalIndent(str_register_bank, "", "")

		register_bank[registerCode] = string(jrbank)
	}

	jregister_bank, _ := json.MarshalIndent(register_bank, "", "")
	want := string(jregister_bank)
	got := pr.GetRegisterBank().GetRegisterBankJson()

	if got != want {
		t.Errorf("GetRegisterBankJson to Program \n got: %v \n want %v \n", got, want)
	}

}
