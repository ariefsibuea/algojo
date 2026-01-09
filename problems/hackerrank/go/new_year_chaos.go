package main

import (
	"fmt"
	"os"

	"github.com/ariefsibuea/algojo/libs/go/cmp"
	"github.com/ariefsibuea/algojo/libs/go/format"
)

/*
 * Problem 			: New Year Chaos
 * Topics           : Arrays, Greedy, Simulation
 * Level            : Medium
 * URL              : https://www.hackerrank.com/challenges/new-year-chaos/problem
 * Description      : Given the final state of a queue where each person can bribe the person directly in front at most
 *                    twice, determine if the configuration is valid and, if so, compute the minimum number of bribes
 *                    that could have produced it. When any person has moved forward more than two positions, the queue
 *                    is deemed “Too chaotic”; otherwise, count how many times individuals jumped ahead.
 * Examples         : Input: [2 1 5 3 4] -> Output: 3 (person 5 bribed twice, person 3 once)
 *              	  Input: [2 5 1 3 4] -> Output: Too chaotic (person 1 moved ahead more than twice)
 */

func minimumBribes(q []int32) int {
	var bribes = 0

	for i, originalPos := range q {
		currentPos := i + 1

		if originalPos-int32(currentPos) > 2 {
			// fmt.Printf("Too chaotic\n")
			return -1
		}

		// we have -2 here because a person can bribe max of 2
		for j := max(originalPos-2, 0); j <= int32(i); j++ {
			if q[j] > originalPos {
				bribes += 1
			}
		}
	}

	return bribes
}

func RunTestNewYearChaos() {
	testCases := map[string]struct {
		q      []int32
		expect int
	}{
		"case-1": {
			q:      []int32{2, 1, 5, 3, 4},
			expect: 3,
		},
		"case-2": {
			q:      []int32{2, 5, 1, 3, 4},
			expect: -1,
		},
	}

	for name, testCase := range testCases {
		fmt.Printf("RUN %s\n", name)

		result := minimumBribes(testCase.q)
		format.PrintInput(map[string]interface{}{"1": testCase.q})

		if !cmp.EqualNumbers(result, testCase.expect) {
			format.PrintFailed("expect = %v - got = %v\n", testCase.expect, result)
			os.Exit(1)
		}
		format.PrintSuccess("test case '%s' passed", name)
	}

	fmt.Printf("\n✅ All tests passed!\n")
}
