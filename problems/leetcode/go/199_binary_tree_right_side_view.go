package main

import (
	"fmt"
	"os"

	"github.com/ariefsibuea/algojo/libs/go/cmp"
)

func init() {
	register("BinaryTreeRightSideView", RunTestBinaryTreeRightSideView)
}

/*
 * Problem 			: Binary Tree Right Side View
 * Topics           : Tree, Depth-First Search, Breadth-First Search, Binary Tree
 * Level            : Medium
 * URL              : https://leetcode.com/problems/binary-tree-right-side-view
 * Description      : Given the root of a binary tree, imagine yourself standing on the right side of it, return the
 * 					  values of the nodes you can see ordered from top to bottom.
 * Examples         :
 * 					  Example 1:
 * 					  Input: root = [1,2,3,null,5,null,4]
 * 					  Output: [1,3,4]
 *
 * 					  Example 2:
 * 					  Input: root = [1,2,3,4,null,null,null,5]
 * 					  Output: [1,3,4,5]
 *
 * 					  Example 3:
 * 					  Input: root = [1,null,3]
 * 					  Output: [1,3]
 *
 * 					  Example 4:
 * 					  Input: root = []
 * 					  Output: []
 */

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	var (
		queue            = []*TreeNode{root}
		currentLevelSize = 0
		currentNode      *TreeNode
		result           = []int{}
	)

	for len(queue) > 0 {
		currentLevelSize = len(queue)

		for i := 0; i < currentLevelSize; i++ {
			// dequeu
			currentNode = queue[0]
			queue[0] = nil
			queue = queue[1:]

			if currentNode.Left != nil {
				queue = append(queue, currentNode.Left)
			}
			if currentNode.Right != nil {
				queue = append(queue, currentNode.Right)
			}
			if i == currentLevelSize-1 {
				result = append(result, currentNode.Val)
			}
		}
	}

	return result
}

func RunTestBinaryTreeRightSideView() {
	testCases := map[string]struct {
		root   *TreeNode
		expect []int
	}{
		"case-1": {
			root:   mockInputP199Case1(),
			expect: []int{1, 3, 4},
		},
		"case-2": {
			root:   mockInputP199Case2(),
			expect: []int{1, 3, 4, 5},
		},
		"case-3": {
			root:   nil,
			expect: []int{},
		},
	}

	for name, testCase := range testCases {
		fmt.Printf("RUN %s\n", name)
		result := rightSideView(testCase.root)
		if !cmp.EqualSlices(result, testCase.expect) {
			fmt.Printf("=== FAILED: expect = %v - got = %v\n", testCase.expect, result)
			os.Exit(1)
		}
		fmt.Printf("=== PASSED\n")
	}

	fmt.Printf("\n✅ All tests passed!\n")
}

func mockInputP199Case1() *TreeNode {
	node4 := &TreeNode{Val: 4}
	node3 := &TreeNode{Val: 3, Right: node4}

	node5 := &TreeNode{Val: 5}
	node2 := &TreeNode{Val: 2, Right: node5}

	return &TreeNode{Val: 1, Left: node2, Right: node3}
}

func mockInputP199Case2() *TreeNode {
	node3 := &TreeNode{Val: 3}

	node5 := &TreeNode{Val: 5}
	node4 := &TreeNode{Val: 4, Left: node5}
	node2 := &TreeNode{Val: 2, Left: node4}

	return &TreeNode{Val: 1, Left: node2, Right: node3}
}
