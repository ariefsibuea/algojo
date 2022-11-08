package reorderlist_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	lib "algojo.ariefsibuea.dev/leetcode/143-reorder_list"
)

func Test_ReorderList(t *testing.T) {
	testcases := []struct {
		got  *lib.ListNode
		want []int
	}{
		{
			got:  mockListNodeI(),
			want: []int{1, 4, 2, 3},
		},
		{
			got:  mockListNodeII(),
			want: []int{1, 5, 2, 4, 3},
		},
		{
			got:  mockListNodeIII(),
			want: []int{1, 2},
		},
		{
			got:  mockListNodeIV(),
			want: []int{1},
		},
		{
			got:  mockListNodeV(),
			want: []int{1, 3, 2},
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprintf("case-%d", i+1), func(t *testing.T) {
			lib.ReorderList(tc.got)

			exp := make([]int, 0)
			head := tc.got
			for head != nil {
				exp = append(exp, head.Val)
				head = head.Next
			}

			require.Equal(t, tc.want, exp)
		})
	}
}

func mockListNodeI() *lib.ListNode {
	head := &lib.ListNode{Val: 1}
	head.Next = &lib.ListNode{Val: 2}
	head.Next.Next = &lib.ListNode{Val: 3}
	head.Next.Next.Next = &lib.ListNode{Val: 4}

	return head
}

func mockListNodeII() *lib.ListNode {
	head := &lib.ListNode{Val: 1}
	head.Next = &lib.ListNode{Val: 2}
	head.Next.Next = &lib.ListNode{Val: 3}
	head.Next.Next.Next = &lib.ListNode{Val: 4}
	head.Next.Next.Next.Next = &lib.ListNode{Val: 5}

	return head
}

func mockListNodeIII() *lib.ListNode {
	head := &lib.ListNode{Val: 1}
	head.Next = &lib.ListNode{Val: 2}

	return head
}

func mockListNodeIV() *lib.ListNode {
	head := &lib.ListNode{Val: 1}

	return head
}

func mockListNodeV() *lib.ListNode {
	head := &lib.ListNode{Val: 1}
	head.Next = &lib.ListNode{Val: 2}
	head.Next.Next = &lib.ListNode{Val: 3}

	return head
}
