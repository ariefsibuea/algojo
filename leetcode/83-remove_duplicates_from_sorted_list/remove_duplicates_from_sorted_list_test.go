package removeduplicatesfromsortedlist_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	rdfsl "github.com/ariefsibuea/dsa/leetcode/83-remove_duplicates_from_sorted_list"
)

func Test_RemoveDuplicatesFromSortedList(t *testing.T) {
	a3 := &rdfsl.ListNode{2, nil}
	a2 := &rdfsl.ListNode{1, a3}
	a1 := &rdfsl.ListNode{1, a2}

	testcases := []struct {
		head   *rdfsl.ListNode
		output []int
	}{
		{
			head:   a1,
			output: []int{1, 2},
		},
	}

	for i, testcase := range testcases {
		t.Run(fmt.Sprintf("case-%d", i+1), func(t *testing.T) {
			out := rdfsl.DeleteDuplicates(testcase.head)

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
