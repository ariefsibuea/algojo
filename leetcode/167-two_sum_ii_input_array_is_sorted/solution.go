package twosumiiinputarrayissorted

/**
 * Problem source: https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/
 * Agenda:
 * 		Study Plan: Algorithm I
 *		Day III Two Pointers
 *		Level: Easy
 * Solution source:
**/

// TwoSum implements two pointers technique to solve two sum with sorted array problem
func TwoSum(numbers []int, target int) []int {
	start, end := 0, len(numbers)-1
	for start < end {
		if numbers[start]+numbers[end] == target {
			break
		} else if numbers[start]+numbers[end] < target {
			start += 1
		} else {
			end -= 1
		}
	}
	return []int{start + 1, end + 1}
}
