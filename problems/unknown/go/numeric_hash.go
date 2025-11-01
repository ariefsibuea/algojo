package main

import (
	"fmt"
	"os"

	"github.com/ariefsibuea/algojo/libs/go/cmp"
)

/**
 * Problem 			: Numeric Hash
 * Description      :
 * 					Given a string `key` and integer `MAX_SIZE`, generate a hash using these steps:
 * 						1. Convert each character to its ASCII value.
 * 						2. Perform exponential operation on ASCII value, where the ASCII value will be the base and
 * 							its position will be the exponent. For example, you have `key = Hello` which ASCII value
 * 							of `H = 72` and `position = 1`. So, you will have $72^1$ (72 to the power of 1).
 * 						3. Multiply all exponential results.
 * 						4. Apply modulo `MAX_SIZE` to prevent too big hash result.
 * Examples         :
 * 					Example 1:
 * 					Input: key = a, MAX_SIZE = 1000000
 * 					Output: 97
 *
 */

func generateNumericHash(key string, MAX_SIZE int) int64 {
	if key == "" {
		return 0
	}

	total := int64(1)
	mod := int64(MAX_SIZE)

	for i, c := range key {
		ascii := int64(c)
		powResult := powMod(ascii, i+1, mod)
		total = (total * powResult) % mod
	}

	return total
}

func powMod(base int64, exp int, mod int64) int64 {
	result := int64(1)
	b := base % mod

	for exp > 0 {
		if (exp & 1) == 1 {
			result = (result * b) % mod
		}
		b = (b * b) % mod
		exp = exp >> 1
	}

	return result
}

func RunTestNumericHash() {
	testCases := map[string]struct {
		key      string
		MAX_SIZE int
		expect   int64
	}{
		"case-1": {
			key:      "a",
			MAX_SIZE: 1000000,
			expect:   97,
		},
		"case-2": {
			key:      "Cool!",
			MAX_SIZE: 1000000,
			expect:   773376,
		},
	}

	for name, testCase := range testCases {
		fmt.Printf("RUN %s\n", name)
		result := generateNumericHash(testCase.key, testCase.MAX_SIZE)
		if !cmp.EqualNumbers(result, testCase.expect) {
			fmt.Printf("=== FAILED: expect = %v - got = %v\n", testCase.expect, result)
			os.Exit(1)
		}
		fmt.Printf("=== PASSED\n")
	}

	fmt.Printf("\n✅ All tests passed!\n")
}
