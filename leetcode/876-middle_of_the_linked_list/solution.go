package middleofthelinkedlist

/**
 * Problem source: https://leetcode.com/problems/middle-of-the-linked-list/
 * Agenda:
 * 		Study Plan: Algorithm I
 *		Day V Two Pointers
 *		Level: Easy
 * Solution source:
**/

type ListNode struct {
	Val  int
	Next *ListNode
}

// MiddleNode implements two pointers technique to find middle of linked list
func MiddleNode(head *ListNode) *ListNode {
	mid, cur := head, head
	for cur != nil && cur.Next != nil {
		cur = cur.Next.Next
		mid = mid.Next
	}

	return mid
}
