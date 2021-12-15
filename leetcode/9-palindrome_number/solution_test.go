package palindromenumber_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	lib "github.com/ariefsibuea/cobra/leetcode/9-palindrome_number"
)

func Test_PalindromeNumber(t *testing.T) {
	testcases := []struct {
		input  int
		result bool
	}{
		{
			input:  121,
			result: true,
		},
		{
			input:  -121,
			result: false,
		},
		{
			input:  10,
			result: false,
		},
		{
			input:  -101,
			result: false,
		},
	}

	for i, testcase := range testcases {
		t.Run(fmt.Sprintf("case-%d", i+1), func(t *testing.T) {
			res := lib.IsPalindrome(testcase.input)
			require.Equal(t, testcase.result, res)
		})
	}
}