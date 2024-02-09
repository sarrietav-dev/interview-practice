from typing import List
import unittest


class Solution:
    def search(self, nums: List[int], target: int) -> int:
        l, r = 0, len(nums) - 1

        while l < r:
            mid = l + ((r - l) // 2)
            if nums[l] > nums[mid]:
                r = mid
            elif nums[mid] > nums[r]:
                l = mid + 1
            else:
                r = mid - 1
        else:
            mid = l

        l, r = 0, len(nums) - 1

        while l <= r:
            if nums[mid] == target:
                return mid
            if nums[mid] <= target <= nums[r]:
                l = mid + 1
            else:
                r = mid - 1
            mid = l + ((r - l) // 2)
        return -1


class Test(unittest.TestCase):
    def test_search(self):
        solution = Solution()
        self.assertEqual(solution.search([4, 5, 6, 7, 0, 1, 2], 0), 4)

    def test_search_2(self):
        solution = Solution()
        self.assertEqual(solution.search([4, 5, 6, 7, 0, 1, 2], 3), -1)

    def test_search_3(self):
        solution = Solution()
        self.assertEqual(solution.search([1], 0), -1)

    def test_search_4(self):
        solution = Solution()
        self.assertEqual(solution.search([1, 3], 3), 1)

    def test_search_5(self):
        solution = Solution()
        self.assertEqual(solution.search([3, 1], 3), 0)

    def test_search_6(self):
        solution = Solution()
        self.assertEqual(solution.search([3, 1], 1), 1)

    def test_search_7(self):
        solution = Solution()
        self.assertEqual(solution.search([3, 1], 0), -1)

    def test_search_8(self):
        solution = Solution()
        self.assertEqual(solution.search([3, 1], 2), -1)

    def test_search_9(self):
        solution = Solution()
        self.assertEqual(solution.search([3, 1], 4), -1)

    def test_search_10(self):
        solution = Solution()
        self.assertEqual(solution.search([3, 1, 2], 2), 2)

    def test_search_11(self):
        solution = Solution()
        self.assertEqual(solution.search([3, 1, 2], 3), 0)

    def test_search_12(self):
        solution = Solution()
        self.assertEqual(solution.search([3, 1, 2], 1), 1)

    def test_search_13(self):
        solution = Solution()
        self.assertEqual(solution.search([3, 1, 2], 0), -1)

    def test_search_14(self):
        solution = Solution()
        self.assertEqual(solution.search([3, 1, 2], 4), -1)

    def test_search_15(self):
        solution = Solution()
        self.assertEqual(solution.search([3, 1, 2], 5), -1)


if __name__ == "__main__":
    unittest.main()
