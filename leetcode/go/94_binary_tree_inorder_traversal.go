package main

import (
	"fmt"
	"os"
)

type P94TreeNode struct {
	Val   int
	Left  *P94TreeNode
	Right *P94TreeNode
}

func inorderTraversal(root *P94TreeNode) []int {
	// Solution using recursive:
	// nodeValues := []int{}
	// return btreeInorderTraverse(root, nodeValues)

	// Solution using iterative:
	return btreeInordertraverseIterative(root)
}

func btreeInorderTraverseRecursive(root *P94TreeNode, nodeValues []int) []int {
	if root == nil {
		return nodeValues
	}

	nodeValues = btreeInorderTraverseRecursive(root.Left, nodeValues)
	nodeValues = append(nodeValues, root.Val)
	nodeValues = btreeInorderTraverseRecursive(root.Right, nodeValues)

	return nodeValues
}

func btreeInordertraverseIterative(root *P94TreeNode) []int {
	result := []int{}
	stack := []*P94TreeNode{}
	current := root

	for current != nil || len(stack) > 0 {
		// push all nodes in the left into stack start from current
		for current != nil {
			stack = append(stack, current)
			current = current.Left
		}

		// pop the top node from the stack
		stackLen := len(stack)
		current = stack[stackLen-1]
		stack = stack[:stackLen-1]

		// append current value
		result = append(result, current.Val)

		// go to right node
		current = current.Right
	}

	return result
}

func RunTestBtreeInorderTraversal() {
	testCases := map[string]struct {
		root   *P94TreeNode
		expect []int
	}{
		"case-1": {
			root:   btreeCase1(),
			expect: []int{1, 3, 2},
		},
		"case-2": {
			root:   btreeCase2(),
			expect: []int{4, 2, 6, 5, 7, 1, 3, 9, 8},
		},
	}

	for name, testCase := range testCases {
		fmt.Printf("RUN %s\n", name)

		result := inorderTraversal(testCase.root)
		if !EqualSlices(result, testCase.expect) {
			fmt.Printf("=== FAILED: expect = %v - got = %v\n", testCase.expect, result)
			os.Exit(1)
		}
		fmt.Printf("=== PASSED\n")

	}
	fmt.Printf("\n✅ All tests passed!\n")
}

func btreeCase1() *P94TreeNode {
	root := &P94TreeNode{Val: 1}
	root.Right = &P94TreeNode{Val: 2}
	root.Right.Left = &P94TreeNode{Val: 3}
	return root
}

func btreeCase2() *P94TreeNode {
	root := &P94TreeNode{Val: 1}
	root.Left = &P94TreeNode{Val: 2}
	root.Left.Left = &P94TreeNode{Val: 4}
	root.Left.Right = &P94TreeNode{Val: 5}
	root.Left.Right.Left = &P94TreeNode{Val: 6}
	root.Left.Right.Right = &P94TreeNode{Val: 7}
	root.Right = &P94TreeNode{Val: 3}
	root.Right.Right = &P94TreeNode{Val: 8}
	root.Right.Right.Left = &P94TreeNode{Val: 9}
	return root
}
