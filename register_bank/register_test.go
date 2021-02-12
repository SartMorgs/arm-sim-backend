package register_bank

import "testing"

func TestRegisterSet(t *testing.T){
	setValue(7)

	got := getValue()
	want := 7

	if got != want{
		t.Errorf("setValue \n got: %v \n want: \n%v", got, want)
	}
}