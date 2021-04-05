package program

import (
	"arm/memmory"
	"arm/registerbank"
	"arm/ula"
	"arm/programcounter"
	"arm/controller"
	"arm/interruption"
	"arm/io"

	"testing"
	"reflect"
	"strconv"
	//"encoding/json"
)

var pr Program
var test_program []string = []string{
		"10111000101000000010001100000000",
		"10111000001000000010011100000000",
		"00000100100001010000100000000000",
		"00001000111001010000100000000000",
		"11010100100000000000000010000000",
		"11010100111000000000000100000000"}
var rom_teste *memmory.CodeMemmory

func TestInitializeCodeMemmory(t *testing.T){
	rom_teste := memmory.NewCodeMemmory()
	var addressMemmoryCode string
	for count := (4 * 1024); count < ((4 * 1024) + len(test_program)); count++{
		addressMemmoryCode = "0x" + strconv.FormatInt(int64(count), 16)
		rom_teste.AddInstructionField(addressMemmoryCode, test_program[count - (4 * 1024)])
	}
	pr := NewProgram(rom_teste)

	want := memmory.NewCodeMemmory()

	for count := (4 * 1024); count < ((4 * 1024) + len(test_program)); count++{
		addressMemmoryCode = "0x" + strconv.FormatInt(int64(count), 16)
		want.AddInstructionField(addressMemmoryCode, test_program[count - (4 * 1024)])
	}

	got := pr.GetCodeMemmory()

	if !reflect.DeepEqual(want, got){
		t.Errorf("Initialize Code Memmory \n got: %v \n want %v \n", got, want)
	}
}

func TestInitializeDataMemmory(t *testing.T){
	rom_teste := memmory.NewCodeMemmory()
	var addressMemmoryCode string
	for count := (4 * 1024); count < ((4 * 1024) + len(test_program)); count++{
		addressMemmoryCode = "0x" + strconv.FormatInt(int64(count), 16)
		rom_teste.AddInstructionField(addressMemmoryCode, test_program[count - (4 * 1024)])
	}
	pr := NewProgram(rom_teste)

	want := memmory.NewDataMemmory()
	got := pr.GetDataMemmory()

	if !reflect.DeepEqual(want, got){
		t.Errorf("Initialize Data Memmory \n got: %v \n want %v \n", got, want)
	}
}

func TestInitializeDeviceMemmory(t *testing.T){
	rom_teste := memmory.NewCodeMemmory()
	var addressMemmoryCode string
	for count := (4 * 1024); count < ((4 * 1024) + len(test_program)); count++{
		addressMemmoryCode = "0x" + strconv.FormatInt(int64(count), 16)
		rom_teste.AddInstructionField(addressMemmoryCode, test_program[count - (4 * 1024)])
	}
	pr := NewProgram(rom_teste)

	want := memmory.NewDeviceMemmory()
	got := pr.GetDeviceMemmory()

	if !reflect.DeepEqual(want, got){
		t.Errorf("Initialize Device Memmory \n got: %v \n want %v \n", got, want)
	}
}

func TestInitializeRegisterBank(t *testing.T){
	rom_teste := memmory.NewCodeMemmory()
	var addressMemmoryCode string
	for count := (4 * 1024); count < ((4 * 1024) + len(test_program)); count++{
		addressMemmoryCode = "0x" + strconv.FormatInt(int64(count), 16)
		rom_teste.AddInstructionField(addressMemmoryCode, test_program[count - (4 * 1024)])
	}
	pr := NewProgram(rom_teste)

	want := registerbank.NewRegisterBank()
	got := pr.GetRegisterBank()

	if !reflect.DeepEqual(want, got){
		t.Errorf("Initialize Register Bank \n got: %v \n want %v \n", got, want)
	}
}

