package programcounter

import(
	"strconv"
)

type ProgramCounter struct{
	currentPointer int
	decAddressMemmory int
	hexAddressMemmory string
	binAddressMemmory string
}

func (pc *ProgramCounter) SetAddressMemmory(addr int){
	pc.decAddressMemmory = addr
}

func (pc *ProgramCounter) GetDecAddressMemmory() int{
	return pc.decAddressMemmory
}

func (pc *ProgramCounter) GetHexAddressMemmory() string{
	return pc.hexAddressMemmory
}

func (pc *ProgramCounter) GetBinAddressMemmory() string{
	return pc.binAddressMemmory
}

func (pc *ProgramCounter) GetPointerPC() int{
	return pc.currentPointer
}

func (pc *ProgramCounter) toHexAddress(){
	pc.hexAddressMemmory = "0x" + strconv.FormatInt(int64(pc.decAddressMemmory), 16)
}

func (pc *ProgramCounter) toBinAddress(){
	pc.binAddressMemmory = strconv.FormatInt(int64(pc.decAddressMemmory), 2)
}

func (pc *ProgramCounter) ResetPC(){
	pc.currentPointer = 0;
}

func (pc *ProgramCounter) nextPointer(){
	pc.currentPointer += 1
}