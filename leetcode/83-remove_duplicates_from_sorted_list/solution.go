package removeduplicatesfromsortedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func DeleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	currentNode := head
	next := currentNode.Next

	for next != nil {
		if currentNode.Val == next.Val {
			next = next.Next
			currentNode.Next = next
			continue
		}

		temp := next
		next = next.Next
		currentNode = temp
	}

	return head
}
