/* Two Sum
Source		: https://leetcode.com/problems/two-sum/
Level		: Easy
Description	: Given an array of integers `nums` and an integer `target`, return indices of the two numbers such that
			they add up to `target`.

Example 1:
Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].

Example 2:
Input: nums = [3,2,4], target = 6
Output: [1,2]

Example 3:
Input: nums = [3,3], target = 6
Output: [0,1]
*/

package leetcode

func TwoSum(nums []int, target int) []int {
	m := map[int]int{}
	for i := range nums {
		complement := target - nums[i]
		if _, ok := m[complement]; ok {
			return []int{m[complement], i}
		}

		m[nums[i]] = i
	}
	return nil
}
