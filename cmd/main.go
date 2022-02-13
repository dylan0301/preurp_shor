package main

import (
	"fmt"
	"preurp/gate"
	"preurp/qsim"
	"preurp/qubit"
)

func main() {
	var c qsim.Circuit
	c.Data = []complex128{0.5, complex(0.5, 0.5), 0, -0.5}
	c.Size = 2
	fmt.Println(c.Measure(0))
	fmt.Println(c)
	qubit.New(c.Data)

	fmt.Println(gate.X)
}
