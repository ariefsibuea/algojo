package rotatearray

/**
 * Problem source: https://leetcode.com/problems/rotate-array/
 * Agenda:
 * 		Study Plan: Algorithm I
 *		Day II Two Pointers
 *		Level: Medium
 * Solution source: https://www.geeksforgeeks.org/array-rotation/
**/

func gcd(x, y int) int {
	if y == 0 {
		return x
	}
	return gcd(y, x%y)
}

// Rotate implements juggling algorithm to rotate array elements to the right
func Rotate(nums []int, k int) {
	n := len(nums)
	k = k % n
	numSet := gcd(k, n)
	for i := n - 1; i >= n-numSet; i-- {
		temp := nums[i]
		curr := i
		for {
			next := curr - k
			if next < 0 {
				next = next + n
			}
			if next == i {
				break
			}
			nums[curr] = nums[next]
			curr = next
		}
		nums[curr] = temp
	}
}
