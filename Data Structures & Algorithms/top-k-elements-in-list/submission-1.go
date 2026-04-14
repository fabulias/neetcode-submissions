func topKFrequent(nums []int, k int) []int {
	count := make(map[int]int)
	freq := make([][]int, len(nums)+1)

	for _, n := range nums {
		count[n]++
	}
	for n, c := range count {
		freq[c] = append(freq[c], n)
	}

	result := []int{}
	for ix := len(freq)-1; ix>=0; ix-- {
		for _, n := range freq[ix] {
			result = append(result, n)
			if len(result) == k {
				return result
			}
		}
	}
	return result
}
