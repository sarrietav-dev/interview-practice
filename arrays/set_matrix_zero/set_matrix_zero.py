from typing import List


class Solution:
    def setZeroes(self, matrix: List[List[int]]) -> None:
        """
        Do not return anything, modify matrix in-place instead.
        """
        cols = []
        rows = []

        for r in range(len(matrix)):
            for c in range(len(matrix[r])):
                if matrix[r][c] == 0:
                    cols += [c]
                    rows += [r]

        for c in cols:
            for r in range(len(matrix)):
                matrix[r][c] = 0

        for r in rows:
            for c in range(len(matrix[r])):
                matrix[r][c] = 0
