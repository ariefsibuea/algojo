package twosumiiinputarrayissorted_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	lib "github.com/ariefsibuea/cobra/leetcode/167-two_sum_ii_input_array_is_sorted"
)

func Test_TwoSum(t *testing.T) {
	type input struct {
		numbers []int
		target  int
	}

	testcases := []struct {
		got  input
		want []int
	}{
		{
			got:  input{[]int{2, 7, 11, 15}, 9},
			want: []int{1, 2},
		},
		{
			got:  input{[]int{2, 3, 4}, 6},
			want: []int{1, 3},
		},
		{
			got:  input{[]int{-1, 0}, -1},
			want: []int{1, 2},
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprintf("case-%d", i+1), func(t *testing.T) {
			require.Equal(t, tc.want, lib.TwoSum(tc.got.numbers, tc.got.target))
		})
	}
}
