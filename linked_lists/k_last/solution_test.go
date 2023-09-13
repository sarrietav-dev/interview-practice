package klast_test

import (
	"testing"

	linkedlists "github.com/sarrietav-dev/interview-practice/linked_lists"
	klast "github.com/sarrietav-dev/interview-practice/linked_lists/k_last"
)

func TestSolution(t *testing.T) {
	ll := linkedlists.LinkedList[int]{
		Head: &linkedlists.Node[int]{
			Value: 1,
			Next: &linkedlists.Node[int]{
				Value: 2,
				Next: &linkedlists.Node[int]{
					Value: 3,
					Next: &linkedlists.Node[int]{
						Value: 4,
						Next: &linkedlists.Node[int]{
							Value: 5,
							Next: &linkedlists.Node[int]{
								Value: 6,
								Next:  nil,
							},
						},
					},
				},
			},
		},
	}

	k := 2

	expected := 5

	actual := klast.KLast(ll.Head, k)

	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}
