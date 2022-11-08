package hackerrank_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	hr "algojo.ariefsibuea.dev/hackerrank"
)

func Test_SolveMeFirst(t *testing.T) {
	testCases := []struct {
		input struct {
			a uint32
			b uint32
		}
		output uint32
	}{
		{
			input: struct {
				a uint32
				b uint32
			}{a: 2, b: 3},
			output: 5,
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("case-%d", i+1), func(t *testing.T) {
			out := hr.SolveMeFirst(testCase.input.a, testCase.input.b)
			require.Equal(t, testCase.output, out)
		})
	}
}
