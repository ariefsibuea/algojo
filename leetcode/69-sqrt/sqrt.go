package sqrt

func MySqrt(x int) int {
	if x < 2 {
		return x
	}

	left, right := 1, x
	for left < right {
		mid := left + ((right - left) / 2)

		if mid*mid == x {
			return mid
		}

		if mid*mid > x {
			right = mid
			continue
		}
		left = mid + 1
	}

	return left - 1
}
