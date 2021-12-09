package binarytreetilt

// Problem source: https://leetcode.com/problems/binary-tree-tilt/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type NodeValue struct {
	SumTilt int
	SumVal  int
}

func absolute(n int) int {
	if n > -1 {
		return n
	}
	return n * -1
}

func ExploreNode(node *TreeNode) NodeValue {
	if node == nil {
		return NodeValue{}
	}

	leftSubtree := ExploreNode(node.Left)
	rightSubtree := ExploreNode(node.Right)
	tilt := absolute(leftSubtree.SumVal - rightSubtree.SumVal)

	nodeVal := NodeValue{
		SumTilt: tilt + leftSubtree.SumTilt + rightSubtree.SumTilt,
		SumVal:  node.Val + leftSubtree.SumVal + rightSubtree.SumVal,
	}

	return nodeVal
}

func FindTilt(root *TreeNode) int {
	if root == nil {
		return 0
	}

	nodeVal := ExploreNode(root)
	return nodeVal.SumTilt
}
