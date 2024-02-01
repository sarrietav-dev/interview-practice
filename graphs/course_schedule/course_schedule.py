from collections import deque
from typing import List
import unittest


class Solution:
    def canFinish(self, numCourses: int, prerequisites: List[List[int]]) -> bool:
        graph = {i: [] for i in range(numCourses)}
        for a, b in prerequisites:
            graph[b].append(a)

        degrees = {i: 0 for i in range(numCourses)}

        for k in graph.keys():
            for v in graph[k]:
                degrees[v] += 1

        queue = deque([])

        for k in degrees.keys():
            if degrees[k] == 0:
                queue.append(k)

        order = []

        while queue:
            node = queue.popleft()
            order.append(node)
            for neighbor in graph[node]:
                degrees[neighbor] -= 1
                if degrees[neighbor] == 0:
                    queue.append(neighbor)

        if len(order) != numCourses:
            return False

        return True


class Tests(unittest.TestCase):
    def test_example1(self):
        self.assertEqual(Solution().canFinish(2, [[1, 0]]), True)

    def test_example2(self):
        self.assertEqual(Solution().canFinish(2, [[1, 0], [0, 1]]), False)


if __name__ == "__main__":
    unittest.main()
