package ula

import (
	"strconv"
	"testing"
)

var ula1 Ula

// Unit Tests
// Test all methods using mock
func TestSetValue1(t *testing.T) {
	want := 20
	ula1.SetValue1(20)
	got := ula1.value1

	if got != want {
		t.Errorf("SetValue1 \n got: %v \n want %v \n", got, want)
	}
}

func TestSetValue2(t *testing.T) {
	want := 5
	ula1.SetValue2(5)
	got := ula1.value2

	if got != want {
		t.Errorf("SetValue2 \n got: %v \n want %v \n", got, want)
	}
}

func TestGetValue1(t *testing.T) {
	want := 20
	ula1.value1 = 20
	got := ula1.GetValue1()

	if got != want {
		t.Errorf("GetValue1 \n got: %v \n want %v \n", got, want)
	}
}

func TestGetValue2(t *testing.T) {
	want := 20
	ula1.value2 = 20
	got := ula1.GetValue2()

	if got != want {
		t.Errorf("GetValue2 \n got: %v \n want %v \n", got, want)
	}
}

func TestAdd(t *testing.T) {
	want := 25

	ula1.value1 = 20
	ula1.value2 = 5

	got := ula1.Add()

	if got != want {
		t.Errorf("Add \n got: %v \n want %v \n", got, want)
	}
}

func TestSub(t *testing.T) {
	want := 15

	ula1.value1 = 20
	ula1.value2 = 5

	got := ula1.Sub()

	if got != want {
		t.Errorf("Sub \n got: %v \n want %v \n", got, want)
	}
}

func TestMult(t *testing.T) {
	want := 100

	ula1.value1 = 20
	ula1.value2 = 5

	got := ula1.Mult()

	if got != want {
		t.Errorf("Mult \n got: %v \n want %v \n", got, want)
	}
}

func TestAnd(t *testing.T) {
	want := 2

	ula1.value1 = 2
	ula1.value2 = 3

	got := ula1.And()

	if got != want {
		t.Errorf("And \n got: %v \n want %v \n", got, want)
	}
}

func TestOr(t *testing.T) {
	want := 3

	ula1.value1 = 2
	ula1.value2 = 3

	got := ula1.Or()

	if got != want {
		t.Errorf("Or \n got: %v \n want %v \n", got, want)
	}
}

func TestEor(t *testing.T) {
	want := 1

	ula1.value1 = 2
	ula1.value2 = 3

	got := ula1.Eor()

	if got != want {
		t.Errorf("Eor \n got: %v \n want %v \n", got, want)
	}
}

func TestBic(t *testing.T) {
	want := 0

	ula1.value1 = 2
	ula1.value2 = 3

	got := ula1.Bic()

	if got != want {
		t.Errorf("Bic \n got: %v \n want %v \n", got, want)
	}
}

func TestAsr(t *testing.T) {
	want := 16

	ula1.value1 = 2
	ula1.value2 = 3

	got := ula1.Lsl()

	if got != want {
		t.Errorf("Lsl \n got: %v \n want %v \n", got, want)
	}
}

func TestLsl(t *testing.T) {
	want := 16

	ula1.value1 = 2
	ula1.value2 = 3

	got := ula1.Lsl()

	if got != want {
		t.Errorf("Lsl \n got: %v \n want %v \n", got, want)
	}
}

func TestLsr(t *testing.T) {
	want := 9

	ula1.value1 = 36
	ula1.value2 = 2

	got := ula1.Lsr()

	if got != want {
		t.Errorf("Lsr \n got: %v \n want %v \n", got, want)
	}
}

func TestRor(t *testing.T) {
	want := 9

	ula1.value1 = 36
	ula1.value2 = 2

	got := ula1.Ror()

	if got != want {
		t.Errorf("Ror \n got: %v \n want %v \n", got, want)
	}
}

func TestCmp(t *testing.T) {
	ula1.value1 = 36
	ula1.value2 = 36

	ula1.Cmp()
	want := (ula1.result == 0)

	ula1.allResultFlag()
	got := ula1.zeroResultFlag

	if got != want {
		t.Errorf("Cmp \n got: %v \n want %v \n", got, want)
	}
}

func TestCmn(t *testing.T) {
	ula1.value1 = -36
	ula1.value2 = -36

	ula1.Cmn()
	want := (ula1.result == 0)

	ula1.allResultFlag()
	got := ula1.zeroResultFlag

	if got != want {
		t.Errorf("Cmn \n got: %v \n want %v \n", got, want)
	}
}

// ------------------------------------------------------------------------------
// REGISTRATOR OF ULA STATE

func TestNegativeResultFlag(t *testing.T) {
	ula1.value1 = 2
	ula1.value2 = 36

	result := ula1.Sub()
	want := (result < 0)

	ula1.negativeResult()
	got := ula1.negativeResultFlag

	if got != want {
		t.Errorf("NegativeResultFlag \n got: %v \n want %v \n", got, want)
	}
}

func TestZeroResultFlag(t *testing.T) {
	ula1.value1 = 36
	ula1.value2 = 36

	result := ula1.Sub()
	want := (result == 0)

	ula1.zeroResult()
	got := ula1.zeroResultFlag

	if got != want {
		t.Errorf("ZeroResultFlag \n got: %v \n want %v \n", got, want)
	}
}

func TestCarryResultFlag(t *testing.T) {
	// 1111
	ula1.value1 = 15
	ula1.value2 = 1

	result := ula1.Add()

	len_value1 := len("1111")
	len_result := len(strconv.FormatInt(int64(result), 2))

	want := (len_value1 > len_result)

	ula1.carryResult()
	got := ula1.carryResultFlag

	if got != want {
		t.Errorf("CarryResultFlag \n got: %v \n want %v \n", got, want)
	}
}

func TestOverflowResultFlag(t *testing.T) {
	ula1.value1 = 2147483647
	ula1.value2 = 3

	result := ula1.Add()

	want := (result > 2147483647)

	ula1.overflowResult()
	got := ula1.overflowResultFlag

	if got != want {
		t.Errorf("OverflowResultFlag \n got: %v \n want %v \n", got, want)
	}
}

func TestAllResultFlag(t *testing.T) {
	ula1.value1 = 36
	ula1.value2 = 36

	len_value1 := len("100100")
	len_result := len("100100")

	result := ula1.Sub()
	want_negative := (result < 0)
	want_zero := (result == 0)
	want_carry := (len_value1 > len_result)
	want_overflow := (result > 2147483647)

	ula1.allResultFlag()
	got_negative := ula1.negativeResultFlag
	got_zero := ula1.zeroResultFlag
	got_carry := ula1.carryResultFlag
	got_overflow := ula1.overflowResultFlag

	if (got_negative != want_negative) && (got_zero != want_zero) && (got_carry != want_carry) && (got_overflow != want_overflow) {
		t.Errorf("AllResultFlag \n got_negative: %v \n want_negative %v \n got_zero %v \n want_zero %v \n got_carry %v \n want_carry %v \n got_overflow %v \n want_overflow %v \n", got_negative, want_negative, got_zero, want_zero, got_carry, want_carry, got_overflow, want_overflow)
	}
}