func TestInitializeUla(t *testing.T){
	rom_teste := memmory.NewCodeMemmory()
	var addressMemmoryCode string
	for count := (4 * 1024); count < ((4 * 1024) + len(test_program)); count++{
		addressMemmoryCode = "0x" + strconv.FormatInt(int64(count), 16)
		rom_teste.AddInstructionField(addressMemmoryCode, test_program[count - (4 * 1024)])
	}
	pr := NewProgram(rom_teste)

	want := ula.NewUla()
	got := pr.GetUla()

	if !reflect.DeepEqual(want, got){
		t.Errorf("Initialize Ula \n got: %v \n want %v \n", got, want)
	}
}

func TestInitializeProgramCounter(t *testing.T){
	rom_teste := memmory.NewCodeMemmory()
	var addressMemmoryCode string
	for count := (4 * 1024); count < ((4 * 1024) + len(test_program)); count++{
		addressMemmoryCode = "0x" + strconv.FormatInt(int64(count), 16)
		rom_teste.AddInstructionField(addressMemmoryCode, test_program[count - (4 * 1024)])
	}
	pr := NewProgram(rom_teste)

	want := programcounter.NewProgramCounter()
	got := pr.GetProgramCounter()

	if !reflect.DeepEqual(want, got){
		t.Errorf("Initialize Program Counter \n got: %v \n want %v \n", got, want)
	}
}

func TestInitializeController(t *testing.T){
	rom_teste := memmory.NewCodeMemmory()
	var addressMemmoryCode string
	for count := (4 * 1024); count < ((4 * 1024) + len(test_program)); count++{
		addressMemmoryCode = "0x" + strconv.FormatInt(int64(count), 16)
		rom_teste.AddInstructionField(addressMemmoryCode, test_program[count - (4 * 1024)])
	}
	pr := NewProgram(rom_teste)

	want := controller.NewController()
	got := pr.GetController()

	if !reflect.DeepEqual(want, got){
		t.Errorf("Initialize Controller \n got: %v \n want %v \n", got, want)
	}
}

func TestInitializeInterruptionBank(t *testing.T){
	rom_teste := memmory.NewCodeMemmory()
	var addressMemmoryCode string
	for count := (4 * 1024); count < ((4 * 1024) + len(test_program)); count++{
		addressMemmoryCode = "0x" + strconv.FormatInt(int64(count), 16)
		rom_teste.AddInstructionField(addressMemmoryCode, test_program[count - (4 * 1024)])
	}
	pr := NewProgram(rom_teste)

	want := interruption.NewInterruptionBank()
	got := pr.GetInterruptionBank()

	if !reflect.DeepEqual(want, got){
		t.Errorf("Initialize Interruption Bank \n got: %v \n want %v \n", got, want)
	}
}
func TestInitializeIo(t *testing.T){
	rom_teste := memmory.NewCodeMemmory()
	var addressMemmoryCode string
	for count := (4 * 1024); count < ((4 * 1024) + len(test_program)); count++{
		addressMemmoryCode = "0x" + strconv.FormatInt(int64(count), 16)
		rom_teste.AddInstructionField(addressMemmoryCode, test_program[count - (4 * 1024)])
	}
	pr := NewProgram(rom_teste)

	want := io.NewDevice("standard", "0x0", "00000000000000000000000000000000")
	got := pr.GetDevice("standard")

	if !reflect.DeepEqual(want, got){
		t.Errorf("Initialize Standard Device \n got: %v \n want %v \n", got, want)
	}
}

//-----------------------------------------------------------------------------------
// Execução do programa
//-----------------------------------------------------------------------------------

func TestBuildCodeMemmoryJson(t *testing.T){
	rom_teste := memmory.NewCodeMemmory()
	var addressMemmoryCode string
	for count := (4 * 1024); count < ((4 * 1024) + len(test_program)); count++{
		addressMemmoryCode = "0x" + strconv.FormatInt(int64(count), 16)
		rom_teste.AddInstructionField(addressMemmoryCode, test_program[count - (4 * 1024)])
	}
	pr := NewProgram(rom_teste)

	for count := (4 * 1024); count < ((4 * 1024) + 10); count++{
		addressMemmoryCode = "0x" + strconv.FormatInt(int64(count), 16)
		print(pr.rom.GetRomList()[addressMemmoryCode].GetDecValue())
	}
}