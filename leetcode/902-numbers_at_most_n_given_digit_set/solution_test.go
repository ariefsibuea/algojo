package numbersatmostngivendigitset_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	lib "algojo.ariefsibuea.dev/leetcode/902-numbers_at_most_n_given_digit_set"
)

func Test_AtMostNGivenDigitSet(t *testing.T) {
	testcases := []struct {
		digits []string
		n      int
		output int
	}{
		{
			digits: []string{"1", "3", "5", "7"},
			n:      100,
			output: 20,
		},
		{
			digits: []string{"1", "4", "9"},
			n:      1000000000,
			output: 29523,
		},
		{
			digits: []string{"3", "4", "8"},
			n:      4,
			output: 2,
		},
	}

	for i, testcase := range testcases {
		t.Run(fmt.Sprintf("case-%d", i+1), func(t *testing.T) {
			require.Equal(t, testcase.output, lib.AtMostNGivenDigitSet(testcase.digits, testcase.n))
		})
	}
}
