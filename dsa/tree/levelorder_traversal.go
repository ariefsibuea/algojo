package main

import "fmt"

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

func RunLevelOrder() {
	root := &BinaryTreeNode{data: 1}
	// left
	root.left = &BinaryTreeNode{data: 2}
	root.left.left = &BinaryTreeNode{data: 4}
	root.left.right = &BinaryTreeNode{data: 5}
	// right2
	root.right = &BinaryTreeNode{data: 3}
	root.right.left = &BinaryTreeNode{data: 6}
	root.right.right = &BinaryTreeNode{data: 7}

	LevelOrder(root)
	fmt.Println()
}
