package program

import(
	"arm/memmory"
	"arm/registerbank"

	"strconv"
	"testing"
	"encoding/json"
)

var sim Simulator

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
			"register_bank": register_bank.GetRegisterBankJson()}

		jstep, _ := json.Marshal(str_steps)

		step_program[count] = string(jstep)
	}

	jprogram, _ := json.Marshal(step_program)
	want := string(jprogram)
	got := sim.GetJsonSimulation()

	if got != want{
		t.Errorf("GetJsonSimulation \n got: %v \n want %v \n", got, want)
	}
}