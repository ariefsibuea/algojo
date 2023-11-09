package hackerrank_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	hr "algojo.ariefsibuea.dev/hackerrank"
)

func Test_DiagonalDifference(t *testing.T) {
	testCases := []struct {
		input  [][]int32
		output int32
	}{
		{
			input: [][]int32{
				{1, 2, 3},
				{4, 5, 6},
				{9, 8, 9},
			},
			output: 2,
		},
		{
			input: [][]int32{
				{11, 2, 4},
				{4, 5, 6},
				{10, 8, -12},
			},
			output: 15,
		},
	}

	for _, testCase := range testCases {
		out := hr.DiagonalDifference(testCase.input)
		require.Equal(t, testCase.output, out)
	}
}
