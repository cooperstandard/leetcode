package main

import (
	"github.com/cooperstandard/leetcode/datastructures"
)

type MinStack struct {
	stack datastructures.Stack[int]
	min   int
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	if this.stack.IsEmpty() {
		this.min = val
		this.stack.Push(val)
	} else {
		if val < this.min {
			this.stack.Push(val - this.min)
			this.min = val
		}

	}
}

func (this *MinStack) Pop() {
	pop, _ := this.stack.Pop()
	if pop < 0 {
		this.min -= pop
	}

}

func (this *MinStack) Top() int {
	top, _ := this.stack.Peek()
	if top > 0 {
		return top + this.min
	}
	return top
}

func (this *MinStack) GetMin() int {
	return this.min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

