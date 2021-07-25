package program

import(
	"arm/memmory"
	"arm/registerbank"
	"arm/controller"

	"strconv"
	"testing"
	"encoding/json"
	//"fmt"
)

var sim Simulator
var st Step

// ---------------------------------------------------------------------------------------
// Tests about steps
// ---------------------------------------------------------------------------------------
func TestSetStepOnStep(t *testing.T){
	want := "1"
	st.SetStepOnStep("1")
	got := st.step 

	if want != got{
		t.Errorf("SetStepOnStep \n got: %v \n want %v \n", got, want)
	}
}

func TestSetCodeMemmoryOnStep(t *testing.T){
	want := "teste"
	st.SetCodeMemmoryOnStep("teste")
	got := st.code_memmory

	if want != got{
		t.Errorf("SetCodeMemmoryOnStep \n got: %v \n want %v \n", got, want)
	}
}

func TestSetDataMemmoryOnStep(t *testing.T){
	want := "teste"
	st.SetDataMemmoryOnStep("teste")
	got := st.data_memmory

	if want != got{
		t.Errorf("SetDataMemmoryOnStep \n got: %v \n want %v \n", got, want)
	}
}

func TestSetDeviceMemmoryOnStep(t *testing.T){
	want := "teste"
	st.SetDeviceMemmoryOnStep("teste")
	got := st.device_memmory

	if want != got{
		t.Errorf("SetDeviceMemmoryOnStep \n got: %v \n want %v \n", got, want)
	}
}

func TestSetRegisterBankOnStep(t *testing.T){
	want := "teste"
	st.SetRegisterBankOnStep("teste")
	got := st.register_bank

	if want != got{
		t.Errorf("SetRegisterBankOnStep \n got: %v \n want %v \n", got, want)
	}
}
// ---------------------------------------------------------------------------------------
// Tests about simulation
// ---------------------------------------------------------------------------------------

func TestSetMaxStepSize(t *testing.T){
	want := 2000
	sim.SetMaxStepSize(2000)
	got := sim.max_step_size

	if want != got{
		t.Errorf("SetMaxStepSize \n got: %v \n want %v \n", got, want)
	}
}

func TestGetMaxStepSize(t *testing.T){
	want := 2000
	sim.max_step_size = 2000
	got := sim.GetMaxStepSize()

	if want != got{
		t.Errorf("GetMaxStepSize \n got: %v \n want %v \n", got, want)
	}
}

func TestGetStepJson(t *testing.T){
	rom_teste := memmory.NewCodeMemmory()
	var addressMemmoryCode string
	for count := (4 * 1024); count < ((4 * 1024) + len(test_program)); count++{
		addressMemmoryCode = "0x" + strconv.FormatInt(int64(count), 16)
		rom_teste.AddInstructionField(addressMemmoryCode, test_program[count - (4 * 1024)])
	}
	pr := NewProgram(rom_teste)
	sim := NewSimulation(pr)

	data_memmory := memmory.NewDataMemmory()
	device_memmory := memmory.NewDeviceMemmory()
	register_bank := registerbank.NewRegisterBank()

	str_step1 := map[string]string{
		"step": "2",
		"code_memmory": rom_teste.GetCodeMemmoryJson(),
		"data_memmory": data_memmory.GetDataMemmoryJson(),
		"device_memmory": device_memmory.GetDeviceMemmoryJson(),
		"register_bank": register_bank.GetRegisterBankJson(),
		"current_instruction_fetch": "10010001010101011010100001011010",
		"current_instruction_decode": "00110010000101011010100001011010",
		"current_instruction_execute": "10110110101001001110100001011010",
		"current_address_fetch": "0x2"}

	jstep, _ := json.MarshalIndent(str_step1, "", "")

	want := string(jstep)

	sim.stepList[2] = NewStep(
		"2",
		rom_teste.GetCodeMemmoryJson(),
		data_memmory.GetDataMemmoryJson(),
		device_memmory.GetDeviceMemmoryJson(),
		register_bank.GetRegisterBankJson(),
		"10010001010101011010100001011010",
		"00110010000101011010100001011010",
		"10110110101001001110100001011010",
		"0x2")

	got := sim.GetStepJson(2)

	if got != want{
		t.Errorf("GetStepJson \n got: %v \n want %v \n", got, want)
	}
}

