package mergetwosortedlists_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	lib "algojo.ariefsibuea.dev/leetcode/21-merge_two_sorted_lists"
)

func Test_MergeTwoSortedLists(t *testing.T) {
	l1c := lib.ListNode{4, nil}
	l1b := lib.ListNode{2, &l1c}
	l1a := lib.ListNode{1, &l1b}

	l2c := lib.ListNode{4, nil}
	l2b := lib.ListNode{3, &l2c}
	l2a := lib.ListNode{1, &l2b}

	testcases := []struct {
		l1     *lib.ListNode
		l2     *lib.ListNode
		result []int
	}{
		{
			l1:     &l1a,
			l2:     &l2a,
			result: []int{1, 1, 2, 3, 4, 4},
		},
	}

	for i, testcase := range testcases {
		t.Run(fmt.Sprintf("case-%d", i+1), func(t *testing.T) {
			res := lib.MergeTwoLists(testcase.l1, testcase.l2)

			sorted := make([]int, 0)
			for res != nil {
				sorted = append(sorted, res.Val)
				res = res.Next
			}

			require.Equal(t, testcase.result, sorted)
		})
	}
}
