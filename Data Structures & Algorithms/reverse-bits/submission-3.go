func reverseBits(n int) int {
	reverseN := 0
	count := 31
	for count >= 0 {
		bit := n & 1
		reverseN = reverseN << 1
		reverseN |=  bit
		n = n >> 1
		count--
	}
	return reverseN
}
