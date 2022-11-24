package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"algojo.ariefsibuea.dev/leetcode"
)

type inputCheckInclusion struct {
	s1 string
	s2 string
}

func Test_CheckInclusion(t *testing.T) {

	testcases := []struct {
		input  inputCheckInclusion
		output bool
	}{
		{
			input:  inputCheckInclusion{"ab", "eidbaooo"},
			output: true,
		},
		{
			input:  inputCheckInclusion{"adc", "dcda"},
			output: true,
		},
	}

	for i, tc := range testcases {
		testName := fmt.Sprintf("case-%d", i+1)
		t.Run(testName, func(t *testing.T) {
			require.Equal(t, tc.output, leetcode.CheckInclusion(tc.input.s1, tc.input.s2))
		})
	}
}
