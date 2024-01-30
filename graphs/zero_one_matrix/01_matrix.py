from collections import deque
import math
from typing import List
import unittest


class Solution:
    def updateMatrixBfs(self, mat: List[List[int]]) -> List[List[int]]:
        if not mat:
            return []

        ROW, COL = 0, 1
        visited = set()
        rows, cols = len(mat), len(mat[0])
        directions = ((0, 1), (0, -1), (1, 0), (-1, 0))

        queue = deque()

        for i in range(rows):
            for j in range(cols):
                if mat[i][j] == 0:
                    queue.append((i, j))
                else:
                    mat[i][j] = -1

        while queue:
            curr_i, curr_j = queue.popleft()

            visited.add((curr_i, curr_j))

            for direction in directions:
                next_i, next_j = (
                    curr_i + direction[ROW],
                    curr_j + direction[COL],
                )

                if (
                    0 <= next_i < rows
                    and 0 <= next_j < cols
                    and mat[next_i][next_j] == -1
                ):
                    mat[next_i][next_j] = mat[curr_i][curr_j] + 1
                    queue.append((next_i, next_j))

        return mat

    def updateMatrix(self, mat: List[List[int]]) -> List[List[int]]:
        rows, cols = len(mat), len(mat[0])

        for r in range(rows):
            for c in range(cols):
                if mat[r][c] > 0:
                    top = mat[r - 1][c] if r > 0 else math.inf
                    left = mat[r][c - 1] if c > 0 else math.inf
                    mat[r][c] = min(top, left) + 1

        for r in range(rows - 1, -1, -1):
            for c in range(cols - 1, -1, -1):
                if mat[r][c] > 0:
                    bottom = mat[r + 1][c] if r < rows - 1 else math.inf
                    right = mat[r][c + 1] if c < cols - 1 else math.inf
                    mat[r][c] = min(mat[r][c], bottom + 1, right + 1)

        return mat


class Tests(unittest.TestCase):
    def setUp(self) -> None:
        self.sol = Solution()

    def test_example_1(self) -> None:
        mat = [[0, 0, 0], [0, 1, 0], [0, 0, 0]]
        expected = [[0, 0, 0], [0, 1, 0], [0, 0, 0]]
        self.assertEqual(self.sol.updateMatrix(mat), expected)

    def test_example_2(self) -> None:
        mat = [[0, 0, 0], [0, 1, 0], [1, 1, 1]]
        expected = [[0, 0, 0], [0, 1, 0], [1, 2, 1]]
        self.assertEqual(self.sol.updateMatrix(mat), expected)

    def test_example_3(self) -> None:
        mat = [
            [1, 1, 0, 0, 1, 0, 0, 1, 1, 0],
            [1, 0, 0, 1, 0, 1, 1, 1, 1, 1],
            [1, 1, 1, 0, 0, 1, 1, 1, 1, 0],
            [0, 1, 1, 1, 0, 1, 1, 1, 1, 1],
            [0, 0, 1, 1, 1, 1, 1, 1, 1, 0],
            [1, 1, 1, 1, 1, 1, 0, 1, 1, 1],
            [0, 1, 1, 1, 1, 1, 1, 0, 0, 1],
            [1, 1, 1, 1, 1, 0, 0, 1, 1, 1],
            [0, 1, 0, 1, 1, 0, 1, 1, 1, 1],
            [1, 1, 1, 0, 1, 0, 1, 1, 1, 1],
        ]
        expected = [
            [2, 1, 0, 0, 1, 0, 0, 1, 1, 0],
            [1, 0, 0, 1, 0, 1, 1, 2, 2, 1],
            [1, 1, 1, 0, 0, 1, 2, 2, 1, 0],
            [0, 1, 2, 1, 0, 1, 2, 3, 2, 1],
            [0, 0, 1, 2, 1, 2, 1, 2, 1, 0],
            [1, 1, 2, 3, 2, 1, 0, 1, 1, 1],
            [0, 1, 2, 3, 2, 1, 1, 0, 0, 1],
            [1, 2, 1, 2, 1, 0, 0, 1, 1, 2],
            [0, 1, 0, 1, 1, 0, 1, 2, 2, 3],
            [1, 2, 1, 0, 1, 0, 1, 2, 3, 4],
        ]
        self.assertEqual(self.sol.updateMatrix(mat), expected)


if __name__ == "__main__":
    unittest.main()
