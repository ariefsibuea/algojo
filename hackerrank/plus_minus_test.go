package hackerrank_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	hr "algojo.ariefsibuea.dev/hackerrank"
)

type outPlusMinus struct {
	ratioPositive float64
	ratioNegative float64
	ratioZero     float64
}

func Test_PlusMinus(t *testing.T) {
	testCases := []struct {
		input  []int32
		output outPlusMinus
	}{
		{
			input: []int32{-4, 3, -9, 0, 4, 1},
			output: outPlusMinus{
				ratioPositive: 0.500000,
				ratioNegative: 0.333333,
				ratioZero:     0.166667,
			},
		},
	}

	for _, testCase := range testCases {
		ratioPositive, ratioNegative, ratioZero := hr.PlusMinus(testCase.input)
		require.Equal(t, testCase.output.ratioPositive, ratioPositive)
		require.Equal(t, testCase.output.ratioNegative, ratioNegative)
		require.Equal(t, testCase.output.ratioZero, ratioZero)
	}
}
