package memmory

import (
	"strconv"
	"encoding/json"
)

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

func (dmem *DeviceMemmory) GetDeviceMemmoryList() map[string]*Memmory{
	return dmem.deviceList
}

func (dmem *DeviceMemmory) GetDeviceMemmoryJson() string{
	device_memmory := make(map[int]string)

	var addressMemmoryCode string
	for count := 0; count < (4 * 1024); count++{
		addressMemmoryCode = "0x" + strconv.FormatInt(int64(count), 16)
		str_memmory := map[string]string{
			"decimal_value": strconv.FormatInt(int64(dmem.deviceList[addressMemmoryCode].GetDecValue()), 10),
			"hexadecimal_value": dmem.deviceList[addressMemmoryCode].GetHexValue(),
			"binary_value": dmem.deviceList[addressMemmoryCode].GetBinValue(),
			"full_binary": dmem.deviceList[addressMemmoryCode].GetFullBinValue(),
			"config_type": dmem.deviceList[addressMemmoryCode].GetConfigType(),
			"memmory_address": dmem.deviceList[addressMemmoryCode].GetAddress(),
			"alias_field": dmem.deviceList[addressMemmoryCode].GetAliasField()}

		jmem, _ := json.Marshal(str_memmory)

		device_memmory[count] = string(jmem)
	}

	jdevice_memmory, _ := json.Marshal(device_memmory)

	return string(jdevice_memmory)
}