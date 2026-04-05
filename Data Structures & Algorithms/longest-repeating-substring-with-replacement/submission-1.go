func maxFreq(in map[byte]int) (maxValue int) {
	for _, freq := range in {
		maxValue = max(maxValue, freq)
	}
	return
}

func characterReplacement(s string, k int) int {
	left := 0
	longest := 0
	countChars := make(map[byte]int, 0)
	for right := 0; right < len(s); right++{
		countChars[s[right]]++
		for (right - left + 1) - maxFreq(countChars) > k {
			countChars[s[left]]--
			left++
		} 
		
		longest = max(longest, right - left + 1)
	}
	return longest
}
