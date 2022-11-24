package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"algojo.ariefsibuea.dev/leetcode"
)

type inputAddBinary struct {
	a string
	b string
}

func Test_AddBinary(t *testing.T) {
	testcases := []struct {
		input  inputAddBinary
		output string
	}{
		{
			input: inputAddBinary{
				a: "11",
				b: "1",
			},
			output: "100",
		},
		{
			input: inputAddBinary{
				a: "1010",
				b: "1011",
			},
			output: "10101",
		},
	}

	for i, tc := range testcases {
		testName := fmt.Sprintf("case-%d", i+1)
		t.Run(testName, func(t *testing.T) {
			out := leetcode.AddBinary(tc.input.a, tc.input.b)
			require.Equal(t, tc.output, out)
		})
	}
}
