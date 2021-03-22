package memmory

import "strconv"

type DeviceMemmory struct{
	deviceList map[string]*Memmory
}

func NewDeviceMemmory() *DeviceMemmory{
	devicememmory := new(DeviceMemmory)
	devicememmory.deviceList = make(map[string]*Memmory)

	var addressMemmoryCode string

	for count := 0; count < (4 * 1024); count++{
		addressMemmoryCode = "0x" + strconv.FormatInt(int64(count), 16)
		devicememmory.deviceList[addressMemmoryCode] = NewMemmory(addressMemmoryCode)
	}

	return devicememmory
}

func (dmem *DeviceMemmory) AddDeviceNameOnMemmory(addr string, name string){
	dmem.deviceList[addr].SetAliasField(name)
}

func (dmem *DeviceMemmory) AddDeviceConfig(addr string, config string){
	dmem.deviceList[addr].SetFullBinValue(config)
}

func (dmem *DeviceMemmory) ResetDeviceMemmory(){
	var addressMemmoryCode string
	for count := 0; count < (4 * 1024); count++{
		addressMemmoryCode = "0x" + strconv.FormatInt(int64(count), 16)
		dmem.deviceList[addressMemmoryCode].SetFullBinValue("00000000000000000000000000000000")
	}
}