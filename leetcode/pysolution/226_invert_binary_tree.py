"""
LeetCode Problem : Invert Binary Tree
Topic            : Tree, Depth-First Search, Breadth-First Search, Binary Tree
Level            : Easy
URL              : https://leetcode.com/problems/invert-binary-tree/description/
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
        """
        Inverts a binary tree.
        Given the root of a binary tree, this function recursively inverts the tree, swapping the left and right children of all nodes.
        Args:
            root (Optional[TreeNode]): The root node of the binary tree.
        Returns:
            Optional[TreeNode]: The root node of the inverted binary tree.
        Solution:
            Recursive: Post-Order Traversal
                - Invert the left and right subtrees recursively
                - Swap the left and right children of the current node
        Time Complexity:
            O(n): Each node is visited once
        Space Complexity:
            O(h): Tree height
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
