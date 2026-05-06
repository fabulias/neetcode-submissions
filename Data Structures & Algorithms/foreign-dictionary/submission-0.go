import "slices"
func foreignDictionary(words []string) string {
    adj := make(map[rune]map[rune]bool)
	for _, word := range words {
		for _, ch := range word {
			if _, ok := adj[ch]; !ok {
				adj[ch] = make(map[rune]bool)
			}
		}
	}

	for ix := 0; ix < len(words)-1;ix++ {
		w1, w2 := words[ix], words[ix+1] 
		minLen := min(len(w1), len(w2))
		if len(w1) > len(w2) && w1[:minLen] == w2[:minLen] {
			return ""
		}
		for iy := range minLen {
			if w1[iy] != w2[iy] {
				adj[rune(w1[iy])][rune(w2[iy])] = true
				break
			}
		}
	}
	visit := make(map[rune]bool, 0)
	res := []rune{}

	var dfs func(ch rune) bool
	dfs = func(ch rune) bool {
		if _, ok := visit[ch]; ok {
			return visit[ch]
		}
		visit[ch] = true

		for neighbor := range adj[ch] {
			if dfs(neighbor) {
				return true
			}
		}
		res = append(res, ch)
		visit[ch] = false
		return false
	}

	for ch := range adj {
		if dfs(ch) {
			return ""
		}
	}
	slices.Reverse(res)
	output := ""
	for _, r := range res {
		output += string(r)
	}
	return output
}
