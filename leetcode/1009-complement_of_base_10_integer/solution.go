package complementofbase10integer

/**
 * Problem source: https://leetcode.com/problems/complement-of-base-10-integer/
**/

func BitwiseComplement(n int) int {
	if n == 1 {
		return 0
	}
	if n == 0 {
		return 1
	}

	rem, res, i := 0, 0, 0
	for n != 0 {
		rem = n % 2
		if rem == 0 {
			res += (1 << i)
		}

		n /= 2
		i += 1
	}
	return res
}
