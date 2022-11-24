package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"algojo.ariefsibuea.dev/leetcode"
)

type inputStrStr struct {
	haystack string
	needle   string
}

func Test_StrStr(t *testing.T) {
	testcases := []struct {
		input  inputStrStr
		output int
	}{
		{
			input: inputStrStr{
				haystack: "hello",
				needle:   "ll",
			},
			output: 2,
		},
		{
			input: inputStrStr{
				haystack: "aaaaa",
				needle:   "bba",
			},
			output: -1,
		},
		{
			input: inputStrStr{
				haystack: "",
				needle:   "",
			},
			output: 0,
		},
	}

	for i, tc := range testcases {
		testName := fmt.Sprintf("case-%d", i+1)
		t.Run(testName, func(t *testing.T) {
			out := leetcode.StrStr(tc.input.haystack, tc.input.needle)
			require.Equal(t, tc.output, out)
		})
	}
}
