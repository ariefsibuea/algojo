package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"algojo.ariefsibuea.dev/leetcode"
)

func Test_PlusOne(t *testing.T) {
	testcases := []struct {
		input  []int
		output []int
	}{
		{
			input:  []int{1, 2, 3},
			output: []int{1, 2, 4},
		},
		{
			input:  []int{4, 3, 2, 1},
			output: []int{4, 3, 2, 2},
		},
		{
			input:  []int{0},
			output: []int{1},
		},
	}

	for i, tc := range testcases {
		testName := fmt.Sprintf("case-%d", i+1)
		t.Run(testName, func(t *testing.T) {
			out := leetcode.PlusOne(tc.input)
			require.Equal(t, tc.output, out)
		})
	}
}
