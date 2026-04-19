func reverseBits(n int) int {
	reverseN := int(0)
	count := 0
	for n > 0 {
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
	for count < 32 {
		reverseN = reverseN << 1
		count++
	}
	return int(reverseN)
}
