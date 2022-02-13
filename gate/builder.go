package gate

import "math"

var h complex128 = complex(1.0/math.Sqrt(2), 0)
var H Gate = New([][]complex128{
	{h, h},
	{h, -h},
})

//test : one => 1
var X Gate = New([][]complex128{
	{0, 1},
	{1, 0},
})

var Y Gate = New([][]complex128{
	{0, complex(0, -1)},
	{complex(0, 1), 0},
})

var Z Gate = New([][]complex128{
	{1, 0},
	{0, complex(0, -1)},
})

//phase shift 360도법으로 입력
func PS(degree float64) Gate {
	theta := math.Pi * degree / 180
	var sp Gate = New([][]complex128{
		{1, 0},
		{0, complex(math.Cos(theta), math.Sin(theta))},
	})
	return sp
}

var CX Gate = New([][]complex128{
	{1, 0, 0, 0},
	{0, 1, 0, 0},
	{0, 0, 0, 1},
	{0, 0, 1, 0},
})

var CZ Gate = New([][]complex128{
	{1, 0, 0, 0},
	{0, 1, 0, 0},
	{0, 0, 1, 0},
	{0, 0, 0, -1},
})

//자료 출처 : 위키피디아
//https://en.wikipedia.org/wiki/Quantum_logic_gate

var SWAP Gate = New([][]complex128{
	{1, 0, 0, 0},
	{0, 0, 1, 0},
	{0, 1, 0, 0},
	{0, 0, 0, 1},
})
