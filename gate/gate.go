package gate

import (
	"fmt"
	"preurp/useful"
)

type Gate struct {
	Data [][]complex128
}

//유효한 Gate인가? 그렇다면 Gate로 반환
func New(data [][]complex128) Gate {
	l := len(data)
	for i := range data {
		if len(data[i]) != l {
			panic("정사각행렬 아님")
		}
	}

	for l > 1 {
		if l%2 == 1 {
			panic("크기가 2^n꼴이 아님")
		}
		l /= 2
	}

	if !IsUnitary(data) {
		panic("unitary 행렬이 아님")
	}

	var g Gate
	g.Data = data
	return g
}

//n * n짜리 0행렬 게이트 만들기
func NewZero(n int) Gate {
	var g Gate
	g.Data = make([][]complex128, n)
	for i := range g.Data {
		g.Data[i] = make([]complex128, n)
	}
	return g
}

//행렬 a b 곱하기
func Mul(a [][]complex128, b [][]complex128) [][]complex128 {
	m := make([][]complex128, len(a))
	for i := range m {
		m[i] = make([]complex128, len(b[0]))
		for j := range m[i] {
			for k := range b {
				m[i][j] += a[i][k] * b[k][j]
			}
		}
	}
	return m
}

//행렬 a의 켤레 전치 행렬 게이트
func Transpose(a [][]complex128) [][]complex128 {
	t := NewZero(len(a)).Data
	for i := range a {
		for j := range a {
			t[j][i] = complex(real(a[i][j]), -1*imag(a[i][j]))
		}
	}
	return t
}

//어떤 행렬이 unitary인지 판단
//전치행렬=역행렬일때 unitary
func IsUnitary(a [][]complex128) bool {
	size := len(a)
	t := Transpose(a)
	u := Mul(a, t)

	var b complex128
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if i == j {
				b = 1
			} else {
				b = 0
			}
			if useful.CmplxAbs2(u[i][j]-b) > 1e-6 {
				return false
			}
		}
	}
	return true
}

// A.Tensor(B) -> Gate A 텐서곱 Gate B
func (a Gate) Tensor(b Gate) Gate {
	var g Gate
	g.Data = make([][]complex128, len(a.Data)*len(b.Data))
	for i := range a.Data {
		for j := range b.Data {
			g.Data[i*len(b.Data)+j] = make([]complex128, len(a.Data[i])*len(b.Data[j]))
		}
	}
	for ai := range a.Data {
		for aj := range a.Data[0] {
			for bi := range b.Data {
				for bj := range b.Data[0] {
					g.Data[ai*len(b.Data)+bi][aj*len(b.Data[0])+bj] = a.Data[ai][aj] * b.Data[bi][bj]
				}
			}
		}
	}
	return g
}

//Gate를 행렬로 출력
//(예쁘게 보임)
func (g Gate) String() string {
	s := ""
	for i := range g.Data {
		s += "["
		for j := range g.Data[0] {
			if real(g.Data[i][j]) >= 0 {
				s += "+"
			}
			s += fmt.Sprintf("%f", real(g.Data[i][j]))
			if imag(g.Data[i][j]) >= 0 {
				s += "+"
			}
			s += fmt.Sprintf("%fi ", imag(g.Data[i][j]))
		}
		s += "]\n"
	}
	return s
}

/*
게이트 출력방식
[1+0i 0+0i 0+0i 0+0i]
[0+0i 1+0i 0+0i 0+0i]
[0+0i 0+0i 1+0i 0+0i]
[0+0i 0+0i 0+0i 1+0i]
*/
