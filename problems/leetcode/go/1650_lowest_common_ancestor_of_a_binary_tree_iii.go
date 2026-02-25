package main

import (
	"fmt"
	"os"

	"github.com/ariefsibuea/algojo/libs/go/cmp"
)

func init() {
	register("LowestCommonAncestorIII", RunTestLowestCommonAncestorIII)
}

/**
 * Problem 			: Lowest Common Ancestor of a Binary Tree III
 * Topics           : Tree, Binary Tree
 * Level            : Medium
 * URL              : https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree-iii
 * Description      :
 * Examples         :
 * 					Example 1:
 * 					Input: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1
 * 					Output: 3
 * 					Explanation: The LCA of nodes 5 and 1 is 3.
 *
 * 					Example 2:
 * 					Input: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 4
 * 					Output = 5
 * 					Explanation: The LCA of nodes 5 and 4 is 5 since a node can be a descendant of itself according to
 * 					the LCA definition.
 */

type Node struct {
	val    int
	left   *Node
	right  *Node
	parent *Node
}

func lowestCommonAncestorIII(p, q *Node) *Node {
	ptr1, ptr2 := p, q

	for ptr1 != ptr2 {
		if ptr1.parent != nil {
			ptr1 = ptr1.parent
		} else {
			ptr1 = q
		}

		if ptr2.parent != nil {
			ptr2 = ptr2.parent
		} else {
			ptr2 = p
		}
	}

	return ptr1
}

func RunTestLowestCommonAncestorIII() {
	constructTreeP1650()

	testCases := map[string]struct {
		p      *Node
		q      *Node
		expect int
	}{
		"case-1": {
			p:      case1P1650()[0],
			q:      case1P1650()[1],
			expect: 3,
		},
		"case-2": {
			p:      case2P1650()[0],
			q:      case2P1650()[1],
			expect: 5,
		},
	}

	for name, testCase := range testCases {
		fmt.Printf("RUN %s\n", name)
		result := lowestCommonAncestorIII(testCase.p, testCase.q)
		if !cmp.EqualNumbers(result.val, testCase.expect) {
			fmt.Printf("=== FAILED: expect = %v - got = %v\n", testCase.expect, result)
			os.Exit(1)
		}
		fmt.Printf("=== PASSED\n")
	}

	fmt.Printf("\n✅ All tests passed!\n")
}

var (
	root  = &Node{val: 3}
	node2 = &Node{val: 5}
	node3 = &Node{val: 1}
	node4 = &Node{val: 6}
	node5 = &Node{val: 2}
	node6 = &Node{val: 0}
	node7 = &Node{val: 8}
	node8 = &Node{val: 7}
	node9 = &Node{val: 4}
)

func constructTreeP1650() {
	root.left = node2
	root.right = node3

	node2.left = node4
	node2.right = node5
	node2.parent = root

	node3.left = node6
	node3.right = node7
	node3.parent = root

	node4.parent = node2

	node5.left = node8
	node5.right = node9
	node5.parent = node2

	node6.parent = node3
	node7.parent = node3
	node8.parent = node5
	node9.parent = node5
}

func case1P1650() []*Node {
	return []*Node{
		node2,
		node3,
	}
}

func case2P1650() []*Node {
	return []*Node{
		node2,
		node9,
	}
}
