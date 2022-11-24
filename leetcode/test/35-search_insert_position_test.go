package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"algojo.ariefsibuea.dev/leetcode"
)

type inputSearchInsert struct {
	nums   []int
	target int
}

func Test_SearchInsert(t *testing.T) {

	testcases := []struct {
		input  inputSearchInsert
		output int
	}{
		{
			input:  inputSearchInsert{[]int{1, 3, 5, 6}, 5},
			output: 2,
		},
		{
			input:  inputSearchInsert{[]int{1, 3, 5, 6}, 2},
			output: 1,
		},
		{
			input:  inputSearchInsert{[]int{1, 3, 5, 6}, 7},
			output: 4,
		},
		{
			input:  inputSearchInsert{[]int{1, 3}, 2},
			output: 1,
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprintf("case-%d", i+1), func(t *testing.T) {
			require.Equal(t, tc.output, leetcode.SearchInsert(tc.input.nums, tc.input.target))
		})
	}
}
