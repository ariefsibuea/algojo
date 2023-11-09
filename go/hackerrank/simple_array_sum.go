package hackerrank

// Description:
// Given an array of integers, find the sum of its elements.
// For example, if the array ar = [1,2,3]; 1+2+3 = 6, return 6.

func SimpleArraySum(ar []int32) int32 {
	var res int32
	for _, v := range ar {
		res += v
	}
	return res
}
