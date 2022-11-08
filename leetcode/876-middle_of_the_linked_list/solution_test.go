package middleofthelinkedlist_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	lib "algojo.ariefsibuea.dev/leetcode/876-middle_of_the_linked_list"
)

func Test_MiddleNode(t *testing.T) {
	testcases := []struct {
		got  *lib.ListNode
		want []int
	}{
		{
			got:  mockListNodeI(),
			want: []int{3, 4, 5},
		},
		{
			got:  mockListNodeII(),
			want: []int{4, 5, 6},
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprintf("case-%d", i+1), func(t *testing.T) {
			mid := lib.MiddleNode(tc.got)
			exp := make([]int, 0)
			for mid != nil {
				exp = append(exp, mid.Val)
				mid = mid.Next
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
	head.Next.Next.Next.Next = &lib.ListNode{Val: 5}

	return head
}

func mockListNodeII() *lib.ListNode {
	head := &lib.ListNode{Val: 1}
	head.Next = &lib.ListNode{Val: 2}
	head.Next.Next = &lib.ListNode{Val: 3}
	head.Next.Next.Next = &lib.ListNode{Val: 4}
	head.Next.Next.Next.Next = &lib.ListNode{Val: 5}
	head.Next.Next.Next.Next.Next = &lib.ListNode{Val: 6}

	return head
}
