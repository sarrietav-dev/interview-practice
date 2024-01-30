package insertinterval

// https://leetcode.com/problems/insert-interval
func Insert(intervals [][]int, newInterval []int) [][]int {
	result := [][]int{}

	START := 0
	END := 1

	for i, interval := range intervals {
		if newInterval[END] < interval[START] {
			result = append(result, newInterval)
			return append(result, intervals[i:]...)
		} else if interval[END] < newInterval[START] {
			result = append(result, interval)
		} else {
			newInterval = []int{
				min(interval[START], newInterval[START]),
				max(interval[END], newInterval[END]),
			}
		}
	}

	return append(result, newInterval)
}