func TestRegisterBankRunProgram(t *testing.T){

	rom_teste := memmory.NewCodeMemmory()
	var addressMemmoryCode string
	for count := (4 * 1024); count < ((4 * 1024) + len(test_program)); count++{
		addressMemmoryCode = "0x" + strconv.FormatInt(int64(count), 16)
		rom_teste.AddInstructionField(addressMemmoryCode, test_program[count - (4 * 1024)])
		rom_teste.AddAliasMemmoryField(addressMemmoryCode, "Program Area")
	}

	pr := NewProgram(rom_teste)
	sim := NewSimulation(pr)

	register_bank := registerbank.NewRegisterBank()

	register_bank_values := map[int][2]string{
		0: {"R5", "70"},
		1: {"R1", "8"},
		2: {"R4", "78"},
		3: {"R7", "62"}}

	var register_bank_value int
	for count := (4 * 1024); count < ((4 * 1024) + 10); count++{
		addressMemmoryCode = "0x" + strconv.FormatInt(int64(count), 16)

		if regbank_value, ok := register_bank_values[count - (4 * 1024)]; ok {
			register_bank_value, _ = strconv.Atoi(regbank_value[1])
			register_bank.ChangeRegister(regbank_value[0], register_bank_value)
		}
	}

	want := register_bank.GetRegisterBankJson()
	sim.SetMaxStepSize(10)
	sim.Simulation()
	got := sim.prog.GetRegisterBank().GetRegisterBankJson()

	if got != want{
		t.Errorf("RegisterBankRunProgram \n got: %v \n want %v \n", len(got), len(want))
	}
}

func TestCodeMemmoryRunProgram(t *testing.T){
	rom_teste := memmory.NewCodeMemmory()
	var addressMemmoryCode string
	for count := (4 * 1024); count < ((4 * 1024) + len(test_program)); count++{
		addressMemmoryCode = "0x" + strconv.FormatInt(int64(count), 16)
		rom_teste.AddInstructionField(addressMemmoryCode, test_program[count - (4 * 1024)])
		rom_teste.AddAliasMemmoryField(addressMemmoryCode, "Program Area")
	}

	pr := NewProgram(rom_teste)
	sim := NewSimulation(pr)

	want := rom_teste.GetCodeMemmoryJson()
	sim.Simulation()
	got := sim.prog.GetCodeMemmory().GetCodeMemmoryJson()

	if got != want{
		t.Errorf("CodeMemmoryRunProgram \n got: %v \n want %v \n", len(got), len(want))
	}
}

func TestDataMemmoryRunPorgram(t *testing.T){
	rom_teste := memmory.NewCodeMemmory()
	var addressMemmoryCode string
	for count := (4 * 1024); count < ((4 * 1024) + len(test_program)); count++{
		addressMemmoryCode = "0x" + strconv.FormatInt(int64(count), 16)
		rom_teste.AddInstructionField(addressMemmoryCode, test_program[count - (4 * 1024)])
		rom_teste.AddAliasMemmoryField(addressMemmoryCode, "Program Area")
	}

	pr := NewProgram(rom_teste)
	sim := NewSimulation(pr)

	data_memmory := memmory.NewDataMemmory()

	data_memmory_values := map[int][2]string{
		5: {"0x1", "78"},
		6: {"0x2", "62"}}

	var data_memmory_value int
	for count := (4 * 1024); count < ((4 * 1024) + 10); count++{
		addressMemmoryCode = "0x" + strconv.FormatInt(int64(count), 16)

		if dtmem_value, ok := data_memmory_values[count - (4 * 1024)]; ok {
			data_memmory_value, _ = strconv.Atoi(dtmem_value[1])
			data_memmory.ChangeField(dtmem_value[0], data_memmory_value)
		}
	}

	want := data_memmory.GetDataMemmoryJson()
	sim.SetMaxStepSize(10)
	sim.Simulation()
	got := sim.prog.GetDataMemmory().GetDataMemmoryJson()

	if got != want{
		t.Errorf("DataMemmoryRunPorgram \n got: %v \n want %v \n", len(got), len(want))
	}
}

func TestDeviceMemmoryRunPorgram(t *testing.T){
	rom_teste := memmory.NewCodeMemmory()
	var addressMemmoryCode string
	for count := (4 * 1024); count < ((4 * 1024) + len(test_program)); count++{
		addressMemmoryCode = "0x" + strconv.FormatInt(int64(count), 16)
		rom_teste.AddInstructionField(addressMemmoryCode, test_program[count - (4 * 1024)])
		rom_teste.AddAliasMemmoryField(addressMemmoryCode, "Program Area")
	}

	pr := NewProgram(rom_teste)
	sim := NewSimulation(pr)

	device_memmory := memmory.NewDeviceMemmory()

	want := device_memmory.GetDeviceMemmoryJson()
	sim.Simulation()
	got := sim.prog.GetDeviceMemmory().GetDeviceMemmoryJson()

	if got != want{
		t.Errorf("DeviceMemmoryRunPorgram \n got: %v \n want %v \n", len(got), len(want))
	}
}

