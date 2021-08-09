package romantointeger_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	rti "github.com/ariefsibuea/dsa/leetcode/13-roman_to_integer"
)

func Test_RomanToInteger(t *testing.T) {
	testcases := []struct {
		input  string
		result int
	}{
		{
			input:  "III",
			result: 3,
		},
		{
			input:  "LVIII",
			result: 58,
		},
		{
			input:  "IV",
			result: 4,
		},
		{
			input:  "IX",
			result: 9,
		},
		{
			input:  "MCMXCIV",
			result: 1994,
		},
	}

	for i, testcase := range testcases {
		t.Run(fmt.Sprintf("case-%d", i+1), func(t *testing.T) {
			res := rti.RomanToInt(testcase.input)
			require.Equal(t, testcase.result, res)
		})
	}
}
