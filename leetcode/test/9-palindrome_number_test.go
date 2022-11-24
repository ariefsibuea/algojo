package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"algojo.ariefsibuea.dev/leetcode"
)

func Test_PalindromeNumber(t *testing.T) {
	testcases := []struct {
		input  int
		output bool
	}{
		{
			input:  121,
			output: true,
		},
		{
			input:  -121,
			output: false,
		},
		{
			input:  10,
			output: false,
		},
		{
			input:  -101,
			output: false,
		},
	}

	for i, tc := range testcases {
		testName := fmt.Sprintf("case-%d", i+1)
		t.Run(testName, func(t *testing.T) {
			out := leetcode.IsPalindrome(tc.input)
			require.Equal(t, tc.output, out)
		})
	}
}
