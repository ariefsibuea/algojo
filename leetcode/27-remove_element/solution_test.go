package removeelement_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	lib "algojo.ariefsibuea.dev/leetcode/27-remove_element"
)

func Test_RemoveElemetn(t *testing.T) {
	testcases := []struct {
		nums   []int
		val    int
		k      int
		output []int
	}{
		{
			nums:   []int{3, 2, 2, 3},
			val:    3,
			k:      2,
			output: []int{2, 2},
		},
		{
			nums:   []int{0, 1, 2, 2, 3, 0, 4, 2},
			val:    2,
			k:      5,
			output: []int{0, 1, 3, 0, 4},
		},
	}

	for i, testcase := range testcases {
		t.Run(fmt.Sprintf("case-%d", i+1), func(t *testing.T) {
			res := lib.RemoveElement(testcase.nums, testcase.val)
			require.Equal(t, testcase.k, res)
			for i := 0; i < res; i++ {
				require.Equal(t, testcase.output[i], testcase.nums[i])
			}
		})
	}
}
