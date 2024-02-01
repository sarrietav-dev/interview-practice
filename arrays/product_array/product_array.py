from typing import List
import unittest


class Solution:
    def productExceptSelf(self, nums: List[int]) -> List[int]:
        prefix = 1
        suffix = 1
        result = [0] * len(nums)
        for i in range(len(nums)):
            result[i] = prefix
            prefix *= nums[i]

        for i in range(len(nums) - 1, -1, -1):
            result[i] *= suffix
            suffix *= nums[i]

        return result


class Tests(unittest.TestCase):
    def test_example1(self):
        self.assertEqual(Solution().productExceptSelf([1, 2, 3, 4]), [24, 12, 8, 6])

    def test_example2(self):
        self.assertEqual(
            Solution().productExceptSelf([-1, 1, 0, -3, 3]), [0, 0, 9, 0, 0]
        )

    def test_example3(self):
        self.assertEqual(
            Solution().productExceptSelf([1, 2, 3, 4, 5]), [120, 60, 40, 30, 24]
        )


if __name__ == "__main__":
    unittest.main()
