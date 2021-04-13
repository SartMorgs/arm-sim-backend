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

func (c *Controller) ChangeInstructionExecute(inst string){
	c.executeUnit.instruction = inst
}

func (c *Controller) GetFetchUnit() *FetchUnit{
	return c.fetchUnit
}

func (c* Controller) GetDecodeUnit() *DecodeUnit{
	return c.decodeUnit
}

func (c *Controller) GetExecuteUnit() *ExecuteUnit{
	return c.executeUnit
}