package qsim

import (
	"fmt"
	"math"
	"math/rand"
	"preurp/gate"
	"preurp/qubit"
	"preurp/useful"
	"time"
)

//아예 Circuit 자체를 열벡터
type Circuit struct {
	Size int
	Data qubit.Qubit
	//a0ㅣ000> + a1ㅣ001> + a2ㅣ010> + a3ㅣ011> + a4ㅣ100> + a5ㅣ101> + a6ㅣ110> + a7ㅣ111> 열벡터
	//Size=3
	//Data=[a0, a1, a2, a3, a4, a5, a6, a7]
}

//유효한 Circuit인가? 그렇다면 Circuit으로 바꿔서 반환
func New(data []complex128) Circuit {
	var c Circuit
	c.Data = qubit.New(data)
	c.Size = useful.Log_2(len(data))
	return c
}

//예를들어 hadamard gate, 1/루트2=h라고 하면 ㅣ0>은 hㅣ0>+hㅣ1>=H의 0열이 되고, ㅣ1>은 hㅣ0>-hㅣ1>=H의 1열이 되는데 그러면 하다마드를 0번째 큐빗에다 적용한다고 치자
//a0ㅣ000> + a1ㅣ001> + a2ㅣ010> + a3ㅣ011> + a4ㅣ100> + a5ㅣ101> + a6ㅣ110> + a7ㅣ111>
//a0ㅣ00>Hㅣ0> + a1ㅣ00>Hㅣ1> + a2ㅣ01>Hㅣ0> + a3ㅣ01>Hㅣ1> + a4ㅣ10>Hㅣ0> + a5ㅣ10>Hㅣ1> + a6ㅣ11>Hㅣ0> + a7ㅣ11>Hㅣ1>
//CNOT gate같이 2 qubit짜리는 그냥ㅣ00>,ㅣ01>,ㅣ10>,ㅣ11>이 어떻게 변하는지 알면 됨
//Circuit c의 a번째 큐빗부터 g의 size 만큼 쭉 gate 적용
func (c *Circuit) Apply(g gate.Gate, a int) {
	temp := make([]complex128, len(c.Data))

	//i=0110100에서 a=2, gSize=3이면 101을 가져와야함. i >> a 하고 2^gsize로 나눈 나머지 =101=part
	//gSize는 Size만큼의 qubit에 적용하는거라서 그 자체로 2^n꼴임
	//part에다가 게이트 적용, ㅣ101>에 gate를 적용하면 그냥 gate의 101번째 열을 가져오면 됨
	//101에 gate 적용했더니 p0ㅣ000> + p1ㅣ001> + p2ㅣ010> + p3ㅣ011> + p4ㅣ100> + p5ㅣ101> + p6ㅣ110> + p7ㅣ111>가 됨, pj=gate[j][101]이 됨
	//temp.Data[01"001"00]+=c.Data[i]*p001
	//0100100구하는법: i=0110100에서 ((i>>a)-part+j)<<a
	for i := range c.Data {
		part := (i >> a) % len(g.Data)
		for j := range g.Data {
			temp[((i>>a)-part+j)<<a] += c.Data[i] * g.Data[j][part]
		}
	}
	for i := range c.Data {
		c.Data[i] = temp[i]
	}
}

//김지원 여기다가해봄 g는 size 상관없는 gate
//Circuit c의 a번째 큐빗부터 g의 size 만큼 쭉 gate 적용, 이건 그냥 Data 텐서곱으로 구현
// func (c *Circuit) Apply(g gate.Gate, a int) {
// 	left := gate.NewZero(1 << a)
// 	right := gate.NewZero(1<<c.Size - a - useful.Log_2(len(g.Data)))
// 	for i := range left.Data {
// 		left.Data[i][i] = 1
// 	}
// 	for i := range right.Data {
// 		right.Data[i][i] = 1
// 	}
// 	newGate := left.Tensor(g).Tensor(right)
// 	newVector := make([]complex128, len(c.Data))
// 	for i := range newVector {
// 		for j := range newGate.Data[0] {
// 			newVector[i] += c.Data[j] * newGate.Data[i][j]
// 		}
// 	}
// 	for i := range newVector {
// 		c.Data[i] = newVector[i]
// 	}
// }

//Circuit을 열벡터로 출력
//(예쁘게 보임)
func (c Circuit) String() string {
	s := ""
	for i := range c.Data {
		s += "["
		if real(c.Data[i]) >= 0 {
			s += "+"
		}
		s += fmt.Sprintf("%f", real(c.Data[i]))
		if imag(c.Data[i]) >= 0 {
			s += "+"
		}
		s += fmt.Sprintf("%fi]\n", imag(c.Data[i]))
	}
	return s
}

/*
Circuit
[1+0i]
[0+0i]
[0+0i]
[0+0i]
*/

//김지원 다만들어놨다
//i번째 qubit을 measure
//현재 Meausre 함수 문제 있음 약간 수정필요
/*
func (c *Circuit) Measure(i int) int {
	var prob0 float64 = 0
	period := 1 << i
	for k := 0; k < (1 << (c.Size - i - 1)); {
		for j := 0; j < period; j++ {
			prob0 += useful.CmplxAbs2(c.Data[i])
			k++
		}
		k += period
	}

	rand.Seed(time.Now().UnixNano())
	random := rand.Float64()

	//0을 measure
	if random < prob0 {
		for k := 0; k < len(c.Data); {
			for j := 0; j < period; j++ {
				c.Data[k] = complex(real(c.Data[k])/math.Sqrt(prob0), imag(c.Data[k])/math.Sqrt(prob0))
				k++
			}
			for j := 0; j < period; j++ {
				c.Data[k] = 0
				k++
			}
		}
		return 0
	}

	for k := 0; k < len(c.Data); {
		for j := 0; j < period; j++ {
			c.Data[k] = 0
			k++
		}
		for j := 0; j < period; j++ {
			c.Data[k] = complex(real(c.Data[k])/math.Sqrt(1-prob0), imag(c.Data[k])/math.Sqrt(1-prob0))
			k++
		}
	}
	return 1
}
*/

//a번째 qubit을 measure
func (c *Circuit) Measure(a int) int {
	var prob0 float64 = 0
	for i := range c.Data {
		//a번째 qubit이 0인가?
		if (i>>a)%2 == 0 {
			prob0 += useful.CmplxAbs2(c.Data[i])
		}
	}
	rand.Seed(time.Now().UnixNano())
	random := rand.Float64()

	var measured int
	var division float64
	if random < prob0 {
		measured = 0
		division = math.Sqrt(prob0)
	} else {
		measured = 1
		division = math.Sqrt(1 - prob0)
	}

	for i := range c.Data {
		if (i>>a)%2 == measured {
			c.Data[i] = complex(real(c.Data[i])/division, imag(c.Data[i])/division)
		} else {
			c.Data[i] = 0
		}
	}
	return measured
}
