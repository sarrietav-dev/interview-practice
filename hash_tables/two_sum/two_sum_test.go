package twosum_test

import (
	"cmp"
	"slices"
	"testing"

	twosum "github.com/sarrietav-dev/interview-practice/hash_tables/two_sum"
)

func TestTwoSumOne(t *testing.T) {
	nums, target := []int{
		2, 7, 11, 15,
	}, 9

	result := twosum.TwoSum(nums, target)

	if !compare(result, []int{0, 1}) {
		t.Errorf("Expected {0, 1} but got %v", result)
	}
}

func TestTwoSumTwo(t *testing.T) {
	nums, target := []int{
		3, 2, 4,
	}, 6

	result := twosum.TwoSum(nums, target)

	if !compare(result, []int{1, 2}) {
		t.Errorf("Expected {1, 2} but got %v", result)
	}
}

func TestTwoSumThree(t *testing.T) {
	nums, target := []int{
		3, 3,
	}, 6

	result := twosum.TwoSum(nums, target)

	if !compare(result, []int{0, 1}) {
		t.Errorf("Expected {0, 1} but got %v", result)
	}
}

// Compares the result slice with the expected slice and its reverse
func compare[T cmp.Ordered](result []T, expected []T) bool {
	reverced := slices.Clone(expected)
	slices.Reverse(reverced)
	return slices.Compare(result, expected) == 0 || slices.Compare(result, reverced) == 0
}
