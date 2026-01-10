package main

import (
	"fmt"
	"os"

	"github.com/ariefsibuea/algojo/libs/go/cmp"
	"github.com/ariefsibuea/algojo/libs/go/format"
)

/*
 * Problem 			: Array Left Rotation
 * Topics           : Arrays, Simulation, Rotations
 * Level            : Easy
 * URL              : https://www.hackerrank.com/challenges/array-left-rotation/problem
 * Description      : Given n integers and a rotation count d, repeatedly move the first element to the end to
 * 					  simulate d left rotations. Input supplies n and d followed by the array values, and the goal is
 * 					  to print or return the array after all rotations so the relative order reflects the left shifts.
 * Examples         : Input: n=5 d=4 arr=[1 2 3 4 5]        -> Output: [5 1 2 3 4]
 *              	  Input: n=5 d=2 arr=[1 2 3 4 5]        -> Output: [3 4 5 1 2]
 *              	  Input: n=1 d=0 arr=[41]               -> Output: [41]
 */

func rotateLeft(d int32, arr []int32) []int32 {
	var n = len(arr)

	k := int(d) % n
	if k == 0 {
		return arr
	}

	newArr := make([]int32, n)
	for i := 0; i < n; i++ {
		newArr[i] = arr[(i+k)%n]
	}

	return newArr
}

func RunTestLeftRotation() {
	testCases := map[string]struct {
		d      int32
		arr    []int32
		expect []int32
	}{
		"case-1": {
			d:      4,
			arr:    []int32{1, 2, 3, 4, 5},
			expect: []int32{5, 1, 2, 3, 4},
		},
	}

	for name, testCase := range testCases {
		fmt.Printf("RUN %s\n", name)

		result := rotateLeft(testCase.d, testCase.arr)
		format.PrintInput(map[string]interface{}{"d": testCase.d, "arr": testCase.arr})

		if !cmp.EqualSlices(result, testCase.expect) {
			format.PrintFailed("expect = %v - got = %v\n", testCase.expect, result)
			os.Exit(1)
		}
		format.PrintSuccess("test case '%s' passed", name)
	}

	fmt.Printf("\n✅ All tests passed!\n")
}
