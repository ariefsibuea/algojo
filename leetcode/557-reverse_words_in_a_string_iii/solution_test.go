package reversewordsinastringiii_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	lib "github.com/ariefsibuea/cobra/leetcode/557-reverse_words_in_a_string_iii"
)

func Test_ReverseWords(t *testing.T) {
	testcases := []struct {
		got  string
		want string
	}{
		{
			got:  "Let's take LeetCode contest",
			want: "s'teL ekat edoCteeL tsetnoc",
		},
		{
			got:  "God Ding",
			want: "doG gniD",
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprintf("case-%d", i+1), func(t *testing.T) {
			require.Equal(t, tc.want, lib.ReverseWords(tc.got))
		})
	}
}
