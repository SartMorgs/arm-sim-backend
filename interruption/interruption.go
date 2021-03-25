package interruption

import "strconv"

type Interruption struct{
	trigger bool
	number int
	memmoryAddress string
}

func NewInterruption(tg bool, nr int) *Interruption{
	interruption := new(Interruption)
	interruption.trigger = tg
	interruption.number = nr

	var address int
	for countInt := 1; countInt < 17; countInt++{
		address = (countInt - 1) * 256
		if interruption.number == countInt{
			interruption.memmoryAddress = "0x" + strconv.FormatInt(int64(address), 16)
		}
	}

	return interruption
}

func (i *Interruption) TurnOnTrigger(){
	i.trigger = true
}

func (i *Interruption) TurnOffTrigger(){
	i.trigger = false
}

func (i *Interruption) GetTrigger() bool{
	return i.trigger
}

func (i *Interruption) SetNumber(nr int){
	i.number = nr
}

func (i *Interruption) GetNumber() int{
	return i.number
}

func (i *Interruption) SetMemmoryAddress(addr string){
	i.memmoryAddress = addr
}

func (i *Interruption) GetMemmoryAddress() string{
	return i.memmoryAddress
}