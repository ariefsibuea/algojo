package main

import (
	"fmt"
)

func PostOrder(root *BinaryTreeNode) {
	if root != nil {
		PostOrder(root.left)
		PostOrder(root.right)
		fmt.Printf("%d ", root.data)
	}
}

func RunPostOrder() {
	root := &BinaryTreeNode{data: 1}
	// left
	root.left = &BinaryTreeNode{data: 2}
	root.left.left = &BinaryTreeNode{data: 4}
	root.left.right = &BinaryTreeNode{data: 5}
	// right2
	root.right = &BinaryTreeNode{data: 3}
	root.right.left = &BinaryTreeNode{data: 6}
	root.right.right = &BinaryTreeNode{data: 7}

	PostOrder(root)
	fmt.Println()
}

func PostOrderNonRecursive(root *BinaryTreeNode) {
	stack := []*BinaryTreeNode{}
	stackLen := 0

	var previous *BinaryTreeNode

	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.left
		}

		stackLen = len(stack)
		for root == nil && stackLen > 0 {
			// top of stack
			root = stack[stackLen-1]

			if root.right == nil || root.right == previous {
				fmt.Printf("%d ", root.data)
				// pop stack
				stack = stack[:stackLen-1]
				stackLen = len(stack)

				previous = root
				root = nil
			} else {
				root = root.right
			}
		}
	}
}

func RunPostOrderNonRecursive() {
	root := &BinaryTreeNode{data: 1}
	// left
	root.left = &BinaryTreeNode{data: 2}
	root.left.left = &BinaryTreeNode{data: 4}
	root.left.right = &BinaryTreeNode{data: 5}
	// right2
	root.right = &BinaryTreeNode{data: 3}
	root.right.left = &BinaryTreeNode{data: 6}
	root.right.right = &BinaryTreeNode{data: 7}

	PostOrderNonRecursive(root)
	fmt.Println()
}
