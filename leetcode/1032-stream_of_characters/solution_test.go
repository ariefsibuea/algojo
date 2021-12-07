package streamofcharacters_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	lib "github.com/ariefsibuea/cobra/leetcode/1032-stream_of_characters"
)

func Test_StreamOfCharacters(t *testing.T) {
	words := []string{"cd", "f", "kl"}
	inputStream := "abcdefghijkl"
	expectedOutputStream := []bool{false, false, false, true, false, true, false, false, false, false, false, true}

	streamChecker := lib.Constructor(words)
	actualOutputStream := make([]bool, 0)
	for _, letter := range []byte(inputStream) {
		actualOutputStream = append(actualOutputStream, streamChecker.Query(letter))
	}

	require.Len(t, actualOutputStream, len(expectedOutputStream))
	for i := range actualOutputStream {
		require.Equal(t, expectedOutputStream[i], actualOutputStream[i])
	}
}
