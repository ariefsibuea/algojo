package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"algojo.ariefsibuea.dev/leetcode"
)

type inputTwoSum struct {
	nums   []int
	target int
}

func Test_TwoSum(t *testing.T) {
	testCases := []struct {
		input  inputTwoSum
		output []int
	}{
		{
			input: inputTwoSum{
				nums:   []int{2, 7, 11, 15},
				target: 9,
			},
			output: []int{0, 1},
		},
	}

	for i, tc := range testCases {
		testName := fmt.Sprintf("case-%d", i+1)
		t.Run(testName, func(t *testing.T) {
			out := leetcode.TwoSum(tc.input.nums, tc.input.target)
			require.Equal(t, tc.output, out)
		})
	}
}
