package main

import (
	"fmt"
	"os"

	"github.com/ariefsibuea/algojo/libs/go/cmp"
	"github.com/ariefsibuea/algojo/libs/go/format"
)

/*
 * Problem 			: Frequency Queries
 * Topics           : Hash Maps, Frequency Counting, Data Structures
 * Level            : Medium
 * URL              : https://www.hackerrank.com/challenges/frequency-queries/problem
 * Description      : Process a sequence of queries on a multiset of integers where each query is a pair [operation,
 *                    value]. Operation 1 inserts the value, operation 2 removes a single occurrence if present, and
 *                    operation 3 checks whether any integer currently occurs exactly 'value' times. For every
 *                    frequency check, append 1 to the answer array when such a frequency exists, otherwise append 0.
 * Examples         : Input: [[1 1], [2 2], [3 2], [1 1], [1 1], [2 1], [3 2]] -> Output: [0 1]
 *              	  Input: [[1 5], [1 6], [3 2], [1 10], [1 10], [1 6], [2 5], [3 2]] -> Output: [0 1]
 */

const (
	OperationInsert = 1
	OperationDelete = 2
	OperationCheck  = 3
)

func freqQuery(queries [][]int32) []int32 {
	if len(queries) == 0 {
		return []int32{}
	}

	var result = []int32{}
	var freqOfData = map[int32]int{}
	var freqCount = map[int]int32{}

	for _, query := range queries {
		operation, data := query[0], query[1]

		switch operation {
		case OperationInsert:
			freq := freqOfData[data]
			if freq > 0 {
				freqCount[freq]--
			}
			freqOfData[data] = freq + 1
			freqCount[freq+1] = freqCount[freq+1] + 1
		case OperationDelete:
			freq := freqOfData[data]
			if freq > 0 {
				freqCount[freq]--
				if freq-1 > 0 {
					freqOfData[data] = freq - 1
					freqCount[freq-1] = freqCount[freq-1] + 1
				} else {
					delete(freqOfData, data)
				}
			}
		case OperationCheck:
			if freqCount[int(data)] > 0 {
				result = append(result, 1)
			} else {
				result = append(result, 0)
			}
		}
	}

	return result
}

func RunTestFrequencyQueries() {
	testCases := map[string]struct {
		queries [][]int32
		expect  []int32
	}{
		"case-1": {
			queries: [][]int32{
				{1, 1},
				{2, 2},
				{3, 2},
				{1, 1},
				{1, 1},
				{2, 1},
				{3, 2},
			},
			expect: []int32{0, 1},
		},
		"case-2": {
			queries: [][]int32{
				{1, 5},
				{1, 6},
				{3, 2},
				{1, 10},
				{1, 10},
				{1, 6},
				{2, 5},
				{3, 2},
			},
			expect: []int32{0, 1},
		},
		"case-3": {
			queries: [][]int32{
				{3, 4},
				{2, 1003},
				{1, 16},
				{3, 1},
			},
			expect: []int32{0, 1},
		},
		"case-4": {
			queries: [][]int32{
				{1, 3},
				{2, 3},
				{3, 2},
				{1, 4},
				{1, 5},
				{1, 5},
				{1, 4},
				{3, 2},
				{2, 4},
				{3, 2},
			},
			expect: []int32{0, 1, 1},
		},
	}

	for name, testCase := range testCases {
		fmt.Printf("RUN %s\n", name)

		result := freqQuery(testCase.queries)
		format.PrintInput(map[string]interface{}{"queries": testCase.queries})

		if !cmp.EqualSlices(result, testCase.expect) {
			format.PrintFailed("expect = %v - got = %v\n", testCase.expect, result)
			os.Exit(1)
		}
		format.PrintSuccess("test case '%s' passed", name)
	}

	fmt.Printf("\n✅ All tests passed!\n")
}
