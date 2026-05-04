func countComponents(n int, edges [][]int) int {
    parent := make([]int, n)
	rank := make([]int, n)
	for ix := range n {
		parent[ix] = ix
		rank[ix] = 1
	}

	var find func(n1 int) int
	find = func(n1 int) int {
		par := n1
		for par != parent[par] {
			parent[par] = parent[parent[par]]
			par = parent[par]
		}
		return par
	}

	var union func(n1, n2 int) int
	union = func(n1, n2 int) int {
		p1, p2 := find(n1), find(n2)
		if p1 == p2 {
			return 0
		}
		if rank[p2] > rank[p1] {
			parent[p1] = p2
			rank[p2] += rank[p1]
		} else {
			parent[p2] = p1
			rank[p1] += rank[p2]			
		}
		return 1	
	}

	result := n
	for _, edge := range edges {
		result -= union(edge[0], edge[1])
	}
	return result
}
