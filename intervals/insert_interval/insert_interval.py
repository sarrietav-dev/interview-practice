import unittest
from typing import List


class Solution:
    def insert(
        self, intervals: List[List[int]], newInterval: List[int]
    ) -> List[List[int]]:
        if len(intervals) == 0:
            return [newInterval]

        START = 0
        END = 1
        newIntervals = []

        for i, interval in enumerate(intervals):
            if newInterval[END] < interval[START]:
                newIntervals.append(newInterval)
                return newIntervals + intervals[i:]
            elif interval[END] < newInterval[START]:
                newIntervals.append(interval)
            else:
                newInterval = [min(newInterval[START], interval[START]), max(newInterval[END], interval[END])]

        return newIntervals + [newInterval]

class Tests(unittest.TestCase):
    def setUp(self):
        self.sol = Solution()

    def test_1(self):
        intervals = [[1, 3], [6, 9]]
        newInterval = [2, 5]
        expected = [[1, 5], [6, 9]]
        self.assertEqual(self.sol.insert(intervals, newInterval), expected)

    def test_2(self):
        intervals = [[1, 2], [3, 5], [6, 7], [8, 10], [12, 16]]
        newInterval = [4, 8]
        expected = [[1, 2], [3, 10], [12, 16]]
        self.assertEqual(self.sol.insert(intervals, newInterval), expected)

    def test_3(self):
        intervals = [[1, 5]]
        newInterval = [5, 7]
        expected = [[1, 7]]
        self.assertEqual(self.sol.insert(intervals, newInterval), expected)

    def test_4(self):
        intervals = [[1, 5]]
        newInterval = [6, 8]
        expected = [[1, 5], [6, 8]]
        self.assertEqual(self.sol.insert(intervals, newInterval), expected)


if __name__ == "__main__":
    unittest.main()
