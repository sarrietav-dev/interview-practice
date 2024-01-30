# Definition for a Node.
class Node:
    def __init__(self, val=0, neighbors=None):
        self.val = val
        self.neighbors = neighbors if neighbors is not None else []


from collections import deque
from typing import Optional


class Solution:
    def cloneGraph(self, node: Optional["Node"]) -> Optional["Node"]:
        if not node:
            return node

        if node.neighbors == []:
            return Node(node.val, None)

        queue, created = deque([node]), {node.val: Node(node.val)}

        while queue:
            current = queue.popleft()
            for neighbor in current.neighbors:
                if neighbor.val not in created:
                    created[neighbor.val] = Node(neighbor.val)
                    queue.append(neighbor)
                created[current.val].neighbors.append(created[neighbor.val])

        return created[node.val]


if __name__ == "__main__":
    node1 = Node(1)
    node2 = Node(2)
    node3 = Node(3)
    node4 = Node(4)

    node1.neighbors = [node2, node4]
    node2.neighbors = [node1, node3]
    node3.neighbors = [node2, node4]
    node4.neighbors = [node1, node3]

    response = Solution().cloneGraph(node1)
