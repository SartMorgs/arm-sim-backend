package interruption

type InterruptionBank struct{
	interruptionList map[int]*Interruption
}

func NewInterruptionBank() *InterruptionBank{
	interruptionbank := new(InterruptionBank)
	interruptionbank.interruptionList = make(map[int]*Interruption)

	for count := 1; count < 17; count++{
		interruptionbank.interruptionList[count] = NewInterruption(false, count)
	}

	return interruptionbank
}

func (ibank *InterruptionBank) TriggerInterruption(intnum int){
	ibank.interruptionList[intnum].TurnOnTrigger()
}

func (ibank *InterruptionBank) TurnOffInterruption(intnum int){
	ibank.interruptionList[intnum].TurnOffTrigger()
}