package climbstairs_test

import (
	"fmt"
	"testing"

	cs "github.com/ariefsibuea/dsa/leetcode/70-climb_stairs"
	"github.com/stretchr/testify/require"
)

func Test_ClimbStairs(t *testing.T) {
	testcases := []struct {
		n      int
		output int
	}{
		{
			n:      2,
			output: 2,
		},
		{
			n:      3,
			output: 3,
		},
	}

	for i, testcase := range testcases {
		t.Run(fmt.Sprintf("case-%d", i+1), func(t *testing.T) {
			out := cs.ClimbStairs(testcase.n)
			require.Equal(t, testcase.output, out)
		})
	}
}
