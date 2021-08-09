package mergetwosortedlists_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	mtsl "github.com/ariefsibuea/dsa/leetcode/21-merge_two_sorted_lists"
)

func Test_MergeTwoSortedLists(t *testing.T) {
	l1c := mtsl.ListNode{4, nil}
	l1b := mtsl.ListNode{2, &l1c}
	l1a := mtsl.ListNode{1, &l1b}

	l2c := mtsl.ListNode{4, nil}
	l2b := mtsl.ListNode{3, &l2c}
	l2a := mtsl.ListNode{1, &l2b}

	testcases := []struct {
		l1     *mtsl.ListNode
		l2     *mtsl.ListNode
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
			res := mtsl.MergeTwoLists(testcase.l1, testcase.l2)

			sorted := make([]int, 0)
			for res != nil {
				sorted = append(sorted, res.Val)
				res = res.Next
			}

			require.Equal(t, testcase.result, sorted)
		})
	}
}
