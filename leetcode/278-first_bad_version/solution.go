package firstbadversion

/**
 * Problem source: https://leetcode.com/problems/binary-search/
 * Agenda:
 * 		Study Plan: Algorithm I
 *		Day I Binary Search
 *		Level: Easy
 * Solution source: https://www.programiz.com/dsa/binary-search
**/

// FirstBadVersion implements binary tree to find the first bad version
func FirstBadVersion(n int) int {
	// list of version
	low, high, mid := 1, n, 0
	firstBad := n

	for low <= high {
		mid = (low + high) / 2
		switch {
		case isBadVersion(mid):
			if mid < firstBad {
				firstBad = mid
			}
			high = mid - 1
		default:
			low = mid + 1
		}
	}

	return firstBad
}

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */
func isBadVersion(version int) bool {
	// this function is not implemented here
	return true
}
