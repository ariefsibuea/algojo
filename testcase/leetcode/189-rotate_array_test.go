package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"algojo.ariefsibuea.dev/leetcode"
)

type inputRotate struct {
	nums []int
	k    int
}

func Test_Rotate(t *testing.T) {
	testcases := []struct {
		input  inputRotate
		output []int
	}{
		{
			input:  inputRotate{[]int{1, 2, 3, 4, 5, 6, 7}, 3},
			output: []int{5, 6, 7, 1, 2, 3, 4},
		},
		{
			input:  inputRotate{[]int{-1, -100, 3, 99}, 2},
			output: []int{3, 99, -1, -100},
		},
	}

	soln := leetcode.Solution{}

	for i, tc := range testcases {
		testName := fmt.Sprintf("case-%d", i+1)
		t.Run(testName, func(t *testing.T) {
			soln.Rotate(tc.input.nums, tc.input.k)
			require.Equal(t, tc.output, tc.input.nums)
		})
	}
}
