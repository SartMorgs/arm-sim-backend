  package controller

type SaveStatus struct{
	lastAddressRom string
	registerValues [17]int64
	registerAspr string 
}

func (s *SaveStatus) SetAddressRom(addr string){
	s.lastAddressRom = addr
}