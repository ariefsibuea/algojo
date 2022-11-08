package problems

func Mystery(a, b int32) int32 {
	x := a
	y := b

	for x != y {
		if x > y {
			x = x - y
		}
		if x < y {
			y = y - x
		}
	}

	return x
}
