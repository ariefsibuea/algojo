package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"algojo.ariefsibuea.dev/leetcode"
)

type inputRemoveElement struct {
	nums []int
	val  int
}

type outputRemoveElement struct {
	nums []int
	k    int
}

func Test_RemoveElemetn(t *testing.T) {
	testcases := []struct {
		input  inputRemoveElement
		output outputRemoveElement
	}{
		{
			input: inputRemoveElement{
				nums: []int{3, 2, 2, 3},
				val:  3,
			},
			output: outputRemoveElement{
				k:    2,
				nums: []int{2, 2},
			},
		},
		{
			input: inputRemoveElement{
				nums: []int{0, 1, 2, 2, 3, 0, 4, 2},
				val:  2,
			},
			output: outputRemoveElement{
				k:    5,
				nums: []int{0, 1, 3, 0, 4},
			},
		},
	}

	for i, tc := range testcases {
		testName := fmt.Sprintf("case-%d", i+1)
		t.Run(testName, func(t *testing.T) {
			out := leetcode.RemoveElement(tc.input.nums, tc.input.val)
			require.Equal(t, tc.output.k, out)
			for i := 0; i < out; i++ {
				require.Equal(t, tc.output.nums[i], tc.input.nums[i])
			}
		})
	}
}
