package program

import(
	//"encoding/json"
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
	json_steps map[int]*Step
	max_step_size int
	current_instruction_fetch string
	current_instruction_decode string
	current_instruction_execute string
	current_address_fetch string
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
	simulation.json_steps = make(map[int]*Step)
	simulation.max_step_size = 1000

	return simulation
}

func (s *Simulator) NextStep(){
	s.current_step += 1
}

func (s *Simulator) FetchInstruction(){
	s.NextStep()
	s.current_address_fetch = s.prog.pc.GetHexAddressMemmory()
	s.current_instruction_fetch = s.prog.rom.GetRomList()[s.current_address_fetch].GetAddress()
	current_program_area := s.prog.rom.GetRomList()[s.current_address_fetch].GetAliasField()
	s.prog.controller.ChangeInstructionFetch(s.current_instruction_fetch)
	s.prog.controller.GetFetchUnit().SetProgramArea(current_program_area)
}

func (s *Simulator) DecodeInstruction(){
	s.current_instruction_decode = s.prog.controller.GetFetchUnit().GetInstruction()
	s.prog.controller.ChangeInstructionDecode(s.current_instruction_decode)
	s.prog.controller.GetDecodeUnit().MapInstruction()
	s.prog.controller.GetDecodeUnit().SplitInstruction()
	s.current_mapinstruction = s.prog.controller.GetDecodeUnit().GetInstructionFormat()
}

func (s *Simulator) ExecuteInstruction(){
	s.current_instruction_execute = s.prog.controller.GetDecodeUnit().GetInstruction()
	current_instruction_alias := s.prog.controller.GetDecodeUnit().GetInstructionName()
	s.prog.controller.ChangeInstructionExecute(s.current_instruction_execute)

	s.current_type_instruction[0] = s.prog.controller.GetDecodeUnit().GetInstructionName()
	s.current_type_instruction[1] = s.prog.controller.GetDecodeUnit().GetInstructionType2()

	for k, v := range s.current_mapinstruction{
		s.prog.controller.GetExecuteUnit().SetFunctionForConfigFormat(k, v)
	}
	s.prog.controller.GetExecuteUnit().ConfigInstruction(s.current_type_instruction[0], s.current_type_instruction[1])	

	s.prog.ExecuteFunctionForInstruction(current_instruction_alias, s.current_type_instruction[1])
}

func (s *Simulator) Simulation(){
	// While simulation in execution
	for s.current_step < s.max_step_size || s.current_instruction_fetch != "0x3000"{

	}
}

func (s *Simulator) GetJsonSimulation(){

}
