"""
LeetCode Problem : Lowest Common Ancestor of a Binary Search Tree
Topic            : Tree, Depth-First Search, Binary Search Tree, Binary Tree
Level            : Medium
URL              : https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree/description/
"""

from typing import Any


# Default definition for a binary tree node.
class TreeNode:
    def __init__(self, x, l=None, r=None):
        self.val = x
        self.left = l
        self.right = r


class Solution:
    def lowestCommonAncestor(self, root: "TreeNode", p: "TreeNode", q: "TreeNode") -> "TreeNode":
        """
        Given a binary search tree (BST), find the lowest common ancestor (LCA) node of two given nodes in the BST.
        The lowest common ancestor is defined between two nodes p and q as the lowest node in T that has both p and q as descendants (where we allow a node to be a descendant of itself).
        Args:
            root (TreeNode): The root node of the binary search tree
            p (TreeNode): First node to find ancestor for
            q (TreeNode): Second node to find ancestor for
        Returns:
            TreeNode: The lowest common ancestor node of p and q
        Example:
            Input: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 8
            Output: 6
            Explanation: The LCA of nodes 2 and 8 is 6.
        Solution:
            Iterative searching in binary search tree
        Time complexity:
            O(h) where h is the height of the tree
        Space complexity:
            O(1)
        """

        current = root

        while current:
            if current.val < p.val and current.val < q.val:
                current = current.right
            elif current.val > p.val and current.val > q.val:
                current = current.left
            else:
                return current

        return current


def run_tests():
    btree_case_1 = TreeNode(
        x=6,
        l=TreeNode(x=2, l=TreeNode(x=0), r=TreeNode(x=4, l=TreeNode(x=3), r=TreeNode(x=5))),
        r=TreeNode(x=8, l=TreeNode(x=7), r=TreeNode(x=9)),
    )

    inputs = {"case_1": [btree_case_1, TreeNode(x=2), TreeNode(x=8)]}
    outputs = {"case_1": 6}

    solution = Solution()

    for case, input in inputs.items():
        result = solution.lowestCommonAncestor(input[0], input[1], input[2])
        assert result.val == outputs[case], f"{case}: expected {outputs[case]}, got {result}"

    print("✅ All tests passed!")


if __name__ == "__main__":
    run_tests()
