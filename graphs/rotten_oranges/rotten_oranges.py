from collections import deque
from typing import List
import unittest


class Solution:
    def orangesRotting(self, grid: List[List[int]]) -> int:
        rows, cols = len(grid), len(grid[0])
        queue = deque(
            [(i, j, 0) for i in range(rows) for j in range(cols) if grid[i][j] == 2]
        )
        seconds = 0
        directions = [(0, 1), (0, -1), (1, 0), (-1, 0)]

        while queue:
            rotten_r, rotten_c, time = queue.popleft()
            grid[rotten_r][rotten_c] = 2
            seconds = max(time, seconds)

            for dir_r, dir_c in directions:
                new_r, new_c = rotten_r + dir_r, rotten_c + dir_c

                if (
                    0 <= new_r < rows
                    and 0 <= new_c < cols
                    and grid[new_r][new_c] == 1
                    and (new_r, new_c, time) not in queue
                ):
                    queue.append((new_r, new_c, time + 1))

        for row in grid:
            if 1 in row:
                return -1

        return seconds


class Tests(unittest.TestCase):
    def test_example1(self):
        grid = [[2, 1, 1], [1, 1, 0], [0, 1, 1]]
        self.assertEqual(Solution().orangesRotting(grid), 4)

    def test_example2(self):
        grid = [[2, 1, 1], [0, 1, 1], [1, 0, 1]]
        self.assertEqual(Solution().orangesRotting(grid), -1)

    def test_example3(self):
        grid = [[0, 2]]
        self.assertEqual(Solution().orangesRotting(grid), 0)

    def test_example4(self):
        grid = [[2, 2], [1, 1], [0, 0], [2, 0]]
        self.assertEqual(Solution().orangesRotting(grid), 1)


if __name__ == "__main__":
    unittest.main()
