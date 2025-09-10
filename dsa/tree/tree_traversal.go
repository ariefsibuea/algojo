package main

import "fmt"

type BinaryTreeNode struct {
	data  int
	left  *BinaryTreeNode
	right *BinaryTreeNode
}

func PreOrder(root *BinaryTreeNode) {
	if root != nil {
		fmt.Printf("%d ", root.data)
		PreOrder(root.left)
		PreOrder(root.right)
	}
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

func InOrder(root *BinaryTreeNode) {
	if root != nil {
		InOrder(root.left)
		fmt.Printf("%d ", root.data)
		InOrder(root.right)
	}
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

func PostOrder(root *BinaryTreeNode) {
	if root != nil {
		PostOrder(root.left)
		PostOrder(root.right)
		fmt.Printf("%d ", root.data)
	}
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

func LevelOrder(root *BinaryTreeNode) {
	var temp *BinaryTreeNode
	var queue = make([]*BinaryTreeNode, 0)

	if root == nil {
		return
	}

	queue = append(queue, root)
	for len(queue) > 0 {
		temp = queue[0]
		queue = queue[1:]

		fmt.Printf("%d ", temp.data)

		if temp.left != nil {
			queue = append(queue, temp.left)
		}
		if temp.right != nil {
			queue = append(queue, temp.right)
		}
	}
}

func main() {
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
	PreOrderNonRecursive(root)
	fmt.Println()

	InOrder(root)
	fmt.Println()
	InOrderNonRecursive(root)
	fmt.Println()

	PostOrder(root)
	fmt.Println()
	PostOrderNonRecursive(root)
	fmt.Println()
}
