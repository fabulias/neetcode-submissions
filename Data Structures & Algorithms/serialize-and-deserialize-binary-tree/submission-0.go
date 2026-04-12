/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {
    
}

func Constructor() Codec {
    return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	output := []string{}
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			output = append(output, "N")
			return
		}
		output = append(output, strconv.Itoa(node.Val))
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	return strings.Join(output, ",")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	values := strings.Split(data, ",")
	ix := 0

	var dfs func() *TreeNode
	dfs = func() *TreeNode {
		if values[ix] == "N" {
			ix++
			return nil
		}
		value, _ := strconv.Atoi(values[ix])
		node := &TreeNode{Val:value}
		ix++
		node.Left = dfs()
		node.Right = dfs()
		return node
	}
	
    return dfs()
}
