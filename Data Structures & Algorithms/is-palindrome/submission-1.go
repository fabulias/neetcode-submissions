func isPalindrome(s string) bool {
	n := len(s)
	left := 0
	right := n-1
	for left < right {
		for left < right && !(unicode.IsLetter(rune(s[left])) || unicode.IsNumber(rune(s[left]))){
			left++
		}
		for left < right && !(unicode.IsLetter(rune(s[right])) || unicode.IsNumber(rune(s[right]))){
			right--
		}
		if left > right {
			return false
		}
		
		lowerLeft := byte(unicode.ToLower(rune(s[left])))
		lowerRight := byte(unicode.ToLower(rune(s[right])))
	
		if lowerLeft != lowerRight {
			return false
		}
		left++
		right--
	}
	return true
}
