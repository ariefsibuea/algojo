/* Squares of a Sorted Array
Source		: https://leetcode.com/problems/squares-of-a-sorted-array/
Level		: Easy
Description	: Given an integer array nums sorted in non-decreasing order, return an array of the squares of each number
			sorted in non-decreasing order.

Example 1:
Input: nums = [-4,-1,0,3,10]
Output: [0,1,9,16,100]
Explanation: After squaring, the array becomes [16,1,0,9,100].
After sorting, it becomes [0,1,9,16,100].

Example 2:
Input: nums = [-7,-3,2,3,11]
Output: [4,9,9,49,121]

Solution source: https://www.callicoder.com/squares-of-a-sorted-array/
*/

package leetcode

// SortedSquares implements two pointers technique to find squares of array
func (soln Solution) SortedSquares(nums []int) []int {
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
