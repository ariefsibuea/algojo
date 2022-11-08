package maximalsquare_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	lib "algojo.ariefsibuea.dev/leetcode/221-maximal-square"
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

	for i, testcase := range testcases {
		t.Run(fmt.Sprintf("case-%d", i+1), func(t *testing.T) {
			require.Equal(t, testcase.output, lib.MaximalSquare(testcase.input))
		})
	}
}
