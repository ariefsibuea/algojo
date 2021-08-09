package implementstrstr_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	is "github.com/ariefsibuea/dsa/leetcode/28-implement_strstr"
)

func Test_StrStr(t *testing.T) {
	testcases := []struct {
		haystack string
		needle   string
		output   int
	}{
		{
			haystack: "hello",
			needle:   "ll",
			output:   2,
		},
		{
			haystack: "aaaaa",
			needle:   "bba",
			output:   -1,
		},
		{
			haystack: "",
			needle:   "",
			output:   0,
		},
	}

	for i, testcase := range testcases {
		t.Run(fmt.Sprintf("case-%d", i+1), func(t *testing.T) {
			res := is.StrStr(testcase.haystack, testcase.needle)
			require.Equal(t, testcase.output, res)
		})
	}
}
