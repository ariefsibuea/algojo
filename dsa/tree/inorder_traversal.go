package main

import "fmt"

func InOrder(root *BinaryTreeNode) {
	if root != nil {
		InOrder(root.left)
		fmt.Printf("%d ", root.data)
		InOrder(root.right)
	}
}

func RunInOrder() {
	root := &BinaryTreeNode{data: 1}
	// left
	root.left = &BinaryTreeNode{data: 2}
	root.left.left = &BinaryTreeNode{data: 4}
	root.left.right = &BinaryTreeNode{data: 5}
	// right
	root.right = &BinaryTreeNode{data: 3}
	root.right.left = &BinaryTreeNode{data: 6}
	root.right.right = &BinaryTreeNode{data: 7}

	InOrder(root)
	fmt.Println()
}

func InOrderNonRecursive(root *BinaryTreeNode) {
	stack := []*BinaryTreeNode{}
	stackLen := 0

	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.left
		}

		stackLen = len(stack)
		if stackLen == 0 {
			break
		}

		// pop stack
		root = stack[stackLen-1]
		stack = stack[:stackLen-1]

		fmt.Printf("%d ", root.data)

		root = root.right
	}
}

func RunInOrderNonRecursive() {
	root := &BinaryTreeNode{data: 1}
	// left
	root.left = &BinaryTreeNode{data: 2}
	root.left.left = &BinaryTreeNode{data: 4}
	root.left.right = &BinaryTreeNode{data: 5}
	// right
	root.right = &BinaryTreeNode{data: 3}
	root.right.left = &BinaryTreeNode{data: 6}
	root.right.right = &BinaryTreeNode{data: 7}

	InOrderNonRecursive(root)
	fmt.Println()
}
