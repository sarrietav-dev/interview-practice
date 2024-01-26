package twosum

// https://leetcode.com/problems/two-sum
func TwoSum(nums []int, target int) []int {
  numMap := make(map[int]int)

  for i, v := range nums {
    needed := target - v

    index, exists := numMap[needed]

    if exists {
      return []int {i, index}
    }

    numMap[v] = i
  }

  return nil
}
