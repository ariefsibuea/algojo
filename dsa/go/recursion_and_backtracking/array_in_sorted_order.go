package main

// IsArrayInAscOrder checks whether the array is in ascending order with recursion.
func IsArrayInAscOrder(n int, nums []int) bool {
	if n == 1 {
		return true
	}

	if nums[n-1] < nums[n-2] {
		return false
	}
	return IsArrayInAscOrder(n-1, nums)
}

// IsArrayInDescOrder checks whether the array is in descending order with recursion.
func IsArrayInDescOrder(n int, nums []int) bool {
	if n == 1 {
		return true
	}

	if nums[n-1] > nums[n-2] {
		return false
	}
	return IsArrayInAscOrder(n-1, nums)
}
