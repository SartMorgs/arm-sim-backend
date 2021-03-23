package memmory

import(
	"testing"
	"strconv"
)

var dvmm DeviceMemmory

func TestAddDeviceNameOnMemmory(t *testing.T){
	dvmm := NewDeviceMemmory()

	want := "Timer01"
	dvmm.AddDeviceNameOnMemmory("0x55", "Timer01")
	got := dvmm.deviceList["0x55"].GetAliasField()

	if got != want{
		t.Errorf("AddDeviceNameOnMemmory \n got: %v \n want %v \n", got, want)
	}
}

func TestAddDeviceConfig(t *testing.T){
	dvmm := NewDeviceMemmory()

	want := "10001101010100001010100001011010"
	dvmm.AddDeviceConfig("0x55", "10001101010100001010100001011010")
	got := dvmm.deviceList["0x55"].GetFullBinValue()
	
	if got != want{
		t.Errorf("AddDeviceConfig \n got: %v \n want %v \n", got, want)
	}	
}

func TestResetDeviceMemmory(t *testing.T){
	dvmm := NewDeviceMemmory()

	for count := 0; count < (4 * 1024); count++{
		dvmm.AddDeviceConfig("0x" + strconv.FormatInt(int64(count), 16), "00000000000000000000000000000000")
	}

	dvmm.ResetDeviceMemmory()
	want := "00000000000000000000000000000000"
	got := dvmm.deviceList["0x55"].GetFullBinValue()

	if got != want{
		t.Errorf("ResetDeviceMemmory \n got: %v \n want %v \n", got, want)
	}
}