package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"algojo.ariefsibuea.dev/leetcode"
)

func Test_MiddleNode(t *testing.T) {
	testcases := []struct {
		input  *leetcode.ListNode
		output []int
	}{
		{
			input:  mockListNodeMiddleNodeI(),
			output: []int{3, 4, 5},
		},
		{
			input:  mockListNodeMiddleNodeII(),
			output: []int{4, 5, 6},
		},
	}

	soln := leetcode.Solution{}

	for i, tc := range testcases {
		testName := fmt.Sprintf("case-%d", i+1)
		t.Run(testName, func(t *testing.T) {
			mid := soln.MiddleNode(tc.input)
			exp := make([]int, 0)
			for mid != nil {
				exp = append(exp, mid.Val)
				mid = mid.Next
			}

			require.Equal(t, tc.output, exp)
		})
	}
}

func mockListNodeMiddleNodeI() *leetcode.ListNode {
	head := &leetcode.ListNode{Val: 1}
	head.Next = &leetcode.ListNode{Val: 2}
	head.Next.Next = &leetcode.ListNode{Val: 3}
	head.Next.Next.Next = &leetcode.ListNode{Val: 4}
	head.Next.Next.Next.Next = &leetcode.ListNode{Val: 5}

	return head
}

func mockListNodeMiddleNodeII() *leetcode.ListNode {
	head := &leetcode.ListNode{Val: 1}
	head.Next = &leetcode.ListNode{Val: 2}
	head.Next.Next = &leetcode.ListNode{Val: 3}
	head.Next.Next.Next = &leetcode.ListNode{Val: 4}
	head.Next.Next.Next.Next = &leetcode.ListNode{Val: 5}
	head.Next.Next.Next.Next.Next = &leetcode.ListNode{Val: 6}

	return head
}
