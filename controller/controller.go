package controller

import(

)

type Controller struct{
	fetchUnit *FetchUnit
	decodeUnit *DecodeUnit
	executeUnit *ExecuteUnit
}

func NewController() *Controller{
	controller := new(Controller)
	controller.fetchUnit = NewFetchUnit()
	controller.decodeUnit = NewDecodeUnit()
	controller.executeUnit = NewExecuteUnit()

	return controller
}

func (c *Controller) ChangeInstructionFetch(inst string){
	c.fetchUnit.instruction = inst
}

func (c *Controller) ChangeInstructionDecode(inst string){
	c.decodeUnit.instructionCode = inst
}