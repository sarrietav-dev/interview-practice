package isbalanced

import (
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// https://leetcode.com/problems/balanced-binary-tree
func IsBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	leftHeight, leftBalanced := evaluate(root.Left)

	if !leftBalanced {
		return false
	}

	rightHeight, rightBalanced := evaluate(root.Right)

	if !rightBalanced {
		return false
	}

	isBalanced := int(math.Abs(float64(leftHeight-rightHeight))) <= 1

	return isBalanced
}

func evaluate(root *TreeNode) (int, bool) {
	if root == nil {
		return -1, true
	}

	leftHeight, leftBalanced := evaluate(root.Left)

	if !leftBalanced {
		return -1, false
	}

	rightHeight, rightBalanced := evaluate(root.Right)

	if !rightBalanced {
		return -1, false
	}

	height := int(max(leftHeight, rightHeight)) + 1
	isBalanced := int(math.Abs(float64(leftHeight-rightHeight))) <= 1

	return height, isBalanced
}
