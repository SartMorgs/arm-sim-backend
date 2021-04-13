package program

import(
	"encoding/json"
)

type Step struct{
	step string
	code_memmory string
	data_memmory string
	device_memmory string
	register_bank string
	current_instruction string
}

type Simulator struct{
	prog Program
	current_step int
	current_mapinstruction map[string]string
	json_steps map[int]Step
}

// Realizar simulação
func NewSimulation(prog *Program) *Simulator{
	simulation := new(Simulator)

	simulator.prog = prog
	simulation.current_step = 0
	simulation.current_mapinstruction = make(map[string]string)

	return simulation
}

func (s *Simulator) NextStep(){
	s.current_step += 1
}

func (s *Simulator) FetchInstruction(){
	s.NextStep()
	current_address := s.prog.pc.GetHexAddressMemmory()
	current_instruction := s.rom.GetRomList().GetAddress()
	current_program_area := s.rom.GetRomList().GetAliasField()
	s.steps[s.current_step].current_instruction := current_instruction
	s.controller.ChangeInstructionFetch(current_instruction)
	s.controller.GetFetchUnit().SetProgramArea(current_program_area)
}

func (s *Simulator) DecodeInstruction(){
	current_instruction := s.steps[s.current_step].current_instruction
	s.controller.ChangeInstructionDecode(current_instruction)
	s.controller.GetDecodeUnit().MapInstruction()
	s.controller.GetDecodeUnit().SplitInstruction()
	s.controller.GetDecodeUnit().SplitInstruction()
	s.current_mapinstruction = 
}