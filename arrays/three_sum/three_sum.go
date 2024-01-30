package threesum

import (
	"slices"
)

// https://leetcode.com/problems/3sum
func ThreeSum(nums []int) [][]int {
	result := [][]int{}
	slices.Sort(nums)

	for i, n := range nums {
		// If the first position is still equal
		if i > 0 && nums[i-1] == n {
			continue
		}

		// Instanciate two pointers
		l, r := i+1, len(nums)-1

		for l < r {
			sum := n + nums[l] + nums[r]

			// If the sum is too big, reduce a number
			if sum > 0 {
				r -= 1
				// If the sum is too small, add a number
			} else if sum < 0 {
				l += 1
			} else {
				result = append(result, []int{n, nums[l], nums[r]})

				l += 1
				// Update until there's no duplicates for the second position
				for nums[l] == nums[l-1] && l < r {
					l += 1
				}
			}
		}
	}

	return result
}
