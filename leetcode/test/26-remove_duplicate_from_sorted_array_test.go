package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"algojo.ariefsibuea.dev/leetcode"
)

type outputRemoveDuplicateFromSortedArray struct {
	k    int
	nums []int
}

func Test_RemoveDuplicateFromSortedArray(t *testing.T) {
	testcases := []struct {
		input  []int
		output outputRemoveDuplicateFromSortedArray
	}{
		{
			input: []int{1, 1, 2},
			output: outputRemoveDuplicateFromSortedArray{
				k:    2,
				nums: []int{1, 2},
			},
		},
		{
			input: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			output: outputRemoveDuplicateFromSortedArray{
				k:    5,
				nums: []int{0, 1, 2, 3, 4},
			},
		},
	}

	for i, tc := range testcases {
		testName := fmt.Sprintf("case-%d", i+1)
		t.Run(testName, func(t *testing.T) {
			out := leetcode.RemoveDuplicates(tc.input)
			require.Equal(t, tc.output.k, out)
			for i := 0; i < out; i++ {
				require.Equal(t, tc.output.nums[i], tc.input[i])
			}
		})
	}
}
