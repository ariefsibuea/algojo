package removenthnodefromendoflist

/**
 * Problem source: https://leetcode.com/problems/remove-nth-node-from-end-of-list/
 * Agenda:
 * 		Study Plan: Algorithm I
 *		Day V Two Pointers
 *		Level: Medium
 * Solution source:
**/

type ListNode struct {
	Val  int
	Next *ListNode
}

// RemoveNthFromEnd implements two pointers technique to solve remove nth
// node from the end of the linked list.
func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	slow, fast := head, head

	i := 0
	for i = 0; i < n && fast != nil; i++ {
		fast = fast.Next
	}
	if i < n-1 {
		return head
	}
	if fast == nil {
		return head.Next
	}

	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}

	slow.Next = slow.Next.Next
	return head
}
