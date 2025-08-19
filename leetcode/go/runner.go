package main

import (
	"fmt"
	"strings"
)

type TestFunc func()

type SolutionRunner struct {
	tests map[string]TestFunc
}

func NewSolutionRunner() SolutionRunner {
	runner := SolutionRunner{
		tests: make(map[string]TestFunc),
	}

	runner.registerSolution()

	return runner
}

func (r SolutionRunner) Run(solutionName string) error {
	testFunc, exists := r.tests[solutionName]
	if !exists {
		return fmt.Errorf("solution '%s' does not exist", solutionName)
	}

	fmt.Printf("Running solution: %s\n", solutionName)
	fmt.Println(strings.Repeat("=", 50))
	testFunc()
	fmt.Println()

	return nil
}

func (r SolutionRunner) List() {
	fmt.Println("Available solutions:")
	for name := range r.tests {
		fmt.Printf("  - %s\n", name)
	}
}

func (r *SolutionRunner) registerSolution() {
	r.tests["TwoSum"] = RunTestTwoSum
	r.tests["LongestSubstringWithoutRepeatingCharacters"] = RunTestLongestSubstringWithoutRepeatingCharacters
	r.tests["MedianOfTwoSortedArrays"] = RunTestFindMedianSortedArrays
	r.tests["MaxArea"] = RunTestMaxArea
	r.tests["RemoveNthNodeFromEndOfList"] = RunTestRemoveNthFromEnd
	r.tests["IsValidParentheses"] = RunTestIsValidParentheses
	r.tests["MergeTwoSortedLists"] = RunTestMergeTwoSortedLists
	r.tests["RemoveDuplicatesFromSortedArray"] = RunTestRemoveDuplicatesFromSortedArray
	r.tests["SearchInRotatedSortedArray"] = RunTestSearchInRotatedSortedArray
	r.tests["BestTimeToBuyAndSellStock"] = RunTestMaxProfit
	r.tests["ValidPalindrome"] = RunTestIsValidPalindrome
	r.tests["InvertBinaryTree"] = RunTestInvertTree
	r.tests["ValidAnagram"] = RunTestIsValidAnagram
	r.tests["BinarySearch"] = RunTestBinarySearch
	r.tests["FloodFill"] = RunTestFloodFill
}
