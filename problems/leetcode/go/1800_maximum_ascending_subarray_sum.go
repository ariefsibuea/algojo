package main

/**
 * LeetCode Problem : Maximum Ascending Subarray Sum
 * Topics           : Array, Weekly Contest 233
 * Level            : Easy
 * URL              : https://leetcode.com/problems/maximum-ascending-subarray-sum
 * Description      :
 * Examples         :
 * 					Example 1:
 * 					Input: nums = [10,20,30,5,10,50]
 * 					Output: 65
 * 					Explanation: [5,10,50] is the ascending subarray with the maximum sum of 65.
 *
 * 					Example 2:
 * 					Input: nums = [10,20,30,40,50]
 * 					Output: 150
 * 					Explanation: [10,20,30,40,50] is the ascending subarray with the maximum sum of 150.
 *
 * 					Example 3:
 * 					Input: nums = [12,17,15,13,10,11,12]
 * 					Output: 33
 * 					Explanation: [10,11,12] is the ascending subarray with the maximum sum of 33.
 */

func maxAscendingSum(nums []int) int {
	sum, maxSum := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] <= nums[i-1] {
			sum = 0
		}

		sum += nums[i]
		maxSum = max(maxSum, sum)
	}
	return maxSum
}
