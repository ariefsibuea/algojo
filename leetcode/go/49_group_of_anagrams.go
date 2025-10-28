package main

import (
	"fmt"
	"os"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
	groups := make(map[string][]string)
	for _, str := range strs {
		signature := generateSignature(str)
		groups[signature] = append(groups[signature], str)
	}

	result := make([][]string, 0, len(groups))
	for _, group := range groups {
		result = append(result, group)
	}

	return result
}

func generateSignature(str string) string {
	alphabets := make([]rune, 26)

	for _, c := range str {
		alphabets[c-'a'] += 1
	}

	var signatureBuilder strings.Builder
	for i, count := range alphabets {
		if count > 0 {
			signatureBuilder.WriteRune('a' + rune(i))
			signatureBuilder.WriteString(fmt.Sprintf("%d", count))
		}
	}

	return signatureBuilder.String()
}

func RunTestGroupAnagrams() {
	// FIX: The test cases still fail because the order of the groups in the result slice result will be randomized
	// as it is converted from a map to a slice.
	testCases := map[string]struct {
		strs   []string
		expect [][]string
	}{
		"case-1": {
			strs: []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			expect: [][]string{
				{"eat", "tea", "ate"},
				{"bat"},
				{"tan", "nat"},
			},
		},
		"case-2": {
			strs: []string{""},
			expect: [][]string{
				{""},
			},
		},
		"case-3": {
			strs: []string{"a"},
			expect: [][]string{
				{"a"},
			},
		},
	}

	for name, testCase := range testCases {
		fmt.Printf("RUN %s\n", name)

		result := groupAnagrams(testCase.strs)
		for i, group := range testCase.expect {
			if !EqualSlices(group, result[i]) {
				fmt.Printf("=== FAILED: expect = %v - got = %v\n", group, result[i])
				os.Exit(1)
			}
		}
		fmt.Printf("=== PASSED\n")
	}

	fmt.Printf("\n✅ All tests passed!\n")
}
