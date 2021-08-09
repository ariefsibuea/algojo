package palindromenumber

func IsPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	return x == reverseNumber(x)
}

func reverseNumber(x int) int {
	reverse := 0

	for x != 0 {
		remainder := x % 10
		reverse = (reverse * 10) + remainder
		x /= 10
	}

	return reverse
}
