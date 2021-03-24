package io

import "strconv"

type Device struct{
	name string
	address string
	config string
	dType string
	decValue int
	binValue string
	hexValue string
}

func NewDevice(nm string, addr string, cfg string) *Device{
	device := new(Device)

	device.name = nm
	device.address = addr
	device.config = cfg
	device.dType = ""
	device.decValue = 0
	device.binValue = "00"
	device.hexValue = "0x0"

	return device
}

func (d *Device) SetName(nm string){
	d.name = nm
}

func (d *Device) GetName() string{
	return d.name
}

func (d *Device) SetAddress(addr string){
	d.address = addr
}

func (d *Device) GetAddress() string{
	return d.address
}

func (d *Device) SetConfig(cfg string){
	d.config = cfg
}

func (d *Device) GetConfig() string{
	return d.config
}

func (d *Device) SetType(t string){
	d.dType = t
}

func (d *Device) GetType() string{
	return d.dType
}

func (d *Device) SetValue(value int){
	d.decValue = value
	d.toBinValue()
	d.toHexValue()
}

func (d *Device) GetDecValue() int{
	return d.decValue
}

func (d *Device) GetBinValue() string{
	return d.binValue
}

func (d *Device) GetHexValue() string{
	return d.hexValue
}

func (d *Device) toBinValue() {
	d.binValue = strconv.FormatInt(int64(d.decValue), 2)
}

func (d *Device) toHexValue(){
	d.hexValue = "0x" + strconv.FormatInt(int64(d.decValue), 16)
}