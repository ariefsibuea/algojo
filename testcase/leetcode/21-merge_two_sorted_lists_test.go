package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"algojo.ariefsibuea.dev/leetcode"
)

type inputMergeTwoSortedLists struct {
	list1 *leetcode.ListNode
	list2 *leetcode.ListNode
}

func Test_MergeTwoSortedLists(t *testing.T) {
	soln := leetcode.Solution{}

	l1c := leetcode.ListNode{Val: 4, Next: nil}
	l1b := leetcode.ListNode{Val: 2, Next: &l1c}
	l1a := leetcode.ListNode{Val: 1, Next: &l1b}

	l2c := leetcode.ListNode{Val: 4, Next: nil}
	l2b := leetcode.ListNode{Val: 3, Next: &l2c}
	l2a := leetcode.ListNode{Val: 1, Next: &l2b}

	testcases := []struct {
		input  inputMergeTwoSortedLists
		result []int
	}{
		{
			input: inputMergeTwoSortedLists{
				list1: &l1a,
				list2: &l2a,
			},
			result: []int{1, 1, 2, 3, 4, 4},
		},
	}

	for i, tc := range testcases {
		testName := fmt.Sprintf("case-%d", i+1)
		t.Run(testName, func(t *testing.T) {
			out := soln.MergeTwoLists(tc.input.list1, tc.input.list2)

			sorted := make([]int, 0)
			for out != nil {
				sorted = append(sorted, out.Val)
				out = out.Next
			}

			require.Equal(t, tc.result, sorted)
		})
	}
}
