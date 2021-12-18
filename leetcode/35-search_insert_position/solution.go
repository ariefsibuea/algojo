package searchinsertposition

import "fmt"

/**
 * Problem source: https://leetcode.com/problems/binary-search/
 * Agenda:
 * 		Study Plan: Algorithm I
 *		Day I Binary Search
 *		Level: Easy
 * Solution source: https://www.programiz.com/dsa/binary-search
**/

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
