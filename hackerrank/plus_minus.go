package hackerrank

import (
	"fmt"
	"math"
)

func PlusMinus(arr []int32) (float64, float64, float64) {
	var positive, negative, zero int32
	for _, v := range arr {
		switch {
		case v == 0:
			zero++
		case v > 0:
			positive++
		case v < 0:
			negative++
		}
	}

	fmt.Printf("%.6f\n", float64(positive)/float64(len(arr)))
	fmt.Printf("%.6f\n", float64(negative)/float64(len(arr)))
	fmt.Printf("%.6f\n", float64(zero)/float64(len(arr)))

	return roundFloat(float64(positive)/float64(len(arr)), 6),
		roundFloat(float64(negative)/float64(len(arr)), 6),
		roundFloat(float64(zero)/float64(len(arr)), 6)
}

func roundFloat(val float64, precision int32) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}
