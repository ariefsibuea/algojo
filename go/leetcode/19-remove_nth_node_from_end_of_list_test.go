package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"algojo.ariefsibuea.dev/leetcode"
)

type inputRemoveNthFromEnd struct {
	head *leetcode.ListNode
	n    int
}

func Test_RemoveNthFromEnd(t *testing.T) {
	soln := leetcode.Solution{}

	testcases := []struct {
		input  inputRemoveNthFromEnd
		output []int
	}{
		{
			input:  inputRemoveNthFromEnd{mockLinkedListI(), 2},
			output: []int{1, 2, 3, 5},
		},
		{
			input:  inputRemoveNthFromEnd{mockLinkedListII(), 1},
			output: []int{},
		},
		{
			input:  inputRemoveNthFromEnd{mockLinkedListIII(), 1},
			output: []int{1},
		},
	}

	for i, tc := range testcases {
		testName := fmt.Sprintf("case-%d", i+1)
		t.Run(testName, func(t *testing.T) {
			head := soln.RemoveNthFromEnd(tc.input.head, tc.input.n)
			exp := make([]int, 0)
			for head != nil {
				exp = append(exp, head.Val)
				head = head.Next
			}
			require.Equal(t, tc.output, exp)
		})
	}
}

func mockLinkedListI() *leetcode.ListNode {
	head := &leetcode.ListNode{Val: 1}
	head.Next = &leetcode.ListNode{Val: 2}
	head.Next.Next = &leetcode.ListNode{Val: 3}
	head.Next.Next.Next = &leetcode.ListNode{Val: 4}
	head.Next.Next.Next.Next = &leetcode.ListNode{Val: 5}
	return head
}

func mockLinkedListII() *leetcode.ListNode {
	head := &leetcode.ListNode{Val: 1}
	return head
}

func mockLinkedListIII() *leetcode.ListNode {
	head := &leetcode.ListNode{Val: 1}
	head.Next = &leetcode.ListNode{Val: 2}
	return head
}
