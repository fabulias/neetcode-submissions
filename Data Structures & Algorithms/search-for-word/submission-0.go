func exist(board [][]byte, word string) bool {
	type pair struct {
		x int
		b int
	}
	rows, cols := len(board), len(board[0])
	path := make(map[pair]bool, 0)

	var dfs func(r, c, index int) bool
	dfs = func(r, c, index int) bool {
		if index == len(word) {
			return true
		}
		if r < 0 || c < 0 ||
			r >= rows || c >= cols ||
			word[index] != board[r][c] ||
			path[pair{r,c}] {
				return false
			}
		path[pair{r,c}] = true
		resp := dfs(r-1, c, index+1) ||
				dfs(r+1, c, index+1) ||
				dfs(r, c-1, index+1) ||
				dfs(r, c+1, index+1) 
		path[pair{r,c}] = false
		return resp
	} 
	for ix := range rows {
		for iy := range cols {
			if dfs(ix, iy, 0) {return true}
		}
	}
	return false
}
