package squaresofasortedarray

/**
 * Problem source: https://leetcode.com/problems/squares-of-a-sorted-array/
 * Agenda:
 * 		Study Plan: Algorithm I
 *		Day II Two Pointers
 *		Level: Easy
 * Solution source: https://www.callicoder.com/squares-of-a-sorted-array/
**/

// SortedSquares implements two pointers technique to find squares of array
func SortedSquares(nums []int) []int {
	sortedSqrs := make([]int, len(nums))
	left, right := 0, len(nums)-1
	// index start from the end
	index := right

	for left <= right {
		leftSqr := nums[left] * nums[left]
		rightSqr := nums[right] * nums[right]
		switch {
		case leftSqr > rightSqr:
			sortedSqrs[index] = leftSqr
			left++
		default:
			sortedSqrs[index] = rightSqr
			right--
		}
		index--
	}

	return sortedSqrs
}
