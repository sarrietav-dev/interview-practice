from typing import List


# https://leetcode.com/problems/find-all-anagrams-in-a-string/
class Solution:
    def findAnagrams(self, s: str, p: str) -> List[int]:
        p_map = {}
        result = []

        # O(p)
        for c in p:
            if c not in p_map:
                p_map[c] = 0
            p_map[c] += 1

        l, r = 0, len(p) - 1
        s_map = {}

        # O(r)
        for c in s[l : r + 1]:
            if c not in s_map:
                s_map[c] = 0
            s_map[c] += 1

        # O(s)
        while r < len(s) and l <= r:
            if s_map == p_map:
                result.append(l)

            s_map[s[l]] -= 1
            if s_map[s[l]] == 0:
                del s_map[s[l]]
            l += 1

            if r < len(s) - 1:
                r += 1
                if s[r] not in s_map:
                    s_map[s[r]] = 0
                s_map[s[r]] += 1

        return result
