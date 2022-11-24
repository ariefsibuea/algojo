package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"algojo.ariefsibuea.dev/leetcode"
)

func Test_BitwiseComplement(t *testing.T) {
	testcases := []struct {
		input  int
		output int
	}{
		{
			input:  5,
			output: 2,
		},
		{
			input:  7,
			output: 0,
		},
		{
			input:  10,
			output: 5,
		},
		{
			input:  1,
			output: 0,
		},
		{
			input:  0,
			output: 1,
		},
	}

	for i, tc := range testcases {
		testName := fmt.Sprintf("case-%d", i+1)
		t.Run(testName, func(t *testing.T) {
			require.Equal(t, tc.output, leetcode.BitwiseComplement(tc.input))
		})
	}
}
