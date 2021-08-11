package ula

import "strconv"

type Ula struct {
	value1 int
	value2 int
	result int

	negativeResultFlag bool
	zeroResultFlag     bool
	carryResultFlag    bool
	overflowResultFlag bool
}

func NewUla() *Ula {
	ula := new(Ula)

	ula.value1 = 0
	ula.value2 = 0
	ula.result = 0

	ula.negativeResultFlag = false
	ula.zeroResultFlag = false
	ula.carryResultFlag = false
	ula.overflowResultFlag = false

	return ula
}

func (u *Ula) GetValue1() int {
	return u.value1
}

func (u *Ula) GetValue2() int {
	return u.value2
}

func (u *Ula) GetNegativeResultFlag() bool {
	return u.negativeResultFlag
}

func (u *Ula) GetZeroResultFlag() bool {
	return u.zeroResultFlag
}

func (u *Ula) GetCarryResultFlag() bool {
	return u.carryResultFlag
}

func (u *Ula) GetOverflowResultFlag() bool {
	return u.overflowResultFlag
}

func (u *Ula) SetValue1(val1 int) {
	u.value1 = val1
}

func (u *Ula) SetValue2(val2 int) {
	u.value2 = val2
}

func (u *Ula) Add() int {
	u.result = u.value1 + u.value2
	return u.result
}

func (u *Ula) Sub() int {
	u.result = u.value1 - u.value2
	return u.result
}

func (u *Ula) Mult() int {
	u.result = u.value1 * u.value2
	return u.result
}

func (u *Ula) And() int {
	u.result = u.value1 & u.value2
	return u.result
}

func (u *Ula) Or() int {
	u.result = u.value1 | u.value2
	return u.result
}

func (u *Ula) Eor() int {
	u.result = u.value1 ^ u.value2
	return u.result
}

func (u *Ula) Bic() int {
	u.result = u.value1 &^ u.value2
	return u.result
}

func (u *Ula) Asr() int {
	u.result = u.value1 << u.value2
	return u.result
}

func (u *Ula) Lsl() int {
	u.result = u.value1 << u.value2
	return u.result
}

func (u *Ula) Lsr() int {
	u.result = u.value1 >> u.value2
	return u.result
}

func (u *Ula) Ror() int {
	u.result = u.value1 >> u.value2
	return u.result
}

func (u *Ula) Cmp() {
	u.result = u.value1 - u.value2
}

func (u *Ula) Cmn() {
	u.result = u.value1 - u.value2
}

func (u *Ula) negativeResult() {
	u.negativeResultFlag = (u.result < 0)
}

func (u *Ula) zeroResult() {
	u.zeroResultFlag = (u.result == 0)
}

func (u *Ula) carryResult() {
	len_value1 := len(strconv.FormatInt(int64(u.value1), 2))
	len_result := len(strconv.FormatInt(int64(u.result), 2))
	u.carryResultFlag = (len_result > len_value1)
}

func (u *Ula) overflowResult() {
	const max_size_word = 2147483647

	u.overflowResultFlag = (u.result > max_size_word)
}

func (u *Ula) AllResultFlag() {
	u.negativeResult()
	u.zeroResult()
	u.carryResult()
	u.overflowResult()
}
