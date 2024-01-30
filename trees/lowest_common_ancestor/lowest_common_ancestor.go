package lowestcommonancestor

import "github.com/sarrietav-dev/interview-practice/trees"

// https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree
// Remember it's a binary SEARCH tree
func LowestCommonAncestor(root, p, q *trees.TreeNode) *trees.TreeNode {
	for root != nil {
		if root.Val > p.Val && root.Val > q.Val {
			root = root.Left
		} else if root.Val < p.Val && root.Val < q.Val {
			root = root.Right
		} else {
			return root
		}
	}

	return root
}
