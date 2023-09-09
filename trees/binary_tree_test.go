package trees_test

import (
	"testing"

	"github.com/sarrietav-dev/interview-practice/trees"
)

func TestHeight(t *testing.T) {
	bt := getTreeFixture()

	if trees.Height(bt) != 4 {
		t.Errorf("Expected height 4, got %v", trees.Height(bt))
	}
}

func TestSize(t *testing.T) {
	bt := getTreeFixture()

	if trees.Size(bt) != 8 {
		t.Errorf("Expected size 8, got %v", trees.Size(bt))
	}
}

func TestInOrderTraversal(t *testing.T) {
	bt := getTreeFixture()

	trees.InOrderTraversal(bt)
}

func TestPreOrderTraversal(t *testing.T) {
	bt := getTreeFixture()

	trees.PreOrderTraversal(bt)
}

func getTreeFixture() *trees.BinaryTree {

	bt := &trees.BinaryTree{Value: 1}
	bt.Left = &trees.BinaryTree{Value: 2}
	bt.Right = &trees.BinaryTree{Value: 3}
	bt.Left.Left = &trees.BinaryTree{Value: 4}
	bt.Left.Right = &trees.BinaryTree{Value: 5}
	bt.Right.Left = &trees.BinaryTree{Value: 6}
	bt.Right.Right = &trees.BinaryTree{Value: 7}
	bt.Right.Right.Right = &trees.BinaryTree{Value: 8}

	return bt
}