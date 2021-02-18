package ula

import "testing"

var ula1 Ula

func TestSetValue1 (t *testing.T){
	want := 20

	ula1.SetValue1(20)

	got := ula1.GetValue1()

	if got != want{
		t.Errorf("SetValue1 \n got: %v \n want %v \n", got, want)
	}
}

func TestSetValue2 (t *testing.T){
	want := 5

	ula1.SetValue2(5)

	got := ula1.GetValue2()

	if got != want{
		t.Errorf("SetValue2 \n got: %v \n want %v \n", got, want)
	}
}

func TestAdd (t *testing.T) {
	want := 25

	ula1.SetValue1(20)
	ula1.SetValue2(5)

	got := ula1.Add()

	if got != want{
		t.Errorf("Add \n got: %v \n want %v \n",got, want)
	}
}

func TestSub(t *testing.T){
	want := 15

	ula1.SetValue1(20)
	ula1.SetValue2(5)

	got := ula1.Sub()

	if got != want{
		t.Errorf("Sub \n got: %v \n want %v \n", got, want)
	}
}

func TestMult(t *testing.T){
	want := 100

	ula1.SetValue1(20)
	ula1.SetValue2(5)

	got := ula1.Mult()

	if got != want{
		t.Errorf("Mult \n got: %v \n want %v \n", got, want)
	}
}