func hammingWeight(n int) int {
	return bits.OnesCount(uint(n))
}
