func twoSum(nums []int, target int) []int {
    // 7-3 ? NO: [3,0]
	// 7/4 ? YES: [3,0] , 4,1 -> 0,1

	// [4,5,6] t=10
	// 10-4 ? NO: [4,0]
	// 10-5 ? NO: [4,0 -- 5,1]
	// 10-6 ? YES: 4,0 6,2 -> 0, 2
	results := make(map[int]int)
	for ix, n := range nums {
		if v, ok := results[target-n]; ok {
			return []int{v, ix}
		}
		results[n] = ix
	}
	return []int{-1, -1}
}
