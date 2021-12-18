package numbersatmostngivendigitset

import (
	"math"
	"sort"
	"strconv"
)

// Problem source: https://leetcode.com/problems/numbers-at-most-n-given-digit-set/submissions/
// Answer source: https://www.youtube.com/watch?v=nqCM8opotjU

func AtMostNGivenDigitSet(digits []string, n int) int {
	sort.Slice(digits, func(i, j int) bool {
		return digits[i] < digits[j]
	})

	numstr := strconv.Itoa(n)
	count := 0
	for i := 1; i <= len(numstr)-1; i++ {
		count += powInt(len(digits), i)
	}

	i := 0
	for i < len(numstr) {
		j := 0
		for j < len(digits) && digits[j][0] < numstr[i] {
			count += powInt(len(digits), len(numstr)-1-i)
			j++
		}
		if j == len(digits) || digits[j][0] > numstr[i] {
			break
		}
		i++
	}

	if i == len(numstr) {
		count++
	}

	return count
}

func powInt(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}
