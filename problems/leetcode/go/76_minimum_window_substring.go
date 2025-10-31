package main

import (
	"fmt"
	"math"
	"os"

	"github.com/ariefsibuea/algojo/libs/go/cmp"
)

func minWindowSubstring(s string, t string) string {
	requiredChars := 0
	target := [128]int{}
	for _, r := range t {
		if target[r] == 0 {
			requiredChars++
		}
		target[r]++
	}

	matchedChars := 0
	minWindowLength := math.MaxInt

	start := 0
	resultStart := 0

	source := [128]int{}
	for end, r := range s {
		source[r]++

		if target[r] > 0 && target[r] == source[r] {
			matchedChars++
		}

		for requiredChars == matchedChars {
			currentLenght := (end - start) + 1
			if currentLenght < minWindowLength {
				minWindowLength = currentLenght
				resultStart = start
			}

			leftChar := s[start]
			source[leftChar]--

			if target[leftChar] > 0 && target[leftChar] > source[leftChar] {
				matchedChars--
			}

			start++
		}
	}

	if minWindowLength == math.MaxInt {
		return ""
	}

	return s[resultStart : resultStart+minWindowLength]
}

func RunTestMinWindowSubstring() {
	testCases := map[string]struct {
		s      string
		t      string
		expect string
	}{
		"case-1": {
			s:      "ADOBECODEBANC",
			t:      "ABC",
			expect: "BANC",
		},
	}

	for name, testCase := range testCases {
		fmt.Printf("RUN %s\n", name)

		result := minWindowSubstring(testCase.s, testCase.t)
		if !cmp.EqualStrings(result, testCase.expect) {
			fmt.Printf("=== FAILED: expect = %v - got = %v\n", testCase.expect, result)
			os.Exit(1)
		}
		fmt.Printf("=== PASSED\n")
	}

	fmt.Printf("\n✅ All tests passed!\n")
}
