package rotatearray_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	lib "algojo.ariefsibuea.dev/leetcode/189-rotate_array"
)

func Test_Rotate(t *testing.T) {
	type input struct {
		nums []int
		k    int
	}

	testcases := []struct {
		got  input
		want []int
	}{
		{
			got:  input{[]int{1, 2, 3, 4, 5, 6, 7}, 3},
			want: []int{5, 6, 7, 1, 2, 3, 4},
		},
		{
			got:  input{[]int{-1, -100, 3, 99}, 2},
			want: []int{3, 99, -1, -100},
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprintf("case-%d", i+1), func(t *testing.T) {
			lib.Rotate(tc.got.nums, tc.got.k)
			require.Equal(t, tc.want, tc.got.nums)
		})
	}
}
