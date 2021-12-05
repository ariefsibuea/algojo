package plusone_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	lib "github.com/ariefsibuea/dsa/leetcode/66-plus_one"
)

func Test_PlusOne(t *testing.T) {
	testcases := []struct {
		digits []int
		output []int
	}{
		{
			digits: []int{1, 2, 3},
			output: []int{1, 2, 4},
		},
		{
			digits: []int{4, 3, 2, 1},
			output: []int{4, 3, 2, 2},
		},
		{
			digits: []int{0},
			output: []int{1},
		},
	}

	for i, testcase := range testcases {
		t.Run(fmt.Sprintf("case-%d", i+1), func(t *testing.T) {
			out := lib.PlusOne(testcase.digits)
			require.Equal(t, testcase.output, out)
		})
	}
}
