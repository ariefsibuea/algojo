package partitionequalsubsetsum_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	lib "github.com/ariefsibuea/cobra/leetcode/416-partition_equal_subset_sum"
)

func Test_CanPartition(t *testing.T) {
	testcases := []struct {
		nums   []int
		output bool
	}{
		{
			nums:   []int{1, 5, 11, 5},
			output: true,
		},
		{
			nums:   []int{1, 2, 3, 5},
			output: false,
		},
		{
			nums:   []int{2, 2, 1, 1},
			output: true,
		},
	}

	for i, testcase := range testcases {
		t.Run(fmt.Sprintf("case-%d", i+1), func(t *testing.T) {
			require.Equal(t, testcase.output, lib.CanPartition(testcase.nums))
		})
	}
}
