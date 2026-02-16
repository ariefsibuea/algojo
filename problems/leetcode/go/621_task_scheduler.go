package main

import (
	"fmt"
	"sort"

	"github.com/ariefsibuea/algojo/libs/go/cmp"
	"github.com/ariefsibuea/algojo/libs/go/format"
	"github.com/ariefsibuea/algojo/libs/go/runner"
)

/*
 * Problem          : 621. Task Scheduler
 * Topics           : Array, Hash Table, Greedy, Sorting, Priority Queue (Heap), Counting
 * Level            : Medium
 * URL              : https://leetcode.com/problems/task-scheduler/
 * Description      : Schedule CPU tasks (labeled A-Z) with a cooldown constraint. Given a list of tasks and a cooldown
 *                    period n, find the minimum time needed to execute all tasks. Each time unit can either process one
 *                    task or be idle. The same task cannot be executed again until n time units have passed since its
 *                    last execution. Tasks can be performed in any order, and idle periods may be necessary to satisfy
 *                    the cooldown constraint. Return the total number of time units required.
 * Constraints      : 1 <= tasks.length <= 10^4
 *                    tasks[i] is an uppercase English letter.
 *                    0 <= n <= 100
 * Examples         : Example 1:
 *                    Input: tasks = ["A","A","A","B","B","B"], n = 2
 *                    Output: 8
 *                    Explanation: A possible sequence is: A -> B -> idle -> A -> B -> idle -> A -> B.
 *                    After completing task A, you must wait two intervals before doing A again.
 *                    The same applies to task B. In the 3rd interval, neither A nor B can be done, so you idle. By the
 * 					  4th interval, you can do A again as 2 intervals have passed.
 *
 *                    Example 2:
 *                    Input: tasks = ["A","C","A","B","D","B"], n = 1
 *                    Output: 6
 *                    Explanation: A possible sequence is: A -> B -> C -> D -> A -> B.
 *                    With a cooling interval of 1, you can repeat a task after just one other task.
 *
 *                    Example 3:
 *                    Input: tasks = ["A","A","A","B","B","B"], n = 3
 *                    Output: 10
 *                    Explanation: A possible sequence is: A -> B -> idle -> idle -> A -> B -> idle -> idle -> A -> B.
 *                    There are only two types of tasks, A and B, which need to be separated by 3 intervals. This leads
 * 					  to idling twice between repetitions of these tasks.
 */

func leastInterval(tasks []byte, n int) int {
	taskFreq := make([]int, 26)
	for _, task := range tasks {
		taskFreq[task-'A']++
	}

	sort.Slice(taskFreq, func(i, j int) bool {
		return taskFreq[i] > taskFreq[j]
	})

	gap := taskFreq[0] - 1
	idle := n * gap

	for i := 1; i < 26; i++ {
		idle = idle - min(gap, taskFreq[i])
	}

	idle = max(0, idle)
	return len(tasks) + idle
}

func RunTestTaskScheduler() {
	runner.InitMetrics("TaskScheduler")

	testCases := map[string]struct {
		tasks  []byte
		n      int
		expect int
	}{
		"example-1-basic-with-idle": {
			tasks:  []byte{'A', 'A', 'A', 'B', 'B', 'B'},
			n:      2,
			expect: 8,
		},
		"example-2-no-idle-needed": {
			tasks:  []byte{'A', 'C', 'A', 'B', 'D', 'B'},
			n:      1,
			expect: 6,
		},
		"example-3-more-idle-required": {
			tasks:  []byte{'A', 'A', 'A', 'B', 'B', 'B'},
			n:      3,
			expect: 10,
		},
		"single-task": {
			tasks:  []byte{'A'},
			n:      2,
			expect: 1,
		},
		"single-task-no-cooldown": {
			tasks:  []byte{'A'},
			n:      0,
			expect: 1,
		},
		"no-cooldown-multiple-tasks": {
			tasks:  []byte{'A', 'A', 'A', 'B', 'B', 'B'},
			n:      0,
			expect: 6,
		},
		"all-same-tasks-large-cooldown": {
			tasks:  []byte{'A', 'A', 'A', 'A', 'A'},
			n:      2,
			expect: 13,
		},
		"many-different-tasks": {
			tasks:  []byte{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J'},
			n:      2,
			expect: 10,
		},
		"edge-case-zero-cooldown": {
			tasks:  []byte{'A', 'A', 'A', 'A', 'A'},
			n:      0,
			expect: 5,
		},
		"edge-case-two-tasks-only": {
			tasks:  []byte{'A', 'B'},
			n:      5,
			expect: 2,
		},
	}

	var passedCount uint16 = 0

	for name, testCase := range testCases {
		fmt.Printf("RUN %s\n", name)

		result := runner.ExecCountMetrics(leastInterval, testCase.tasks, testCase.n).(int)
		format.PrintInput(map[string]interface{}{"tasks": string(testCase.tasks), "n": testCase.n})

		if !cmp.EqualNumbers(result, testCase.expect) {
			format.PrintFailed("expect = %v - got = %v", testCase.expect, result)
			continue
		}

		format.PrintSuccess("test case '%s' passed", name)
		passedCount++
	}

	fmt.Printf("\n📊 Test Summary: %d/%d passed\n", passedCount, len(testCases))
	runner.PrintMetrics()
}
