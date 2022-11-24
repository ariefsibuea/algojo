/* Remove Duplicates from Sorted Array
Source		: https://leetcode.com/problems/remove-duplicates-from-sorted-array/
Level		: Easy
Description	: Given an integer array nums sorted in non-decreasing order, remove the duplicates in-place such that each
			unique element appears only once. The relative order of the elements should be kept the same. Do not
			allocate extra space for another array. You must do this by modifying the input array in-place with O(1)
			extra memory.

Example 1:
Input: nums = [1,1,2]
Output: 2, nums = [1,2,_]
Explanation: Your function should return k = 2, with the first two elements of nums being 1 and 2 respectively.
It does not matter what you leave beyond the returned k (hence they are underscores).

Example 2:
Input: nums = [0,0,1,1,1,2,2,3,3,4]
Output: 5, nums = [0,1,2,3,4,_,_,_,_,_]
Explanation: Your function should return k = 5, with the first five elements of nums being 0, 1, 2, 3, and 4 respectively.
It does not matter what you leave beyond the returned k (hence they are underscores).
*/

package leetcode

func RemoveDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return 1
	}

	index, lastNum := 1, nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] == lastNum {
			continue
		}

		lastNum = nums[i]
		nums[index] = lastNum
		index++
	}

	return index
}
