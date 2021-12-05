package climbstairs

func ClimbStairs(n int) int {
	if n == 0 || n == 1 {
		return 1
	}

	last1, last2, sum := 1, 1, 0
	for i := 2; i <= n; i++ {
		sum = last1 + last2
		last2 = last1
		last1 = sum
	}

	return sum
}
