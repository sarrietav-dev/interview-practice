package deletemiddlenode

import linkedlists "github.com/sarrietav-dev/interview-practice/linked_lists"

func Delete[T comparable](node *linkedlists.Node[T]) T {
	if node.Next.Next == nil {
		nextValue := node.Next.Value
		node.Next = nil
		v := node.Value
		node.Value = nextValue
		return v
	}

	nextValue := Delete(node.Next)
	v := node.Value
	node.Value = nextValue
	return v
}