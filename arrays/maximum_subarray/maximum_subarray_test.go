package maximumsubarray_test

import (
	"testing"

	maximumsubarray "github.com/sarrietav-dev/interview-practice/arrays/maximum_subarray"
)

func TestMaximumSubarrayOne(t *testing.T) {
	input := []int{
		-2, 1, -3, 4, -1, 2, 1, -5, 4,
	}
	expected := 6
	result := maximumsubarray.MaxSubArray(input)

	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestMaximumSubarrayTwo(t *testing.T) {
	input := []int{
		1,
	}
	expected := 1
	result := maximumsubarray.MaxSubArray(input)

	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestMaximumSubarrayThree(t *testing.T) {
	input := []int{
		5, 4, -1, 7, 8,
	}
	expected := 23
	result := maximumsubarray.MaxSubArray(input)

	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}
