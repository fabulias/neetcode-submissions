

func threeSum(nums []int) [][]int {
	n := len(nums)
	results := make([][]int, 0)
	
	sort.Ints(nums)

	for a := 0; a < n; a++ {
		if a > 0 && nums[a] == nums[a-1] {
			continue
		}
		aValue := nums[a]
		left, right := a + 1, n-1
		for left < right {
			sum := nums[a] + nums[left] + nums[right]
			if sum > 0 {
				right--
			} else if sum < 0 {
				left++
			} else {
				results = append(results,[]int{aValue, nums[left], nums[right]})
				left++

				for left < right && nums[left] == nums[left-1] {
					left++
				}
			}

		}
	}

	return results
}
