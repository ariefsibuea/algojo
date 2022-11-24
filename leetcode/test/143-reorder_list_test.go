package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"algojo.ariefsibuea.dev/leetcode"
)

func Test_ReorderList(t *testing.T) {
	testcases := []struct {
		input  *leetcode.ListNode
		output []int
	}{
		{
			input:  mockListNodeI(),
			output: []int{1, 4, 2, 3},
		},
		{
			input:  mockListNodeII(),
			output: []int{1, 5, 2, 4, 3},
		},
		{
			input:  mockListNodeIII(),
			output: []int{1, 2},
		},
		{
			input:  mockListNodeIV(),
			output: []int{1},
		},
		{
			input:  mockListNodeV(),
			output: []int{1, 3, 2},
		},
	}

	for i, tc := range testcases {
		testName := fmt.Sprintf("case-%d", i+1)
		t.Run(testName, func(t *testing.T) {
			leetcode.ReorderList(tc.input)

			exp := make([]int, 0)
			head := tc.input
			for head != nil {
				exp = append(exp, head.Val)
				head = head.Next
			}

			require.Equal(t, tc.output, exp)
		})
	}
}

func mockListNodeI() *leetcode.ListNode {
	head := &leetcode.ListNode{Val: 1}
	head.Next = &leetcode.ListNode{Val: 2}
	head.Next.Next = &leetcode.ListNode{Val: 3}
	head.Next.Next.Next = &leetcode.ListNode{Val: 4}

	return head
}

func mockListNodeII() *leetcode.ListNode {
	head := &leetcode.ListNode{Val: 1}
	head.Next = &leetcode.ListNode{Val: 2}
	head.Next.Next = &leetcode.ListNode{Val: 3}
	head.Next.Next.Next = &leetcode.ListNode{Val: 4}
	head.Next.Next.Next.Next = &leetcode.ListNode{Val: 5}

	return head
}

func mockListNodeIII() *leetcode.ListNode {
	head := &leetcode.ListNode{Val: 1}
	head.Next = &leetcode.ListNode{Val: 2}

	return head
}

func mockListNodeIV() *leetcode.ListNode {
	head := &leetcode.ListNode{Val: 1}

	return head
}

func mockListNodeV() *leetcode.ListNode {
	head := &leetcode.ListNode{Val: 1}
	head.Next = &leetcode.ListNode{Val: 2}
	head.Next.Next = &leetcode.ListNode{Val: 3}

	return head
}
