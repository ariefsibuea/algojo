/* Minimum Absolute Difference
Source		: https://leetcode.com/problems/minimum-absolute-difference/
Level		: Easy
Description	:
	Given an array of distinct integers arr, find all pairs of elements with the minimum absolute difference of any two
	elements.
	Return a list of pairs in ascending order(with respect to pairs), each pair [a, b] follows
    - a, b are from arr
    - a < b
    - b - a equals to the minimum absolute difference of any two elements in arr

Example 1:
Input: arr = [4,2,1,3]
Output: [[1,2],[2,3],[3,4]]
Explanation: The minimum absolute difference is 1. List all pairs with difference equal to 1 in ascending order.

Example 2:
Input: arr = [1,3,6,10,15]
Output: [[1,3]]

Example 3:
Input: arr = [3,8,-10,23,19,-4,-14,27]
Output: [[-14,-10],[19,23],[23,27]]
*/

package leetcode

import "sort"

func MinimumAbsDifference(arr []int) [][]int {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	res := [][]int{}
	minDiff := arr[len(arr)-1]

	for i, j := 0, 1; j < len(arr); i, j = i+1, j+1 {
		if arr[j]-arr[i] == minDiff {
			res = append(res, []int{arr[i], arr[j]})
		}
		if arr[j]-arr[i] < minDiff {
			minDiff = arr[j] - arr[i]
			res = [][]int{}
			res = append(res, []int{arr[i], arr[j]})
		}
	}
	return res
}
