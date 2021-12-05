package reverseinteger_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	lib "github.com/ariefsibuea/dsa/leetcode/7-reverse_integer"
)

func Test_ReverseInteger(t *testing.T) {
	testcases := []struct {
		x      int
		result int
	}{
		{
			x:      123,
			result: 321,
		},
		{
			x:      -123,
			result: -321,
		},
		{
			x:      120,
			result: 21,
		},
		{
			x:      0,
			result: 0,
		},
		{
			x:      1534236469,
			result: 0,
		},
	}

	for i, testcase := range testcases {
		t.Run(fmt.Sprintf("case-%d", i+1), func(t *testing.T) {
			res := lib.Reverse(testcase.x)
			require.Equal(t, testcase.result, res)
		})
	}
}
