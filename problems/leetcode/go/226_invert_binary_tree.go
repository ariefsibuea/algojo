package main

import (
	"fmt"
	"os"

	"github.com/ariefsibuea/algojo/libs/go/cmp"
)

/*
 * LeetCode Problem : Invert Binary Tree
 * Topics           : Tree, Depth-First Search, Breadth-First Search, Binary Tree
 * Level            : Easy
 * URL              : https://leetcode.com/problems/invert-binary-tree/
 * Description      : Given the root of a binary tree, invert the tree, and return its root.
 * Examples         :
 *        			Example 1:
 *        			Input: root = [4,2,7,1,3,6,9]
 *        			Output: [4,7,2,9,6,3,1]
 *
 *        			Example 2:
 *        			Input: root = [2,1,3]
 *        			Output: [2,3,1]
 *
 *        			Example 3:
 *        			Input: root = []
 *        			Output: []
 */

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)
	return root
}

func RunTestInvertTree() {
	treeCase1 := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 1},
			Right: &TreeNode{Val: 3},
		},
		Right: &TreeNode{
			Val:   7,
			Left:  &TreeNode{Val: 6},
			Right: &TreeNode{Val: 9},
		},
	}

	treeCase2 := &TreeNode{
		Val:   2,
		Left:  &TreeNode{Val: 1},
		Right: &TreeNode{Val: 3},
	}

	testCases := map[string]struct {
		root   *TreeNode
		expect []int
	}{
		"case-1": {
			root:   treeCase1,
			expect: []int{4, 7, 2, 9, 6, 3, 1},
		},
		"case-2": {
			root:   treeCase2,
			expect: []int{2, 3, 1},
		},
		"case-3": {
			root:   nil,
			expect: []int{},
		},
	}

	for name, testCase := range testCases {
		fmt.Printf("RUN %s\n", name)

		head := invertTree(testCase.root)
		result := TraverseTreeP266(head)

		if !cmp.EqualSlices(result, testCase.expect) {
			fmt.Printf("=== FAILED: expect = %v - got = %v\n", testCase.expect, result)
			os.Exit(1)
		}
		fmt.Printf("=== PASSED\n")
	}

	fmt.Printf("\n✅ All tests passed!\n")
}

func TraverseTreeP266(head *TreeNode) []int {
	result := []int{}

	if head == nil {
		return result
	}

	queue := []TreeNode{*head}

	for len(queue) != 0 {
		levelValues := []int{}

		for _, currentNode := range queue {
			// dequeue
			queue = queue[1:]

			levelValues = append(levelValues, currentNode.Val)

			if currentNode.Left != nil {
				queue = append(queue, *currentNode.Left)
			}
			if currentNode.Right != nil {
				queue = append(queue, *currentNode.Right)
			}
		}

		result = append(result, levelValues...)
	}

	return result
}
