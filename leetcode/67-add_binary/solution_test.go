package addbinary_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	lib "github.com/ariefsibuea/cobra/leetcode/67-add_binary"
)

func Test_AddBinary(t *testing.T) {
	testcases := []struct {
		a      string
		b      string
		output string
	}{
		{
			a:      "11",
			b:      "1",
			output: "100",
		},
		{
			a:      "1010",
			b:      "1011",
			output: "10101",
		},
	}

	for i, testcase := range testcases {
		t.Run(fmt.Sprintf("case-%d", i+1), func(t *testing.T) {
			out := lib.AddBinary(testcase.a, testcase.b)
			require.Equal(t, testcase.output, out)
		})
	}
}
