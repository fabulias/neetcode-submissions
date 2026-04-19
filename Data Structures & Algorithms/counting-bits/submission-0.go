func countBits(n int) []int {
	answer := make([]int, n+1)
	for ix := 0; ix <= n; ix++ {
		result := ix & ((1 << ix) - 1)
		answer[ix] = bits.OnesCount(uint(result))
	}
	return answer
}
