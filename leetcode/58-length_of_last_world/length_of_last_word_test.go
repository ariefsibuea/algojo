package lengthoflastworld_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	lolw "github.com/ariefsibuea/dsa/leetcode/58-length_of_last_world"
)

func Test_LengthOfLastWord(t *testing.T) {
	testcases := []struct {
		s      string
		output int
	}{
		{
			s:      "Hello World",
			output: 5,
		},
		{
			s:      "   fly me   to   the moon  ",
			output: 4,
		},
		{
			s:      "luffy is still joyboy",
			output: 6,
		},
	}

	for i, testcase := range testcases {
		t.Run(fmt.Sprintf("case-%d", i+1), func(t *testing.T) {
			out := lolw.LengthOfLastWord(testcase.s)
			require.Equal(t, testcase.output, out)
		})
	}
}
