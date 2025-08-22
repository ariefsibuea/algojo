package main

import (
	"fmt"
	"os"
)

/**
 * LeetCode Problem : First Bad Version
 * Topic            : Binary Search, Interactive
 * Level            : Easy
 * URL              : https://leetcode.com/problems/first-bad-version
 * Description      : You are a product manager and currently leading a team to develop a new product. Unfortunately,
 * 			the
 * 			latest version of your product fails the quality check. Since each version is developed based on the
 * 			previous version, all the versions after a bad version are also bad. Given n versions [1, 2, ..., n], you
 * 			want to find out the first bad one, which causes all the following ones to be bad. You are given an API
 * 			bool isBadVersion(version) which returns whether version is bad.
 * Examples         :
 *         Example 1:
 *         Input: n = 5, bad = 4
 *         Output: 4
 *         Explanation: call isBadVersion(3) -> false
 *                      call isBadVersion(5) -> true
 *                      call isBadVersion(4) -> true
 *                      Then 4 is the first bad version.
 *
 *         Example 2:
 *         Input: n = 1, bad = 1
 *         Output: 1
 */

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool
 */

func firstBadVersion(n int, isBadVersion func(version int) bool) int {
	first, last := 1, n
	for first < last {
		mid := (first + last) / 2

		if isBadVersion(mid) {
			last = mid
		} else {
			first = mid + 1
		}
	}
	return first
}

func RunTestFirstBadVersion() {
	testCases := map[string]struct {
		n      int
		fn     func(version int) bool
		expect int
	}{
		"case-1": {
			n: 5,
			fn: func(version int) bool {
				return version >= 4
			},
			expect: 4,
		},
		"case-2": {
			n: 1,
			fn: func(version int) bool {
				return version >= 1
			},
			expect: 1,
		},
		"case-3": {
			n: 5,
			fn: func(version int) bool {
				return version >= 2
			},
			expect: 2,
		},
		"case-4": {
			n: 3,
			fn: func(version int) bool {
				return version >= 2
			},
			expect: 2,
		},
	}

	for name, testCase := range testCases {
		fmt.Printf("RUN %s\n", name)

		result := firstBadVersion(testCase.n, testCase.fn)
		if !EqualNumbers(result, testCase.expect) {
			fmt.Printf("=== FAILED: expect = %v - got = %v\n", testCase.expect, result)
			os.Exit(1)
		}
		fmt.Printf("=== PASSED\n")
	}

	fmt.Printf("\n✅ All tests passed!\n")
}
