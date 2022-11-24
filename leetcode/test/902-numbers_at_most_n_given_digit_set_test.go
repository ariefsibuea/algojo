package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"algojo.ariefsibuea.dev/leetcode"
)

type inputAtMostNGivenDigitSet struct {
	digits []string
	n      int
}

func Test_AtMostNGivenDigitSet(t *testing.T) {
	testcases := []struct {
		input  inputAtMostNGivenDigitSet
		output int
	}{
		{
			input: inputAtMostNGivenDigitSet{
				digits: []string{"1", "3", "5", "7"},
				n:      100,
			},
			output: 20,
		},
		{
			input: inputAtMostNGivenDigitSet{
				digits: []string{"1", "4", "9"},
				n:      1000000000,
			},
			output: 29523,
		},
		{
			input: inputAtMostNGivenDigitSet{
				digits: []string{"3", "4", "8"},
				n:      4,
			},
			output: 2,
		},
	}

	for i, tc := range testcases {
		testName := fmt.Sprintf("case-%d", i+1)
		t.Run(testName, func(t *testing.T) {
			require.Equal(t, tc.output, leetcode.AtMostNGivenDigitSet(tc.input.digits, tc.input.n))
		})
	}
}
