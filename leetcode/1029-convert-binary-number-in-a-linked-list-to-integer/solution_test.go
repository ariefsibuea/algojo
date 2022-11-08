package convertbinarynumberinalinkedlisttointeger_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	lib "algojo.ariefsibuea.dev/leetcode/1029-convert-binary-number-in-a-linked-list-to-integer"
)

func Test_GetDecimalValue(t *testing.T) {
	testcases := []struct {
		input  *lib.ListNode
		output int
	}{
		{mockLinkedListI(), 5},
		{mockLinkedListII(), 0},
		{mockLinkedListIII(), 1},
		{mockLinkedListIV(), 18880},
		{mockLinkedListV(), 0},
	}

	for i, testcase := range testcases {
		t.Run(fmt.Sprintf("case-%d", i+1), func(t *testing.T) {
			res := lib.GetDecimalValue(testcase.input)
			require.Equal(t, testcase.output, res)
		})
	}
}

func mockLinkedListI() *lib.ListNode {
	elements := []int{1, 0, 1}
	return insertElements(elements)
}

func mockLinkedListII() *lib.ListNode {
	elements := []int{0}
	return insertElements(elements)
}

func mockLinkedListIII() *lib.ListNode {
	elements := []int{1}
	return insertElements(elements)
}

func mockLinkedListIV() *lib.ListNode {
	elements := []int{1, 0, 0, 1, 0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 0}
	return insertElements(elements)
}

func mockLinkedListV() *lib.ListNode {
	elements := []int{0, 0}
	return insertElements(elements)
}

func insertElements(elements []int) *lib.ListNode {
	head := &lib.ListNode{Val: elements[0]}
	currentNode := head
	for _, element := range elements[1:] {
		if currentNode.Next == nil {
			currentNode.Next = &lib.ListNode{}
		}
		currentNode.Next.Val = element
		currentNode = currentNode.Next
	}
	return head
}
