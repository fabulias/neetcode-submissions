func reverseBits(n int) int {
	reverseN := 0
	count := 0
	for count < 32 {
		bit := n & 1
		reverseN = reverseN << 1
		reverseN |=  bit
		n = n >> 1
		count ++
	}
	return reverseN
}
