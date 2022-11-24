package hackerrank_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	hr "algojo.ariefsibuea.dev/hackerrank"
)

func Test_SimpleArraySum(t *testing.T) {
	testCases := []struct {
		input  []int32
		output int32
	}{
		{
			input:  []int32{1, 2, 3, 4, 10, 11},
			output: 31,
		},
	}

	for _, testCase := range testCases {
		out := hr.SimpleArraySum(testCase.input)
		require.Equal(t, testCase.output, out)
	}
}
