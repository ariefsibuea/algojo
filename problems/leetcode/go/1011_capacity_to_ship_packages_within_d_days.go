package main

import (
	"fmt"

	"github.com/ariefsibuea/algojo/libs/go/cmp"
	"github.com/ariefsibuea/algojo/libs/go/format"
	"github.com/ariefsibuea/algojo/libs/go/runner"
)

func init() {
	register("CapacityToShipPackagesWithinDDays", RunTestCapacityToShipPackagesWithinDDays)
}

/*
 * Problem	: Capacity To Ship Packages Within D Days
 * Topics	: Array, Binary Search
 * Level	: Medium
 * URL		: https://leetcode.com/problems/capacity-to-ship-packages-within-d-days/
 *
 * Description:
 *		A conveyor belt has packages that must be shipped within a given number of days. Each package has a weight,
 *		and packages must be loaded in order (no skipping or rearranging). The ship has a weight capacity limit - the
 *		total weight loaded on any single day cannot exceed this capacity. Find the minimum ship capacity required to
 *		ship all packages within the specified number of days.
 *
 * Constraints:
 *		- 1 <= days <= weights.length <= 5 * 10^4
 *		- 1 <= weights[i] <= 500
 *
 * Examples:
 *		Example 1:
 *		Input: weights = [1,2,3,4,5,6,7,8,9,10], days = 5
 *		Output: 15
 *		Explanation: With capacity 15, we can ship as follows:
 *		Day 1: [1,2,3,4,5] (weight 15), Day 2: [6,7] (weight 13), Day 3: [8] (weight 8),
 *		Day 4: [9] (weight 9), Day 5: [10] (weight 10). Total: 5 days.
 *
 *		Example 2:
 *		Input: weights = [3,2,2,4,1,4], days = 3
 *		Output: 6
 *		Explanation: With capacity 6, we can ship as follows:
 *		Day 1: [3,2] (weight 5), Day 2: [2,4] (weight 6), Day 3: [1,4] (weight 5). Total: 3 days.
 */

func shipWithinDays(weights []int, days int) int {
	left, right := 0, 0
	for _, w := range weights {
		left = max(left, w)
		right += w
	}

	feasible := func(cap int) bool {
		currentWeight := 0
		daysNeeded := 1

		for _, w := range weights {
			currentWeight += w
			if currentWeight > cap {
				daysNeeded++
				currentWeight = w
			}
		}

		return daysNeeded <= days
	}

	fitCapacity := 0
	for left <= right {
		mid := (left + right) / 2
		if feasible(mid) {
			fitCapacity = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return fitCapacity
}

func RunTestCapacityToShipPackagesWithinDDays() {
	runner.InitMetrics("CapacityToShipPackagesWithinDDays")

	testCases := map[string]struct {
		weights []int
		days    int
		expect  int
	}{
		"example-1-basic-case": {
			weights: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			days:    5,
			expect:  15,
		},
		"example-2-multiple-days": {
			weights: []int{3, 2, 2, 4, 1, 4},
			days:    3,
			expect:  6,
		},
		"single-package": {
			weights: []int{10},
			days:    1,
			expect:  10,
		},
		"days-equal-packages": {
			weights: []int{1, 2, 3, 4, 5},
			days:    5,
			expect:  5,
		},
		"single-day-all-packages": {
			weights: []int{1, 2, 3, 4, 5},
			days:    1,
			expect:  15,
		},
		"same-weights": {
			weights: []int{5, 5, 5, 5, 5},
			days:    3,
			expect:  10,
		},
		"large-weights": {
			weights: []int{500, 500, 500, 500, 500},
			days:    2,
			expect:  1500,
		},
		"increasing-weights": {
			weights: []int{1, 10, 100, 1000, 10000},
			days:    3,
			expect:  10000,
		},
		"minimum-constraint": {
			weights: []int{1},
			days:    1,
			expect:  1,
		},
		"capacity-equals-max-weight": {
			weights: []int{10, 50, 100, 100, 50, 10},
			days:    4,
			expect:  100,
		},
	}

	var passedCount uint16 = 0

	for name, testCase := range testCases {
		fmt.Printf("RUN %s\n", name)
		format.PrintInput(map[string]any{"weights": testCase.weights, "days": testCase.days})

		result := runner.ExecCountMetrics(shipWithinDays, testCase.weights, testCase.days).(int)
		if !cmp.EqualNumbers(result, testCase.expect) {
			format.PrintFailed("expect = %v - got = %v", testCase.expect, result)
			continue
		}

		format.PrintSuccess("test case '%s' passed", name)
		passedCount++
	}

	fmt.Printf("\n📊 Test Summary: %d/%d passed\n", passedCount, len(testCases))
	runner.PrintMetrics()
}
