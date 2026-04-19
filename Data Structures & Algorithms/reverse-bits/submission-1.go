func reverseBits(n int) int {
	reverseN := 0
	count := 0
	for count < 32 {
		one := 1
		if n & one == one {
			fmt.Println("it's 1")
			reverseN = reverseN << 1
			reverseN = reverseN | 1
		} else {
			fmt.Println("it's 0")
			reverseN = reverseN << 1
			reverseN = reverseN | 0
		}
		n = n >> 1
		count ++
	}
	return reverseN
}
