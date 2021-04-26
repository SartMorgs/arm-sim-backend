package program

import(
	"encoding/json"
	"strings"
	"strconv"
	//"fmt"
)

type Step struct{
	step string
	code_memmory string
	data_memmory string
	device_memmory string
	register_bank string
	current_instruction_fetch string
	current_instruction_decode string
	current_instruction_execute string
	current_address_fetch string
}

type Simulator struct{
	prog *Program
	current_step int
	current_mapinstruction map[string]string
	current_type_instruction [2]string
	stepList map[int]*Step
	jsonStepList map[int]string
	max_step_size int
	current_instruction_fetch string
	current_instruction_decode string
	current_instruction_execute string
	current_address_fetch string
	current_address_decode string
	current_address_execute string
}

func NewStep(stp string, cdmem string, dtmem string, dvmem string, rgbank string, finst string, dinst string, einst string, addr string) *Step{
	step := new(Step)
	step.step = stp
	step.code_memmory = cdmem
	step.data_memmory = dtmem
	step.device_memmory = dvmem
	step.register_bank = rgbank 
	step.current_instruction_fetch = finst
	step.current_instruction_decode = dinst 
	step.current_instruction_execute = einst
	step.current_address_fetch = addr

	return step
}

// Step's functions
func (st *Step) SetStepOnStep(step string){
	st.step = step
}

func (st *Step) SetCodeMemmoryOnStep(cd_mem string){
	st.code_memmory = cd_mem
}

func (st *Step) SetDataMemmoryOnStep(data_mem string){
	st.data_memmory = data_mem
}

func (st *Step) SetDeviceMemmoryOnStep(device_mem string){
	st.device_memmory = device_mem
}

func (st *Step) SetRegisterBankOnStep(reg_bank string){
	st.register_bank = reg_bank
}

func (s *Simulator) SetMaxStepSize(max_step int){
	s.max_step_size = max_step
}

func (s *Simulator) GetMaxStepSize() int{
	return s.max_step_size
}

// Realizar simulação
func NewSimulation(prog *Program) *Simulator{
	simulation := new(Simulator)

	simulation.prog = prog
	simulation.current_step = 0
	simulation.current_mapinstruction = make(map[string]string)
	simulation.stepList = make(map[int]*Step)
	simulation.jsonStepList = make(map[int]string)

	simulation.max_step_size = 10

	for count := 0; count < simulation.max_step_size; count++{
		simulation.stepList[count] = NewStep("", "", "", "", "", "", "", "", "")
		simulation.jsonStepList[count] = ""
	}

	return simulation
}

func (s *Simulator) NextStep(){
	s.current_step += 1
}

func (s *Simulator) FetchInstruction(){
	s.NextStep()
	s.current_address_fetch = s.prog.pc.GetHexAddressMemmory()
	s.current_instruction_fetch = s.prog.rom.GetRomList()[s.current_address_fetch].GetFullBinValue()
	current_program_area := s.prog.rom.GetRomList()[s.current_address_fetch].GetAliasField()
	s.prog.controller.ChangeInstructionFetch(s.current_instruction_fetch)
	s.prog.controller.GetFetchUnit().SetProgramArea(current_program_area)
}

func (s *Simulator) DecodeInstruction(){
	s.current_address_decode = s.current_address_fetch
	s.current_instruction_decode = s.prog.controller.GetFetchUnit().GetInstruction()
	s.prog.controller.ChangeInstructionDecode(s.current_instruction_decode)
	s.prog.controller.GetDecodeUnit().SplitInstruction()
	s.current_mapinstruction = s.prog.controller.GetDecodeUnit().GetInstructionFormat()
}

func (s *Simulator) ExecuteInstruction(){
	s.current_address_execute = s.current_address_decode
	s.current_instruction_execute = s.prog.controller.GetDecodeUnit().GetInstruction()
	current_instruction_alias := s.prog.controller.GetDecodeUnit().GetInstructionName()
	flag_address := s.prog.controller.GetDecodeUnit().GetGapAddressFlag()

	s.prog.controller.ChangeInstructionExecute(s.current_instruction_execute)

	s.current_type_instruction[0] = s.prog.controller.GetDecodeUnit().GetInstructionName()
	s.current_type_instruction[1] = s.prog.controller.GetDecodeUnit().GetInstructionType2()

	for k, v := range s.current_mapinstruction{
		s.prog.controller.GetExecuteUnit().SetFunctionForConfigFormat(k, v)
	}
	s.prog.controller.GetExecuteUnit().ConfigInstruction(s.current_type_instruction[0], s.current_type_instruction[1])	

	s.prog.ExecuteFunctionForInstruction(current_instruction_alias, s.current_type_instruction[1])

	// Next Address Program Counter
	if !flag_address{
		s.prog.pc.NextPointer()
		next_address_hex := s.current_address_execute
		next_address_string := strings.Replace(next_address_hex, "0x", "", 1)
		next_address_int, _ := strconv.ParseInt(next_address_string, 16, 64)
		s.prog.pc.SetAddressMemmory(int(next_address_int) + 1)
	}

	s.current_mapinstruction = make(map[string]string)
}

func (s *Simulator) Simulation(){
	// While simulation in execution
	for s.current_step < s.max_step_size{ //|| s.current_instruction_fetch != "0x10"{
		s.FetchInstruction()
		s.DecodeInstruction()
		s.ExecuteInstruction()

		// Set Steps
		s.stepList[s.current_step - 1].step = strconv.FormatInt(int64(s.current_step - 1), 10)
		s.stepList[s.current_step - 1].code_memmory = s.prog.GetCodeMemmory().GetCodeMemmoryJson()
		s.stepList[s.current_step - 1].data_memmory = s.prog.GetDataMemmory().GetDataMemmoryJson()
		s.stepList[s.current_step - 1].device_memmory = s.prog.GetDeviceMemmory().GetDeviceMemmoryJson()
		s.stepList[s.current_step - 1].register_bank = s.prog.GetRegisterBank().GetRegisterBankJson()
		s.stepList[s.current_step - 1].current_instruction_fetch = s.current_instruction_fetch
		s.stepList[s.current_step - 1].current_instruction_decode = s.current_instruction_decode
		s.stepList[s.current_step - 1].current_instruction_execute = s.current_instruction_execute
		s.stepList[s.current_step - 1].current_address_fetch = s.current_address_fetch

		s.jsonStepList[s.current_step - 1] = s.GetStepJson(s.current_step - 1)
	}
}

func (s *Simulator) GetStepJson(step int) string{
	step_for_json := map[string]string{
		"step": s.stepList[step].step,
		"code_memmory": s.stepList[step].code_memmory,
		"data_memmory": s.stepList[step].data_memmory,
		"device_memmory": s.stepList[step].device_memmory,
		"register_bank": s.stepList[step].register_bank,
		"current_instruction_fetch": s.stepList[step].current_instruction_fetch,
		"current_instruction_decode": s.stepList[step].current_instruction_decode,
		"current_instruction_execute": s.stepList[step].current_instruction_execute,
		"current_address_fetch": s.stepList[step].current_address_fetch}

	jstep, _ := json.Marshal(step_for_json)
	return string(jstep)
}

func (s *Simulator) GetJsonSimulation() string{
	jsteps, _ := json.Marshal(s.jsonStepList)
	return string(jsteps)
}
