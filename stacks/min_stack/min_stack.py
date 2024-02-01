from collections import deque
import unittest


class MinStack:
    def __init__(self):
        self.stack = deque([])

    def push(self, val: int) -> None:
        if not self.stack:
            self.stack.append((val, val))
            return

        curr_min = self.getMin()

        if val <= curr_min:
            self.stack.append((val, val))
        else:
            self.stack.append((val, curr_min))

    def pop(self) -> None:
        self.stack.pop()

    def top(self) -> int:
        val, _ = self.stack[-1]
        return val

    def getMin(self) -> int:
        _, curr_min = self.stack[-1]
        return curr_min


class Tests(unittest.TestCase):
    def test_example1(self):
        minStack = MinStack()
        minStack.push(-2)
        minStack.push(0)
        minStack.push(-3)
        self.assertEqual(minStack.getMin(), -3)
        minStack.pop()
        self.assertEqual(minStack.top(), 0)
        self.assertEqual(minStack.getMin(), -2)

    def test_example2(self):
        minStack = MinStack()
        minStack.push(0)
        minStack.push(1)
        minStack.push(0)
        self.assertEqual(minStack.getMin(), 0)
        minStack.pop()
        self.assertEqual(minStack.getMin(), 0)

    def test_example3(self):
        minStack = MinStack()
        minStack.push(2147483646)
        minStack.push(2147483646)
        minStack.push(2147483647)
        self.assertEqual(minStack.top(), 2147483647)
        minStack.pop()
        self.assertEqual(minStack.getMin(), 2147483646)
        minStack.pop()
        self.assertEqual(minStack.getMin(), 2147483646)
        minStack.pop()
        minStack.push(2147483647)
        self.assertEqual(minStack.top(), 2147483647)
        self.assertEqual(minStack.getMin(), 2147483647)
        minStack.push(-2147483648)
        self.assertEqual(minStack.top(), -2147483648)
        self.assertEqual(minStack.getMin(), -2147483648)
        minStack.pop()
        self.assertEqual(minStack.getMin(), 2147483647)


if __name__ == "__main__":
    unittest.main()
