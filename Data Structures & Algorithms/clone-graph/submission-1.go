/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}
    oldToNew := make(map[*Node]*Node)

	var dfs func(node *Node) *Node
	dfs = func(node *Node) *Node {
		if v, ok := oldToNew[node]; ok {
			return v
		}
		copy := &Node{Val: node.Val}
		oldToNew[node] = copy
		for _, n := range node.Neighbors {
			copy.Neighbors = append(copy.Neighbors,dfs(n))
		}
		return copy
	}

	return dfs(node)
}
