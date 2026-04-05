func lengthOfLongestSubstring(s string) int {
	if len(s) <= 1 {
		return len(s)
	}

	longest := 0
	left, right := 0, 0
	setChars := make(map[byte]bool, 0)
	for ;right < len(s);right++ {
		for setChars[s[right]] {
			delete(setChars, s[left])
			left++
		}

		setChars[s[right]] = true
		longest = max(longest, len(setChars))
	}
	return longest
}
