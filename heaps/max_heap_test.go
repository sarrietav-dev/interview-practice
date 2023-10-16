package heaps_test

import (
	"testing"

	"github.com/sarrietav-dev/interview-practice/heaps"
)

func TestParent(t *testing.T) {
	i := 5
	expected := 2
	actual := heaps.Parent(i)

	if heaps.Parent(i) != expected {
		t.Errorf("Right of %d should be %d but got %d", i, expected, actual)
	}
}

func TestLeft(t *testing.T) {
	i := 5
	expected := 10
	actual := heaps.Left(i)
	if actual != expected {
		t.Errorf("Right of %d should be %d but got %d", i, expected, actual)
	}
}

func TestRight(t *testing.T) {
	i := 5
	expected := 11
	actual := heaps.Right(i)
	if actual != expected {
		t.Errorf("Right of %d should be %d but got %d", i, expected, actual)
	}
}

func TestSiftUp(t *testing.T) {
	arr := []int{7, 6, 5, 4, 3, 2, 1, 8}
	i := 7
	expected := []int{8, 7, 5, 6, 3, 2, 1, 4}
	heaps.SiftUp(arr, i)
	if !heaps.Equal(arr, expected) {
		t.Errorf("Expected %v but got %v", expected, arr)
	}
}

func TestSiftDown(t *testing.T) {
	arr := []int{4, 8, 7, 6, 5, 3, 2, 1}
	i := 0
	expected := []int{8, 6, 7, 4, 5, 3, 2, 1}
	heaps.SiftDown(arr, i)
	if !heaps.Equal(arr, expected) {
		t.Errorf("Expected %v but got %v", expected, arr)
	}
}

func TestExtractMaxNumber(t *testing.T) {
	arr := []int{8, 7, 5, 6, 3, 2, 1, 4}
	expected := 8
	actual, _ := heaps.ExtractMax(arr)
	if expected != actual {
		t.Errorf("Expected %v but got %v", expected, arr)
	}
}

func TestExtractMaxArray(t *testing.T) {
	arr := []int{8, 6, 7, 4, 5, 3, 2, 1}
	expected := []int{7, 6, 3, 4, 5, 1, 2}
	_, newArr := heaps.ExtractMax(arr)
	if !heaps.Equal(newArr, expected) {
		t.Errorf("Expected %v but got %v", expected, arr)
	}
}

func TestInsert(t *testing.T) {
	arr := []int{8, 7, 5, 6, 3, 2, 1, 4}
	key := 9
	expected := []int{9, 8, 5, 7, 3, 2, 1, 4, 6}
	actual := heaps.Insert(arr, key)
	if !heaps.Equal(actual, expected) {
		t.Errorf("Expected %v but got %v", expected, arr)
	}
}

func TestRemove(t *testing.T) {
	arr := []int{8, 7, 6, 5, 4, 3, 2, 1}
	key := 5
	expected := []int{8, 7, 6, 5, 4, 1, 2}
	value, actual := heaps.Remove(arr, key)

	if !heaps.Equal(actual, expected) {
		t.Errorf("Expected %v but got %v", expected, arr)
	}

	if value != 3 {
		t.Errorf("Expected %v but got %v", 5, value)
	}
}

func TestChangePriority(t *testing.T) {
	arr := []int{8, 7, 6, 5, 4, 3, 2, 1}
	key := 5
	priority := 10
	expected := []int{10, 7, 8, 5, 4, 6, 2, 1}
	actual := heaps.ChangePriority(arr, key, priority)

	if !heaps.Equal(actual, expected) {
		t.Errorf("Expected %v but got %v", expected, arr)
	}
}

func TestFindIndex(t *testing.T) {
	arr := []int{8, 7, 6, 5, 4, 3, 2, 1}
	key := 5
	expected := 3
	actual := heaps.FindIndex(arr, key)

	if actual != expected {
		t.Errorf("Expected %v but got %v", expected, arr)
	}
}
