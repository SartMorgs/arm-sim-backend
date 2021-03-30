package controller

import "arm/programcounter"

type FetchUnit struct{
	programcounter.ProgramCounter
	instruction string
	programArea string
}

func NewFetchUnit() *FetchUnit{
	fetchUnit := new(FetchUnit)
	fetchUnit.instruction = ""
	fetchUnit.programArea = "main"

	return fetchUnit
}

func (ft *FetchUnit) SetInstruction(inst string){
	ft.instruction = inst
}

func (ft *FetchUnit) SetProgramArea(prg_area string){
	ft.programArea = prg_area
}

func (ft *FetchUnit) GetInstruction() string{
	return ft.instruction
}

func (ft *FetchUnit) GetProgramArea() string{
	return ft.programArea
}