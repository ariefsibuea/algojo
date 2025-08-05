"""
LeetCode Problem : Invert Binary Tree
Topic            : Tree, Depth-First Search, Breadth-First Search, Binary Tree
Level            : Easy
URL              : https://leetcode.com/problems/invert-binary-tree/description/
Description      : Given the root of a binary tree, invert the tree, and return its root.
Examples         :
        Example 1:
        Input: root = [4,2,7,1,3,6,9]
        Output: [4,7,2,9,6,3,1]

        Example 2:
        Input: root = [2,1,3]
        Output: [2,3,1]

        Example 3:
        Input: root = []
        Output: []
"""

from collections import deque
from typing import List, Optional


class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    def invertTree(self, root: Optional[TreeNode]) -> Optional[TreeNode]:
        """Inverts a binary tree by recursively swapping the left and right children of all nodes.

        Args:
            root (Optional[TreeNode]): The root node of the binary tree.

        Returns:
            Optional[TreeNode]: The root node of the inverted binary tree.

        Time Complexity:
            O(n): Where n is the number of nodes in the tree, as each node is visited once.

        Space Complexity:
            O(h): Where h is the height of the tree due to the recursion stack.
        """

        if not root:
            return None

        root.left, root.right = self.invertTree(root.right), self.invertTree(root.left)
        return root


def run_tests():
    btree_case_1 = TreeNode(
        val=4,
        left=TreeNode(val=2, left=TreeNode(val=1), right=TreeNode(val=3)),
        right=TreeNode(val=7, left=TreeNode(val=6), right=TreeNode(val=9)),
    )

    inputs = {"case_1": [btree_case_1]}
    outputs = {"case_1": [4, 7, 2, 9, 6, 3, 1]}

    solution = Solution()

    for case, input in inputs.items():
        result = solution.invertTree(input[0])

        array_result = tree_to_list_bfs(result)
        assert array_result == outputs[case], f"{case}: expected {outputs[case]}, got {result}"

    print("✅ All tests passed!")


def tree_to_list_bfs(root: Optional[TreeNode]) -> List:
    if not root:
        return []

    result = []
    queue = deque([root])

    while queue:
        node = queue.popleft()
        result.append(node.val)

        if node.left:
            queue.append(node.left)
        if node.right:
            queue.append(node.right)

    return result


if __name__ == "__main__":
    run_tests()
