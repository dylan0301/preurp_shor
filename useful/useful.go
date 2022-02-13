package useful

//a^b
func Power(a, b int) int {
	now := 1
	for b > 0 {
		if b%2 == 1 {
			now *= a
		}
		b /= 2
		a *= a
	}
	return now
}

// a^b mod n
func ModPow(a, b, n int) int {
	now := 1
	for b > 0 {
		if b%2 == 1 {
			now = now * a % n
		}
		b /= 2
		a = a * a % n
	}
	return now
}

//분수를 연분수로 바꿔주는 함수
func ContinuedFrct(numerator, denominator int) []int {
	continuous_fraction := []int{}
	i := 0
	for denominator != 0 {
		c := numerator / denominator
		numerator -= c * denominator
		continuous_fraction = append(continuous_fraction, c)
		temp := denominator
		denominator = numerator
		numerator = temp
		i += 1
	}
	return continuous_fraction
}

//log2 n값을 반환. 자연수로
func Log_2(n int) int {
	result := 0
	for n != 1 {
		n /= 2
		result += 1
	}
	return result
}

//복소수의 절댓값 제곱
func CmplxAbs2(x complex128) float64 {
	return real(x)*real(x) + imag(x)*imag(x)
}

//GCD 유클리드 호제법
func GCD(a int, b int) int {
	for b > 0 {
		if a < b {
			temp := a
			a = b
			b = temp
		}
		a %= b
	}
	return b
}
