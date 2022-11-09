package hackerrank_test

import (
	"testing"

	"algojo.ariefsibuea.dev/hackerrank"
	"github.com/stretchr/testify/require"
)

type outMiniMaxSum struct {
	totalMin int64
	totalMax int64
}

func Test_MiniMaxSum(t *testing.T) {
	testCases := []struct {
		input  []int32
		output outMiniMaxSum
	}{
		{
			input: []int32{1, 3, 5, 7, 9},
			output: outMiniMaxSum{
				totalMin: 16,
				totalMax: 24,
			},
		},
		{
			input: []int32{793810624, 895642170, 685903712, 623789054, 468592370},
			output: outMiniMaxSum{
				totalMin: 2572095760,
				totalMax: 2999145560,
			},
		},
	}

	for _, testCase := range testCases {
		totalMin, totalMax := hackerrank.MiniMaxSum(testCase.input)
		require.Equal(t, testCase.output.totalMin, totalMin)
		require.Equal(t, testCase.output.totalMax, totalMax)
	}
}
