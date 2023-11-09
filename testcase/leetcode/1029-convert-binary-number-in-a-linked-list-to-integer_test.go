package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"algojo.ariefsibuea.dev/leetcode"
)

func Test_GetDecimalValue(t *testing.T) {
	testcases := []struct {
		input  *leetcode.ListNode
		output int
	}{
		{mockInputGetDecimalValueI(), 5},
		{mockInputGetDecimalValueII(), 0},
		{mockInputGetDecimalValueIII(), 1},
		{mockInputGetDecimalValueIV(), 18880},
		{mockInputGetDecimalValueV(), 0},
	}

	soln := leetcode.Solution{}

	for i, tc := range testcases {
		testName := fmt.Sprintf("case-%d", i+1)
		t.Run(testName, func(t *testing.T) {
			out := soln.GetDecimalValue(tc.input)
			require.Equal(t, tc.output, out)
		})
	}
}

func mockInputGetDecimalValueI() *leetcode.ListNode {
	elements := []int{1, 0, 1}
	return insertElements(elements)
}

func mockInputGetDecimalValueII() *leetcode.ListNode {
	elements := []int{0}
	return insertElements(elements)
}

func mockInputGetDecimalValueIII() *leetcode.ListNode {
	elements := []int{1}
	return insertElements(elements)
}

func mockInputGetDecimalValueIV() *leetcode.ListNode {
	elements := []int{1, 0, 0, 1, 0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 0}
	return insertElements(elements)
}

func mockInputGetDecimalValueV() *leetcode.ListNode {
	elements := []int{0, 0}
	return insertElements(elements)
}

func insertElements(elements []int) *leetcode.ListNode {
	head := &leetcode.ListNode{Val: elements[0]}
	currentNode := head
	for _, element := range elements[1:] {
		if currentNode.Next == nil {
			currentNode.Next = &leetcode.ListNode{}
		}
		currentNode.Next.Val = element
		currentNode = currentNode.Next
	}
	return head
}
