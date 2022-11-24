/* Search Insert Position
Source		: https://leetcode.com/problems/search-insert-position/
Level		: Easy
Description	: Given a sorted array of distinct integers and a target value, return the index if the target is found. If
			not, return the index where it would be if it were inserted in order. You must write an algorithm with
			O(log n) runtime complexity.

Example 1:
Input: nums = [1,3,5,6], target = 5
Output: 2

Example 2:
Input: nums = [1,3,5,6], target = 2
Output: 1

Example 3:
Input: nums = [1,3,5,6], target = 7
Output: 4

Source Solution: https://www.programiz.com/dsa/binary-search
*/

package leetcode

import "fmt"

func SearchInsert(nums []int, target int) int {
	low, high, mid := 0, len(nums)-1, 0
	for low <= high {
		mid = (low + high) / 2
		fmt.Println("test mid =", mid)
		switch {
		case nums[mid] == target:
			return mid
		case target > nums[mid]:
			low = mid + 1
		default:
			high = mid - 1
		}
	}

	if nums[mid] > target {
		return mid
	}
	return mid + 1
}
