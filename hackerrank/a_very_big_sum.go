package hackerrank

// Description:
// Given an array, return the sum of array elements.
// Example:
// 	ar = [1000000001, 1000000002, 1000000003, 1000000004, 1000000005]
// 	result = 5000000015

func AVeryBigSum(ar []int64) int64 {
	var res int64
	for _, v := range ar {
		res += v
	}
	return res
}
