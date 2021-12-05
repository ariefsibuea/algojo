package removeduplicatefromsortedarray_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	lib "github.com/ariefsibuea/dsa/leetcode/26-remove_duplicate_from_sorted_array"
)

func Test_RemoveDuplicateFromSortedArray(t *testing.T) {
	testcases := []struct {
		nums   []int
		k      int
		output []int
	}{
		{
			nums:   []int{1, 1, 2},
			k:      2,
			output: []int{1, 2},
		},
		{
			nums:   []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			k:      5,
			output: []int{0, 1, 2, 3, 4},
		},
	}

	for i, testcase := range testcases {
		t.Run(fmt.Sprintf("case-%d", i+1), func(t *testing.T) {
			res := lib.RemoveDuplicates(testcase.nums)
			require.Equal(t, testcase.k, res)
			for i := 0; i < res; i++ {
				require.Equal(t, testcase.output[i], testcase.nums[i])
			}
		})
	}
}
