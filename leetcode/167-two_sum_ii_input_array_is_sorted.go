/* Two Sum II - Input Array Is Sorted
Source		: https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/
Level		: Medium
Description	: Given a 1-indexed array of integers numbers that is already sorted in non-decreasing order, find two
			numbers such that they add up to a specific target number. Let these two numbers be numbers[index1] and
			numbers[index2] where 1 <= index1 < index2 <= numbers.length. Return the indices of the two numbers, index1
			and index2, added by one as an integer array [index1, index2] of length 2.

Example 1:
Input: numbers = [2,7,11,15], target = 9
Output: [1,2]
Explanation: The sum of 2 and 7 is 9. Therefore, index1 = 1, index2 = 2. We return [1, 2].

Example 2:
Input: numbers = [2,3,4], target = 6
Output: [1,3]
Explanation: The sum of 2 and 4 is 6. Therefore index1 = 1, index2 = 3. We return [1, 3].

Example 3:
Input: numbers = [-1,0], target = -1
Output: [1,2]
Explanation: The sum of -1 and 0 is -1. Therefore index1 = 1, index2 = 2. We return [1, 2].
*/

package leetcode

// TwoSum implements two pointers technique to solve two sum with sorted array problem
func TwoSumII(numbers []int, target int) []int {
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
