package twosum_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	ts "github.com/ariefsibuea/dsa/leetcode/1-two_sum"
)

func Test_TwoSum(t *testing.T) {
	testCases := []struct {
		nums   []int
		target int
		result []int
	}{
		{
			nums:   []int{2, 7, 11, 15},
			target: 9,
			result: []int{0, 1},
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("case-%d", i+1), func(t *testing.T) {
			res := ts.TwoSum(testCase.nums, testCase.target)
			require.Equal(t, testCase.result, res)
		})
	}
}
