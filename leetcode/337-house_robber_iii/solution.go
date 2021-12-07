package houserobberiii

// Problem source: https://leetcode.com/problems/house-robber-iii/

type TotalValue struct {
	WithRoot    int
	WithoutRoot int
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func MaxNodeVal(totalValue TotalValue) int {
	if totalValue.WithoutRoot > totalValue.WithRoot {
		return totalValue.WithoutRoot
	}
	return totalValue.WithRoot
}

func ExploreNode(node *TreeNode) TotalValue {
	if node == nil {
		return TotalValue{}
	}

	leftNodeVal := ExploreNode(node.Left)
	rightNodeVal := ExploreNode(node.Right)

	totalValue := TotalValue{
		WithRoot:    node.Val + leftNodeVal.WithoutRoot + rightNodeVal.WithoutRoot,
		WithoutRoot: MaxNodeVal(leftNodeVal) + MaxNodeVal(rightNodeVal),
	}

	return totalValue
}

func Rob(root *TreeNode) int {
	if root == nil {
		return 0
	}

	totalValue := ExploreNode(root)
	if totalValue.WithRoot > totalValue.WithoutRoot {
		return totalValue.WithRoot
	}
	return totalValue.WithoutRoot
}
