"""
LeetCode Problem : Construct Binary Tree from Preorder and Inorder Traversal
Topic            : Tree, Array, Hash Table, Divide and Conquer, Binary Tree
Level            : Medium
URL              : https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/
Description      : Given two integer arrays preorder and inorder where preorder is the preorder traversal of a binary
        tree and inorder is the inorder traversal of the same tree, construct and return the binary tree.
Examples         :
        Example 1:
        Input: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
        Output: [3,9,20,null,null,15,7]

        Example 2:
        Input: preorder = [-1], inorder = [-1]
        Output: [-1]
"""

# Definition for a binary tree node.
from typing import List, Optional


class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    def buildTree(self, preorder: List[int], inorder: List[int]) -> Optional[TreeNode]:
        """Constructs a binary tree from preorder and inorder traversal lists.

        Algorithm: Divide and Conquer with Hash Map Optimization

        The algorithm works by using the properties of tree traversals:
        1. In preorder traversal, the first element is always the root of the tree
        2. In inorder traversal, elements to the left of the root are in the left subtree, and elements to the right are
            in the right subtree
        3. Uses a hash map to quickly find the position of root in inorder traversal
        4. Recursively builds left and right subtrees by calculating appropriate subarray indices

        Args:
            preorder: A list of integers representing the preorder traversal of a binary tree.
            inorder: A list of integers representing the inorder traversal of the same binary tree.

        Returns:
            The root node of the constructed binary tree.

        Time Complexity:
            O(n) where n is the number of nodes in the tree. Each node is processed exactly once,
            and the hash map lookup is O(1).

        Space Complexity:
            O(n) for the recursive call stack and the hash map storing inorder indices.
        """
        # Create a hash map of inorder values to their indices for O(1) lookup
        inorder_map = {}
        for i, v in enumerate(inorder):
            inorder_map[v] = i

        def buildTreeFromPreorderInorder(preorder_start, preorder_end, inorder_start, inorder_end):
            # Base case: if no elements to process, return None
            if preorder_start > preorder_end:
                return None

            # In preorder traversal, the first element is always the root
            root_value = preorder[preorder_start]
            root = TreeNode(root_value)

            # Find the position of the root in the inorder traversal
            inorder_root_index = inorder_map[root_value]

            # Calculate the size of the left subtree
            # All elements before root in inorder are in the left subtree
            left_subtree_size = inorder_root_index - inorder_start

            # Recursively build the left subtree
            # Preorder: starts right after root, ends after processing left_subtree_size elements
            # Inorder: starts at the beginning, ends right before the root
            root.left = buildTreeFromPreorderInorder(
                preorder_start + 1,  # Skip the root in preorder
                preorder_start + left_subtree_size,  # Process only left subtree elements
                inorder_start,  # Start of left subtree in inorder
                inorder_root_index - 1,  # End right before root in inorder
            )

            # Recursively build the right subtree
            # Preorder: starts after left subtree, ends at the end
            # Inorder: starts right after root, ends at the end
            root.right = buildTreeFromPreorderInorder(
                preorder_start + left_subtree_size + 1,  # Skip root and left subtree
                preorder_end,  # Process until the end
                inorder_root_index + 1,  # Start right after root in inorder
                inorder_end,  # Process until the end
            )

            return root

        # Start the recursive construction with the full arrays
        return buildTreeFromPreorderInorder(0, len(preorder) - 1, 0, len(inorder) - 1)
