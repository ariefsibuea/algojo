package main

import (
	"fmt"
	"os"
	"sort"

	"github.com/ariefsibuea/algojo/libs/go/cmp"
	"github.com/ariefsibuea/algojo/libs/go/format"
)

func init() {
	register("RepeatedDNASequence", RunTestRepeatedDNASequences)
}

/*
 * Problem          : Repeated DNA Sequences
 * Topics           : Hash Table, String, Bit Manipulation
 * Level            : Medium
 * URL              : https://leetcode.com/problems/repeated-dna-sequences
 * Description      : The DNA sequence is composed of a series of nucleotides abbreviated as 'A', 'C', 'G', and 'T'.
 *                    When studying DNA, it is useful to identify repeated sequences within the DNA. Given a string s
 *                    that represents a DNA sequence, return all the 10-letter-long sequences (substrings) that occur
 *                    more than once in a DNA molecule. You may return the answer in any order.
 * Constraints      :
 *                    - 1 <= s.length <= 10^5
 *                    - s[i] is either 'A', 'C', 'G', or 'T'.
 * Examples         :
 *                    Example 1:
 *                    Input: s = "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"
 *                    Output: ["AAAAACCCCC","CCCCCAAAAA"]
 *
 *                    Example 2:
 *                    Input: s = "AAAAAAAAAAAAA"
 *                    Output: ["AAAAAAAAAA"]
 */

func findRepeatedDnaSequences(s string) []string {
	n := len(s)
	if n < 10 {
		return []string{}
	}

	w := 10
	base := map[byte]uint32{'A': 0, 'C': 1, 'G': 2, 'T': 3}

	existed := make(map[uint32]bool)
	repeated := make(map[string]bool)

	// NOTE: To solve the challenge we can create 10-length substrings from the DNA sequence and check for repeating
	// using a hash map. However, it is more optimal to use rolling hash as the hash map key instead of the original
	// substring. Below, The Bitwise Window is implemented to count the hash value for each "10-length substring".

	// Create a mask to keep only the bits within 10-length. We need 20 bits set to 1.
	mask := uint32(1<<(w*2)) - 1

	currentHash := uint32(0)

	for i := 0; i < w; i++ {
		currentHash = (currentHash << 2) | base[s[i]]
	}

	existed[currentHash] = true
	start := 1

	for end := w; end < n; end++ {
		currentHash = ((currentHash << 2) | base[s[end]]) & mask
		if existed[currentHash] {
			repeated[s[start:end+1]] = true
		}

		existed[currentHash] = true
		start++
	}

	repeatedDnaSeq := make([]string, 0, len(repeated))
	for dnaSeq := range repeated {
		repeatedDnaSeq = append(repeatedDnaSeq, dnaSeq)
	}

	return repeatedDnaSeq
}

func RunTestRepeatedDNASequences() {
	testCases := map[string]struct {
		s      string
		expect []string
	}{
		"example-1": {
			s:      "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT",
			expect: []string{"AAAAACCCCC", "CCCCCAAAAA"},
		},
		"example-2": {
			s:      "AAAAAAAAAAAAA",
			expect: []string{"AAAAAAAAAA"},
		},
		"short-sequence": {
			s:      "AAAAAAAAA",
			expect: []string{},
		},
		"no-repeats": {
			s:      "ACGTACGTACGT",
			expect: []string{},
		},
		"all-same-10": {
			s:      "AAAAAAAAAA",
			expect: []string{},
		},
	}

	for name, testCase := range testCases {
		fmt.Printf("RUN %s\n", name)

		result := findRepeatedDnaSequences(testCase.s)
		format.PrintInput(map[string]interface{}{"s": testCase.s})

		// Sort both to ensure comparison is order-independent
		sort.Strings(result)
		sort.Strings(testCase.expect)

		if !cmp.EqualSlices(result, testCase.expect) {
			format.PrintFailed("expect = %v - got = %v\n", testCase.expect, result)
			os.Exit(1)
		}
		format.PrintSuccess("test case '%s' passed", name)
	}

	fmt.Printf("\n✅ All tests passed!\n")
}
