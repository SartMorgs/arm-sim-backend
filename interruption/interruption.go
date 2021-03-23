package interruption

type Interruption struct{
	trigger bool
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