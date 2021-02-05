package power

import (
	"strconv"
)

func SolutionSimple(a float64, n int) float64 {
	p := 1.0
	for i := 1; i <= n; i++ {
		p *= a
	}

	return p
}

func SolutionPowerOfTwoWithMultiplication(a float64, n int) float64 {
	if n == 0 {
		return 1.0
	}

	p := a
	step := 1
	for i := 1; i < n; i++ {
		if step*2 > n {
			break
		}

		a *= a
		step *= 2
	}

	for i := 1; i <= n-step; i++ {
		a *= p
	}

	return a
}

func SolutionBinaryExpansionOfExponent(a float64, n int) float64 {
	binary := strconv.FormatInt(int64(n), 2)
	res := 1.0

	for i := len(binary) - 1; i >= 0; i-- {
		if binary[i] == '1' {
			res *= a
		}

		a *= a
	}

	return res
}
