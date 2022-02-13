package qsim

import (
	"math/rand"
	"preurp/gate"
	"preurp/useful"
)

//3n qubit짜리 circuit, 000000001111 상태
func StartShor(n int) Circuit {
	data := make([]complex128, 3*n)
	data[1<<n-1] = 1
	c := New(data)
	return c
}

//앞쪽 2n에다가 hadamard 적용
func (c *Circuit) ApplyHadamard() {
	for i := 0; i < c.Size/3*2; i++ {
		c.Apply(gate.H, i)
	}
}

//shor oracle 적용
func (c *Circuit) ApplyShororacle() {

}

//QFT

//연분수 근사로 주기 알아내기
func FindperiodbyContinuedFrct(c Circuit, a, n int) int {
	d := 1
	b := 2
	confrct := useful.ContinuedFrct(d, b)
	L := len(confrct)
	var last int
	for i := 0; i < L; i++ {
		if confrct[i] > (1 << n) {
			last = i - 1
			break
		}
	}
	s := 1             //분자
	m := confrct[last] //분모
	for i := last - 1; i >= 0; i++ {
		m = m*confrct[i] + s
		s = confrct[i+1]
	}
	return m
}

//이거 N=pq입력 받아서 p,q랑 리턴하는거 아님?
func Factorization(n int) int {
	a := rand.Intn(n-3) + 2
	g := useful.GCD(a, n)
	if g != 1 {
		return g
	}
	r := findPeriodByModPow(n, a)

	hp := 1 << 2 / r

	g1 := useful.GCD(N)
}

func findPeriodByModPow(n, a int) int {
	for x := 1; x < n; x++ {
		if useful.ModPow(a, x, n) == 1 {
			return x
		}
	}
	return -1
}
