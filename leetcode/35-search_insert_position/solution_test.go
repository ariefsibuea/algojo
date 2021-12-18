package searchinsertposition_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	lib "github.com/ariefsibuea/cobra/leetcode/35-search_insert_position"
)

func Test_SearchInsert(t *testing.T) {
	type input struct {
		nums   []int
		target int
	}
	testcases := []struct {
		got  input
		want int
	}{
		{
			got:  input{[]int{1, 3, 5, 6}, 5},
			want: 2,
		},
		{
			got:  input{[]int{1, 3, 5, 6}, 2},
			want: 1,
		},
		{
			got:  input{[]int{1, 3, 5, 6}, 7},
			want: 4,
		},
		{
			got:  input{[]int{1, 3}, 2},
			want: 1,
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprintf("case-%d", i+1), func(t *testing.T) {
			require.Equal(t, tc.want, lib.SearchInsert(tc.got.nums, tc.got.target))
		})
	}
}
