package linkedlists

type LinkedList[T comparable] struct {
	Head *Node[T]
	Tail *Node[T]
}

type Node[T comparable] struct {
	Value T
	Next  *Node[T]
}

func (ll *LinkedList[T]) Add(value T) *Node[T] {
	node := &Node[T]{Value: value}

	if ll.Head == nil {
		ll.Head = node
		ll.Tail = node
		return node
	}

	ll.Tail.Next = node
	ll.Tail = node
	return node
}