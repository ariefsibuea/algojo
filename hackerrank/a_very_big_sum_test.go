package hackerrank_test

import (
	"testing"

	hr "algojo.ariefsibuea.dev/hackerrank"
	"github.com/stretchr/testify/require"
)

func Test_AVeryBigSum(t *testing.T) {
	testCases := []struct {
		input  []int64
		output int64
	}{
		{
			input:  []int64{1000000001, 1000000002, 1000000003, 1000000004, 1000000005},
			output: 5000000015,
		},
	}

	for _, testCase := range testCases {
		out := hr.AVeryBigSum(testCase.input)
		require.Equal(t, testCase.output, out)
	}
}
