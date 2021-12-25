package poweroftwo_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	lib "github.com/ariefsibuea/cobra/leetcode/231-power_of_two"
)

func Test_IsPowerOfTwo(t *testing.T) {
	testcases := []struct {
		got  int
		want bool
	}{
		{1, true},
		{16, true},
		{3, false},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprintf("case-%d", i+1), func(t *testing.T) {
			require.Equal(t, tc.want, lib.IsPowerOfTwo(tc.got))
		})
	}
}
