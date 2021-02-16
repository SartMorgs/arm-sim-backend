package ula
 
type Ula struct {
	Value1 int
	Value2 int
	Result int
}

func (u *Ula) GetValue1() int{
	return u.Value1
}

func (u *Ula) GetValue2() int{
	return u.Value2
}

func (u *Ula) SetValue1(value1 int){
	u.Value1 = value1
}

func (u *Ula) SetValue2(value2 int){
	u.Value2 = value2
}

func (u *Ula) Add() int{
	u.Result = u.Value1 + u.Value2
	return u.Result
}

func (u *Ula) Sub() int{
	u.Result = u.Value1 - u.Value2
	return u.Result
}