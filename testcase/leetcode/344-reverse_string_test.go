package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"algojo.ariefsibuea.dev/leetcode"
)

func Test_ReverseString(t *testing.T) {
	testcases := []struct {
		input  []byte
		output []byte
	}{
		{
			input:  []byte{'h', 'e', 'l', 'l', 'o'},
			output: []byte{'o', 'l', 'l', 'e', 'h'},
		},
		{
			input:  []byte{'H', 'a', 'n', 'n', 'a', 'h'},
			output: []byte{'h', 'a', 'n', 'n', 'a', 'H'},
		},
	}

	soln := leetcode.Solution{}

	for i, tc := range testcases {
		testName := fmt.Sprintf("case-%d", i+1)
		t.Run(testName, func(t *testing.T) {
			soln.ReverseString(tc.input)
			require.Equal(t, tc.output, tc.input)
		})
	}
}
