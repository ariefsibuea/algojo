package hackerrank

func BirthdayCakeCandles(candles []int32) int32 {
	tallest := candles[0]
	total := int32(1)

	for i := 1; i < len(candles); i++ {
		switch {
		case tallest < candles[i]:
			tallest = candles[i]
			total = 1
		case tallest == candles[i]:
			total++
		}
	}

	return total
}
