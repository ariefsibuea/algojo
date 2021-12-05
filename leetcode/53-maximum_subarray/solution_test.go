package maximumsubarray_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	lib "github.com/ariefsibuea/dsa/leetcode/53-maximum_subarray"
)

func Test_MaximumSubarray(t *testing.T) {
	testcases := []struct {
		nums   []int
		output int
	}{
		{
			nums:   []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			output: 6,
		},
		{
			nums:   []int{1},
			output: 1,
		},
		{
			nums:   []int{5, 4, -1, 7, 8},
			output: 23,
		},
	}

	for i, testcase := range testcases {
		t.Run(fmt.Sprintf("case-%d", i+1), func(t *testing.T) {
			res := lib.MaxSubArray(testcase.nums)
			require.Equal(t, testcase.output, res)
		})
	}
}
