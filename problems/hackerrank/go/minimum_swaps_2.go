package main

import (
	"fmt"
	"os"

	"github.com/ariefsibuea/algojo/libs/go/cmp"
	"github.com/ariefsibuea/algojo/libs/go/format"
)

/*
 * Problem 			: Minimum Swaps 2
 * Topics           : Arrays, Greedy, Cycle Detection
 * Level            : Medium
 * URL              : https://www.hackerrank.com/challenges/minimum-swaps-2/problem
 * Description      : Given an array of consecutive integers from 1..n in arbitrary order,
 *                    determine the minimum number of swaps needed to sort it in ascending
 *                    order by swapping any two elements.
 * Examples         : Input: [4, 3, 1, 2]        -> 3
 *                    Input: [2, 3, 4, 1, 5]     -> 3
 *                    Input: [1, 3, 5, 2, 4, 6]  -> 3
 */

func minimumSwaps(arr []int32) int32 {
	n := len(arr)

	pos := make([]int, n+1)
	for i, v := range arr {
		pos[v] = i
	}

	var swaps int32

	for i := 0; i < n; i++ {
		currentVal := arr[i]
		expectVal := int32(i + 1)
		if currentVal == expectVal {
			continue
		}

		j := pos[expectVal]

		arr[i], arr[j] = arr[j], arr[i]

		pos[currentVal] = j
		pos[expectVal] = i

		swaps++
	}

	return swaps
}

func RunTestMinimumSwaps2() {
	testCases := map[string]struct {
		arr    []int32
		expect int32
	}{
		"case-1": {
			arr:    []int32{7, 1, 3, 2, 4, 5, 6},
			expect: 5,
		},
		"case-2": {
			arr:    []int32{4, 3, 1, 2},
			expect: 3,
		},
		"case-3": {
			arr:    []int32{2, 3, 4, 1, 5},
			expect: 3,
		},
		"case-4": {
			arr:    []int32{1, 3, 5, 2, 4, 6, 7},
			expect: 3,
		},
	}

	for name, testCase := range testCases {
		fmt.Printf("RUN %s\n", name)

		result := minimumSwaps(testCase.arr)
		format.PrintInput(map[string]interface{}{"arr": testCase.arr})

		if !cmp.EqualNumbers(result, testCase.expect) {
			format.PrintFailed("expect = %v - got = %v\n", testCase.expect, result)
			os.Exit(1)
		}
		format.PrintSuccess("test case '%s' passed", name)
	}

	fmt.Printf("\n✅ All tests passed!\n")
}
