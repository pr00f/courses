package tickets

var cache map[int]int

func Solution(n int) int {
	var res int

	cache = make(map[int]int)
	fillCache(0, n, 0)

	for _, v := range cache {
		res += v * v
	}

	return res
}

func fillCache(nn, n, s int) {
	if nn == n {
		cache[s]++
		return
	}

	for i := 0; i <= 9; i++ {
		fillCache(nn+1, n, s+i)
	}
}
