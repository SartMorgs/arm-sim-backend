package lifo

type SystemContext struct {
	registerBank string
	ram          string
	deviceMem    string
}

type Lifo struct {
	lifo []*SystemContext
}

func NewSystemContext() *SystemContext {
	syscont := new(SystemContext)

	syscont.registerBank = ""
	syscont.ram = ""
	syscont.deviceMem = ""

	return syscont
}

func NewLifo() *Lifo {
	lf := new(Lifo)

	lf.lifo = nil

	return lf
}

func (sc *SystemContext) SetRegisterBank(rgbk string) {
	sc.registerBank = rgbk
}

func (sc *SystemContext) GetRegisterBank() string {
	return sc.registerBank
}

func (sc *SystemContext) SetDataMemmory(ram string) {
	sc.ram = ram
}

func (sc *SystemContext) GetDataMemmory() string {
	return sc.ram
}

func (sc *SystemContext) SetDeviceMemmory(devicemem string) {
	sc.deviceMem = devicemem
}

func (sc *SystemContext) GetDeviceMemmory() string {
	return sc.deviceMem
}

func (lf *Lifo) GetLifo() []*SystemContext {
	return lf.lifo
}

func (lf *Lifo) IsEmpty() bool {
	return len(lf.lifo) == 0
}

func (lf *Lifo) Push(sc *SystemContext) {
	lf.lifo = append(lf.lifo, sc)
}

func (lf *Lifo) Pop() (*SystemContext, bool) {
	if lf.IsEmpty() {
		return nil, false
	} else {
		index := len(lf.lifo) - 1
		element := (lf.lifo)[index]
		lf.lifo = (lf.lifo)[:index]
		return element, true
	}
}
