package main

import (
	"fmt"
	"os"
)

func isValidAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	charCount := make(map[rune]int)

	for _, c := range s {
		charCount[c] = charCount[c] + 1
	}

	for _, c := range t {
		count, ok := charCount[c]
		if !ok || count == 0 {
			return false
		}
		charCount[c] = count - 1
	}

	return true
}

func RunTestIsValidAnagram() {
	testCases := map[string]struct {
		s      string
		t      string
		expect bool
	}{
		"case-1": {
			s:      "anagram",
			t:      "nagaram",
			expect: true,
		},
		"case-2": {
			s:      "rat",
			t:      "car",
			expect: false,
		},
	}

	for name, testCase := range testCases {
		fmt.Printf("RUN %s\n", name)

		result := isValidAnagram(testCase.s, testCase.t)
		if !EqualBooleans(result, testCase.expect) {
			fmt.Printf("=== FAILED: expect = %v - got = %v\n", testCase.expect, result)
			os.Exit(1)
		}
		fmt.Printf("=== PASSED\n")
	}

	fmt.Printf("\n✅ All tests passed!\n")
}
