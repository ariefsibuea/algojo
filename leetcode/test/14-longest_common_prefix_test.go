package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"algojo.ariefsibuea.dev/leetcode"
)

func Test_LongestCommonPrefix(t *testing.T) {
	testcases := []struct {
		input  []string
		output string
	}{
		{
			input:  []string{"flower", "flow", "flight"},
			output: "fl",
		},
		{
			input:  []string{"dog", "racecar", "car"},
			output: "",
		},
		{
			input:  []string{"c", "acc", "ccc"},
			output: "",
		},
	}

	for i, tc := range testcases {
		testName := fmt.Sprintf("case-%d", i+1)
		t.Run(testName, func(t *testing.T) {
			out := leetcode.LongestCommonPrefix(tc.input)
			require.Equal(t, tc.output, out)
		})
	}
}
