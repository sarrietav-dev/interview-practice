package threesum_test

import (
	"reflect"
	"testing"

	threesum "github.com/sarrietav-dev/interview-practice/arrays/three_sum"
)

func TestThreeSumOne(t *testing.T) {
	input := []int{
		-1, 0, 1, 2, -1, -4,
	}

	expected := [][]int{
		{-1, -1, 2},
		{-1, 0, 1},
	}

	actual := threesum.ThreeSum(input)

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}
}

func TestThreeSumTwo(t *testing.T) {
	input := []int{
		0, 0, 0,
	}

	expected := [][]int{
		{0, 0, 0},
	}

	actual := threesum.ThreeSum(input)

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}
}

func TestThreeSumThree(t *testing.T) {
	input := []int{
		0, 1, 1,
	}

  expected := [][]int{}

  actual := threesum.ThreeSum(input)

  if !reflect.DeepEqual(actual, expected) {
    t.Errorf("Expected %v, but got %v", expected, actual)
  }
}
