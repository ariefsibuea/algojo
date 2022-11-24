package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"algojo.ariefsibuea.dev/leetcode"
)

type inputSearch struct {
	nums   []int
	target int
}

func Test_Search(t *testing.T) {

	testcases := []struct {
		input  inputSearch
		output int
	}{
		{
			input:  inputSearch{[]int{-1, 0, 3, 5, 9, 12}, 9},
			output: 4,
		},
		{
			input:  inputSearch{[]int{-1, 0, 3, 5, 9, 12}, 2},
			output: -1,
		},
	}

	for i, tc := range testcases {
		testName := fmt.Sprintf("case-%d", i+1)
		t.Run(testName, func(t *testing.T) {
			require.Equal(t, tc.output, leetcode.Search(tc.input.nums, tc.input.target))
		})
	}
}
