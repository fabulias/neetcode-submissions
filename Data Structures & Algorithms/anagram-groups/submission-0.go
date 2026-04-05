import "slices"

func groupAnagrams(strs []string) [][]string {
	solutionsMap := make(map[string][]string)
	for _, s := range strs {
		runes := []rune(s)
		slices.Sort(runes)
		solution := string(runes)
		solutionsMap[solution] = append(solutionsMap[solution], s)
	}
	solutions := make([][]string,0)
	
	for _, v := range solutionsMap {
		solutions = append(solutions, v)
	}
	return solutions
}
