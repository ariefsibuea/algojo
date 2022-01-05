package longestsubstringwithoutrepeatingcharacters_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	lib "github.com/ariefsibuea/cobra/leetcode/3-longest_substring_without_repeating_characters"
)

func Test_LengthOfLongestSubstring(t *testing.T) {
	testcases := []struct {
		got  string
		want int
	}{
		{
			got:  "abcabcbb",
			want: 3,
		},
		{
			got:  "bbbbb",
			want: 1,
		},
		{
			got:  "pwwkew",
			want: 3,
		},
		{
			got:  "dvdf",
			want: 3,
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprintf("case-%d", i+1), func(t *testing.T) {
			require.Equal(t, tc.want, lib.LengthOfLongestSubstring(tc.got))
		})
	}
}
