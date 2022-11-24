package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"algojo.ariefsibuea.dev/leetcode"
)

type inputCanReach struct {
	arr   []int
	start int
}

func Test_CanReach(t *testing.T) {
	testcases := []struct {
		input  inputCanReach
		output bool
	}{
		{
			input: inputCanReach{
				arr:   []int{4, 2, 3, 0, 3, 1, 2},
				start: 5,
			},
			output: true,
		},
		{
			input: inputCanReach{
				arr:   []int{4, 2, 3, 0, 3, 1, 2},
				start: 0,
			},
			output: true,
		},
		{
			input: inputCanReach{
				arr:   []int{3, 0, 2, 1, 2},
				start: 2,
			},
			output: false,
		},
	}

	for i, tc := range testcases {
		testName := fmt.Sprintf("case-%d", i+1)
		t.Run(testName, func(t *testing.T) {
			require.Equal(t, tc.output, leetcode.CanReach(tc.input.arr, tc.input.start))
		})
	}
}
