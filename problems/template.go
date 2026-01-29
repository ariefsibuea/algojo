package main

import (
	"fmt"

	"github.com/ariefsibuea/algojo/libs/go/format"
)

/*
 * Problem 			: <Title>
 * Topics           : <Algorithm Categories>
 * Level            : <Easy | Medium | Hard>
 * URL              : <URL>
 * Description      : <Description>
 * Constraints      : <Constraints>
 * Examples         : <Examples>
 */

func RunTestXxx() {
	testCases := map[string]struct{}{
		"case-1": {
			// test case 1
		},
		"case-2": {
			// test case 2
		},
	}

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
