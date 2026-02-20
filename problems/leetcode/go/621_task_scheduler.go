package main

import (
	"container/heap"
	"fmt"
	"sort"

	"github.com/ariefsibuea/algojo/libs/go/cmp"
	"github.com/ariefsibuea/algojo/libs/go/format"
	"github.com/ariefsibuea/algojo/libs/go/runner"
)

/*
 * Problem          : Task Scheduler
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

type PriorityCPUTasks []int

func (p PriorityCPUTasks) Len() int           { return len(p) }
func (p PriorityCPUTasks) Less(i, j int) bool { return p[i] > p[j] }
func (p PriorityCPUTasks) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func (p *PriorityCPUTasks) Push(x interface{}) {
	*p = append(*p, x.(int))
}

func (p *PriorityCPUTasks) Pop() interface{} {
	old := *p
	n := len(old)
	x := old[n-1]
	old[n-1] = 0
	*p = old[0 : n-1]
	return x
}

func leastInterval_PriorityQueue(tasks []byte, n int) int {
	taskFreq := make([]int, 26)
	for _, task := range tasks {
		taskFreq[task-'A']++
	}

	priorityTasks := new(PriorityCPUTasks)
	heap.Init(priorityTasks)

	for i := 0; i < 26; i++ {
		if taskFreq[i] > 0 {
			heap.Push(priorityTasks, taskFreq[i])
		}
	}

	var numInterval = 0

	for priorityTasks.Len() > 0 {
		remainingTasks := make([]int, 0, priorityTasks.Len())
		cycle := 0
		taskCount := 0

		for cycle < n+1 && priorityTasks.Len() > 0 {
			task := heap.Pop(priorityTasks).(int)
			remainingTasks = append(remainingTasks, task-1)

			taskCount++
			cycle++
		}

		for _, task := range remainingTasks {
			if task > 0 {
				heap.Push(priorityTasks, task)
			}
		}

		if priorityTasks.Len() > 0 {
			numInterval += n + 1
		} else {
			numInterval += taskCount
		}
	}

	return numInterval
}

func leastInterval_FillTheSlotAndSort(tasks []byte, n int) int {
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

		result := runner.ExecCountMetrics(leastInterval_PriorityQueue, testCase.tasks, testCase.n).(int)
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
