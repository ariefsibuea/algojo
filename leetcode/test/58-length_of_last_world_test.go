package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"algojo.ariefsibuea.dev/leetcode"
)

func Test_LengthOfLastWord(t *testing.T) {
	testcases := []struct {
		input  string
		output int
	}{
		{
			input:  "Hello World",
			output: 5,
		},
		{
			input:  "   fly me   to   the moon  ",
			output: 4,
		},
		{
			input:  "luffy is still joyboy",
			output: 6,
		},
	}

	for i, tc := range testcases {
		testName := fmt.Sprintf("case-%d", i+1)
		t.Run(testName, func(t *testing.T) {
			out := leetcode.LengthOfLastWord(tc.input)
			require.Equal(t, tc.output, out)
		})
	}
}
