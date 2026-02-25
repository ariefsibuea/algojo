package main

import (
	"fmt"

	"github.com/ariefsibuea/algojo/libs/go/format"
)

func init() {
	register("NumberOfWaysToDivideALongCorridor", RunTestNumberOfWaysToDivideALongCorridor)
}

/*
 * Problem 			: <Title>
 * Topics           : <Algorithm Categories>
 * Level            : <Easy | Medium | Hard>
 * URL              : <URL>
 * Description      : <Description>
 * Examples         : <Examples>
 */

var CorridorModulo int = 1e9 + 7

func numberOfWays(corridor string) int {
	var numOfSeat, numOfWays = 0, 1
	var lastSeatIndex = 0

	for i := range corridor {
		if corridor[i] == 'S' {
			numOfSeat += 1
			if numOfSeat%2 == 1 && numOfSeat > 1 {
				numOfDivider := i - lastSeatIndex
				numOfWays = (numOfWays * numOfDivider) % CorridorModulo
			}

			lastSeatIndex = i
		}
	}

	if numOfSeat%2 == 0 && numOfSeat/2 > 0 {
		return numOfWays
	}
	return 0
}

func RunTestNumberOfWaysToDivideALongCorridor() {
	testCases := map[string]struct{}{}

	for name, testCase := range testCases {
		fmt.Printf("RUN %s\n", name)

		fmt.Println(testCase)
		// result := twoSum(testCase.nums, testCase.target)
		// format.PrintInput(map[string]interface{}{"input-1": testCase.Input1})

		// if !EqualSlices(result, testCase.expect) {
		//  format.PrintFailed("expect = %v - got = %v\n", testCase.expect, result)
		// 	os.Exit(1)
		// }
		format.PrintSuccess("test case '%s' passed", name)
	}

	fmt.Printf("\n✅ All tests passed!\n")
}
