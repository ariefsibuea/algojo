package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/ariefsibuea/algojo/libs/go/cmp"
	"github.com/ariefsibuea/algojo/libs/go/format"
)

/*
 * Problem 			: Rank String Parsing
 * Topics           : String Parsing
 * Level            : Easy
 * Description      : You are given a string s representing a rank in a leaderboard. The rank can be in one of the
 * 					  following forms:
 * 						1. A single positive integer, e.g. "10"
 * 						2. A single positive integer prefixed by "=", e.g. "=5"
 * 						3. A range of positive integers "A-B", e.g. "201-250"
 * 					  You need to write a function that returns a single integer rankValue defined as:
 * 						- If s is "10" → rankValue = 10
 * 						- If s is "=5" → rankValue = 5
 * 						- If s is "201-250" → rankValue = 201 (the lower bound of the range)
 * 					  If the string s is not in a valid format, you may assume it won’t appear in the input (you don’t
 * 					  need to handle invalid formats beyond the three defined forms).
 * Examples         :
 * 					  Example 1:
 * 					  s = "10"
 * 					  Output: 10
 *
 * 					  Example 2:
 * 					  s = "=7"
 * 					  Output: 7
 *
 * 					  Example 3:
 * 					  s = "201-250"
 * 					  Output: 201
 *
 * 					  Example 4:
 * 					  s = "=1-5"
 * 					  This is not a valid input according to our formats, and you can assume it will not appear.
 */

func rankStringParsing(s string) int {
	// Split the string by "-". If the length of the result > 1 with prefix "=", then the format will be "=n-m". Return
	// 0 as expression for invalid format.
	strRanks := strings.Split(s, "-")
	if len(strRanks) > 2 {
		return 0
	}
	// BUG: how if the input is "-5"?
	if len(strRanks) > 1 && strRanks[0][0] == '=' {
		return 0
	}

	// Remove prefix "=" if any
	validStrRank := strRanks[0]
	validStrRank, _ = strings.CutPrefix(validStrRank, "=")

	// Parse the string rank using built-in function
	rank, err := strconv.ParseInt(validStrRank, 10, 32)
	if err != nil {
		return 0
	}

	return int(rank)
}

func RunTestRankStringParsing() {
	testCases := map[string]struct {
		s      string
		expect int
	}{
		"case-1": {
			s:      "10",
			expect: 10,
		},
		"case-2": {
			s:      "=7",
			expect: 7,
		},
		"case-3": {
			s:      "201-250",
			expect: 201,
		},
		"case-4": {
			s:      "=1-5",
			expect: 0,
		},
		"case-5": {
			s:      "1-5-10",
			expect: 0,
		},
	}

	for name, testCase := range testCases {
		fmt.Printf("RUN %s\n", name)

		result := rankStringParsing(testCase.s)
		format.PrintInput(map[string]interface{}{"s": testCase.s})

		if !cmp.EqualNumbers(result, testCase.expect) {
			format.PrintFailed("expect = %v - got = %v\n", testCase.expect, result)
			os.Exit(1)
		}
		format.PrintSuccess("test case '%s' passed", name)
	}

	fmt.Printf("\n✅ All tests passed!\n")
}
