package movezeroes

/**
 * Problem source: https://leetcode.com/problems/move-zeroes/
 * Agenda:
 * 		Study Plan: Algorithm I
 *		Day III Two Pointers
 *		Level: Easy
 * Solution source:
**/

// MoveZeroes implements two pointers technique to solve move zeroes problem
func MoveZeroes(nums []int) {
	lastNonZero := 0
	for index := range nums {
		if nums[index] == 0 {
			continue
		}
		nums[lastNonZero] = nums[index]
		lastNonZero += 1
	}
	for i := lastNonZero; i < len(nums); i++ {
		nums[i] = 0
	}
}
