package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"algojo.ariefsibuea.dev/leetcode"
)

type inputTwoSumII struct {
	numbers []int
	target  int
}

func Test_TwoSumII(t *testing.T) {
	testcases := []struct {
		input  inputTwoSumII
		output []int
	}{
		{
			input:  inputTwoSumII{[]int{2, 7, 11, 15}, 9},
			output: []int{1, 2},
		},
		{
			input:  inputTwoSumII{[]int{2, 3, 4}, 6},
			output: []int{1, 3},
		},
		{
			input:  inputTwoSumII{[]int{-1, 0}, -1},
			output: []int{1, 2},
		},
	}

	for i, tc := range testcases {
		testName := fmt.Sprintf("case-%d", i+1)
		t.Run(testName, func(t *testing.T) {
			require.Equal(t, tc.output, leetcode.TwoSumII(tc.input.numbers, tc.input.target))
		})
	}
}
