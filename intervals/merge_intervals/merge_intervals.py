from typing import List
import unittest


class Solution:
    def merge(self, intervals: List[List[int]]) -> List[List[int]]:
        intervals = sorted(intervals, key=lambda x: x[0])
        result = [intervals[0]]

        for interval in intervals[1:]:
            if self.does_overlap(interval, result[-1]):
                merged = self.merge_overlap(interval, result[-1])
                result[-1] = merged
            else:
                result.append(interval)

        return result

    def does_overlap(self, a, b):
        return a[0] <= b[1] and b[0] <= a[1]

    def merge_overlap(self, a, b):
        return [min(a[0], b[0]), max(a[1], b[1])]


class Tests(unittest.TestCase):
    def setUp(self):
        self.sol = Solution()

    def test_example1(self):
        intervals = [[1, 3], [2, 6], [8, 10], [15, 18]]
        self.assertEqual(self.sol.merge(intervals), [[1, 6], [8, 10], [15, 18]])

    def test_example2(self):
        intervals = [[1, 4], [4, 5]]
        self.assertEqual(self.sol.merge(intervals), [[1, 5]])

    def test_example3(self):
        intervals = [[1, 4], [0, 4]]
        self.assertEqual(self.sol.merge(intervals), [[0, 4]])

    def test_example4(self):
        intervals = [[1, 4], [2, 3]]
        self.assertEqual(self.sol.merge(intervals), [[1, 4]])

    def test_example5(self):
        intervals = [[1, 4], [0, 0]]
        self.assertEqual(self.sol.merge(intervals), [[0, 0], [1, 4]])

    def test_example6(self):
        intervals = [[1, 4], [0, 2], [3, 5]]
        self.assertEqual(self.sol.merge(intervals), [[0, 5]])

    def test_example7(self):
        intervals = [[1, 4], [0, 5], [3, 5]]
        self.assertEqual(self.sol.merge(intervals), [[0, 5]])


if __name__ == "__main__":
    unittest.main()
