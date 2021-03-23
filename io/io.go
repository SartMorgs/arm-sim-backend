package io

type Device struct{
	name string
	address string
	config string
}

func NewDevice(nm string, addr string, cfg string) *Device{
	device := new(Device)

	device.name = nm
	device.address = addr
	device.config = cfg

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