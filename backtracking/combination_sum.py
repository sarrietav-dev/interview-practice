from typing import List
import unittest


class Solution:
    def combinationSum(self, candidates: List[int], target: int) -> List[List[int]]:
        candidates = sorted(candidates)

        def dfs(i: int, res: List[int]):
            if sum(res) == target:
                return [res]

            if sum(res) > target:
                return []

            result = []
            for j in self.children(candidates, i):
                result += dfs(j, res + [candidates[j]])

            return result

        res = []

        for i in range(len(candidates)):
            res += dfs(i, [candidates[i]])

        return res

    def children(self, candidates: List[int], i: int):
        return range(i, len(candidates))


class Test(unittest.TestCase):
    def test_combinationSum(self):
        solution = Solution()
        self.assertEqual(solution.combinationSum([2, 3, 6, 7], 7), [[2, 2, 3], [7]])

    def test_combinationSum_2(self):
        solution = Solution()
        self.assertEqual(
            solution.combinationSum([2, 3, 5], 8), [[2, 2, 2, 2], [2, 3, 3], [3, 5]]
        )

    def test_combinationSum_3(self):
        solution = Solution()
        self.assertEqual(solution.combinationSum([2], 1), [])


if __name__ == "__main__":
    unittest.main()
