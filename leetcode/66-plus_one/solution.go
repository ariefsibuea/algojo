package plusone

func PlusOne(digits []int) []int {
	remainder, index := 1, len(digits)-1

	for index >= 0 {
		digits[index] += remainder
		if digits[index] < 10 {
			remainder = 0
			break
		}

		digits[index] = 0
		remainder = 1
		index--
	}

	if remainder == 1 {
		digits = append([]int{1}, digits...)
	}
	return digits
}
