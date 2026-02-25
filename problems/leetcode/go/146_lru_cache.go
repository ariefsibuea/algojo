package main

import (
	"fmt"

	"github.com/ariefsibuea/algojo/libs/go/format"
)

func init() {
	register("LRUCache", RunTestLRUCache)
}

/*
 * Problem 			: LRU Cache
 * Topics           : Hash Table, Linked List, Design, Doubly-Linked List
 * Level            : Medium
 * URL              : https://leetcode.com/problems/lru-cache
 * Description      : <Description>
 * Examples         : <Examples>
 */

type LRUCache struct {
	capacity int
	size     int
	cache    map[int]*DLLNode
	head     *DLLNode
	tail     *DLLNode
}

func LRUCacheConstructor(capacity int) LRUCache {
	headNode := new(DLLNode)
	tailNode := new(DLLNode)
	headNode.Next = tailNode
	tailNode.Prev = headNode

	return LRUCache{
		capacity: capacity,
		size:     0,
		cache:    map[int]*DLLNode{},
		head:     headNode,
		tail:     tailNode,
	}
}

func (this *LRUCache) Get(key int) int {
	n, ok := this.cache[key]
	if !ok {
		return -1
	}

	this.removeNode(n)
	this.addToHead(n)

	return n.Val
}

func (this *LRUCache) Put(key int, value int) {
	n, ok := this.cache[key]
	if ok {
		this.removeNode(n)
		n.Val = value
		this.addToHead(n)
	} else {
		newNode := &DLLNode{Key: key, Val: value}
		this.cache[newNode.Key] = newNode
		this.addToHead(newNode)
		this.size += 1

		if this.size > this.capacity {
			remove := this.tail.Prev
			delete(this.cache, remove.Key)
			this.removeNode(remove)
			this.size -= 1
		}
	}
}

func (this *LRUCache) removeNode(n *DLLNode) {
	n.Prev.Next = n.Next
	n.Next.Prev = n.Prev
}

func (this *LRUCache) addToHead(n *DLLNode) {
	n.Next = this.head.Next
	n.Prev = this.head
	this.head.Next = n
	n.Next.Prev = n
}

func RunTestLRUCache() {
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
