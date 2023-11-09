package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"algojo.ariefsibuea.dev/leetcode"
)

func Test_MinimumAbsDifference(t *testing.T) {
	testcases := []struct {
		input  []int
		output [][]int
	}{
		{
			input:  []int{4, 2, 1, 3},
			output: [][]int{{1, 2}, {2, 3}, {3, 4}},
		},
		{
			input:  []int{1, 3, 6, 10, 15},
			output: [][]int{{1, 3}},
		},
		{
			input:  []int{3, 8, -10, 23, 19, -4, -14, 27},
			output: [][]int{{-14, -10}, {19, 23}, {23, 27}},
		},
	}

	soln := leetcode.Solution{}

	for i, tc := range testcases {
		testName := fmt.Sprintf("case-%d", i+1)
		t.Run(testName, func(t *testing.T) {
			require.Equal(t, tc.output, soln.MinimumAbsDifference(tc.input))
		})
	}
}
