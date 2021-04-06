package program

import(
	"testing"
)

var sim Simulator

func TestRunProgram(t testing.T){
	rom_teste := memmory.NewCodeMemmory()
	var addressMemmoryCode string
	for count := (4 * 1024); count < ((4 * 1024) + len(test_program)); count++{
		addressMemmoryCode = "0x" + strconv.FormatInt(int64(count), 16)
		rom_teste.AddInstructionField(addressMemmoryCode, test_program[count - (4 * 1024)])
	}
	pr := NewProgram(rom_teste)

	data_memmory_values := map[int][2]string{
		5: {"4", "0x1"},
		6: {"7", "0x2"}}

	register_bank_values := map[int][2]string{
		0: {"5", "70"},
		1: {"1", "8"},
		2: {"4", "78"},
		3: {"7", "62"}}

	step_code_memmory := make(map[int]string)

	// Build code memmory for each step
	for count := (4 * 1024); count < (12 * 1024); count++{
		// Add

		str_steps := map[string]string{
			"step": strconv.FormatInt(int64(count), 10),
			"code_memmory": pr.GetCodeMemmory().GetCodeMemmoryJson(),
			"data_memmory": "",
			"device_memmory": "",
			"register_bank": ""}

	}

}