/*
func missingNumber(nums []int) int {
	xorr := len(nums)
	for ix := 0; ix < len(nums); ix++ {
		xorr ^= nums[ix]
		xorr ^= ix
	}
	return xorr
}
*/

func missingNumber(nums []int) int {
	sum := 0
	for ix := 0; ix < len(nums); ix++ {
		sum -= nums[ix]
		sum += ix
	}
	
	return sum + len(nums)
}
