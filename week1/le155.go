package leetcode

import (
	"container/list"
	"math"
)

// https://leetcode-cn.com/problems/min-stack/
// 155. 最小栈

type MinStack struct {
	values, min *list.List
}

func MinStackConstructor() MinStack {
	return MinStack{
		list.New(), list.New(),
	}
}

func (this *MinStack) Push(val int) {
	this.values.PushBack(val)
	e := this.min.Back()
	if e == nil || val < e.Value.(int) {
		this.min.PushBack(val)
	} else {
		this.min.PushBack(e.Value)
	}
}

func (this *MinStack) Pop() {
	this.min.Remove(this.min.Back())
	this.values.Remove(this.values.Back())
}

func (this *MinStack) Top() int {
	e := this.values.Back()
	if e == nil {
		return math.MinInt
	}
	return e.Value.(int)
}

func (this *MinStack) GetMin() int {
	e := this.min.Back()
	if e == nil {
		return math.MinInt
	}
	return e.Value.(int)
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
