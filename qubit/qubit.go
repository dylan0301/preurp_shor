package qubit

import (
	"fmt"
	"math"
	"preurp/useful"
)

type Qubit []complex128

//1 0 0 ... 0 (2^n개) 반환
func Zero(n int) Qubit {
	var size int = 1 << n
	q := make([]complex128, size)
	q[0] = 1
	return q
}

//근데 이렇게 하면 로그 씌웠다가 다시 power하는거라서 시간 오래걸리지않아?
// if useful.Power(2, useful.Log_2(size)) != size {
// 	panic("qubit 크기 2^n꼴 아님")
// }

//유효한 Qubit인가? 그렇다면 Qubit으로 반환
func New(data []complex128) Qubit {
	size := len(data)
	for size > 1 {
		if size%2 == 1 {
			panic("qubit 크기 2^n꼴 아님")
		}
		size /= 2
	}
	var sum float64 = 0
	for i := range data {
		sum += useful.CmplxAbs2(data[i])
	}
	if math.Abs(sum-1) > 1e-5 {
		panic("절댓값 제곱 합이 1이 아님")
	}
	return Qubit(data)
}

//Qubit을 열벡터로 출력
//(예쁘게 보임)
func (q Qubit) String() string {
	s := ""
	for i := range q {
		s += "["
		if real(q[i]) >= 0 {
			s += "+"
		}
		s += fmt.Sprintf("%f ", real(q[i]))
		if imag(q[i]) >= 0 {
			s += "+"
		}
		s += fmt.Sprintf("%f\n", imag(q[i]))
	}
	return s
}

/*
|00>: ~~~
|01>: ~~
..
*/

/*
Qubit
[1+0i]
[0+0i]
[0+0i]
[0+0i]
*/
