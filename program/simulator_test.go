package program

import(
	"arm/memmory"
	"arm/registerbank"

	"strconv"
	"testing"
	"encoding/json"
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

	jstep, _ := json.Marshal(str_step1)

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

func TestRunProgram(t *testing.T){
	rom_teste := memmory.NewCodeMemmory()
	var addressMemmoryCode string
	for count := (4 * 1024); count < ((4 * 1024) + len(test_program)); count++{
		addressMemmoryCode = "0x" + strconv.FormatInt(int64(count), 16)
		rom_teste.AddInstructionField(addressMemmoryCode, test_program[count - (4 * 1024)])
	}
	pr := NewProgram(rom_teste)

	data_memmory := memmory.NewDataMemmory()
	device_memmory := memmory.NewDeviceMemmory()
	register_bank := registerbank.NewRegisterBank()

	data_memmory_values := map[int][2]string{
		5: {"0x1", "4"},
		6: {"0x2", "7"}}

	register_bank_values := map[int][2]string{
		0: {"R5", "70"},
		1: {"R1", "8"},
		2: {"R4", "78"},
		3: {"R7", "62"}}

	step_program := make(map[int]string)

	var data_memmory_value int
	var register_bank_value int
	// Build code memmory for each step
	for count := (4 * 1024); count < (12 * 1024); count++{
		addressMemmoryCode = "0x" + strconv.FormatInt(int64(count - (4 * 1024)), 16)
		data_memmory_value, _ = strconv.Atoi(data_memmory_values[count][1])
		register_bank_value, _ = strconv.Atoi(register_bank_values[count - (4 * 1024)][1])
		data_memmory.ChangeField(data_memmory_values[count][0], data_memmory_value)
		register_bank.ChangeRegister(register_bank_values[count - (4 * 1024)][0], register_bank_value)

		str_steps := map[string]string{
			"step": strconv.FormatInt(int64(count), 10),
			"code_memmory": rom_teste.GetCodeMemmoryJson(),
			"data_memmory": data_memmory.GetDataMemmoryJson(),
			"device_memmory": device_memmory.GetDeviceMemmoryJson(),
			"register_bank": register_bank.GetRegisterBankJson(),
			"current_instruction_fetch": pr.GetController().GetFetchUnit().GetInstruction(),
			"current_instruction_decode": pr.GetController().GetDecodeUnit().GetInstruction(),
			"current_instruction_execute": pr.GetController().GetExecuteUnit().GetInstruction(),
			"current_address_fetch": addressMemmoryCode}

		jstep, _ := json.Marshal(str_steps)

		step_program[count] = string(jstep)
	}

	jprogram, _ := json.Marshal(step_program)
	want := string(jprogram)
	got := sim.GetJsonSimulation()

	if got != want{
		t.Errorf("GetJsonSimulation \n got: \n want \n")
	}
}
