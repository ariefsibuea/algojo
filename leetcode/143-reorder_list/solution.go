package reorderlist

/**
 * Problem source: https://leetcode.com/problems/reorder-list/
**/

type ListNode struct {
	Val  int
	Next *ListNode
}

func ReorderList(head *ListNode) {
	if head.Next == nil {
		return
	}

	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	mid := slow

	next, curr := mid.Next, mid.Next.Next
	mid.Next, next.Next = nil, nil
	for curr != nil {
		temp := curr.Next
		curr.Next = next
		next = curr
		curr = temp
	}

	firstHead, secondHead := head, next
	for secondHead != nil {
		temp := firstHead.Next
		firstHead.Next = secondHead
		firstHead = temp

		temp = secondHead.Next
		secondHead.Next = firstHead
		secondHead = temp
	}
}
