package removenthnodefromendoflist_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	lib "algojo.ariefsibuea.dev/leetcode/19-remove_nth_node_from_end_of_list"
)

func Test_RemoveNthFromEnd(t *testing.T) {
	type input struct {
		head *lib.ListNode
		n    int
	}

	testcases := []struct {
		got  input
		want []int
	}{
		{
			got:  input{mockLinkedListI(), 2},
			want: []int{1, 2, 3, 5},
		},
		{
			got:  input{mockLinkedListII(), 1},
			want: []int{},
		},
		{
			got:  input{mockLinkedListIII(), 1},
			want: []int{1},
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprintf("case-%d", i+1), func(t *testing.T) {
			head := lib.RemoveNthFromEnd(tc.got.head, tc.got.n)
			exp := make([]int, 0)
			for head != nil {
				exp = append(exp, head.Val)
				head = head.Next
			}
			require.Equal(t, tc.want, exp)
		})
	}
}

func mockLinkedListI() *lib.ListNode {
	head := &lib.ListNode{Val: 1}
	head.Next = &lib.ListNode{Val: 2}
	head.Next.Next = &lib.ListNode{Val: 3}
	head.Next.Next.Next = &lib.ListNode{Val: 4}
	head.Next.Next.Next.Next = &lib.ListNode{Val: 5}
	return head
}

func mockLinkedListII() *lib.ListNode {
	head := &lib.ListNode{Val: 1}
	return head
}

func mockLinkedListIII() *lib.ListNode {
	head := &lib.ListNode{Val: 1}
	head.Next = &lib.ListNode{Val: 2}
	return head
}
