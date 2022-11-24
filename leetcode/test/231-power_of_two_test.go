package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"algojo.ariefsibuea.dev/leetcode"
)

func Test_IsPowerOfTwo(t *testing.T) {
	testcases := []struct {
		input  int
		output bool
	}{
		{1, true},
		{16, true},
		{3, false},
	}

	for i, tc := range testcases {
		testName := fmt.Sprintf("case-%d", i+1)
		t.Run(testName, func(t *testing.T) {
			require.Equal(t, tc.output, leetcode.IsPowerOfTwo(tc.input))
		})
	}
}
