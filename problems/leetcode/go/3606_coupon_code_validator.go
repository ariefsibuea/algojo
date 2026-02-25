package main

import (
	"fmt"
	"os"
	"sort"
	"unicode"

	"github.com/ariefsibuea/algojo/libs/go/cmp"
	"github.com/ariefsibuea/algojo/libs/go/format"
)

func init() {
	register("CouponCodeValidator", RunTestCouponCodeValidator)
}

/*
 * Problem 			: Coupon Code Validator
 * Topics           : Array, Hash Table, String, Sorting
 * Level            : Easy
 * URL              : https://leetcode.com/problems/coupon-code-validator
 * Description      : <Description>
 * Examples         : <Examples>
 */

const (
	CategoryElectronics = "electronics"
	CategoryGrocery     = "grocery"
	CategoryPharmacy    = "pharmacy"
	CategoryRestaurant  = "restaurant"

	CouponPattern = "^[a-zA-Z0-9_]+$"
)

var ValidCategory = map[string]int{
	CategoryElectronics: 0,
	CategoryGrocery:     1,
	CategoryPharmacy:    2,
	CategoryRestaurant:  3,
}

type Coupon struct {
	Code         string
	BusinessLine string
}

func validateCoupons(code []string, businessLine []string, isActive []bool) []string {
	var validCoupons = make([]Coupon, 0, len(code))

	for i := range code {
		_, validCategory := ValidCategory[businessLine[i]]
		if validCategory && isValidCoupon(code[i]) && isActive[i] {
			validCoupons = append(validCoupons, Coupon{
				Code:         code[i],
				BusinessLine: businessLine[i],
			})
		}
	}

	sort.Slice(validCoupons, func(i, j int) bool {
		if validCoupons[i].BusinessLine != validCoupons[j].BusinessLine {
			return ValidCategory[validCoupons[i].BusinessLine] < ValidCategory[validCoupons[j].BusinessLine]
		}
		return validCoupons[i].Code < validCoupons[j].Code
	})

	result := make([]string, len(validCoupons))
	for i := range validCoupons {
		result[i] = validCoupons[i].Code
	}

	return result
}

func isValidCoupon(code string) bool {
	for _, c := range code {
		if !unicode.IsLetter(c) && !unicode.IsNumber(c) && c != '_' {
			return false
		}
	}
	return code != ""
}

func RunTestCouponCodeValidator() {
	testCases := map[string]struct {
		code         []string
		businessLine []string
		isActive     []bool
		expect       []string
	}{
		"case-1": {
			code:         []string{"SAVE20", "", "PHARMA5", "SAVE@20"},
			businessLine: []string{"restaurant", "grocery", "pharmacy", "restaurant"},
			isActive:     []bool{true, true, true, true},
			expect:       []string{"PHARMA5", "SAVE20"},
		},
		"case-2": {
			code:         []string{"GROCERY15", "ELECTRONICS_50", "DISCOUNT10"},
			businessLine: []string{"grocery", "electronics", "invalid"},
			isActive:     []bool{false, true, true},
			expect:       []string{"ELECTRONICS_50"},
		},
		"case-3": {
			code:         []string{"TsCwKhY", "qCeVkfb", "ZGjX07H"},
			businessLine: []string{"restaurant", "electronics", "pharmacy"},
			isActive:     []bool{true, true, true},
			expect:       []string{"qCeVkfb", "ZGjX07H", "TsCwKhY"},
		},
	}

	for name, testCase := range testCases {
		fmt.Printf("RUN %s\n", name)

		result := validateCoupons(testCase.code, testCase.businessLine, testCase.isActive)
		format.PrintInput(map[string]interface{}{
			"code":         testCase.code,
			"businessLine": testCase.businessLine,
			"isActive":     testCase.isActive},
		)

		if !cmp.EqualSlices(result, testCase.expect) {
			format.PrintFailed("expect = %v - got = %v\n", testCase.expect, result)
			os.Exit(1)
		}
		format.PrintSuccess("test case '%s' passed", name)
	}

	fmt.Printf("\n✅ All tests passed!\n")
}
