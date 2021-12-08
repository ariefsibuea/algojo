package convertbinarynumberinalinkedlisttointeger

// Problem source: https://leetcode.com/problems/convert-binary-number-in-a-linked-list-to-integer/

type ListNode struct {
	Val  int
	Next *ListNode
}

func GetDecimalValue(head *ListNode) int {
	if head == nil {
		return 0
	}

	binaryNums := make([]int, 0)
	currentNode := head
	for {
		binaryNums = append([]int{currentNode.Val}, binaryNums...)
		if currentNode.Next == nil {
			break
		}
		currentNode = currentNode.Next
	}

	decimalNum := 0
	for i, n := range binaryNums {
		num := n << i
		decimalNum += num
	}
	return decimalNum
}
