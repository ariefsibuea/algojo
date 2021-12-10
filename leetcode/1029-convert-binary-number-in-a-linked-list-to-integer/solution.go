package convertbinarynumberinalinkedlisttointeger

// Problem source: https://leetcode.com/problems/convert-binary-number-in-a-linked-list-to-integer/

type ListNode struct {
	Val  int
	Next *ListNode
}

func GetDecimalValue(head *ListNode) int {
	decimalNum := head.Val
	for head.Next != nil {
		decimalNum = (decimalNum * 2) + head.Next.Val
		head = head.Next
	}
	return decimalNum
}
