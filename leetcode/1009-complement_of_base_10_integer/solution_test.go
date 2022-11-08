package complementofbase10integer_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	lib "algojo.ariefsibuea.dev/leetcode/1009-complement_of_base_10_integer"
)

func Test_BitwiseComplement(t *testing.T) {
	testcases := []struct {
		got  int
		want int
	}{
		{
			got:  5,
			want: 2,
		},
		{
			got:  7,
			want: 0,
		},
		{
			got:  10,
			want: 5,
		},
		{
			got:  1,
			want: 0,
		},
		{
			got:  0,
			want: 1,
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprintf("case-%d", i+1), func(t *testing.T) {
			require.Equal(t, tc.want, lib.BitwiseComplement(tc.got))
		})
	}
}
