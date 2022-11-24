package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"algojo.ariefsibuea.dev/leetcode"
)

func Test_Sqrt(t *testing.T) {
	testcases := []struct {
		input  int
		output int
	}{
		{
			input:  4,
			output: 2,
		},
		{
			input:  8,
			output: 2,
		},
	}

	for i, tc := range testcases {
		testName := fmt.Sprintf("case-%d", i+1)
		t.Run(testName, func(t *testing.T) {
			out := leetcode.MySqrt(tc.input)
			require.Equal(t, tc.output, out)
		})
	}
}
