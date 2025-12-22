package main

import (
	"fmt"

	"github.com/ariefsibuea/algojo/libs/go/format"
)

/*
 * Problem 			: Maximum Profit from Trading Stocks with Discounts
 * Topics           : Array, Dynamic Programming, Tree, Depth-First Search
 * Level            : Hard
 * URL              : https://leetcode.com/problems/maximum-profit-from-trading-stocks-with-discounts
 * Description      : <Description>
 * Examples         : <Examples>
 */

func maxProfitfromTradingStocks(n int, present []int, future []int, hierarchy [][]int, budget int) int {
	type result struct {
		profit        []int
		profitBuyDisc []int
		cost          int
	}

	var dfs func(emp int) result
	var group = make([][]int, n)

	for i := range group {
		group[i] = make([]int, 0)
	}
	for _, h := range hierarchy {
		// Since the employee number is 1-based number, we need to
		// subtract each number by 1 to adjust to 0-based array.
		group[h[0]-1] = append(group[h[0]-1], h[1]-1)
	}

	dfs = func(emp int) result {
		cost := present[emp]
		discCost := present[emp] / 2

		profit := make([]int, budget+1)
		profitBuyDisc := make([]int, budget+1)

		subProfit := make([]int, budget+1)
		subProfitBuyDisc := make([]int, budget+1)

		empBuyCost := cost
		for _, subEmp := range group[emp] {
			subEmpResult := dfs(subEmp)
			empBuyCost += subEmpResult.cost

			for i := budget; i >= 0; i-- {
				for sub := 0; sub <= min(subEmpResult.cost, i); sub++ {
					if i-sub >= 0 {
						subProfit[i] = max(subProfit[i], subProfit[i-sub]+subEmpResult.profit[sub])

						subProfitBuyDisc[i] = max(
							subProfitBuyDisc[i],
							subProfitBuyDisc[i-sub]+subEmpResult.profitBuyDisc[sub],
						)
					}
				}
			}
		}

		for i := 0; i <= budget; i++ {
			profit[i] = subProfit[i]
			profitBuyDisc[i] = subProfit[i]

			if i >= discCost {
				profitBuyDisc[i] = max(subProfit[i], subProfitBuyDisc[i-discCost]+future[emp]-discCost)
			}
			if i >= cost {
				profit[i] = max(subProfit[i], subProfitBuyDisc[i-cost]+future[emp]-cost)
			}
		}

		return result{profit, profitBuyDisc, empBuyCost}
	}

	return dfs(0).profit[budget]
}

func RunTestMaximumProfitFromTradingStocksWithDiscounts() {
	testCases := map[string]struct{}{}

	for name, testCase := range testCases {
		fmt.Printf("RUN %s\n", name)

		fmt.Println(testCase)
		// result := twoSum(testCase.nums, testCase.target)
		// format.PrintInput(map[string]interface{}{"input-1": testCase.Input1})

		// if !EqualSlices(result, testCase.expect) {
		//  format.PrintFailed("expect = %v - got = %v\n", testCase.expect, result)
		// 	os.Exit(1)
		// }
		format.PrintSuccess("test case '%s' passed", name)
	}

	fmt.Printf("\n✅ All tests passed!\n")
}
