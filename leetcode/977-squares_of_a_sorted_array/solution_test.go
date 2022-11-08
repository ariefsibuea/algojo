package squaresofasortedarray_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	lib "algojo.ariefsibuea.dev/leetcode/977-squares_of_a_sorted_array"
)

func Test_SortedSquares(t *testing.T) {
	testcases := []struct {
		got  []int
		want []int
	}{
		{
			got:  []int{-4, -1, 0, 3, 10},
			want: []int{0, 1, 9, 16, 100},
		},
		{
			got:  []int{-7, -3, 2, 3, 11},
			want: []int{4, 9, 9, 49, 121},
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprintf("case-%d", i+1), func(t *testing.T) {
			require.Equal(t, tc.want, lib.SortedSquares(tc.got))
		})
	}
}
