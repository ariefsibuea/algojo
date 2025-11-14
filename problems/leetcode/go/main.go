package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/ariefsibuea/algojo/libs/go/runner"
)

func main() {
	var solution = flag.String("solution", "", "Name of the solution to run (e.g., TwoSum)")
	flag.StringVar(solution, "s", *solution, "Alias for -solution")

	var list = flag.Bool("list", false, "List all available solutions")
	flag.BoolVar(list, "l", *list, "Alias for -list")

	flag.Parse()

	r := runner.NewSolutionRunner()
	registerSolutions(&r)

	if *list {
		r.List()
		return
	}

	if *solution == "" {
		fmt.Println("Please specify a solution to run with -solution flag")
		fmt.Println("Example: go run . -solution TwoSum")
		fmt.Println("Use -list to see available solutions")
		os.Exit(1)
	}

	if err := r.Run(*solution); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}

func registerSolutions(r *runner.SolutionRunner) {
	solutions := map[string]runner.TestFunc{
		"3Sum":          RunTestThreeSum,
		"AddTwoNumbers": RunTestAddTwoNumbers,
		"AppendCharactersToStringToMakeSubsequence":     RunTestAppendCharactersToString,
		"BalancedBinaryTree":                            RunTestBinarySearch,
		"BestTimeToBuyAndSellStock":                     RunTestMaxProfit,
		"BinarySearch":                                  RunTestBinarySearch,
		"BtreeInorderTraversal":                         RunTestBtreeInorderTraversal,
		"CircularArrayLoop":                             RunTestCircularArrayLoop,
		"ClimbingStairs":                                RunTestClimbStairs,
		"ContainsDuplicate":                             RunTestContainsDuplicate,
		"CourseSchedule":                                RunTestCourseSchedule,
		"DesignTwitter":                                 RunTestDesignTwitter,
		"FindTheDuplicateNumber":                        RunTestFindTheDuplicateNumber,
		"FirstBadVersion":                               RunTestFirstBadVersion,
		"FirstUniqueCharacterInAString":                 RunTestFirstUniqChar,
		"FloodFill":                                     RunTestFloodFill,
		"FruitIntoBasket":                               RunTestTotalFruit,
		"GroupAnagrams":                                 RunTestGroupAnagrams,
		"HappyNumber":                                   RunTestHappyNumber,
		"ImplementQueueUsingStacks":                     RunTestImplementQueueUsingStacks,
		"InsertInterval":                                RunTestInsertInterval,
		"InvertBinaryTree":                              RunTestInvertTree,
		"IsPalindrome":                                  RunTestIsPalindrome,
		"IsValidParentheses":                            RunTestIsValidParentheses,
		"LinkedListCycle":                               RunTestHasCycle,
		"LongestIncreasingSubsequence":                  RunTestLongestIncreasingSubsequence,
		"LongestRepeatingCharacterReplacement":          RunTestCharacterReplacement,
		"LongestSubstringWithAtMostKDistinctCharacters": RunTestLengthOfLongestSubstringKDistinct,
		"LongestSubstringWithoutRepeatingCharacters":    RunTestLongestSubstringWithoutRepeatingCharacters,
		"LowestCommonAncestor":                          RunTestLowestCommonAncestor,
		"LowestCommonAncestorIII":                       RunTestLowestCommonAncestorIII,
		"MaxArea":                                       RunTestMaxArea,
		"MaximumSubarray":                               RunTestMaxSubArray,
		"MedianOfTwoSortedArrays":                       RunTestFindMedianSortedArrays,
		"MeetingRooms":                                  RunTestMeetingRooms,
		"MergeIntervals":                                RunTestMergeIntervals,
		"MergeTwoSortedLists":                           RunTestMergeTwoSortedLists,
		"MiddleOfTheLinkedList":                         RunTestMiddleOfTheLinkedList,
		"MinimumWindowSubstring":                        RunTestMinWindowSubstring,
		"MoveZeroes":                                    RunTestMoveZeroes,
		"NumberOfZeroFilledSubarrays":                   RunTestZeroFilledSubarray,
		"PalindromeLinkedList":                          RunTestPalindromeLinkedList,
		"RansomNote":                                    RunTestCanConstructRansomNote,
		"RemoveDuplicatesFromSortedArray":               RunTestRemoveDuplicatesFromSortedArray,
		"RemoveNthNodeFromEndOfList":                    RunTestRemoveNthFromEnd,
		"ReverseString":                                 RunTestReverseString,
		"ReverseWordsInAString":                         RunTestReverseWordsInAString,
		"SearchInRotatedSortedArray":                    RunTestSearchInRotatedSortedArray,
		"SquaresOfASortedArray":                         RunTestSquaresOfASortedArray,
		"StrobogrammaticNumber":                         RunTestStrobogrammaticNumber,
		"TopKFrequentElements":                          RunTestTopKFrequentElements,
		"TrappingRainWater":                             RunTestTrappingRainWater,
		"TwoSum":                                        RunTestTwoSum,
		"TwoSumII":                                      RunTestTwoSumII,
		"TwoSumLessThanK":                               RunTestTwoSumLessThanK,
		"ValidAnagram":                                  RunTestIsValidAnagram,
		"ValidPalindrome":                               RunTestIsValidPalindrome,
		"ValidPalindromeII":                             RunTestValidPalindromeII,
		"ValidTriangleNumber":                           RunTestValidTriangleNumber,
		"ValidWordAbbreviation":                         RunTestValidWordAbbreviation,
		"ValidateBinarySearchTree":                      RunTestIsValidBST,
		"WordSearch":                                    RunTestWordSearch,
	}
	r.RegisterSolutions(solutions)
}
