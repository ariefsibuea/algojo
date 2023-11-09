/* Partition Equal Subset Sum
Source		: https://leetcode.com/problems/partition-equal-subset-sum/
Level		: Medium
Description	: Given a non-empty array nums containing only positive integers, find if the array can be partitioned into
			two subsets such that the sum of elements in both subsets is equal.

Example 1:
Input: nums = [1,5,11,5]
Output: true
Explanation: The array can be partitioned as [1, 5, 5] and [11].

Example 2:
Input: nums = [1,2,3,5]
Output: false
Explanation: The array cannot be partitioned into equal sum subsets.

Solve this problem using Dynamic Programming
*/

package leetcode

func (soln Solution) CanPartition(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}

	if sum%2 != 0 {
		return false
	}

	// bottom up manner to find subset
	part := make([]bool, (sum/2)+1)
	for i := 0; i < len(nums); i++ {
		for j := sum / 2; j >= nums[i]; j-- {
			if part[j-nums[i]] || j == nums[i] {
				part[j] = true
			}
		}
	}

	return part[sum/2]
}
