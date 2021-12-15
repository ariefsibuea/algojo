package rangesumofbst

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func traverseTree(node *TreeNode, low, hight int) int {
	if node == nil {
		return 0
	}

	leftSum, rightSum := 0, 0
	if node.Val >= low {
		leftSum = traverseTree(node.Left, low, hight)
	}
	if node.Val <= hight {
		rightSum = traverseTree(node.Right, low, hight)
	}

	if node.Val >= low && node.Val <= hight {
		return node.Val + leftSum + rightSum
	}
	return leftSum + rightSum
}

func RangeSumBST(root *TreeNode, low int, high int) int {
	return traverseTree(root, low, high)
}
