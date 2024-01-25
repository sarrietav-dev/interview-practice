package binarysearch

import (
	"math"
)

// Solution for https://leetcode.com/problems/binary-search/
func Search(nums []int, target int) int {
	return bSearch(nums, target, 0, len(nums)-1)
}

func bSearch(nums []int, target int, low int, high int) int {
	if low > high {
		return -1
	}

	middle := low + int(math.Floor(float64(high-low)/2))

	if nums[middle] == target {
		return int(middle)
	} else if nums[middle] < target {
		return bSearch(nums, target, middle+1, high)
	} else {
		return bSearch(nums, target, low, middle-1)
	}
}
