package ula
 
type Ula struct {
	value1 int
	value2 int
	result int
}

func (u *Ula) GetValue1() int{
	return u.value1
}

func (u *Ula) GetValue2() int{
	return u.value2
}

func (u *Ula) SetValue1(val1 int){
	u.value1 = val1
}

func (u *Ula) SetValue2(val2 int){
	u.value2 = val2
}

func (u *Ula) Add() int{
	u.result = u.value1 + u.value2
	return u.result
}

func (u *Ula) Sub() int{
	u.result = u.value1 - u.value2
	return u.result
}

func (u *Ula) Mult() int{
	u.result = u.value1 * u.value2 
	return u.result
}