func reverseBits(n int) int {
	reverseN := 0
	count := 0
	for count < 32 {
		bit := (n >> count) & 1
		reverseN = reverseN << 1
		reverseN |=  bit
		count++
	}
	return reverseN
}
