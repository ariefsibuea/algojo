package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"algojo.ariefsibuea.dev/leetcode"
)

func Test_RemoveDuplicatesFromSortedList(t *testing.T) {
	a3 := &leetcode.ListNode{Val: 2, Next: nil}
	a2 := &leetcode.ListNode{Val: 1, Next: a3}
	a1 := &leetcode.ListNode{Val: 1, Next: a2}

	testcases := []struct {
		input  *leetcode.ListNode
		output []int
	}{
		{
			input:  a1,
			output: []int{1, 2},
		},
	}

	soln := leetcode.Solution{}

	for i, testcase := range testcases {
		testName := fmt.Sprintf("case-%d", i+1)
		t.Run(testName, func(t *testing.T) {
			out := soln.DeleteDuplicates(testcase.input)

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
