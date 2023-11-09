package hackerrank

import "fmt"

// MiniMaxSum receives a parameter array with length exactly 5.
func MiniMaxSum(arr []int32) (int64, int64) {
	min := arr[0]
	max := min
	totalMin := int64(arr[0])
	totalMax := totalMin

	for i := 1; i < len(arr); i++ {
		if min > arr[i] {
			min = arr[i]
		}
		if max < arr[i] {
			max = arr[i]
		}
		totalMax += int64(arr[i])
		totalMin += int64(arr[i])
	}

	fmt.Printf("%d %d\n", totalMin-int64(max), totalMax-int64(min))

	return totalMin - int64(max), totalMax - int64(min)
}
