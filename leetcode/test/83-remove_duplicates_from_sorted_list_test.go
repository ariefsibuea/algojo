package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"algojo.ariefsibuea.dev/leetcode"
)

func Test_RemoveDuplicatesFromSortedList(t *testing.T) {
	a3 := &leetcode.ListNode{2, nil}
	a2 := &leetcode.ListNode{1, a3}
	a1 := &leetcode.ListNode{1, a2}

	testcases := []struct {
		input  *leetcode.ListNode
		output []int
	}{
		{
			input:  a1,
			output: []int{1, 2},
		},
	}

	for i, testcase := range testcases {
		testName := fmt.Sprintf("case-%d", i+1)
		t.Run(testName, func(t *testing.T) {
			out := leetcode.DeleteDuplicates(testcase.input)

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
