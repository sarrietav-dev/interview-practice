class Solution:
    def longestPalindrome(self, s: str) -> int:
        count = {}
        for c in s:
            if c not in count:
                count[c] = 0
            count[c] += 1

        longest = 0
        odd_flag = False

        for v in count.values():
            if v % 2 == 0:
                longest += v
            else:
                odd_flag = True
                longest += v - 1

        if odd_flag:
            return longest + 1
        else:
            return longest
