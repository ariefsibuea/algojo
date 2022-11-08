package insertionsortlist_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	lib "algojo.ariefsibuea.dev/leetcode/147-insertion_sort_list"
)

func Test_InsertionSortList(t *testing.T) {
	testcases := []struct {
		input  *lib.ListNode
		output []int
	}{
		{
			input:  mockLinkedListI(),
			output: []int{1, 2, 3, 4},
		},
		{
			input:  mockLinkedListII(),
			output: []int{-1, 0, 3, 4, 5},
		},
	}

	for i, testcase := range testcases {
		t.Run(fmt.Sprintf("case-%d", i+1), func(t *testing.T) {
			res := lib.InsertionSortList(testcase.input)
			for _, n := range testcase.output {
				require.Equal(t, n, res.Val)
				res = res.Next
			}
		})
	}
}

func mockLinkedListI() *lib.ListNode {
	// 4,2,1,3
	root := &lib.ListNode{Val: 4}
	root.Next = &lib.ListNode{Val: 2}
	root.Next.Next = &lib.ListNode{Val: 1}
	root.Next.Next.Next = &lib.ListNode{Val: 3}
	return root
}

func mockLinkedListII() *lib.ListNode {
	// -1,5,3,4,0
	root := &lib.ListNode{Val: -1}
	root.Next = &lib.ListNode{Val: 5}
	root.Next.Next = &lib.ListNode{Val: 3}
	root.Next.Next.Next = &lib.ListNode{Val: 4}
	root.Next.Next.Next.Next = &lib.ListNode{Val: 0}
	return root
}
