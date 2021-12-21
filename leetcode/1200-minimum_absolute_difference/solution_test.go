package minimumabsolutedifference_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	lib "github.com/ariefsibuea/cobra/leetcode/1200-minimum_absolute_difference"
)

func Test_MinimumAbsDifference(t *testing.T) {
	testcases := []struct {
		got  []int
		want [][]int
	}{
		{
			got:  []int{4, 2, 1, 3},
			want: [][]int{{1, 2}, {2, 3}, {3, 4}},
		},
		{
			got:  []int{1, 3, 6, 10, 15},
			want: [][]int{{1, 3}},
		},
		{
			got:  []int{3, 8, -10, 23, 19, -4, -14, 27},
			want: [][]int{{-14, -10}, {19, 23}, {23, 27}},
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprintf("case-%d", i+1), func(t *testing.T) {
			require.Equal(t, tc.want, lib.MinimumAbsDifference(tc.got))
		})
	}
}
