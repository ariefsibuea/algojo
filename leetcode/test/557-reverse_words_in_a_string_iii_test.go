package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"algojo.ariefsibuea.dev/leetcode"
)

func Test_ReverseWords(t *testing.T) {
	testcases := []struct {
		input  string
		output string
	}{
		{
			input:  "Let's take LeetCode contest",
			output: "s'teL ekat edoCteeL tsetnoc",
		},
		{
			input:  "God Ding",
			output: "doG gniD",
		},
	}

	for i, tc := range testcases {
		testName := fmt.Sprintf("case-%d", i+1)
		t.Run(testName, func(t *testing.T) {
			require.Equal(t, tc.output, leetcode.ReverseWords(tc.input))
		})
	}
}
