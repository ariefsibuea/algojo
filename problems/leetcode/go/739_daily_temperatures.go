package main

import (
	"fmt"

	"github.com/ariefsibuea/algojo/libs/go/cmp"
	"github.com/ariefsibuea/algojo/libs/go/format"
	"github.com/ariefsibuea/algojo/libs/go/runner"
)

func init() {
	register("DailyTemperatures", RunTestDailyTemperatures)
}

/*
 * Problem	: Daily Temperatures
 * Topics	: Array, Stack, Monotonic Stack
 * Level	: Medium
 * URL		: https://leetcode.com/problems/daily-temperatures/
 *
 * Description:
 * 		Given an array of integers representing daily temperatures, return an array where each element indicates how
 * 		many days you must wait after that day to encounter a warmer temperature. If no warmer temperature exists in
 * 		the future, that element should be 0.
 *
 * Constraints:
 * 		- 1 <= temperatures.length <= 10^5
 * 		- 30 <= temperatures[i] <= 100
 *
 * Examples:
 * 		Example 1:
 * 		Input: temperatures = [73,74,75,71,69,72,76,73]
 * 		Output: [1,1,4,2,1,1,0,0]
 * 		Explanation: For the first day (73°F), the next warmer day is the very next day (74°F), so answer[0] = 1. For
 * 		day 3 (75°F), the next warmer day is 4 days later (76°F), so answer[2] = 4. The last two days have no warmer
 * 		future days.
 *
 * 		Example 2:
 * 		Input: temperatures = [30,40,50,60]
 * 		Output: [1,1,1,0]
 * 		Explanation: Each day has a warmer temperature the next day, except the last day which has no future days.
 *
 * 		Example 3:
 * 		Input: temperatures = [30,60,90]
 * 		Output: [1,1,0]
 * 		Explanation: Each day has a warmer temperature the next day, except the last day which has no future days.
 */

func dailyTemperatures(temperatures []int) []int {
	stack := make([]int, 0, len(temperatures))
	answer := make([]int, len(temperatures))

	for i := len(temperatures) - 1; i >= 0; i-- {
		for len(stack) > 0 && temperatures[stack[len(stack)-1]] <= temperatures[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			answer[i] = stack[len(stack)-1] - i
		}
		stack = append(stack, i)
	}

	return answer
}

func RunTestDailyTemperatures() {
	runner.InitMetrics("DailyTemperatures")

	testCases := map[string]struct {
		temperatures []int
		expect       []int
	}{
		"example-1-mixed-temperatures": {
			temperatures: []int{73, 74, 75, 71, 69, 72, 76, 73},
			expect:       []int{1, 1, 4, 2, 1, 1, 0, 0},
		},
		"example-2-increasing-temperatures": {
			temperatures: []int{30, 40, 50, 60},
			expect:       []int{1, 1, 1, 0},
		},
		"example-3-ascending-sequence": {
			temperatures: []int{30, 60, 90},
			expect:       []int{1, 1, 0},
		},
		"single-day": {
			temperatures: []int{50},
			expect:       []int{0},
		},
		"all-decreasing-temperatures": {
			temperatures: []int{100, 90, 80, 70, 60},
			expect:       []int{0, 0, 0, 0, 0},
		},
		"all-increasing-temperatures": {
			temperatures: []int{30, 31, 32, 33},
			expect:       []int{1, 1, 1, 0},
		},
		"same-temperature-throughout": {
			temperatures: []int{50, 50, 50, 50, 50},
			expect:       []int{0, 0, 0, 0, 0},
		},
		"large-gap-warm-day": {
			temperatures: []int{30, 31, 32, 33, 90},
			expect:       []int{1, 1, 1, 1, 0},
		},
		"two-days-same-temperature": {
			temperatures: []int{50, 50},
			expect:       []int{0, 0},
		},
		"fluctuating-temperatures": {
			temperatures: []int{50, 40, 45, 35, 55, 30},
			expect:       []int{4, 1, 2, 1, 0, 0},
		},
	}

	var passedCount uint16 = 0

	for name, testCase := range testCases {
		fmt.Printf("RUN %s\n", name)

		temperaturesCopy := make([]int, len(testCase.temperatures))
		copy(temperaturesCopy, testCase.temperatures)

		format.PrintInput(map[string]any{"temperatures": temperaturesCopy})

		result := runner.ExecCountMetrics(dailyTemperatures, testCase.temperatures).([]int)
		if !cmp.EqualSlices(result, testCase.expect) {
			format.PrintFailed("expect = %v - got = %v", testCase.expect, result)
			continue
		}

		format.PrintSuccess("test case '%s' passed", name)
		passedCount++
	}

	fmt.Printf("\n📊 Test Summary: %d/%d passed\n", passedCount, len(testCases))
	runner.PrintMetrics()
}
