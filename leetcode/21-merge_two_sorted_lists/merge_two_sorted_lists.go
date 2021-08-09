package mergetwosortedlists

type ListNode struct {
	Val  int
	Next *ListNode
}

func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	switch {
	case l1 == nil && l2 == nil:
		return nil
	case l1 == nil:
		return l2
	case l2 == nil:
		return l1
	}

	var head, currentNode *ListNode
	switch {
	case l1.Val < l2.Val:
		head = l1
		currentNode = head
		l1 = currentNode.Next
	default:
		head = l2
		currentNode = head
		l2 = currentNode.Next
	}

	for l1 != nil || l2 != nil {
		if l1 == nil || l2 == nil {
			if l1 != nil {
				currentNode.Next = l1
				currentNode = l1
				l1 = currentNode.Next
				continue
			}

			currentNode.Next = l2
			currentNode = l2
			l2 = currentNode.Next
			continue
		}

		if l1.Val < l2.Val {
			currentNode.Next = l1
			currentNode = l1
			l1 = currentNode.Next
			continue
		}

		currentNode.Next = l2
		currentNode = l2
		l2 = currentNode.Next
	}

	return head
}
