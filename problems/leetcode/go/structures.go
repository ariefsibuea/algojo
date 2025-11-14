package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func getListNodeValue(node *ListNode) int {
	if node == nil {
		return 0
	}
	return node.Val
}

func nextNode(node *ListNode) *ListNode {
	if node == nil {
		return nil
	}
	return node.Next
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Interval struct {
	Start int
	End   int
}
