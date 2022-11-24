package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"algojo.ariefsibuea.dev/leetcode"
)

func Test_ReverseInteger(t *testing.T) {
	testcases := []struct {
		input  int
		output int
	}{
		{
			input:  123,
			output: 321,
		},
		{
			input:  -123,
			output: -321,
		},
		{
			input:  120,
			output: 21,
		},
		{
			input:  0,
			output: 0,
		},
		{
			input:  1534236469,
			output: 0,
		},
	}

	for i, tc := range testcases {
		testName := fmt.Sprintf("case-%d", i+1)
		t.Run(testName, func(t *testing.T) {
			out := leetcode.Reverse(tc.input)
			require.Equal(t, tc.output, out)
		})
	}
}
