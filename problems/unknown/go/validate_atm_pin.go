package main

import (
	"fmt"
	"os"

	"github.com/ariefsibuea/algojo/libs/go/cmp"
)

/**
 * Problem 			: Validate ATM PIN
 * Description      :
 * 					Implement a function to validate an ATM PIN with the following rules:
 * 						- Must be 4 or 6 digits long
 * 						- Must contain digits only
 * 						- Must not contain repeated digits
 * 						- Must not contain ascending or descending sequences (e.g., 1234, 4321)
 * 					Return codes:
 * 						- 0: valid 4-digit PIN
 * 						- 1: valid 6-digit PIN
 * 						- 2: invalid length
 * 						- 3: contains non-digit characters
 * 						- 4: contains duplicate digits
 * 						- 5: contains ascending or descending digits
 * Examples         :
 * 					Example 1:
 * 					Input: atmpin = 1245
 * 					Output: 0
 *
 * 					Example 2:
 * 					Input: atmpin = 1234
 * 					Output: 5
 */

const (
	Valid4DigitPin          = 0
	Valid6DigitPin          = 1
	InvalidLength           = 2
	ContainsNonDigit        = 3
	ContainsDuplicateDigit  = 4
	ContainsSequentialDigit = 5

	AscSeqMode  = 1
	DescSeqMode = -1
	NoSeqMode   = 0
)

func validatePin(atmpin string) int {
	if len(atmpin) != 4 && len(atmpin) != 6 {
		return InvalidLength
	}

	digitHasExist := map[rune]bool{}
	sequentialCount, sequentialMode := 0, NoSeqMode

	for i, c := range atmpin {
		if c < '0' || c > '9' {
			return ContainsNonDigit
		}

		if digitHasExist[c] {
			return ContainsDuplicateDigit
		}
		digitHasExist[c] = true

		if i > 0 {
			switch c {
			case rune(atmpin[i-1] + 1):
				if sequentialMode != AscSeqMode {
					sequentialMode = AscSeqMode
					sequentialCount = 0
				}
				sequentialCount += 1
			case rune(atmpin[i-1] - 1):
				if sequentialMode != DescSeqMode {
					sequentialMode = DescSeqMode
					sequentialCount = 0
				}
				sequentialCount += 1
			default:
				sequentialMode = NoSeqMode
				sequentialCount = 0
			}
		}
	}

	if sequentialMode != NoSeqMode && sequentialCount == len(atmpin)-1 {
		return ContainsSequentialDigit
	}

	if len(atmpin) == 4 {
		return Valid4DigitPin
	}
	return Valid6DigitPin
}

func RunTestValidatePin() {
	testCases := map[string]struct {
		atmpin string
		expect int
	}{
		"valid-4-digit": {
			atmpin: "1245",
			expect: 0,
		},
		"valid-6-digit": {
			atmpin: "124578",
			expect: 1,
		},
		"invalid-length": {
			atmpin: "124",
			expect: 2,
		},
		"contains-non-digit": {
			atmpin: "124a",
			expect: 3,
		},
		"contains-duplicate-digit": {
			atmpin: "1241",
			expect: 4,
		},
		"contains-sequential-digit": {
			atmpin: "1234",
			expect: 5,
		},
	}

	for name, testCase := range testCases {
		fmt.Printf("RUN %s\n", name)
		result := validatePin(testCase.atmpin)
		if !cmp.EqualNumbers(result, testCase.expect) {
			fmt.Printf("=== FAILED: expect = %v - got = %v\n", testCase.expect, result)
			os.Exit(1)
		}
		fmt.Printf("=== PASSED\n")
	}

	fmt.Printf("\n✅ All tests passed!\n")
}
