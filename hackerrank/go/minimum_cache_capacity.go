package main

import (
	"container/list"
	"fmt"
	"os"
)

/**
 * Problem 			: Minimum Cache Capacity
 * Level            : Medium
 * URL              :
 * Description      : Given a list of n items accessed in a sequence from a system, and a cache with Least Recently
 * 					Used (LRU) replacement policy, the goal is to determine the minimum cache size required to ensure
 * 					at least k requests hit the cache. If it is not possible to achieve k hits, return -1.
 *
 * 					Assume each item in the list has a size of 1 unit.
 *
 * 					Note: A cache following Least Recently Used (LRU) replacement policy is of fixed size and when it
 * 					becomes full we will delete the value which is least recently used and insert a new value.
 * Examples         :
 * 					n = 5
 * 					requests = ["item1", "item1", "item3", "item1", "item3"]
 * 					k = 1
 */

func minCacheCapacity(requests []string, k int) int {
	if k <= 0 {
		return 1
	}

	n := len(requests)

	low, high := 1, n
	result := -1

	for low <= high {
		mid := (low + high) / 2
		if hitsWithCapacity(requests, mid) >= k {
			result = mid
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return result
}

func hitsWithCapacity(requests []string, cap int) int {
	lru := list.New()
	elems := make(map[string]*list.Element)
	hitCount := 0

	for _, req := range requests {
		if elem, found := elems[req]; found {
			hitCount += 1
			lru.MoveToFront(elem)
		} else {
			if lru.Len() == cap {
				// remove back/tail element => least recently used
				backElem := lru.Back()
				if backElem != nil {
					oldElem := backElem.Value.(string)
					delete(elems, oldElem)
					lru.Remove(backElem)
				}
			}

			newElem := lru.PushFront(req)
			elems[req] = newElem
		}
	}

	return hitCount
}

func RunTestMinCacheCapacity() {
	testCases := map[string]struct {
		requests []string
		k        int
		expect   int
	}{
		"case-1": {
			requests: []string{"item1", "item1", "item3", "item1", "item3"},
			k:        1,
			expect:   1,
		},
	}

	for name, testCase := range testCases {
		fmt.Printf("RUN %s\n", name)
		result := minCacheCapacity(testCase.requests, testCase.k)
		if !EqualNumbers(result, testCase.expect) {
			fmt.Printf("=== FAILED: expect = %v - got = %v\n", testCase.expect, result)
			os.Exit(1)
		}
		fmt.Printf("=== PASSED\n")
	}

	fmt.Printf("\n✅ All tests passed!\n")
}
