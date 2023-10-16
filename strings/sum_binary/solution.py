import unittest


class Solution:
    def addBinary(self, a: str, b: str) -> str:
        def sum_bit(a, b):
            if a == "1" and b == "1":
                return "0", "1"
            elif a == "1" or b == "1":
                return "1", None
            else:
                return "0", None

        if len(a) != len(b):
            max_len = max(len(a), len(b))
            if max_len == len(a):
                zeroes = len(a) - len(b)
                b = "0" * zeroes + b
            else:
                zeroes = len(b) - len(a)
                a = "0" * zeroes + a

        carry = None
        result = ""
        a = a[::-1]
        b = b[::-1]

        print(a, b)

        for num1, num2 in zip(a, b):
            bit, carrying = sum_bit(num1, num2)
            if carry:
                bit2, carrying2 = sum_bit(bit, carry)
                if carrying2:
                    carry = carrying2
                result += bit2
            else:
                carry = carrying
                result += bit

        if carry:
            result += carry
        return result[::-1]


class TestSolution(unittest.TestCase):
    def test_addBinary1(self):
        solution = Solution()
        self.assertEqual(solution.addBinary("11", "1"), "100")

    def test_addBinary2(self):
        solution = Solution()
        self.assertEqual(solution.addBinary("1010", "1011"), "10101")

    def test_addBinary3(self):
        solution = Solution()
        self.assertEqual(solution.addBinary("1111", "1111"), "11110")



if __name__ == "__main__":
    unittest.main()
