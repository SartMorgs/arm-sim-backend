package interruption

type Interruption struct{
	trigger bool
	systemSaved string
}

func NewInterruption(tg bool) *Interruption{
	interruption := new(Interruption)
	interruption.trigger = tg

	return interruption
}

func (i *Interruption) TurnOnTrigger(){
	i.trigger = true
}

func (i *Interruption) TurnOffTrigger(){
	i.trigger = false
}

func (i *Interruption) SaveSystemStatus(sv string){
	i.systemSaved = sv
}