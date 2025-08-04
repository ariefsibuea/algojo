"""
LeetCode Problem : Balanced Binary Tree
Topic            : Tree, Depth-First Search, Binary Tree
Level            : Easy
URL              : https://leetcode.com/problems/balanced-binary-tree/description/
"""

from typing import Optional


# Default definition for a binary tree node.
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    def isBalanced(self, root: Optional[TreeNode]) -> bool:
        """Determines if a binary tree is height-balanced (subtrees of every node differ in height by no more than one).

        Args:
            root (Optional[TreeNode]): Root node of binary tree.

        Returns:
            bool: True if tree is height-balanced, False otherwise.

        Time Complexity:
            O(n): Where n is the number of nodes in the tree, as we visit each node once.

        Space Complexity:
            O(h): Where h is the height of the tree due to the recursion stack.
        """

        def dfs(node: Optional[TreeNode]):
            if not node:
                return (True, 0)

            left = dfs(node.left)
            right = dfs(node.right)

            height = max(left[1], right[1]) + 1
            balanced = abs(left[1] - right[1]) <= 1 and left[0] and right[0]

            return (balanced, height)

        return dfs(root)[0]


def run_tests():
    btree_case_1 = TreeNode(3)
    btree_case_1.left = TreeNode(9)
    btree_case_1.right = TreeNode(20)
    btree_case_1.right.left = TreeNode(15)
    btree_case_1.right.right = TreeNode(7)

    btree_case_2 = TreeNode(1)
    btree_case_2.left = TreeNode(2)
    btree_case_2.left.left = TreeNode(3)
    btree_case_2.left.right = TreeNode(3)
    btree_case_2.left.left.left = TreeNode(4)
    btree_case_2.left.left.right = TreeNode(4)
    btree_case_2.right = TreeNode(2)

    inputs = {"case_1": [btree_case_1], "case_2": [btree_case_2], "case_3": [None]}
    outputs = {"case_1": True, "case_2": False, "case_3": True}

    solution = Solution()

    for case, input in inputs.items():
        result = solution.isBalanced(input[0])
        assert result == outputs[case], f"{case}: expected {outputs[case]}, got {result}"

    print("✅ All tests passed!")


if __name__ == "__main__":
    run_tests()
