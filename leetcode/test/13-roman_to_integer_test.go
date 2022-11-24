package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"algojo.ariefsibuea.dev/leetcode"
)

func Test_RomanToInteger(t *testing.T) {
	testcases := []struct {
		input  string
		output int
	}{
		{
			input:  "III",
			output: 3,
		},
		{
			input:  "LVIII",
			output: 58,
		},
		{
			input:  "IV",
			output: 4,
		},
		{
			input:  "IX",
			output: 9,
		},
		{
			input:  "MCMXCIV",
			output: 1994,
		},
	}

	for i, tc := range testcases {
		testName := fmt.Sprintf("case-%d", i+1)
		t.Run(testName, func(t *testing.T) {
			out := leetcode.RomanToInt(tc.input)
			require.Equal(t, tc.output, out)
		})
	}
}
