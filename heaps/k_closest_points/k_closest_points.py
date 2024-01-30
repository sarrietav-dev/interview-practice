import heapq
import math
from typing import List
import unittest


class Solution:
    def kClosest(self, points: List[List[int]], k: int) -> List[List[int]]:
        min_heap = [[x**2 + y**2, x, y] for x, y in points]

        heapq.heapify(min_heap)

        res = [[x, y] for dist, x, y in heapq.nsmallest(k, min_heap)]

        return res


class Tests(unittest.TestCase):
    def setUp(self):
        self.sol = Solution()

    def test_1(self):
        points = [[1, 3], [-2, 2]]
        k = 1
        expected = [[-2, 2]]
        self.assertEqual(self.sol.kClosest(points, k), expected)

    def test_2(self):
        points = [[3, 3], [5, -1], [-2, 4]]
        k = 2
        expected = [[3, 3], [-2, 4]]
        self.assertEqual(self.sol.kClosest(points, k), expected)

    def test_3(self):
        points = [[0, 1], [1, 0]]
        k = 2
        expected = [[0, 1], [1, 0]]
        self.assertEqual(self.sol.kClosest(points, k), expected)

    def test_4(self):
        points = [[0, 1], [1, 0], [1, 1]]
        k = 2
        expected = [[0, 1], [1, 0]]
        self.assertEqual(self.sol.kClosest(points, k), expected)

    def test_5(self):
        points = [[0, 1], [1, 0], [1, 1]]
        k = 3
        expected = [[0, 1], [1, 0], [1, 1]]
        self.assertEqual(self.sol.kClosest(points, k), expected)

    def test_6(self):
        points = [[0, 1], [1, 0], [1, 1]]
        k = 1
        expected = [[0, 1]]
        self.assertEqual(self.sol.kClosest(points, k), expected)

    def test_7(self):
        points = [[0, 1], [1, 0], [1, 1]]
        k = 0
        expected = []
        self.assertEqual(self.sol.kClosest(points, k), expected)

    def test_8(self):
        points = [[0, 1], [1, 0], [1, 1]]
        k = 4
        expected = [[0, 1], [1, 0], [1, 1]]
        self.assertEqual(self.sol.kClosest(points, k), expected)

    def test_9(self):
        points = [[0, 1], [1, 0], [1, 1]]
        k = 5
        expected = [[0, 1], [1, 0], [1, 1]]
        self.assertEqual(self.sol.kClosest(points, k), expected)


if __name__ == "__main__":
    unittest.main()
