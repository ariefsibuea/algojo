package insertionsortlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func InsertionSortList(head *ListNode) *ListNode {
	root := head
	head = head.Next
	root.Next = nil

	for head != nil {
		if head.Val < root.Val {
			temp := root
			root = head
			head = head.Next
			root.Next = temp
			continue
		}

		currentNode := root
		for {
			if currentNode.Next == nil || currentNode.Next.Val > head.Val {
				temp := currentNode.Next
				currentNode.Next = head
				head = head.Next
				currentNode.Next.Next = temp
				break
			}
			currentNode = currentNode.Next
		}
	}

	return root
}
