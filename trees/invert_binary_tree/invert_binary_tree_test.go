package invertbinarytree_test

import (
	"slices"
	"testing"

	invertbinarytree "github.com/sarrietav-dev/interview-practice/trees/invert_binary_tree"
)

func TestInvertTreeOne(t *testing.T) {
	input := listToTreeNode([]int{4, 2, 7, 1, 3, 6, 9})
	expected := []int{4, 7, 2, 9, 6, 3, 1}
	result := treeNodeToList(invertbinarytree.InvertTree(input))

	if !slices.Equal(expected, result) {
		t.Errorf("Expected %v to equal %v", result, expected)
	}
}

func TestInvertTreeTwo(t *testing.T) {
	input := listToTreeNode([]int{2, 1, 3})
	expected := []int{2, 3, 1}
	result := treeNodeToList(invertbinarytree.InvertTree(input))

	if !slices.Equal(expected, result) {
		t.Errorf("Expected %v to equal %v", result, expected)
	}
}

func TestInvertTreeThree(t *testing.T) {
	input := listToTreeNode([]int{})
	expected := []int{}
	result := treeNodeToList(invertbinarytree.InvertTree(input))

	if !slices.Equal(expected, result) {
		t.Errorf("Expected %v to equal %v", result, expected)
	}
}

func listToTreeNode(tree []int) *invertbinarytree.TreeNode {
	if len(tree) == 0 {
		return nil
	}

	root := &invertbinarytree.TreeNode{Val: tree[0]}
	queue := []*invertbinarytree.TreeNode{root}
	for i := 1; i < len(tree); i++ {
		node := queue[0]
		queue = queue[1:]

		if tree[i] != -1 {
			node.Left = &invertbinarytree.TreeNode{Val: tree[i]}
			queue = append(queue, node.Left)
		}

		i++
		if i < len(tree) && tree[i] != -1 {
			node.Right = &invertbinarytree.TreeNode{Val: tree[i]}
			queue = append(queue, node.Right)
		}
	}

	return root
}

func treeNodeToList(node *invertbinarytree.TreeNode) []int {
	if node == nil {
		return []int{}
	}

	result := []int{node.Val}
	queue := []*invertbinarytree.TreeNode{node}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node.Left != nil {
			result = append(result, node.Left.Val)
			queue = append(queue, node.Left)
		} else {
			result = append(result, -1)
		}

		if node.Right != nil {
			result = append(result, node.Right.Val)
			queue = append(queue, node.Right)
		} else {
			result = append(result, -1)
		}
	}

	for result[len(result)-1] == -1 {
		result = result[:len(result)-1]
	}

	return result
}
