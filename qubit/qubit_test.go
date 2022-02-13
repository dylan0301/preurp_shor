package qubit_test

import (
	"preurp/qubit"
	"testing"
)

func TestZero(t *testing.T) {
	a := []complex128{1, 0, 0, 0}
	b := qubit.Zero(2)
	for i := 0; i < 4; i++ {
		if a[i] != b[i] {
			t.Fail()
		}
	}
}

func TestNew(t *testing.T) {
	a := []complex128{0.5, complex(0.5, 0.5), 0, -0.5}
	qubit.New(a)
	qubit.New([]complex128{1, 0, 0, 0})
}

// func Benchmark1{}

// func Benchmark2 {
// 	useful.Power(2, useful.Log_2(size))
// }
