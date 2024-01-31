from typing import Optional
import unittest


class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    def isValidBST(self, root: Optional[TreeNode]) -> bool:
        return self.isValid(root, [])

    def isValid(self, root, nodes=[]):
        if not root:
            return True

        is_valid = self.isValid(root.left, nodes)

        if not is_valid:
            return False

        if len(nodes) > 0:
            if nodes[-1].val >= root.val:
                return False

        nodes.append(root)

        is_valid = self.isValid(root.right, nodes)

        if not is_valid:
            return False

        return True


class Tests(unittest.TestCase):
    def setUp(self) -> None:
        self.s = Solution()
        return super().setUp()

    def test_isValidBST_1(self):
        node = TreeNode(2, TreeNode(1), TreeNode(3))

        actual = self.s.isValidBST(node)
        expected = True

        self.assertEqual(actual, expected)

    def test_isValidBST_2(self):
        node = TreeNode(0)

        actual = self.s.isValidBST(node)
        expected = True

        self.assertEqual(actual, expected)

    def test_isValidBST_3(self):
        node = TreeNode(1, TreeNode(1))

        actual = self.s.isValidBST(node)
        expected = False

        self.assertEqual(actual, expected)

    def test_isValidBST_4(self):
        node = TreeNode(2, TreeNode(2), TreeNode(2))

        actual = self.s.isValidBST(node)
        expected = False

        self.assertEqual(actual, expected)


if __name__ == "__main__":
    unittest.main()
