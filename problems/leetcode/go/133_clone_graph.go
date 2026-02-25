package main

import (
	"fmt"

	"github.com/ariefsibuea/algojo/libs/go/format"
)

func init() {
	register("CloneGraph", RunTestCloneGraph)
}

/*
 * Problem          : Clone Graph
 * Topics           : Graph, Depth-First Search, Breadth-First Search, Hash Map
 * Level            : Medium
 * URL              : https://leetcode.com/problems/clone-graph
 * Description      : Given a reference to a node in a connected undirected graph, return a deep copy (clone) of the
 * 					  graph. Each node in the graph contains a value (int) and a list of its neighbors.
 *
 *                    A deep copy means creating entirely new node objects with the same values and neighbor
 * 					  relationships as the original graph. The cloned graph should be completely independent of the
 * 					  original, meaning no nodes in the cloned graph should reference any nodes in the original graph.
 *
 *                    The graph is represented in the test cases using an adjacency list. For simplicity, each node's
 * 					  value is the same as its 1-indexed position (e.g., the first node has val = 1, the second has
 * 					  val = 2, and so on). The given node will always be the first node with val = 1.
 * Constraints      : The number of nodes in the graph is in the range [0, 100].
 *                    1 <= Node.val <= 100.
 *                    Node.val is unique for each node.
 *                    There are no repeated edges and no self-loops in the graph.
 *                    The graph is connected, and all nodes can be visited starting from the given node.
 * Examples         :
 *                    Example 1:
 *                    Input: adjList = [[2,4],[1,3],[2,4],[1,3]]
 *                    Output: [[2,4],[1,3],[2,4],[1,3]]
 *                    Explanation: This input represents a graph with 4 nodes.
 *                    Node 1 (val = 1) has neighbors Node 2 and Node 4.
 *                    Node 2 (val = 2) has neighbors Node 1 and Node 3.
 *                    Node 3 (val = 3) has neighbors Node 2 and Node 4.
 *                    Node 4 (val = 4) has neighbors Node 1 and Node 3.
 *
 *                    Example 2:
 *                    Input: adjList = [[]]
 *                    Output: [[]]
 *                    Explanation: The graph consists of only one node with val = 1, and it has no
 *                    neighbors.
 *
 *                    Example 3:
 *                    Input: adjList = []
 *                    Output: []
 *                    Explanation: This is an empty graph with no nodes.
 */

func cloneGraph(node *GraphNode) *GraphNode {
	if node == nil {
		return nil
	}

	visited := make(map[*GraphNode]*GraphNode)

	var dfs func(current *GraphNode) *GraphNode
	dfs = func(current *GraphNode) *GraphNode {
		if _, hasVisited := visited[current]; hasVisited {
			return visited[current]
		}

		clone := &GraphNode{Val: current.Val}
		neighbors := make([]*GraphNode, len(current.Neighbors))

		visited[current] = clone

		for i, n := range current.Neighbors {
			neighbors[i] = dfs(n)
		}

		clone.Neighbors = neighbors
		return clone
	}

	return dfs(node)
}

func RunTestCloneGraph() {
	// TODO: complete the test function!
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
