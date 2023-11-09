package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"algojo.ariefsibuea.dev/leetcode"
)

func Test_MoveZeroes(t *testing.T) {
	testcases := []struct {
		input  []int
		output []int
	}{
		{
			input:  []int{0, 1, 0, 3, 12},
			output: []int{1, 3, 12, 0, 0},
		},
		{
			input:  []int{0},
			output: []int{0},
		},
	}

	soln := leetcode.Solution{}

	for i, tc := range testcases {
		testName := fmt.Sprintf("case-%d", i+1)
		t.Run(testName, func(t *testing.T) {
			soln.MoveZeroes(tc.input)
			require.Equal(t, tc.output, tc.input)
		})
	}
}
