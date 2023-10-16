package heaps

import (
	"cmp"
	"math"
)

func Parent(i int) int {
	return int(math.Floor(float64(i-1) / 2))
}

func Left(i int) int {
	return (i+1)*2 - 1
}

func Right(i int) int {
	return (i + 1) * 2
}

func Swap[T cmp.Ordered](arr []T, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func SiftUp[T cmp.Ordered](arr []T, i int) {
	for i > 0 && arr[Parent(i)] < arr[i] {
		Swap(arr, i, Parent(i))
		i = Parent(i)
	}
}

func SiftDown[T cmp.Ordered](arr []T, i int) {
	n := len(arr)
	maxIndex := i
	l := Left(i)
	if l < n && arr[l] > arr[maxIndex] {
		maxIndex = l
	}
	r := Right(i)
	if r < n && arr[r] > arr[maxIndex] {
		maxIndex = r
	}
	if i != maxIndex {
		Swap(arr, i, maxIndex)
		SiftDown(arr, maxIndex)
	}
}

func Insert[T cmp.Ordered](arr []T, key T) []T {
	arr = append(arr, key)
	SiftUp(arr, len(arr)-1)
	return arr
}

func ExtractMax[T cmp.Ordered](arr []T) (T, []T) {
	max := arr[0]
	arr[0] = arr[len(arr)-1]
	arr = arr[:len(arr)-1]
	SiftDown(arr, 0)
	return max, arr
}

func Equal[T cmp.Ordered](arr1 []T, arr2 []T) bool {
	if len(arr1) != len(arr2) {
		return false
	}
	for i, v := range arr1 {
		if v != arr2[i] {
			return false
		}
	}
	return true
}

func Remove[T cmp.Ordered](arr []T, i int) (T, []T) {
	Swap(arr, i, len(arr)-1)
	num := arr[len(arr)-1]
	arr = arr[:len(arr)-1]
	SiftDown(arr, i)
	return num, arr
}

func ChangePriority[T cmp.Ordered](arr []T, i int, key T) []T {
	oldKey := arr[i]
	arr[i] = key
	if key > oldKey {
		SiftUp(arr, i)
	} else {
		SiftDown(arr, i)
	}
	return arr
}

func FindIndex[T cmp.Ordered](arr []T, key T) int {
	l, h := 0, len(arr)-1

	for l < h {
		middle := (l + h) / 2

		if arr[middle] == key {
			return middle
		} else if arr[middle] < key {
			l = middle + 1
		} else {
			h = middle - 1
		}
	}

	return -1
}
