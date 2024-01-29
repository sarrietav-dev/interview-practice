package diameter

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// https://leetcode.com/problems/diameter-of-binary-tree
func DiameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	maxDiameter := [1]int{0}

	longestPath(root, &maxDiameter)

	return maxDiameter[0]
}

func longestPath(root *TreeNode, maxDiameter *[1]int) int {
	if root == nil {
		return -1
	}

	leftHeight := longestPath(root.Left, maxDiameter)
	rightHeight := longestPath(root.Right, maxDiameter)

	height := 1 + max(leftHeight, rightHeight)

	maxDiameter[0] = max(maxDiameter[0], leftHeight+rightHeight+2)

	return height
}
