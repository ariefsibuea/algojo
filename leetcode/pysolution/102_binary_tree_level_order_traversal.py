"""
LeetCode Problem : Binary Tree Level Order Traversal
Topic            : Tree, Breadth-First Search, Binary Tree
Level            : Medium
URL              : https://leetcode.com/problems/binary-tree-level-order-traversal/description/
"""

from collections import deque
from typing import List, Optional


# Default definition for a binary tree node.
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    def levelOrder(self, root: Optional[TreeNode]) -> List[List[int]]:
        """Returns the level order traversal of a binary tree's nodes' values.
        
        Args:
            root (Optional[TreeNode]): Root node of binary tree, can be None.
            
        Returns:
            List[List[int]]: List of lists containing node values at each level from top to bottom.
            
        Time Complexity:
            O(n): Where n is the number of nodes in the binary tree.
            
        Space Complexity:
            O(n): Where the queue stores up to n nodes in the worst case.
        """

        if not root:
            return []

        queue = deque([root])
        result = []

        while queue:
            level_size = len(queue)
            current_level_values = []

            for _ in range(level_size):
                node = queue.popleft()
                current_level_values.append(node.val)
                if node.left:
                    queue.append(node.left)
                if node.right:
                    queue.append(node.right)

            result.append(current_level_values)

        return result


def run_tests():
    btree_case_1 = TreeNode(
        val=3, left=TreeNode(val=9), right=TreeNode(val=20, left=TreeNode(val=15), right=TreeNode(val=7))
    )

    inputs = {"case_1": [btree_case_1]}
    outputs = {"case_1": [[3], [9, 20], [15, 7]]}

    solution = Solution()

    for case, input in inputs.items():
        result = solution.levelOrder(input[0])
        assert result == outputs[case], f"{case}: expected {outputs[case]}, got {result}"

    print("✅ All tests passed!")


if __name__ == "__main__":
    run_tests()
