func isAnagram(s string, t string) bool {
	sMap := make(map[rune]int)
	tMap := make(map[rune]int)
	for _, b := range s {
		sMap[b]++
	}
	for _, b := range t {
		tMap[b]++
	}
	if len(sMap) != len(tMap) {
		return false
	}
	for k, v := range sMap {
		if tMap[k] != v {
			return false
		}
	}
	return true
}
