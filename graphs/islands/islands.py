from collections import deque
from typing import List
import unittest


class Solution:
    def numIslands(self, grid: List[List[str]]) -> int:
        islands = [0]
        parents = {}
        rows, cols = len(grid), len(grid[0])
        directions = [(0, 1), (0, -1), (1, 0), (-1, 0)]
        queue = deque([])

        def traverse(r, c):
            queue.append((r, c))
            # parents[(r, c)] = (r, c)
            islands[0] += 1

            while queue:
                curr_row, curr_col = queue.popleft()
                grid[curr_row][curr_col] = "-1"

                for dir_row, dir_col in directions:
                    new_row, new_col = curr_row + dir_row, curr_col + dir_col

                    if (
                        0 <= new_row < rows
                        and 0 <= new_col < cols
                        and grid[new_row][new_col] == "1"
                        and (new_row, new_col) not in queue
                    ):
                        # parents[(new_row, new_col)] = (curr_row, curr_col)
                        queue.append((new_row, new_col))

        for r in range(rows):
            for c in range(cols):
                if grid[r][c] == "1":
                    traverse(r, c)

        return islands[0]


class Tests(unittest.TestCase):
    def setUp(self) -> None:
        self.s = Solution()
        return super().setUp()

    def test_numIslands_1(self):
        grid = [
            ["1", "1", "1", "1", "0"],
            ["1", "1", "0", "1", "0"],
            ["1", "1", "0", "0", "0"],
            ["0", "0", "0", "0", "0"],
        ]

        actual = self.s.numIslands(grid)
        expected = 1

        self.assertEqual(actual, expected)

    def test_numIslands_2(self):
        grid = [
            ["1", "1", "0", "0", "0"],
            ["1", "1", "0", "0", "0"],
            ["0", "0", "1", "0", "0"],
            ["0", "0", "0", "1", "1"],
        ]

        actual = self.s.numIslands(grid)
        expected = 3

        self.assertEqual(actual, expected)

    def test_numIslands_3(self):
        grid = [["1", "1", "1"], ["0", "1", "0"], ["1", "1", "1"]]

        actual = self.s.numIslands(grid)
        expected = 1

        self.assertEqual(actual, expected)

    def test_numIslands_4(self):
        grid = [
            ["1", "0", "1", "1", "1"],
            ["1", "0", "1", "0", "1"],
            ["1", "1", "1", "0", "1"],
        ]

        actual = self.s.numIslands(grid)
        expected = 1

        self.assertEqual(actual, expected)

    def test_numIslands_5(self):
        grid = [["1", "1", "1"], ["0", "1", "0"], ["1", "1", "1"]]

        actual = self.s.numIslands(grid)
        expected = 1

        self.assertEqual(actual, expected)

    def test_numIslands_6(self):
        grid = [["1", "1", "1", "1", "1", "1", "1"]]

        actual = self.s.numIslands(grid)
        expected = 1

        self.assertEqual(actual, expected)

    def test_numIslands_7(self):
        grid = [
            [
                "1",
                "1",
                "1",
                "1",
                "1",
                "0",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "0",
                "1",
                "0",
                "1",
                "1",
            ],
            [
                "0",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "0",
                "1",
                "1",
                "1",
                "1",
                "1",
                "0",
            ],
            [
                "1",
                "0",
                "1",
                "1",
                "1",
                "0",
                "0",
                "1",
                "1",
                "0",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
            ],
            [
                "1",
                "1",
                "1",
                "1",
                "0",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
            ],
            [
                "1",
                "0",
                "0",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
            ],
            [
                "1",
                "0",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "0",
                "1",
                "1",
                "1",
                "0",
                "1",
                "1",
                "1",
                "0",
                "1",
                "1",
                "1",
            ],
            [
                "0",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "0",
                "1",
                "1",
                "0",
                "1",
                "1",
                "1",
                "1",
            ],
            [
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "0",
                "1",
                "1",
                "1",
                "1",
                "0",
                "1",
                "1",
            ],
            [
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "0",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
            ],
            [
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
            ],
            [
                "0",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "0",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
            ],
            [
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
            ],
            [
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
            ],
            [
                "1",
                "1",
                "1",
                "1",
                "1",
                "0",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "0",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
            ],
            [
                "1",
                "0",
                "1",
                "1",
                "1",
                "1",
                "1",
                "0",
                "1",
                "1",
                "1",
                "0",
                "1",
                "1",
                "1",
                "1",
                "0",
                "1",
                "1",
                "1",
            ],
            [
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "0",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "0",
            ],
            [
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "0",
                "1",
                "1",
                "1",
                "1",
                "0",
                "0",
            ],
            [
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
            ],
            [
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
            ],
            [
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
                "1",
            ],
        ]

        actual = self.s.numIslands(grid)
        expected = 1

        self.assertEqual(actual, expected)


if __name__ == "__main__":
    unittest.main(verbosity=2)
