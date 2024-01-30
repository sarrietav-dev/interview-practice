# Definition for a binary tree node.
from collections import deque
from typing import List, Optional
import unittest


class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    def levelOrder(self, root: Optional[TreeNode]) -> List[List[int]]:
        if not root:
            return []

        queue = deque([(root, 0)])
        result = []

        while queue:
            node, height = queue.popleft()
            if len(result) - 1 < height:
                result.append([node.val])
            else:
                result[height].append(node.val)

            if node.left:
                queue.append((node.left, height + 1))
            if node.right:
                queue.append((node.right, height + 1))

        return result


class Test(unittest.TestCase):
    def test_1(self):
        s = Solution()
        root = TreeNode(3, TreeNode(9), TreeNode(20, TreeNode(15), TreeNode(7)))
        self.assertEqual(s.levelOrder(root), [[3], [9, 20], [15, 7]])

    def test_2(self):
        s = Solution()
        root = TreeNode(1)
        self.assertEqual(s.levelOrder(root), [[1]])

    def test_3(self):
        s = Solution()
        self.assertEqual(s.levelOrder([]), [])


if __name__ == "__main__":
    unittest.main()
