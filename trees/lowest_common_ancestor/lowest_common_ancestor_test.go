package lowestcommonancestor_test

import (
	"testing"

	"github.com/sarrietav-dev/interview-practice/trees"
	lowestcommonancestor "github.com/sarrietav-dev/interview-practice/trees/lowest_common_ancestor"
)

func TestLowestAncestorOne(t *testing.T) {
	tree := trees.ListToTreeNode([]int{
		6, 2, 8, 0, 4, 7, 9, -1, -1, 3, 5,
	})

	p := trees.ListToTreeNode([]int{2})
	q := trees.ListToTreeNode([]int{8})

	expected := 6

	result := lowestcommonancestor.LowestCommonAncestor(tree, p, q)

	if result.Val != expected {
		t.Errorf("Expected %d, got %d", expected, result.Val)
	}
}

func TestLowestAncestorTwo(t *testing.T) {
	tree := trees.ListToTreeNode([]int{
		6, 2, 8, 0, 4, 7, 9, -1, -1, 3, 5,
	})

	p := trees.ListToTreeNode([]int{2})
	q := trees.ListToTreeNode([]int{4})

	expected := 2

	result := lowestcommonancestor.LowestCommonAncestor(tree, p, q)

	if result.Val != expected {
		t.Errorf("Expected %d, got %d", expected, result.Val)
	}
}

func TestLowestAncestorThree(t *testing.T) {
	tree := trees.ListToTreeNode([]int{
		2, 1,
	})

	p := trees.ListToTreeNode([]int{2})
	q := trees.ListToTreeNode([]int{1})

	expected := 2

	result := lowestcommonancestor.LowestCommonAncestor(tree, p, q)

	if result.Val != expected {
		t.Errorf("Expected %d, got %d", expected, result.Val)
	}
}
