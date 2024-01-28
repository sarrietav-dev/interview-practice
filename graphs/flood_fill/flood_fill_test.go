package floodfill_test

import (
	"testing"

	floodfill "github.com/sarrietav-dev/interview-practice/graphs/flood_fill"
)

func TestFloodFillOne(t *testing.T) {
	image := [][]int{
		{1, 1, 1}, {1, 1, 0}, {1, 0, 1},
	}
	sr := 1
	sc := 1
	newColor := 2
	expected := [][]int{
		{2, 2, 2}, {2, 2, 0}, {2, 0, 1},
	}

	result := floodfill.FloodFill(image, sr, sc, newColor)

	if !equal(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestFloodFillTwo(t *testing.T) {
	image := [][]int{
		{0, 0, 0}, {0, 1, 1},
	}
	sr := 1
	sc := 1
	newColor := 1
	expected := [][]int{
		{0, 0, 0}, {0, 1, 1},
	}

	result := floodfill.FloodFill(image, sr, sc, newColor)

	if !equal(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestFloodFillThree(t *testing.T) {
	image := [][]int{
		{0, 0, 0}, {0, 0, 0},
	}

	sr := 0
	sc := 0
	newColor := 0

	expected := [][]int{
		{0, 0, 0}, {0, 0, 0},
	}

	result := floodfill.FloodFill(image, sr, sc, newColor)

	if !equal(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func equal(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if len(a[i]) != len(b[i]) {
			return false
		}

		for j := 0; j < len(a[i]); j++ {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}

	return true
}
