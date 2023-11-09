/* First Bad Version
Source		: https://leetcode.com/problems/first-bad-version/
Level		: Easy
Description	:

Example 1:
Input: n = 5, bad = 4
Output: 4
Explanation:
call isBadVersion(3) -> false
call isBadVersion(5) -> true
call isBadVersion(4) -> true
Then 4 is the first bad version.

Example 2:
Input: n = 1, bad = 1
Output: 1

Solution source: https://www.programiz.com/dsa/binary-search
*/

package leetcode

// FirstBadVersion implements binary tree to find the first bad version
func (soln Solution) FirstBadVersion(n int) int {
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
