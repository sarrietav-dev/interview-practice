from collections import deque
import math
from typing import List
import unittest


class Solution:
    def evalRPN(self, tokens: List[str]) -> int:
        stack = deque([])
        for token in tokens:
            if token.lstrip("-").isdigit():
                stack.append(token)
            else:
                n1 = stack.pop()
                n2 = stack.pop()

                result = self.operate(n1, n2, token)

                stack.append(result)

        return int(stack.pop())

    def operate(self, n1, n2, operation):
        if operation == "+":
            return str(int(n1) + int(n2))
        if operation == "-":
            return str(int(n2) - int(n1))
        if operation == "*":
            return str(int(n1) * int(n2))
        if operation == "/":
            return str(math.trunc(int(n2) / int(n1)))


class Test(unittest.TestCase):
    def test_evalRPN_1(self):
        actual = Solution().evalRPN(["2", "1", "+", "3", "*"])
        expected = 9
        self.assertEqual(actual, expected)

    def test_evalRPN_2(self):
        actual = Solution().evalRPN(["4", "13", "5", "/", "+"])
        expected = 6
        self.assertEqual(actual, expected)

    def test_evalRPN_3(self):
        actual = Solution().evalRPN(
            ["10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"]
        )
        expected = 22
        self.assertEqual(actual, expected)

    def test_evalRPN_2(self):
        actual = Solution().evalRPN(["4","3","-"])
        expected = 1
        self.assertEqual(actual, expected)


if __name__ == "__main__":
    unittest.main(verbosity=2)
