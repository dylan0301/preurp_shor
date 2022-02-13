package gate_test

import (
	"fmt"
	"preurp/gate"
	"testing"
)

func TestNewZero(t *testing.T) {
	a := [][]complex128{{0, 0}, {0, 0}}
	b := [][]complex128{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}
	c := gate.NewZero(2)
	d := gate.NewZero(3)
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			if c.Data[i][j] != a[i][j] {
				t.Fail()
			}
		}
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if d.Data[i][j] != b[i][j] {
				t.Fail()
			}
		}
	}
}

func TestMul(t *testing.T) {
	a := [][]complex128{{1, 3}, {2, 4}}
	b := [][]complex128{{0, -1, 3}, {1, 4, 0}}
	c := [][]complex128{{3, 11, 3}, {4, 14, 6}}
	d := gate.Mul(a, b)
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			if d[i][j] != c[i][j] {
				t.Fail()
			}
		}
	}
}

func TestTranspose(t *testing.T) {
	a := [][]complex128{{complex(1, 2), complex(3, 4)}, {complex(5, 6), complex(7, 8)}}
	b := [][]complex128{{complex(1, -2), complex(5, -6)}, {complex(3, -4), complex(7, -8)}}
	c := gate.Transpose(a)
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			if b[i][j] != c[i][j] {
				t.Fail()
			}
		}
	}
}

func TestIsUnitary(t *testing.T) {
	a := [][]complex128{{0, complex(0, -1)}, {complex(0, 1), 0}}
	if !gate.IsUnitary(a) {
		t.Fail()
	}
}

func TestNew(t *testing.T) {
	a := [][]complex128{{1, 0, 0, 0}, {0, 1, 0, 0}, {0, 0, 0, 1}, {0, 0, 1, 0}}
	gate.New(a)
}

func TestTensor(t *testing.T) {
	var a, b, c gate.Gate
	a.Data = [][]complex128{{1, 3}, {2, 4}}
	b.Data = [][]complex128{{1, 4}, {2, 5}, {3, 6}}
	c.Data = [][]complex128{{1, 4, 3, 12}, {2, 5, 6, 15}, {3, 6, 9, 18}, {2, 8, 4, 16}, {4, 10, 8, 20}, {6, 12, 12, 24}}
	d := a.Tensor(b)
	for i := range c.Data {
		for j := range c.Data[i] {
			if c.Data[i][j] != d.Data[i][j] {
				t.Fail()
			}
		}
	}
}

func BenchmarkTensor(b *testing.B) {
	var a, e, c gate.Gate
	a.Data = [][]complex128{{1, 3}, {2, 4}}
	e.Data = [][]complex128{{1, 4}, {2, 5}, {3, 6}}
	c.Data = [][]complex128{{1, 4, 3, 12}, {2, 5, 6, 15}, {3, 6, 9, 18}, {2, 8, 4, 16}, {4, 10, 8, 20}, {6, 12, 12, 24}}
	d := a.Tensor(e)
	for i := range c.Data {
		for j := range c.Data[i] {
			if c.Data[i][j] != d.Data[i][j] {
				fmt.Println(i, j)
			}
		}
	}
}
