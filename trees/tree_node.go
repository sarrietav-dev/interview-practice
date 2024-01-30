package trees

// TreeNode is a binary tree node
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Utility function to create a new tree node from a list of integers
func ListToTreeNode(list []int) *TreeNode {
	if len(list) == 0 {
		return nil
	}

	root := &TreeNode{Val: list[0]}
	queue := []*TreeNode{root}
	i := 1

	for len(queue) > 0 && i < len(list) {
		node := queue[0]
		queue = queue[1:]

		if list[i] != -1 {
			node.Left = &TreeNode{Val: list[i]}
			queue = append(queue, node.Left)
		}
		i++

		if i < len(list) && list[i] != -1 {
			node.Right = &TreeNode{Val: list[i]}
			queue = append(queue, node.Right)
		}
		i++
	}

	return root
}
