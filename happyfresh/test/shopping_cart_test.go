package happyfresh_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	hf "algojo.ariefsibuea.dev/happyfresh"
)

func Test_FindLowestPrice(t *testing.T) {
	type input struct {
		products [][]string
		discouns [][]string
	}
	testcases := []struct {
		got  input
		want int32
	}{
		{
			got: input{
				products: [][]string{
					{"10", "sale", "january-sale"},
					{"200", "sale", "EMPTY"},
				},
				discouns: [][]string{
					{"sale", "0", "10"},
					{"january-sale", "1", "10"},
				},
			},
			want: 19,
		},
		{
			got: input{
				products: [][]string{
					{"10", "d0", "d1"},
					{"15", "EMPTY", "EMPTY"},
					{"20", "d1", "EMPTY"},
				},
				discouns: [][]string{
					{"d0", "1", "27"},
					{"d1", "2", "5"},
				},
			},
			want: 35,
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprintf("case-%d", i+1), func(t *testing.T) {
			require.Equal(t, tc.want, hf.FindLowestPrice(tc.got.products, tc.got.discouns))
		})
	}
}
