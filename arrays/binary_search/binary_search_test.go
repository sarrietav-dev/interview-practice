package binarysearch_test

import (
	"testing"

	binarysearch "github.com/sarrietav-dev/interview-practice/arrays/binary_search"
)

func TestBinarySearch1(t *testing.T) {
	nums := []int{
		-1, 0, 3, 5, 9, 12,
	}
	target := 9

	result := binarysearch.Search(nums, target)

	if result != 4 {
		t.Errorf("Expected %d but got %d", 4, result)
	}
}

func TestBinarySearch2(t *testing.T) {
	nums := []int{
		-1, 0, 3, 5, 9, 12,
	}

	target := 2

	result := binarysearch.Search(nums, target)

	if result != -1 {
		t.Errorf("Expected %d but got %d", -1, result)
	}
}
