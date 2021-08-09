package sqrt_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	sqrt "github.com/ariefsibuea/dsa/leetcode/69-sqrt"
)

func Test_Sqrt(t *testing.T) {
	testcases := []struct {
		x      int
		output int
	}{
		{
			x:      4,
			output: 2,
		},
		{
			x:      8,
			output: 2,
		},
	}

	for i, testcase := range testcases {
		t.Run(fmt.Sprintf("case-%d", i+1), func(t *testing.T) {
			out := sqrt.MySqrt(testcase.x)
			require.Equal(t, testcase.output, out)
		})
	}
}
