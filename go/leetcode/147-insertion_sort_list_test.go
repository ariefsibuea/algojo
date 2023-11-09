package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"algojo.ariefsibuea.dev/leetcode"
)

func Test_InsertionSortList(t *testing.T) {
	testcases := []struct {
		input  *leetcode.ListNode
		output []int
	}{
		{
			input:  mockLinkedListInsertionSortListI(),
			output: []int{1, 2, 3, 4},
		},
		{
			input:  mockLinkedListInsertionSortListII(),
			output: []int{-1, 0, 3, 4, 5},
		},
	}

	soln := leetcode.Solution{}

	for i, tc := range testcases {
		testName := fmt.Sprintf("case-%d", i+1)
		t.Run(testName, func(t *testing.T) {
			res := soln.InsertionSortList(tc.input)
			for _, n := range tc.output {
				require.Equal(t, n, res.Val)
				res = res.Next
			}
		})
	}
}

func mockLinkedListInsertionSortListI() *leetcode.ListNode {
	// 4,2,1,3
	root := &leetcode.ListNode{Val: 4}
	root.Next = &leetcode.ListNode{Val: 2}
	root.Next.Next = &leetcode.ListNode{Val: 1}
	root.Next.Next.Next = &leetcode.ListNode{Val: 3}
	return root
}

func mockLinkedListInsertionSortListII() *leetcode.ListNode {
	// -1,5,3,4,0
	root := &leetcode.ListNode{Val: -1}
	root.Next = &leetcode.ListNode{Val: 5}
	root.Next.Next = &leetcode.ListNode{Val: 3}
	root.Next.Next.Next = &leetcode.ListNode{Val: 4}
	root.Next.Next.Next.Next = &leetcode.ListNode{Val: 0}
	return root
}
