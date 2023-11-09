/* House Robber III
Source		: https://leetcode.com/problems/house-robber-iii/
Level		: Medium
Description	: The thief has found himself a new place for his thievery again. There is only one entrance to this area,
			called root. Besides the root, each house has one and only one parent house. After a tour, the smart thief
			realized that all houses in this place form a binary tree. It will automatically contact the police if two
			directly-linked houses were broken into on the same night. Given the root of the binary tree, return the
			maximum amount of money the thief can rob without alerting the police.

Example 1:
Input: root = [3,2,3,null,3,null,1]
Output: 7
Explanation: Maximum amount of money the thief can rob = 3 + 3 + 1 = 7.

Example 2:
Input: root = [3,4,5,1,3,null,1]
Output: 9
Explanation: Maximum amount of money the thief can rob = 4 + 5 = 9.
*/

package leetcode

type TotalValue struct {
	WithRoot    int
	WithoutRoot int
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

func (soln Solution) Rob(root *TreeNode) int {
	if root == nil {
		return 0
	}

	totalValue := ExploreNode(root)
	if totalValue.WithRoot > totalValue.WithoutRoot {
		return totalValue.WithRoot
	}
	return totalValue.WithoutRoot
}
