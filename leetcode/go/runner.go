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
	r.tests["3Sum"] = RunTestThreeSum
	r.tests["BalancedBinaryTree"] = RunTestBinarySearch
	r.tests["BestTimeToBuyAndSellStock"] = RunTestMaxProfit
	r.tests["BinarySearch"] = RunTestBinarySearch
	r.tests["BtreeInorderTraversal"] = RunTestBtreeInorderTraversal
	r.tests["ClimbingStairs"] = RunTestClimbStairs
	r.tests["ContainsDuplicate"] = RunTestContainsDuplicate
	r.tests["CourseSchedule"] = RunTestCourseSchedule
	r.tests["DesignTwitter"] = RunTestDesignTwitter
	r.tests["FirstBadVersion"] = RunTestFirstBadVersion
	r.tests["FloodFill"] = RunTestFloodFill
	r.tests["GroupAnagrams"] = RunTestGroupAnagrams
	r.tests["ImplementQueueUsingStacks"] = RunTestImplementQueueUsingStacks
	r.tests["InvertBinaryTree"] = RunTestInvertTree
	r.tests["IsValidParentheses"] = RunTestIsValidParentheses
	r.tests["LinkedListCycle"] = RunTestHasCycle
	r.tests["LongestSubstringWithoutRepeatingCharacters"] = RunTestLongestSubstringWithoutRepeatingCharacters
	r.tests["LowestCommonAncestor"] = RunTestLowestCommonAncestor
	r.tests["MaxArea"] = RunTestMaxArea
	r.tests["MedianOfTwoSortedArrays"] = RunTestFindMedianSortedArrays
	r.tests["MergeTwoSortedLists"] = RunTestMergeTwoSortedLists
	r.tests["MinimumWindowSubstring"] = RunTestMinWindowSubstring
	r.tests["RansomNote"] = RunTestCanConstructRansomNote
	r.tests["RemoveDuplicatesFromSortedArray"] = RunTestRemoveDuplicatesFromSortedArray
	r.tests["RemoveNthNodeFromEndOfList"] = RunTestRemoveNthFromEnd
	r.tests["SearchInRotatedSortedArray"] = RunTestSearchInRotatedSortedArray
	r.tests["TopKFrequentElements"] = RunTestTopKFrequentElements
	r.tests["TwoSum"] = RunTestTwoSum
	r.tests["TwoSumII"] = RunTestTwoSumII
	r.tests["ValidAnagram"] = RunTestIsValidAnagram
	r.tests["ValidPalindrome"] = RunTestIsValidPalindrome
	r.tests["ValidateBinarySearchTree"] = RunTestIsValidBST
	r.tests["ValidTriangleNumber"] = RunTestValidTriangleNumber
}
