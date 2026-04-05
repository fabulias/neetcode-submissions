func productExceptSelf(nums []int) []int {
	n := len(nums)
	saveResults := make([]int, n)

	prefix := 1
	for ix := 0; ix < n; ix++ {
		saveResults[ix] = prefix
		prefix *= nums[ix]
	}
	fmt.Println("from left", saveResults)

	suffix := 1
	for ix := n-1; ix >= 0; ix-- {
		saveResults[ix] *= suffix
		suffix *= nums[ix]
	}

	fmt.Println("from right", saveResults)

	return saveResults
}