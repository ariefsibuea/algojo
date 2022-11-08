package jumpgameiii_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	lib "algojo.ariefsibuea.dev/leetcode/1306-jump_game_iii"
)

func Test_CanReach(t *testing.T) {
	testcases := []struct {
		arr    []int
		start  int
		output bool
	}{
		{
			arr:    []int{4, 2, 3, 0, 3, 1, 2},
			start:  5,
			output: true,
		},
		{
			arr:    []int{4, 2, 3, 0, 3, 1, 2},
			start:  0,
			output: true,
		},
		{
			arr:    []int{3, 0, 2, 1, 2},
			start:  2,
			output: false,
		},
	}

	for i, testcase := range testcases {
		t.Run(fmt.Sprintf("case-%d", i+1), func(t *testing.T) {
			require.Equal(t, testcase.output, lib.CanReach(testcase.arr, testcase.start))
		})
	}
}
