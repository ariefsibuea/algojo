package movezeroes_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	lib "algojo.ariefsibuea.dev/leetcode/283-move_zeroes"
)

func Test_MoveZeroes(t *testing.T) {
	testcases := []struct {
		got  []int
		want []int
	}{
		{
			got:  []int{0, 1, 0, 3, 12},
			want: []int{1, 3, 12, 0, 0},
		},
		{
			got:  []int{0},
			want: []int{0},
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprintf("case-%d", i+1), func(t *testing.T) {
			lib.MoveZeroes(tc.got)
			require.Equal(t, tc.want, tc.got)
		})
	}
}
