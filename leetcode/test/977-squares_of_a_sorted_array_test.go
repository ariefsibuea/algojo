package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"algojo.ariefsibuea.dev/leetcode"
)

func Test_SortedSquares(t *testing.T) {
	testcases := []struct {
		input  []int
		output []int
	}{
		{
			input:  []int{-4, -1, 0, 3, 10},
			output: []int{0, 1, 9, 16, 100},
		},
		{
			input:  []int{-7, -3, 2, 3, 11},
			output: []int{4, 9, 9, 49, 121},
		},
	}

	for i, tc := range testcases {
		testName := fmt.Sprintf("case-%d", i+1)
		t.Run(testName, func(t *testing.T) {
			require.Equal(t, tc.output, leetcode.SortedSquares(tc.input))
		})
	}
}
