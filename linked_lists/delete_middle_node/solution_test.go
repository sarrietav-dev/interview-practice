package deletemiddlenode_test

import (
	"testing"

	linkedlists "github.com/sarrietav-dev/interview-practice/linked_lists"
	deletemiddlenode "github.com/sarrietav-dev/interview-practice/linked_lists/delete_middle_node"
)

func TestSolution(t *testing.T) {
	ll := linkedlists.LinkedList[string]{}

	ll.Add("A")
	ll.Add("B")
	node := ll.Add("C")
	ll.Add("D")
	ll.Add("E")
	ll.Add("F")

	deletemiddlenode.Delete(node)

	lenght := 0

	for n := ll.Head; n != nil; n = n.Next {
		if n.Value == "C" {
			t.Errorf("Expected node with value 'C' to be deleted")
		}
		lenght++
	}

	if lenght != 5 {
		t.Errorf("Expected lenght to be 5, got %d", lenght)
	}
}
