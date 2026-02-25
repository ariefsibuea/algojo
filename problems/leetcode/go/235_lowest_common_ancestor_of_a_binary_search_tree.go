package main

/*
LeetCode Problem : Lowest Common Ancestor of a Binary Search Tree
Topic            : Tree, Depth-First Search, Binary Search Tree, Binary Tree
Level            : Medium
URL              : https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree/
Description      : Given a binary search tree (BST), find the lowest common ancestor (LCA) node of two given nodes in
        the BST. According to the definition of LCA on Wikipedia: "The lowest common ancestor is defined between two
        nodes p and q as the lowest node in T that has both p and q as descendants (where we allow a node to be a
        descendant of itself)."
Examples         :
        Example 1:
        Input: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 8
        Output: 6
        Explanation: The LCA of nodes 2 and 8 is 6.

        Example 2:
        Input: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 4
        Output: 2
        Explanation: The LCA of nodes 2 and 4 is 2, since a node can be a descendant of itself according to the LCA
                definition.

        Example 3:
        Input: root = [2,1], p = 2, q = 1
        Output: 2
*/

import (
	"fmt"
	"os"

	"github.com/ariefsibuea/algojo/libs/go/cmp"
)

func init() {
	register("LowestCommonAncestor", RunTestLowestCommonAncestor)
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	current := root

	for current != nil {
		if current.Val < p.Val && current.Val < q.Val {
			current = current.Right
		} else if current.Val > p.Val && current.Val > q.Val {
			current = current.Left
		} else {
			return current
		}
	}

	return current
}

func RunTestLowestCommonAncestor() {
	treeCase := &TreeNode{
		Val: 6,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 0,
			},
			Right: &TreeNode{
				Val:   4,
				Left:  &TreeNode{Val: 3},
				Right: &TreeNode{Val: 5},
			},
		},
		Right: &TreeNode{
			Val: 8,
			Left: &TreeNode{
				Val: 7,
			},
			Right: &TreeNode{
				Val: 9,
			},
		},
	}

	testCases := map[string]struct {
		root   *TreeNode
		p      *TreeNode
		q      *TreeNode
		expect int
	}{
		"case-1": {
			root:   treeCase,
			p:      &TreeNode{Val: 2},
			q:      &TreeNode{Val: 8},
			expect: 6,
		},
		"case-2": {
			root:   treeCase,
			p:      &TreeNode{Val: 2},
			q:      &TreeNode{Val: 4},
			expect: 2,
		},
	}

	for name, testCase := range testCases {
		fmt.Printf("RUN %s\n", name)

		result := lowestCommonAncestor(testCase.root, testCase.p, testCase.q)
		if !cmp.EqualNumbers(result.Val, testCase.expect) {
			fmt.Printf("=== FAILED: expect = %v - got = %v\n", testCase.expect, result)
			os.Exit(1)
		}
		fmt.Printf("=== PASSED\n")
	}

	fmt.Printf("\n✅ All tests passed!\n")
}
