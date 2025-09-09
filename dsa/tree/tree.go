package main

type BinaryTreeNode struct {
	data  int
	left  *BinaryTreeNode
	right *BinaryTreeNode
}

func main() {
	// RunPreOrder()
	// RunPreOrderNonRecursive()

	// RunInOrder()
	// RunInOrderNonRecursive()

	// RunPostOrder()
	// RunPostOrderNonRecursive()

	RunLevelOrder()
}
