package useful_test

import (
	"fmt"
	"preurp/useful"
	"testing"
)

func TestContinuedFrct(t *testing.T) {
	numerator := 13547165517
	denominator := 1638906596546
	a := []int{0, 120, 1, 44, 10, 1, 4, 1, 1, 22, 1, 2, 1, 1, 2, 4, 1, 2, 1, 6, 2, 2, 2, 3}
	b := useful.ContinuedFrct(numerator, denominator)
	fmt.Println(b)
	for i := range a {
		if a[i] != b[i] {
			t.Fail()
		}
	}
}

func TestLog_2(t *testing.T) {
	if useful.Log_2(1024) != 10 {
		t.Fail()
	}
}

func TestPower(t *testing.T) {
	if useful.Power(3, 4) != 81 {
		t.Fail()
	}
}
