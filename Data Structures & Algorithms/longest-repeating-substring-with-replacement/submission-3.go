func characterReplacement(s string, k int) int {
	left := 0
	longest := 0
	countChars := make(map[byte]int, 0)
	maxFreq := 0
	for right := 0; right < len(s); right++{
		countChars[s[right]]++
		maxFreq = max(maxFreq, countChars[s[right]])

		for (right - left + 1) - maxFreq > k {
			countChars[s[left]]--
			left++
		} 
		
		longest = max(longest, right - left + 1)
	}
	return longest
}
