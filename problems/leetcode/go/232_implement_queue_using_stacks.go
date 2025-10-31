package main

import (
	"fmt"
	"os"

	"github.com/ariefsibuea/algojo/libs/go/cmp"
)

type MyQueue struct {
	stackPush []int
	stackPop  []int
}

func Constructor() MyQueue {
	return MyQueue{}
}

func (this *MyQueue) Push(x int) {
	if len(this.stackPush) == 0 {
		for i := len(this.stackPop) - 1; i >= 0; i-- {
			this.stackPush = append(this.stackPush, this.stackPop[i])
		}
		this.stackPop = this.stackPop[:0]
	}
	this.stackPush = append(this.stackPush, x)
}

func (this *MyQueue) Pop() int {
	if len(this.stackPop) == 0 {
		for i := len(this.stackPush) - 1; i >= 0; i-- {
			this.stackPop = append(this.stackPop, this.stackPush[i])
		}
		this.stackPush = this.stackPush[:0]
	}

	n := len(this.stackPop) - 1
	result := this.stackPop[n]

	this.stackPop[n] = 0
	this.stackPop = this.stackPop[:n]

	return result
}

func (this *MyQueue) Peek() int {
	if len(this.stackPop) == 0 {
		for i := len(this.stackPush) - 1; i >= 0; i-- {
			this.stackPop = append(this.stackPop, this.stackPush[i])
		}
		this.stackPush = this.stackPush[:0]
	}

	n := len(this.stackPop)
	if n == 0 {
		return 0
	}
	return this.stackPop[n-1]
}

func (this *MyQueue) Empty() bool {
	return len(this.stackPush) == 0 && len(this.stackPop) == 0
}

func RunTestImplementQueueUsingStacks() {
	var expect []int
	var obj = Constructor()

	obj.Push(1)
	expect = append(expect, 1)
	if !cmp.EqualSlices(expect, obj.stackPush) {
		fmt.Printf("=== FAILED: expect = %v - got = %v\n", expect, obj.stackPush)
		os.Exit(1)
	}

	obj.Push(2)
	expect = append(expect, 2)
	if !cmp.EqualSlices(expect, obj.stackPush) {
		fmt.Printf("=== FAILED: expect = %v - got = %v\n", expect, obj.stackPush)
		os.Exit(1)
	}

	result := obj.Peek()
	expect = nil
	expect = []int{2, 1}
	if !cmp.EqualNumbers(1, result) {
		fmt.Printf("=== FAILED: expect = %v - got = %v\n", 1, result)
		os.Exit(1)
	}
	if !cmp.EqualSlices(expect, obj.stackPop) {
		fmt.Printf("=== FAILED: expect = %v - got = %v\n", expect, obj.stackPop)
		os.Exit(1)
	}

	result = obj.Pop()
	expect = nil
	expect = []int{2}
	if !cmp.EqualNumbers(1, result) {
		fmt.Printf("=== FAILED: expect = %v - got = %v\n", 1, result)
		os.Exit(1)
	}
	if !cmp.EqualSlices(expect, obj.stackPop) {
		fmt.Printf("=== FAILED: expect = %v - got = %v\n", expect, obj.stackPop)
		os.Exit(1)
	}

	if !cmp.EqualBooleans(false, obj.Empty()) {
		fmt.Printf("=== FAILED: expect = %v - got = %v\n", false, obj.Empty())
		os.Exit(1)
	}

	fmt.Printf("\n✅ All tests passed!\n")
}
