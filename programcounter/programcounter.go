package programcounter

import(
	"strconv"
)

type ProgramCounter struct{
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

func (pc *ProgramCounter) toHexAddress(){
	pc.hexAddressMemmory = "0x" + strconv.FormatInt(int64(pc.decAddressMemmory), 16)
}

func (pc *ProgramCounter) toBinAddress(){
	pc.binAddressMemmory = strconv.FormatInt(int64(pc.decAddressMemmory), 2)
}