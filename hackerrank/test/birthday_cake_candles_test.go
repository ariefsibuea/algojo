package hackerrank_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	hr "algojo.ariefsibuea.dev/hackerrank"
)

func Test_BirthdayCakeCandles(t *testing.T) {
	testCases := []struct {
		input  []int32
		output int32
	}{
		{
			input:  []int32{3, 2, 1, 3},
			output: 2,
		},
	}

	for _, testCase := range testCases {
		out := hr.BirthdayCakeCandles(testCase.input)
		require.Equal(t, testCase.output, out)
	}
}
