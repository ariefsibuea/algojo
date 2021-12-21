package minimumabsolutedifference

/**
 * Problem source: https://leetcode.com/problems/minimum-absolute-difference/
**/

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
