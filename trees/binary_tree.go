package trees

import (
	"fmt"

	"github.com/sarrietav-dev/interview-practice/queues"
)

type BinaryTree struct {
	Value int
	Left  *BinaryTree
	Right *BinaryTree
}

func Height(bt *BinaryTree) int {
	if bt == nil {
		return 0
	}

	return 1 + max(Height(bt.Left), Height(bt.Right))
}

func Size(bt *BinaryTree) int {
	if bt == nil {
		return 0
	}

	return 1 + Size(bt.Left) + Size(bt.Right)
}

func InOrderTraversal(bt *BinaryTree) {
	if bt == nil {
		return
	}

	InOrderTraversal(bt.Left)
	fmt.Println(bt.Value)
	InOrderTraversal(bt.Right)
}

func PreOrderTraversal(bt *BinaryTree) {
	if bt == nil {
		return
	}

	fmt.Println(bt.Value)
	PreOrderTraversal(bt.Left)
	PreOrderTraversal(bt.Right)
}

func BreadthFirstSearch(bt *BinaryTree) {
	if bt == nil {
		return
	}

	queue := queues.Queue[*BinaryTree]{}
	queue.Enqueue(bt)

	for !queue.IsEmpty() {
		node := queue.Dequeue()
		fmt.Println(node.Value)
		if node.Left != nil {
			queue.Enqueue(node.Left)
		}
		if node.Right != nil {
			queue.Enqueue(node.Right)
		}
	}
}
