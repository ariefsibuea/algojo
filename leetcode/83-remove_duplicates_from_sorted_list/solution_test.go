package removeduplicatesfromsortedlist_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	lib "algojo.ariefsibuea.dev/leetcode/83-remove_duplicates_from_sorted_list"
)

func Test_RemoveDuplicatesFromSortedList(t *testing.T) {
	a3 := &lib.ListNode{2, nil}
	a2 := &lib.ListNode{1, a3}
	a1 := &lib.ListNode{1, a2}

	testcases := []struct {
		head   *lib.ListNode
		output []int
	}{
		{
			head:   a1,
			output: []int{1, 2},
		},
	}

	for i, testcase := range testcases {
		t.Run(fmt.Sprintf("case-%d", i+1), func(t *testing.T) {
			out := lib.DeleteDuplicates(testcase.head)

			length := 0
			for out != nil {
				require.Equal(t, testcase.output[length], out.Val)

				length++
				out = out.Next
			}
			require.Len(t, testcase.output, length)
		})
	}
}
