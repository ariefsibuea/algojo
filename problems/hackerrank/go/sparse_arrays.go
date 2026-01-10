package main

import (
	"fmt"
	"os"

	"github.com/ariefsibuea/algojo/libs/go/cmp"
	"github.com/ariefsibuea/algojo/libs/go/format"
)

/*
 * Problem 			: Sparse Arrays
 * Topics           : Arrays, Hash Maps, Frequency Counting
 * Level            : Medium
 * URL              : https://www.hackerrank.com/challenges/sparse-arrays/problem
 * Description      : Given an initial list of strings and a set of query strings, count how many times each query
 * 					  occurs in the list. The input first provides the string count and each string, followed by the
 * 					  number of queries and the query strings. Output should list one integer per query showing the
 * 					  frequency in the original list, which can be efficiently tracked with a map keyed by string.
 * Examples         : Input: strings=[aba baba aba xzxb], queries=[aba xzxb ab]		-> Output: [2 1 0]
 *              	  Input: strings=[def de fgh], queries=[de lmn fgh]				-> Output: [1 0 1]
 *              	  Input: strings=[abcde sdaklfj asdjf na basdn sdaklfj asdjf],
 * 							 queries=[abcde sdaklfj asdjf na basdn] 				-> Output: [1 2 2 1 1]
 */

func matchingStrings(stringList []string, queries []string) []int32 {
	stringCounter := make(map[string]int32)

	for _, s := range stringList {
		stringCounter[s]++
	}

	var result = make([]int32, 0, len(queries))

	for _, q := range queries {
		result = append(result, stringCounter[q])
	}

	return result
}

func RunTestSparseArrays() {
	testCases := map[string]struct {
		stringList []string
		queries    []string
		expect     []int32
	}{
		"case-1": {
			stringList: []string{
				"aba",
				"baba",
				"aba",
				"xzxb",
			},
			queries: []string{
				"aba",
				"xzxb",
				"ab",
			},
			expect: []int32{2, 1, 0},
		},
		"case-2": {
			stringList: []string{
				"def",
				"de",
				"fgh",
			},
			queries: []string{
				"de",
				"lmn",
				"fgh",
			},
			expect: []int32{1, 0, 1},
		},
	}

	for name, testCase := range testCases {
		fmt.Printf("RUN %s\n", name)

		result := matchingStrings(testCase.stringList, testCase.queries)
		format.PrintInput(map[string]interface{}{"stringList": testCase.stringList, "queries": testCase.queries})

		if !cmp.EqualSlices(result, testCase.expect) {
			format.PrintFailed("expect = %v - got = %v\n", testCase.expect, result)
			os.Exit(1)
		}
		format.PrintSuccess("test case '%s' passed", name)
	}

	fmt.Printf("\n✅ All tests passed!\n")
}
