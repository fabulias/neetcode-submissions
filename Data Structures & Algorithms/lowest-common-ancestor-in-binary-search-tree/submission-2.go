/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {
	curr := root
	for curr != nil {
		if p.Val > curr.Val && q.Val > curr.Val {
			// right subtree
			curr = curr.Right
		} else if p.Val < curr.Val && q.Val < curr.Val {
			// left subtree
			curr = curr.Left
		} else {
			return curr
		}
	}
	return nil
}