func TestRunProgram(t *testing.T){
	rom_teste := memmory.NewCodeMemmory()
	var addressMemmoryCode string
	for count := (4 * 1024); count < ((4 * 1024) + len(test_program)); count++{
		addressMemmoryCode = "0x" + strconv.FormatInt(int64(count), 16)
		rom_teste.AddInstructionField(addressMemmoryCode, test_program[count - (4 * 1024)])
		rom_teste.AddAliasMemmoryField(addressMemmoryCode, "Program Area")
	}
	pr := NewProgram(rom_teste)
	sim := NewSimulation(pr)

	data_memmory := memmory.NewDataMemmory()
	device_memmory := memmory.NewDeviceMemmory()
	register_bank := registerbank.NewRegisterBank()
	controller := controller.NewController()

	data_memmory_values := map[int][2]string{
		4: {"0x1", "78"},
		5: {"0x2", "62"}}

	register_bank_values := map[int][2]string{
		0: {"R5", "70"},
		1: {"R1", "8"},
		2: {"R4", "78"},
		3: {"R7", "62"}}

	instruction_values := map[int][3]string{
		0: {"11010000101000000010001100000000", "11010000101000000010001100000000", "11010000101000000010001100000000"},
		1: {"11010000001000000000010000000000", "11010000001000000000010000000000", "11010000001000000000010000000000"},
		2: {"00000100100001010000100000000000", "00000100100001010000100000000000", "00000100100001010000100000000000"},
		3: {"00001000111001010000100000000000", "00001000111001010000100000000000", "00001000111001010000100000000000"},
		4: {"11010100100000000000000010000000", "11010100100000000000000010000000", "11010100100000000000000010000000"},
		5: {"11010100111000000000000100000000", "11010100111000000000000100000000", "11010100111000000000000100000000"}}

	step_program := make(map[int]string)

	var data_memmory_value int
	var register_bank_value int
	// Build code memmory for each step
	for count := (4 * 1024); count < ((4 * 1024) + 10); count++{
		addressMemmoryCode = "0x" + strconv.FormatInt(int64(count), 16)

		if dtmem_value, ok := data_memmory_values[count - (4 * 1024)]; ok {
			data_memmory_value, _ = strconv.Atoi(dtmem_value[1])
			data_memmory.ChangeField(dtmem_value[0], data_memmory_value)
		}

		if regbank_value, ok := register_bank_values[count - (4 * 1024)]; ok {
			register_bank_value, _ = strconv.Atoi(regbank_value[1])
			register_bank.ChangeRegister(regbank_value[0], register_bank_value)
		}

		if instruction_value, ok := instruction_values[count - (4 * 1024)]; ok {
			controller.ChangeInstructionFetch(instruction_value[0])
			controller.ChangeInstructionDecode(instruction_value[1])
			controller.ChangeInstructionExecute(instruction_value[2])
		} else{
			controller.ChangeInstructionFetch("00000000000000000000000000000000")
			controller.ChangeInstructionDecode("00000000000000000000000000000000")
			controller.ChangeInstructionExecute("00000000000000000000000000000000")
		}

		str_steps := map[string]string{
			"step": strconv.FormatInt(int64(count - (4 * 1024)), 10),
			"code_memmory": rom_teste.GetCodeMemmoryJson(),
			"data_memmory": data_memmory.GetDataMemmoryJson(),
			"device_memmory": device_memmory.GetDeviceMemmoryJson(),
			"register_bank": register_bank.GetRegisterBankJson(),
			"current_instruction_fetch": controller.GetFetchUnit().GetInstruction(),
			"current_instruction_decode": controller.GetDecodeUnit().GetInstruction(),
			"current_instruction_execute": controller.GetExecuteUnit().GetInstruction(),
			"current_address_fetch": addressMemmoryCode}

		jstep, _ := json.MarshalIndent(str_steps, "", "")

		step_program[count - (4 * 1024)] = string(jstep)
	}

	jprogram, _ := json.MarshalIndent(step_program, "", "")
	want := string(jprogram)
	sim.current_step = 0
	sim.SetMaxStepSize(10)
	sim.Simulation()
	got := sim.GetJsonSimulation()

	if want != got{
		t.Errorf("GetJsonSimulation \n got: %v \n want %v \n", len(got), len(want))
	}
	
}