package maximumsubarray

import "math"

func MaxSubArray(nums []int) int {
	sum, max := 0, math.MinInt32

	for _, v := range nums {
		if sum < 0 {
			sum = 0
		}
		sum += v

		if max < sum {
			max = sum
		}
	}

	return max
}
