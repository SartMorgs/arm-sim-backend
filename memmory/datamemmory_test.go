package memmory

import(
	"testing"
	"strconv"
	"encoding/json"
)

var dmm DataMemmory

// Unit Tests
// Test all methods using mock
func TestChangeDataMemmoryField(t *testing.T){
	dmm := NewDataMemmory()

	want := 77
	dmm.ChangeField("0x55", 77)
	got := dmm.memmoryList["0x55"].GetDecValue()

	if got != want{
		t.Errorf("ChangeDataMemmoryField \n got: %v \n want %v \n", got, want)
	}
}

func TestResetDataMemmory(t *testing.T){
	dmm := NewDataMemmory()

	for count := 0; count < (4 * 1024); count++{
		dmm.ChangeField("0x" + strconv.FormatInt(int64(count), 16), count)
	}

	dmm.ResetDataMemmory()
	want := 0
	got := dmm.memmoryList["0x55"].GetDecValue()

	if got != want{
		t.Errorf("ResetDataMemmory \n got: %v \n want %v \n", got, want)
	}
}

func TestGetDataMemmoryJson(t *testing.T){
	dmm := NewDataMemmory()
	test := NewDataMemmory()

	data_memmory := make(map[int]string)

	var addressMemmoryCode string
	for count := 0; count < (4 * 1024); count++{
		addressMemmoryCode = "0x" + strconv.FormatInt(int64(count), 16)
		str_memmory := map[string]string{
			"decimal_value": strconv.FormatInt(int64(test.GetDataMemmoryList()[addressMemmoryCode].GetDecValue()), 10),
			"hexadecimal_value": test.GetDataMemmoryList()[addressMemmoryCode].GetHexValue(),
			"binary_value": test.GetDataMemmoryList()[addressMemmoryCode].GetBinValue(),
			"full_binary": test.GetDataMemmoryList()[addressMemmoryCode].GetFullBinValue(),
			"config_type": test.GetDataMemmoryList()[addressMemmoryCode].GetConfigType(),
			"memmory_address": test.GetDataMemmoryList()[addressMemmoryCode].GetAddress(),
			"alias_field": test.GetDataMemmoryList()[addressMemmoryCode].GetAliasField()}
		
		jmem, _ := json.Marshal(str_memmory)

		data_memmory[count] = string(jmem)
	}

	jdata_memmory, _ := json.Marshal(data_memmory)
	want := string(jdata_memmory)
	got := dmm.GetDataMemmoryJson()

	if got != want{
		t.Errorf("GetDataMemmoryJson \n got: %v \n want %v \n", got, want)
	}
}
