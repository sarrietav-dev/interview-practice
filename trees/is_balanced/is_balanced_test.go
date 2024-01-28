package isbalanced_test

import (
	"testing"

	isbalanced "github.com/sarrietav-dev/interview-practice/trees/is_balanced"
)

func TestIsBalancedOne(t *testing.T) {
	tree := listToTreeNode([]int{3, 9, 20, -1, -1, 15, 7})

	if !isbalanced.IsBalanced(tree) {
		t.Errorf("Expected tree to be balanced")
	}
}

func TestIsBalancedTwo(t *testing.T) {
	tree := listToTreeNode([]int{1, 2, 2, 3, 3, -1, -1, 4, 4})

	if isbalanced.IsBalanced(tree) {
		t.Errorf("Expected tree to not be balanced")
	}
}

func TestIsBalancedThree(t *testing.T) {
	tree := listToTreeNode([]int{})

	if !isbalanced.IsBalanced(tree) {
		t.Errorf("Expected tree to be balanced")
	}
}

func TestIsBalancedFour(t *testing.T) {
	tree := listToTreeNode([]int{1})

	if !isbalanced.IsBalanced(tree) {
		t.Errorf("Expected tree to be balanced")
	}
}

func TestIsBalancedFive(t *testing.T) {
	tree := listToTreeNode([]int{1, -1, 2, -1, 3})

	if isbalanced.IsBalanced(tree) {
		t.Errorf("Expected tree to not be balanced")
	}
}

func TestIsBalancedSix(t *testing.T) {
	tree := listToTreeNode([]int{1, 2, 2, 3, -1, -1, 3, 4, -1, -1, 4})

	if isbalanced.IsBalanced(tree) {
		t.Errorf("Expected tree to not be balanced")
	}
}

func TestIsBalancedSeven(t *testing.T) {
	test := listToTreeNode([]int{
		1, 2, -1, 3, -1, 4, -1, 5,
	})

	if isbalanced.IsBalanced(test) {
		t.Errorf("Expected tree to not be balanced")
	}
}

func TestIsBalancedEight(t *testing.T) {
	test := listToTreeNode([]int{
		2, 1, 3, -1, -1, -1, 4, -1, 5, -1, 6,
	})

	if isbalanced.IsBalanced(test) {
		t.Errorf("Expected tree to not be balanced")
	}
}

func listToTreeNode(list []int) *isbalanced.TreeNode {
	if len(list) == 0 {
		return nil
	}

	root := &isbalanced.TreeNode{Val: list[0]}
	queue := []*isbalanced.TreeNode{root}
	i := 1

	for len(queue) > 0 && i < len(list) {
		node := queue[0]
		queue = queue[1:]

		if list[i] != -1 {
			node.Left = &isbalanced.TreeNode{Val: list[i]}
			queue = append(queue, node.Left)
		}
		i++

		if i < len(list) && list[i] != -1 {
			node.Right = &isbalanced.TreeNode{Val: list[i]}
			queue = append(queue, node.Right)
		}
		i++
	}

	return root
}
