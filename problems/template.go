package main

import (
	"fmt"

	"github.com/ariefsibuea/algojo/libs/go/format"
	"github.com/ariefsibuea/algojo/libs/go/runner"
)

/*
 * Problem	: <Title>
 * Topics	: <Algorithm Categories>
 * Level	: <Easy | Medium | Hard>
 * URL		: <URL>
 *
 * Description:
 * 		<Description>
 *
 * Constraints:
 * 		<Constraints>
 *
 * Examples:
 * 		<Examples>
 */

func RunTestXxx() {
	runner.InitMetrics("ProblemTitle")

	testCases := map[string]struct{}{
		"case-1": {
			// test case 1
		},
		"case-2": {
			// test case 2
		},
	}

	var passedCount uint16 = 0

	for name, testCase := range testCases {
		fmt.Printf("RUN %s\n", name)
		fmt.Println(testCase)
		// format.PrintInput(map[string]interface{}{"input-1": testCase.Input1})

		// result := twoSum(testCase.nums, testCase.target)
		// if !EqualSlices(result, testCase.expect) {
		//  format.PrintFailed("expect = %v - got = %v\n", testCase.expect, result)
		// 	continue
		// }

		format.PrintSuccess("test case '%s' passed", name)
		passedCount++
	}

	fmt.Printf("\n📊 Test Summary: %d/%d passed\n", passedCount, len(testCases))
	runner.PrintMetrics()
}
