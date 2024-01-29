package maximumsubarray

import "math"

// https://leetcode.com/problems/maximum-subarray
// Using Kadane's algorithm
func MaxSubArray(nums []int) int {
  bestSum := math.Inf(-1)
  currentSum := 0

  for _, n := range nums {
    currentSum = max(n, currentSum + n)
    bestSum = max(bestSum, float64(currentSum))
  }

  return int(bestSum)
}
