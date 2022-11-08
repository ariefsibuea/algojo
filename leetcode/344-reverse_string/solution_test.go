package reversestring_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	lib "algojo.ariefsibuea.dev/leetcode/344-reverse_string"
)

func Test_ReverseString(t *testing.T) {
	testcases := []struct {
		got  []byte
		want []byte
	}{
		{
			got:  []byte{'h', 'e', 'l', 'l', 'o'},
			want: []byte{'o', 'l', 'l', 'e', 'h'},
		},
		{
			got:  []byte{'H', 'a', 'n', 'n', 'a', 'h'},
			want: []byte{'h', 'a', 'n', 'n', 'a', 'H'},
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprintf("case-%d", i+1), func(t *testing.T) {
			lib.ReverseString(tc.got)
			require.Equal(t, tc.want, tc.got)
		})
	}
}
