package qsim_test

import (
	"fmt"
	"math"
	"preurp/gate"
	"preurp/qsim"
	"testing"
)

func TestMeasure(t *testing.T) {
	var c qsim.Circuit
	c.Data = []complex128{0.5, 0.5, complex(0, 0.5), -0.5}
	c.Size = 2
	fmt.Println(c.Measure(1))
	fmt.Println(c)
	qsim.New(c.Data)
}

func TestEntangle(t *testing.T) {
	h := complex(1.0/math.Sqrt(2), 0)
	c := qsim.New([]complex128{1, 0, 0, 0})
	c.Apply(gate.H, 1)
	c.Apply(gate.CX, 0)
	// 1 / sqrt(2) |00> + 1/sqrt(2) |11>
	d := qsim.New([]complex128{h, 0, 0, h})
	for i := range c.Data {
		if c.Data[i] != d.Data[i] {
			t.Fail()
		}
	}
}
