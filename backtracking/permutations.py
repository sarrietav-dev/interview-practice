from typing import List
import unittest


class Solution:
    def permute(self, nums: List[int]) -> List[List[int]]:
        permutations = []

        def dfs(perm: List[int]):
            result = []
            for x in [x for x in range(len(nums)) if nums[x] not in perm]:
                result += dfs(perm + [nums[x]])

            if len(result) == 0:
                return [perm]
            else:
                return result

        for num in nums:
            permutations += dfs([num])

        return permutations


class Test(unittest.TestCase):
    def test_permute(self):
        solution = Solution()
        self.assertEqual(
            solution.permute([1, 2, 3]),
            [[1, 2, 3], [1, 3, 2], [2, 1, 3], [2, 3, 1], [3, 1, 2], [3, 2, 1]],
        )

    def test_permute_2(self):
        solution = Solution()
        self.assertEqual(solution.permute([0, 1]), [[0, 1], [1, 0]])

    def test_permute_3(self):
        solution = Solution()
        self.assertEqual(solution.permute([1]), [[1]])

    def test_permute_4(self):
        solution = Solution()
        self.assertEqual(solution.permute([1, 2]), [[1, 2], [2, 1]])


if __name__ == "__main__":
    unittest.main()
