package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"algojo.ariefsibuea.dev/leetcode"
)

func Test_MaximalSquare(t *testing.T) {
	testcases := []struct {
		input  [][]byte
		output int
	}{
		{
			input: [][]byte{
				{'1', '0', '1', '0', '0'},
				{'1', '0', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '0', '0', '1', '0'},
			},
			output: 4,
		},
		{
			input: [][]byte{
				{'0', '1'},
				{'1', '0'},
			},
			output: 1,
		},
	}

	soln := leetcode.Solution{}

	for i, tc := range testcases {
		testName := fmt.Sprintf("case-%d", i+1)
		t.Run(testName, func(t *testing.T) {
			require.Equal(t, tc.output, soln.MaximalSquare(tc.input))
		})
	}
}
