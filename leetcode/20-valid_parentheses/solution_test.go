package validparentheses_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	lib "github.com/ariefsibuea/dsa/leetcode/20-valid_parentheses"
)

func Test_ValidParentheses(t *testing.T) {
	testcases := []struct {
		input  string
		result bool
	}{
		{
			input:  "()",
			result: true,
		},
		{
			input:  "()[]{}",
			result: true,
		},
		{
			input:  "(]",
			result: false,
		},
		{
			input:  "([)]",
			result: false,
		},
		{
			input:  "{[]}",
			result: true,
		},
	}

	for i, testcase := range testcases {
		t.Run(fmt.Sprintf("case-%d", i+1), func(t *testing.T) {
			res := lib.IsValid(testcase.input)
			require.Equal(t, testcase.result, res)
		})
	}
}
