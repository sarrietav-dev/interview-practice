package diameter_test

import (
	"testing"

	"github.com/sarrietav-dev/interview-practice/trees/diameter"
)

func TestDiameterOne(t *testing.T) {
	tree := listToTreeNode([]int{1, 2, 3, 4, 5})
	expected := 3

	actual := diameter.DiameterOfBinaryTree(tree)

	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

func TestDiameterTwo(t *testing.T) {
	tree := listToTreeNode([]int{1, 2})
	expected := 1

	actual := diameter.DiameterOfBinaryTree(tree)

	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

func TestDiameterThree(t *testing.T) {
	tree := listToTreeNode([]int{3, 1, -1, -1, 2})
	expected := 2

	actual := diameter.DiameterOfBinaryTree(tree)

	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

func TestDiameterFour(t *testing.T) {
	tree := listToTreeNode([]int{3, 2, 4, -1, -1, 1})
	expected := 3

	actual := diameter.DiameterOfBinaryTree(tree)

	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

func TestDiameterFive(t *testing.T) {
	tree := listToTreeNode([]int{4,2,-1,3,1,-1,-1,5})
	expected := 3

	actual := diameter.DiameterOfBinaryTree(tree)

	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

func listToTreeNode(list []int) *diameter.TreeNode {
	if len(list) == 0 {
		return nil
	}

	root := &diameter.TreeNode{Val: list[0]}
	queue := []*diameter.TreeNode{root}
	i := 1

	for len(queue) > 0 && i < len(list) {
		node := queue[0]
		queue = queue[1:]

		if list[i] != -1 {
			node.Left = &diameter.TreeNode{Val: list[i]}
			queue = append(queue, node.Left)
		}
		i++

		if i < len(list) && list[i] != -1 {
			node.Right = &diameter.TreeNode{Val: list[i]}
			queue = append(queue, node.Right)
		}
		i++
	}

	return root
}
