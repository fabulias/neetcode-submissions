func hammingWeight(n int) int {
	count := 0
	for n > 0 {
		one := 1
		if one & n == one{
			count++
		}
		n = n >> 1
	}
	return count
}
