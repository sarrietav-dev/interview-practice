import unittest


class Solution:
    def isPalindrome(self, s: str) -> bool:
        s = ''.join(c for c in s if c.isalnum()).lower()    

        for i in range(len(s) // 2):
            if s[len(s) - 1 - i] != s[i]:
                return False
        return True
    

class TestSolution(unittest.TestCase):
    def setUp(self):
        self.solution = Solution()

    def test_panama(self):
        s = "A man, a plan, a canal: Panama"
        self.assertTrue(self.solution.isPalindrome(s))


if __name__ == "__main__":
    unittest.main()