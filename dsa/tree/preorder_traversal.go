package main

import "fmt"

func PreOrder(root *BinaryTreeNode) {
	if root != nil {
		fmt.Printf("%d ", root.data)
		PreOrder(root.left)
		PreOrder(root.right)
	}
}

func RunPreOrder() {
	root := &BinaryTreeNode{data: 1}
	// left
	root.left = &BinaryTreeNode{data: 2}
	root.left.left = &BinaryTreeNode{data: 4}
	root.left.right = &BinaryTreeNode{data: 5}
	// right
	root.right = &BinaryTreeNode{data: 3}
	root.right.left = &BinaryTreeNode{data: 6}
	root.right.right = &BinaryTreeNode{data: 7}

	PreOrder(root)
	fmt.Println()
}

func PreOrderNonRecursive(root *BinaryTreeNode) {
	stack := []*BinaryTreeNode{}
	stackLen := 0

	for root != nil || len(stack) > 0 {
		for root != nil {
			// process current node
			fmt.Printf("%d ", root.data)
			// add root to stack
			stack = append(stack, root)
			// get the left node
			root = root.left
		}

		stackLen = len(stack)
		if stackLen == 0 {
			break
		}

		// pop stack
		root = stack[stackLen-1]
		stack = stack[:stackLen-1]

		// get the right node
		root = root.right
	}
}

func RunPreOrderNonRecursive() {
	root := &BinaryTreeNode{data: 1}
	// left
	root.left = &BinaryTreeNode{data: 2}
	root.left.left = &BinaryTreeNode{data: 4}
	root.left.right = &BinaryTreeNode{data: 5}
	// right
	root.right = &BinaryTreeNode{data: 3}
	root.right.left = &BinaryTreeNode{data: 6}
	root.right.right = &BinaryTreeNode{data: 7}

	PreOrderNonRecursive(root)
	fmt.Println()
}
