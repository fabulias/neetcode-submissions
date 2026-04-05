import "maps"
func longestConsecutive(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	if length == 1 {
		return 1
	}
	myMap := make(map[int]bool, 0)
	
	for _, n := range nums {
		myMap[n] = true
	}

	longest := 0
	theLongest := 0
	for key := range maps.Keys(myMap) {
		if myMap[key-1] {
			continue
		}
		
		for ix := key; myMap[ix]; ix++ {
			longest++
		}
		theLongest = max(theLongest, longest)
		longest = 0
	}

	return theLongest
}
