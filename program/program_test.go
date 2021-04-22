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
	"strconv"
	"encoding/json"
)

var pr Program
var test_program []string = []string{
		"10111000101000000010001100000000", // LDR 5, #70
		"10111000001000000000011100000000", // LDR 1, #8
		"00000100100001010000100000000000", // ADDS 4, 5, 1
		"00001000111001010000100000000000", // SUBS 7, 5, 1
		"11010100100000000000000010000000", // STR 4, #1
		"11010100111000000000000100000000"} // STR 7, #2
var rom_teste *memmory.CodeMemmory

/*
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
*/

//-----------------------------------------------------------------------------------
// Program execution
//-----------------------------------------------------------------------------------
func TestExecuteFunctionForInstruction(t *testing.T){
	rom_teste := memmory.NewCodeMemmory()
	var addressMemmoryCode string
	for count := (4 * 1024); count < ((4 * 1024) + len(test_program)); count++{
		addressMemmoryCode = "0x" + strconv.FormatInt(int64(count), 16)
		rom_teste.AddInstructionField(addressMemmoryCode, test_program[count - (4 * 1024)])
	}
	pr := NewProgram(rom_teste)

	// Set for execute instruction
	pr.controller.GetExecuteUnit().SetTargetRegister("00010")
	pr.controller.GetExecuteUnit().SetAddressInstruction("1000110")

	// LDR
	want := 70
	pr.ExecuteFunctionForInstruction("LDR", "2")
	got := pr.GetRegisterBank().GetRegisterBank()["R2"].GetDecValue()

	if got != want{
		t.Errorf("ExecuteFunctionForInstruction \n got: %v \n want %v \n", got, want)
	}
}

func TestExecuteLdr2(t *testing.T){
	rom_teste := memmory.NewCodeMemmory()
	var addressMemmoryCode string
	for count := (4 * 1024); count < ((4 * 1024) + len(test_program)); count++{
		addressMemmoryCode = "0x" + strconv.FormatInt(int64(count), 16)
		rom_teste.AddInstructionField(addressMemmoryCode, test_program[count - (4 * 1024)])
	}
	pr := NewProgram(rom_teste)

	pr.controller.GetExecuteUnit().SetTargetRegister("00101")
	pr.controller.GetExecuteUnit().SetAddressInstruction("1000110")

	// LDR 5, #70
	want := 70
	pr.ExecuteLdr2()
	got := pr.GetRegisterBank().GetRegisterBank()["R5"].GetDecValue()

	if got != want{
		t.Errorf("ExecuteLdr2 \n got: %v \n want %v \n", got, want)
	}
}

