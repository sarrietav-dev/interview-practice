import unittest


class Solution:
    def search(self, nums: list[int], target: int) -> int:
        low, up = 0, len(nums) - 1
        while low <= up:
            mid = (up + low) // 2
            
            if nums[mid] == target:
                return mid
            elif target < nums[mid]:
                up = mid - 1
            else:
                low = mid + 1
        
        return -1

class TestSolution(unittest.TestCase):
    def setUp(self):
        self.solution = Solution()

    def test_1(self):
        nums = [-1,0,3,5,9,12]
        target = 9
        self.assertEqual(self.solution.search(nums, target), 4)

    def test_2(self):
        nums = [-1,0,3,5,9,12]
        target = 2
        self.assertEqual(self.solution.search(nums, target), -1)

    def test_3(self):
        nums = [5]
        target = 5
        self.assertEqual(self.solution.search(nums, target), 0)

    def test_4(self):
        nums = [5]
        target = -5
        self.assertEqual(self.solution.search(nums, target), -1)

    def test_5(self):
        nums = [5]
        target = 6
        self.assertEqual(self.solution.search(nums, target), -1)

    def test_6(self):
        nums = [5, 6]
        target = 5
        self.assertEqual(self.solution.search(nums, target), 0)

    def test_7(self):
        nums = [5, 6]
        target = 6
        self.assertEqual(self.solution.search(nums, target), 1)

    def test_8(self):
        nums = [5, 6]
        target = 4
        self.assertEqual(self.solution.search(nums, target), -1)

    def test_9(self):
        nums = [5, 6]
        target = 7
        self.assertEqual(self.solution.search(nums, target), -1)

    def test_10(self):
        nums = [5, 6, 7]
        target = 5
        self.assertEqual(self.solution.search(nums, target), 0)

    def test_11(self):
        nums = [5, 6, 7]
        target = 6
        self.assertEqual(self.solution.search(nums, target), 1)

    def test_12(self):
        nums = [5, 6, 7]
        target = 7
        self.assertEqual(self.solution.search(nums, target), 2)

    def test_13(self):
        nums = [5, 6, 7]
        target = 4
        self.assertEqual(self.solution.search(nums, target), -1)

    def test_14(self):
        nums = [5, 6, 7]
        target = 8
        self.assertEqual(self.solution.search(nums, target), -1)


if __name__ == "__main__":
    unittest.main()