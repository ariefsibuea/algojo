package binarysearch_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	lib "algojo.ariefsibuea.dev/leetcode/704-binary_search"
)

func Test_Search(t *testing.T) {
	type input struct {
		nums   []int
		target int
	}

	testcases := []struct {
		got  input
		want int
	}{
		{
			got:  input{[]int{-1, 0, 3, 5, 9, 12}, 9},
			want: 4,
		},
		{
			got:  input{[]int{-1, 0, 3, 5, 9, 12}, 2},
			want: -1,
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprintf("case-%d", i+1), func(t *testing.T) {
			require.Equal(t, tc.want, lib.Search(tc.got.nums, tc.got.target))
		})
	}
}
