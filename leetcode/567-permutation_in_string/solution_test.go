package permutationinstring_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	lib "github.com/ariefsibuea/cobra/leetcode/567-permutation_in_string"
)

func Test_CheckInclusion(t *testing.T) {
	type input struct {
		s1 string
		s2 string
	}

	testcases := []struct {
		got  input
		want bool
	}{
		{
			got:  input{"ab", "eidbaooo"},
			want: true,
		},
		{
			got:  input{"adc", "dcda"},
			want: true,
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprintf("case-%d", i+1), func(t *testing.T) {
			require.Equal(t, tc.want, lib.CheckInclusion(tc.got.s1, tc.got.s2))
		})
	}
}
