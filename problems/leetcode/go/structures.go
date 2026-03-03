package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListFromSlice(vals []int) *ListNode {
	if len(vals) == 0 {
		return nil
	}

	head := &ListNode{Val: vals[0]}
	curr := head
	for i := 1; i < len(vals); i++ {
		curr.Next = &ListNode{Val: vals[i]}
		curr = curr.Next
	}

	return head
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

func listNodeToSlice(head *ListNode) []int {
	if head == nil {
		return nil
	}

	result := make([]int, 0)
	current := head
	for current != nil {
		result = append(result, current.Val)
		current = current.Next
	}
	return result
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

type DLLNode struct {
	Key  int
	Val  int
	Prev *DLLNode
	Next *DLLNode
}

type GraphNode struct {
	Val       int
	Neighbors []*GraphNode
}
