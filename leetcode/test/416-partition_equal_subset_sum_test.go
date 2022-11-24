package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"algojo.ariefsibuea.dev/leetcode"
)

func Test_CanPartition(t *testing.T) {
	testcases := []struct {
		input  []int
		output bool
	}{
		{
			input:  []int{1, 5, 11, 5},
			output: true,
		},
		{
			input:  []int{1, 2, 3, 5},
			output: false,
		},
		{
			input:  []int{2, 2, 1, 1},
			output: true,
		},
	}

	for i, tc := range testcases {
		testName := fmt.Sprintf("case-%d", i+1)
		t.Run(testName, func(t *testing.T) {
			require.Equal(t, tc.output, leetcode.CanPartition(tc.input))
		})
	}
}
