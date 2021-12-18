package binarysearch

/**
 * Problem source: https://leetcode.com/problems/binary-search/
 * Agenda:
 * 		Study Plan: Algorithm I
 *		Day I Binary Search
 *		Level: Easy
 * Solution source: https://www.programiz.com/dsa/binary-search
**/

// Search implements binary search to find target and returns the target index
func Search(nums []int, target int) int {
	// list of index
	low, high, mid := 0, len(nums)-1, 0

	// repeat until the pointers low and high meet each other
	for low <= high {
		mid = (low + high) / 2
		switch {
		case target == nums[mid]:
			return mid
		case target > nums[mid]:
			low = mid + 1
		default:
			high = mid - 1
		}
	}

	return -1
}
