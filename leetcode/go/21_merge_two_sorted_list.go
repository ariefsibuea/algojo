package main

import (
	"fmt"
	"os"
)

type P21ListNode struct {
	Val  int
	Next *P21ListNode
}

func mergeTwoLists(list1 *P21ListNode, list2 *P21ListNode) *P21ListNode {
	dummy := &P21ListNode{Val: 0}
	current := dummy

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			current.Next = list1
			list1 = list1.Next
		} else {
			current.Next = list2
			list2 = list2.Next
		}
		current = current.Next
	}

	if list1 != nil {
		current.Next = list1
	} else {
		current.Next = list2
	}

	return dummy.Next
}

func RunTestMergeTwoSortedLists() {
	list1Case1 := &P21ListNode{
		Val: 1,
		Next: &P21ListNode{
			Val: 2,
			Next: &P21ListNode{
				Val: 4,
			},
		},
	}
	list2Case1 := &P21ListNode{
		Val: 1,
		Next: &P21ListNode{
			Val: 3,
			Next: &P21ListNode{
				Val: 4,
			},
		},
	}

	list2Case3 := &P21ListNode{
		Val: 0,
	}

	testCases := map[string]struct {
		list1  *P21ListNode
		list2  *P21ListNode
		expect []int
	}{
		"case-1": {
			list1:  list1Case1,
			list2:  list2Case1,
			expect: []int{1, 1, 2, 3, 4, 4},
		},
		"case-2": {
			list1:  nil,
			list2:  nil,
			expect: []int{},
		},
		"case-3": {
			list1:  nil,
			list2:  list2Case3,
			expect: []int{0},
		},
	}

	for name, testCase := range testCases {
		fmt.Printf("RUN %s\n", name)

		head := mergeTwoLists(testCase.list1, testCase.list2)
		result := []int{}
		for head != nil {
			result = append(result, head.Val)
			head = head.Next
		}

		if !EqualSlices(result, testCase.expect) {
			fmt.Printf("=== FAILED: expect = %v - got = %v\n", testCase.expect, result)
			os.Exit(1)
		}
		fmt.Printf("=== PASSED\n")
	}

	fmt.Printf("\n✅ All tests passed!\n")

}
