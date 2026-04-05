func topKFrequent(nums []int, k int) []int {
	type pair struct {
		value int
		freq  int
	}
	frequenciesSlice := make([]pair, 2001)
	for _, n := range nums {
		frequenciesSlice[n+1000].value = n
		frequenciesSlice[n+1000].freq++
	}
	sort.Slice(frequenciesSlice, func(i, j int) bool {
		return frequenciesSlice[i].freq > frequenciesSlice[j].freq
	})
	result := make([]int, k)
	for i := 0; i < k; i++ {
		result[i] = frequenciesSlice[i].value
	}
	return result
}
