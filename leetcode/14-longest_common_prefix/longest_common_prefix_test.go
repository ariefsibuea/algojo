package longestcommonprefix_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	lcp "github.com/ariefsibuea/dsa/leetcode/14-longest_common_prefix"
)

func Test_LongetCommonPrefix(t *testing.T) {
	testcases := []struct {
		input  []string
		result string
	}{
		{
			input:  []string{"flower", "flow", "flight"},
			result: "fl",
		},
		{
			input:  []string{"dog", "racecar", "car"},
			result: "",
		},
		{
			input:  []string{"c", "acc", "ccc"},
			result: "",
		},
	}

	for i, testcase := range testcases {
		t.Run(fmt.Sprintf("case-%d", i+1), func(t *testing.T) {
			res := lcp.LongestCommonPrefix(testcase.input)
			require.Equal(t, testcase.result, res)
		})
	}
}
