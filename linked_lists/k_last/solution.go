package klast

import linkedlists "github.com/sarrietav-dev/interview-practice/linked_lists"

// KLast returns the kth to last element of a singly linked list
func KLast[T comparable](head *linkedlists.Node[T], k int) T {
	p1, p2 := head, head
	

	for i := 0; i < k - 1; i++ {
		p2 = p2.Next
	}

	for p2.Next != nil {
		p1 = p1.Next
		p2 = p2.Next
	}

	return p1.Value
}
