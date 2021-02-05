package fibo

import (
	"math"
)

func SolutionRecursion(n int) int {
	if n <= 1 {
		return n
	}

	return SolutionRecursion(n-1) + SolutionRecursion(n-2)
}

func SolutionIteration(n int) int {
	f := 0
	s := 1

	for i := 1; i <= n; i++ {
		f, s = s, f+s
	}

	return f
}

func SolutionGoldenRatio(n int) int {
	res := (math.Pow(1.618, float64(n)) / math.Sqrt(float64(5))) + 0.5

	return int(res)
}

func SolutionMatrix(n int) int {
	if n <= 1 {
		return n
	}

	n--

	matrix := [][]int{{1, 1}, {1, 0}}

	step := 1
	for i := 1; i < n; i++ {
		if step*2 > n {
			break
		}

		matrix = mulMatrix(matrix, matrix)
		step *= 2
	}

	for i := 1; i <= n-step; i++ {
		matrix = mulMatrix(matrix, [][]int{{1, 1}, {1, 0}})
	}

	return matrix[0][0]
}

func mulMatrix(m1, m2 [][]int) [][]int {
	return [][]int{
		{
			m1[0][0]*m2[0][0] + m1[0][1]*m2[1][0], m1[0][0]*m2[1][0] + m1[0][1]*m2[1][1],
		},
		{
			m1[1][0]*m2[0][0] + m1[1][1]*m2[1][0], m1[1][0]*m2[0][1] + m1[1][1]*m2[1][1],
		},
	}
}
