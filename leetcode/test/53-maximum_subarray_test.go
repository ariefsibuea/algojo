package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"algojo.ariefsibuea.dev/leetcode"
)

func Test_MaximumSubarray(t *testing.T) {
	testcases := []struct {
		input  []int
		output int
	}{
		{
			input:  []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			output: 6,
		},
		{
			input:  []int{1},
			output: 1,
		},
		{
			input:  []int{5, 4, -1, 7, 8},
			output: 23,
		},
	}

	for i, tc := range testcases {
		testName := fmt.Sprintf("case-%d", i+1)
		t.Run(testName, func(t *testing.T) {
			res := leetcode.MaxSubArray(tc.input)
			require.Equal(t, tc.output, res)
		})
	}
}
