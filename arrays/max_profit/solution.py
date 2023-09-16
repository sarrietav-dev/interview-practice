import unittest


class Solution:
    def maxProfit(self, prices: list[int]) -> int:
        min_day_found = float('Inf')
        max_profit_found = 0

        for price in prices:
            if price < min_day_found:
                min_day_found = price
                continue

            profit = price - min_day_found

            if profit > max_profit_found:
                max_profit_found = profit

        return max_profit_found
    
class TestSolution(unittest.TestCase):
    def setUp(self):
        self.s = Solution()

    def test_one(self):
        self.assertEqual(self.s.maxProfit([7,1,5,3,6,4]), 5)

    def test_two(self):
        self.assertEqual(self.s.maxProfit([7,6,4,3,1]), 0)

if __name__ == '__main__':
    unittest.main()