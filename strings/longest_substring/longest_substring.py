from collections import deque
import math
import unittest


class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        if s == "":
            return 0

        chars = {}
        max_value = 0

        l, h = 0, 0

        while l <= h and h < len(s):
            if s[h] not in chars:
                chars[s[h]] = 1
                curr_len = len(s[l:h+1])
                if curr_len > max_value:
                    max_value = curr_len
            else:
                chars[s[h]] += 1
                while sum(chars.values()) > len(chars.keys()):
                    chars[s[l]] -= 1

                    if chars[s[l]] == 0:
                        del chars[s[l]]

                    l += 1
            h += 1

        return max_value


class Tests(unittest.TestCase):
    def test_1(self):
        s = Solution()
        self.assertEqual(s.lengthOfLongestSubstring("abcabcbb"), 3)

    def test_2(self):
        s = Solution()
        self.assertEqual(s.lengthOfLongestSubstring("bbbbb"), 1)

    def test_3(self):
        s = Solution()
        self.assertEqual(s.lengthOfLongestSubstring("pwwkew"), 3)

    def test_4(self):
        s = Solution()
        self.assertEqual(s.lengthOfLongestSubstring(""), 0)

    def test_5(self):
        s = Solution()
        self.assertEqual(s.lengthOfLongestSubstring(" "), 1)

    def test_6(self):
        s = Solution()
        self.assertEqual(s.lengthOfLongestSubstring("au"), 2)

    def test_7(self):
        s = Solution()
        self.assertEqual(s.lengthOfLongestSubstring("dvdf"), 3)

    def test_8(self):
        s = Solution()
        self.assertEqual(s.lengthOfLongestSubstring("abba"), 2)

    def test_9(self):
        s = Solution()
        self.assertEqual(s.lengthOfLongestSubstring("tmmzuxt"), 5)

    def test_10(self):
        s = Solution()
        self.assertEqual(s.lengthOfLongestSubstring("ohvhjdml"), 6)

    def test_11(self):
        s = Solution()
        self.assertEqual(s.lengthOfLongestSubstring("jbpnbwwd"), 4)


if __name__ == "__main__":
    unittest.main()
