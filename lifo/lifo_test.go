package Lifo

import (
	"arm/memmory"
	"arm/registerbank"
	"testing"
	//"encoding/json"
)

var syscont SystemContext
var lifo Lifo

func TestSetRegisterBank(t *testing.T) {
	syscont := NewSystemContext()

	rg_bank := registerbank.NewRegisterBank()
	want := rg_bank.GetRegisterBankJson()

	syscont.SetRegisterBank(rg_bank.GetRegisterBankJson())
	got := syscont.registerBank

	if got != want {
		t.Errorf("SetRegisterBank \n got: %v \n want %v \n", got, want)
	}
}

func TestGetRegisterBank(t *testing.T) {
	syscont := NewSystemContext()

	rg_bank := registerbank.NewRegisterBank()
	want := rg_bank.GetRegisterBankJson()

	syscont.registerBank = rg_bank.GetRegisterBankJson()
	got := syscont.GetRegisterBank()

	if got != want {
		t.Errorf("GetRegisterBank \n got: %v \n want %v \n", got, want)
	}
}

func TestSetDataMemmory(t *testing.T) {
	syscont := NewSystemContext()

	ram := memmory.NewDataMemmory()
	want := ram.GetDataMemmoryJson()

	syscont.SetDataMemmory(ram.GetDataMemmoryJson())
	got := syscont.ram

	if got != want {
		t.Errorf("SetDataMemmory \n got: %v \n want %v \n", got, want)
	}
}

func TestGetDataMemmory(t *testing.T) {
	syscont := NewSystemContext()

	ram := memmory.NewDataMemmory()
	want := ram.GetDataMemmoryJson()

	syscont.ram = ram.GetDataMemmoryJson()
	got := syscont.GetDataMemmory()

	if got != want {
		t.Errorf("GetDataMemmory \n got: %v \n want %v \n", got, want)
	}
}

func TestSetDeviceMemmory(t *testing.T) {
	syscont := NewSystemContext()

	devicemem := memmory.NewDeviceMemmory()
	want := devicemem.GetDeviceMemmoryJson()

	syscont.SetDeviceMemmory(devicemem.GetDeviceMemmoryJson())
	got := syscont.deviceMem

	if got != want {
		t.Errorf("SetDeviceMemmory \n got: %v \n want %v \n", got, want)
	}
}

func TestGetDeviceMemmory(t *testing.T) {
	syscont := NewSystemContext()

	devicemem := memmory.NewDeviceMemmory()
	want := devicemem.GetDeviceMemmoryJson()

	syscont.deviceMem = devicemem.GetDeviceMemmoryJson()
	got := syscont.GetDeviceMemmory()

	if got != want {
		t.Errorf("GetDeviceMemmory \n got: %v \n want %v \n", got, want)
	}
}

func TestCheckLifoIsEmpty(t *testing.T) {
	lifo := NewLifo()

	want := true
	got := lifo.IsEmpty()

	if got != want {
		t.Errorf("IsEmpty \n got: %v \n want %v \n", got, want)
	}
}

func TestPushLifo(t *testing.T) {
	lifo := NewLifo()
	rg_bank := registerbank.NewRegisterBank()
	ram := memmory.NewDataMemmory()
	devicemem := memmory.NewDeviceMemmory()

	want_rg_bank := rg_bank.GetRegisterBankJson()
	want_ram := ram.GetDataMemmoryJson()
	want_devicemem := devicemem.GetDeviceMemmoryJson()

	sv := NewSystemContext()
	sv.SetRegisterBank(rg_bank.GetRegisterBankJson())
	sv.SetDataMemmory(ram.GetDataMemmoryJson())
	sv.SetDeviceMemmory(devicemem.GetDeviceMemmoryJson())

	lifo.Push(sv)

	got_rg_bank := lifo.lifo[0].registerBank
	got_ram := lifo.lifo[0].ram
	got_devicemem := lifo.lifo[0].deviceMem

	if (got_rg_bank != want_rg_bank) || (got_ram != want_ram) || (got_devicemem != want_devicemem) {
		t.Errorf("PushLifo \n got_rg_bank: %v \n want_rg_bank %v \n got_ram: %v \n want_ram %v \n got_devicemem: %v \n want_devicemem %v \n", got_rg_bank, want_rg_bank, got_ram, want_ram, got_devicemem, want_devicemem)
	}
}

func TestPopLifo(t *testing.T) {
	lifo := NewLifo()
	rg_bank := registerbank.NewRegisterBank()
	ram := memmory.NewDataMemmory()
	devicemem := memmory.NewDeviceMemmory()

	want_rg_bank := rg_bank.GetRegisterBankJson()
	want_ram := ram.GetDataMemmoryJson()
	want_devicemem := devicemem.GetDeviceMemmoryJson()

	sv := NewSystemContext()
	sv.SetRegisterBank(rg_bank.GetRegisterBankJson())
	sv.SetDataMemmory(ram.GetDataMemmoryJson())
	sv.SetDeviceMemmory(devicemem.GetDeviceMemmoryJson())

	lifo.Push(sv)

	got_rg_bank := ""
	got_ram := ""
	got_devicemem := ""

	sv_test, result := lifo.Pop()

	if result {
		got_rg_bank = sv_test.registerBank
		got_ram = sv_test.ram
		got_devicemem = sv_test.deviceMem
	}

	if (got_rg_bank != want_rg_bank) || (got_ram != want_ram) || (got_devicemem != want_devicemem) {
		t.Errorf("PopLifo \n got_rg_bank: %v \n want_rg_bank %v \n got_ram: %v \n want_ram %v \n got_devicemem: %v \n want_devicemem %v \n", got_rg_bank, want_rg_bank, got_ram, want_ram, got_devicemem, want_devicemem)
	}
}