func TestExecuteAdds1(t *testing.T){
	rom_teste := memmory.NewCodeMemmory()
	var addressMemmoryCode string
	for count := (4 * 1024); count < ((4 * 1024) + len(test_program)); count++{
		addressMemmoryCode = "0x" + strconv.FormatInt(int64(count), 16)
		rom_teste.AddInstructionField(addressMemmoryCode, test_program[count - (4 * 1024)])
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

	if got != want{
		t.Errorf("ExecuteAdds1 \n got: %v \n want %v \n", got, want)
	}
}

func TestExecuteSubs1(t *testing.T){
	rom_teste := memmory.NewCodeMemmory()
	var addressMemmoryCode string
	for count := (4 * 1024); count < ((4 * 1024) + len(test_program)); count++{
		addressMemmoryCode = "0x" + strconv.FormatInt(int64(count), 16)
		rom_teste.AddInstructionField(addressMemmoryCode, test_program[count - (4 * 1024)])
	}
	pr := NewProgram(rom_teste)

	pr.registerBank.ChangeRegister("R5", 70)
	pr.registerBank.ChangeRegister("R1", 8)
	pr.controller.GetExecuteUnit().SetTargetRegister("00110")
	pr.controller.GetExecuteUnit().SetSourceRegister1("00101")
	pr.controller.GetExecuteUnit().SetSourceRegister2("00001")

	// SUBS 7, 5, 1
	want := 62
	pr.GetRegisterBank().GetRegisterBank()["R5"].SetValue(70)
	pr.GetRegisterBank().GetRegisterBank()["R1"].SetValue(8)
	pr.ExecuteSubs1()
	got := pr.GetRegisterBank().GetRegisterBank()["R6"].GetDecValue()

	if got != want{
		t.Errorf("ExecuteSubs1 \n got: %v \n want %v \n", got, want)
	}
}

func TestExecuteStr2(t *testing.T){
	rom_teste := memmory.NewCodeMemmory()
	var addressMemmoryCode string
	for count := (4 * 1024); count < ((4 * 1024) + len(test_program)); count++{
		addressMemmoryCode = "0x" + strconv.FormatInt(int64(count), 16)
		rom_teste.AddInstructionField(addressMemmoryCode, test_program[count - (4 * 1024)])
	}
	pr := NewProgram(rom_teste)

	pr.registerBank.ChangeRegister("R4", 50)
	pr.controller.GetExecuteUnit().SetTargetRegister("00100")
	pr.controller.GetExecuteUnit().SetAddressInstruction("001")

	// STR 4, #1
	want := 50
	pr.ExecuteStr2()
	got := pr.GetDataMemmory().GetDataMemmoryList()["0x1"].GetDecValue()

	if got != want{
		t.Errorf("ExecuteStr2 \n got: %v \n want %v \n", got, want)
	}
}

//-----------------------------------------------------------------------------------
// Json generates
//-----------------------------------------------------------------------------------

func TestBuildCodeMemmoryJson(t *testing.T){
	rom_teste := memmory.NewCodeMemmory()
	var addressMemmoryCode string
	for count := (4 * 1024); count < ((4 * 1024) + len(test_program)); count++{
		addressMemmoryCode = "0x" + strconv.FormatInt(int64(count), 16)
		rom_teste.AddInstructionField(addressMemmoryCode, test_program[count - (4 * 1024)])
	}
	pr := NewProgram(rom_teste)

	code_memmory := make(map[int]string)

	for count := 0; count < (12 * 1024); count++{
		addressMemmoryCode = "0x" + strconv.FormatInt(int64(count), 16)
		str_memmory := map[string]string{
			"decimal_value": strconv.FormatInt(int64(rom_teste.GetRomList()[addressMemmoryCode].GetDecValue()), 10),
			"hexadecimal_value": rom_teste.GetRomList()[addressMemmoryCode].GetHexValue(),
			"binary_value": rom_teste.GetRomList()[addressMemmoryCode].GetBinValue(),
			"full_binary": rom_teste.GetRomList()[addressMemmoryCode].GetFullBinValue(),
			"config_type": rom_teste.GetRomList()[addressMemmoryCode].GetConfigType(),
			"memmory_address": rom_teste.GetRomList()[addressMemmoryCode].GetAddress(),
			"alias_field": rom_teste.GetRomList()[addressMemmoryCode].GetAliasField()}

		jmem, _ := json.Marshal(str_memmory)

		code_memmory[count] = string(jmem)
	}

	jcode_memmory, _ := json.Marshal(code_memmory)
	want := string(jcode_memmory)
	got := pr.GetCodeMemmory().GetCodeMemmoryJson()

	if got != want{
		t.Errorf("GetCodeMemmoryJson to Program \n got: %v \n want %v \n", got, want)
	}
}

func TestBuildDataMemmoryJson(t *testing.T){
	rom_teste := memmory.NewCodeMemmory()
	var addressMemmoryCode string
	for count := (4 * 1024); count < ((4 * 1024) + len(test_program)); count++{
		addressMemmoryCode = "0x" + strconv.FormatInt(int64(count), 16)
		rom_teste.AddInstructionField(addressMemmoryCode, test_program[count - (4 * 1024)])
	}
	pr := NewProgram(rom_teste)

	data_memmory_test := memmory.NewDataMemmory()
	data_memmory := make(map[int]string)

	for count := 0; count < (4 * 1024); count++{
		addressMemmoryCode = "0x" + strconv.FormatInt(int64(count), 16)
		str_memmory := map[string]string{
			"decimal_value": strconv.FormatInt(int64(data_memmory_test.GetDataMemmoryList()[addressMemmoryCode].GetDecValue()), 10),
			"hexadecimal_value": data_memmory_test.GetDataMemmoryList()[addressMemmoryCode].GetHexValue(),
			"binary_value": data_memmory_test.GetDataMemmoryList()[addressMemmoryCode].GetBinValue(),
			"full_binary": data_memmory_test.GetDataMemmoryList()[addressMemmoryCode].GetFullBinValue(),
			"config_type": data_memmory_test.GetDataMemmoryList()[addressMemmoryCode].GetConfigType(),
			"memmory_address": data_memmory_test.GetDataMemmoryList()[addressMemmoryCode].GetAddress(),
			"alias_field": data_memmory_test.GetDataMemmoryList()[addressMemmoryCode].GetAliasField()}

		jmem, _ := json.Marshal(str_memmory)

		data_memmory[count] = string(jmem)
	}

	jdata_memmory, _ := json.Marshal(data_memmory)
	want := string(jdata_memmory)
	got := pr.GetDataMemmory().GetDataMemmoryJson()

	if got != want{
		t.Errorf("GetDataMemmoryJson to Program \n got: %v \n want %v \n", got, want)
	}
}

func TestBuildDeviceMemmoryJson(t *testing.T){
	rom_teste := memmory.NewCodeMemmory()
	var addressMemmoryCode string
	for count := (4 * 1024); count < ((4 * 1024) + len(test_program)); count++{
		addressMemmoryCode = "0x" + strconv.FormatInt(int64(count), 16)
		rom_teste.AddInstructionField(addressMemmoryCode, test_program[count - (4 * 1024)])
	}
	pr := NewProgram(rom_teste)

	device_memmory_test := memmory.NewDeviceMemmory()
	device_memmory := make(map[int]string)

	for count := 0; count < (4 * 1024); count++{
		addressMemmoryCode = "0x" + strconv.FormatInt(int64(count), 16)
		str_memmory := map[string]string{
			"decimal_value": strconv.FormatInt(int64(device_memmory_test.GetDeviceMemmoryList()[addressMemmoryCode].GetDecValue()), 10),
			"hexadecimal_value": device_memmory_test.GetDeviceMemmoryList()[addressMemmoryCode].GetHexValue(),
			"binary_value": device_memmory_test.GetDeviceMemmoryList()[addressMemmoryCode].GetBinValue(),
			"full_binary": device_memmory_test.GetDeviceMemmoryList()[addressMemmoryCode].GetFullBinValue(),
			"config_type": device_memmory_test.GetDeviceMemmoryList()[addressMemmoryCode].GetConfigType(),
			"memmory_address": device_memmory_test.GetDeviceMemmoryList()[addressMemmoryCode].GetAddress(),
			"alias_field": device_memmory_test.GetDeviceMemmoryList()[addressMemmoryCode].GetAliasField()}

		jmem, _ := json.Marshal(str_memmory)

		device_memmory[count] = string(jmem)
	}

	jdevice_memmory, _ := json.Marshal(device_memmory)
	want := string(jdevice_memmory)
	got := pr.GetDataMemmory().GetDataMemmoryJson()

	if got != want{
		t.Errorf("GetDeviceMemmoryJson to Program \n got: %v \n want %v \n", got, want)
	}
}

func TestBuildRegisterBankJson(t *testing.T){
	rom_teste := memmory.NewCodeMemmory()
	var addressMemmoryCode string
	for count := (4 * 1024); count < ((4 * 1024) + len(test_program)); count++{
		addressMemmoryCode = "0x" + strconv.FormatInt(int64(count), 16)
		rom_teste.AddInstructionField(addressMemmoryCode, test_program[count - (4 * 1024)])
	}
	pr := NewProgram(rom_teste)

	register_bank_test := registerbank.NewRegisterBank()
	register_bank := make(map[string]string)

	var registerCode string
	for count := 0; count < 17; count++{
		registerCode = "R" + strconv.Itoa(count)
		str_register_bank := map[string]string{
			"decimal_value": strconv.FormatInt(int64(register_bank_test.GetRegisterBank() [registerCode].GetDecValue()), 10),
			"hexadecimal_value": register_bank_test.GetRegisterBank() [registerCode].GetHexValue(),
			"binary_value": register_bank_test.GetRegisterBank() [registerCode].GetBinValue(),
			"register_name": register_bank_test.GetRegisterBank() [registerCode].GetRegisterName(),
			"register_function": register_bank_test.GetRegisterBank() [registerCode].GetRegisterFunction()}

		jrbank, _ := json.Marshal(str_register_bank)

		register_bank[registerCode] = string(jrbank)
	}

	jregister_bank, _ := json.Marshal(register_bank)
	want := string(jregister_bank)
	got := pr.GetRegisterBank().GetRegisterBankJson()
	
	if got != want{
		t.Errorf("GetRegisterBankJson to Program \n got: %v \n want %v \n", got, want)
	}

}