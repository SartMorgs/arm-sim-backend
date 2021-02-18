package programcounter

type ProgramCounter struct{
	decAddressMemmory int 
	hexAddressMemmory int
}

func (pc *ProgramCounter) SetAddressMemmory(addr int){
	pc.decAddressMemmory = addr
}

func (pc *ProgramCounter) GetHexAddressMemmory() int{
	return pc.hexAddressMemmory
}

func (pc *ProgramCounter) decAddressToHexAddress(){

}