package searchindexposition_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	sip "github.com/ariefsibuea/dsa/leetcode/35-search_index_position"
)

func Test_SearchInsert(t *testing.T) {
	testcases := []struct {
		nums   []int
		target int
		output int
	}{
		{
			nums:   []int{1, 3, 5, 6},
			target: 5,
			output: 2,
		},
		{
			nums:   []int{1, 3, 5, 6},
			target: 2,
			output: 1,
		},
		{
			nums:   []int{1, 3, 5, 6},
			target: 7,
			output: 4,
		},
		{
			nums:   []int{1, 3, 5, 6},
			target: 0,
			output: 0,
		},
		{
			nums:   []int{1},
			target: 0,
			output: 0,
		},
	}

	for i, testcase := range testcases {
		t.Run(fmt.Sprintf("case-%d", i+1), func(t *testing.T) {
			res := sip.SearchInsert(testcase.nums, testcase.target)
			require.Equal(t, testcase.output, res)
		})
	}
}
