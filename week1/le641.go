package leetcode

// https://leetcode-cn.com/problems/design-circular-deque/
// 641. 设计循环双端队列
type DequeNode struct {
	val       int
	pre, next *DequeNode
}

type MyCircularDeque struct {
	head, tail *DequeNode
	cap, size  int
}

func DequeConstructor(k int) MyCircularDeque {
	head, tail := &DequeNode{val: -1}, &DequeNode{val: -1}
	head.next = tail
	tail.pre = head
	return MyCircularDeque{
		head: head,
		tail: tail,
		cap:  k,
		size: 0,
	}
}

func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.size >= this.cap {
		return false
	}
	this.size++
	node := &DequeNode{val: value}
	node.next = this.head.next
	node.next.pre = node
	node.pre = this.head
	node.pre.next = node
	return true
}

func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.size >= this.cap {
		return false
	}
	this.size += 1
	node := &DequeNode{val: value}
	node.next = this.tail
	node.pre = this.tail.pre
	node.next.pre = node
	node.pre.next = node
	return true
}

func (this *MyCircularDeque) DeleteFront() bool {
	if this.size == 0 {
		return false
	}

	this.size--
	node := this.head.next
	this.head.next = node.next
	node.next.pre = this.head
	return true
}

func (this *MyCircularDeque) DeleteLast() bool {
	if this.size == 0 {
		return false
	}
	this.size--
	this.tail.pre = this.tail.pre.pre
	this.tail.pre.next = this.tail
	return true
}

func (this *MyCircularDeque) GetFront() int {
	if this.size < 0 {
		return -1
	}
	return this.head.next.val
}

func (this *MyCircularDeque) GetRear() int {
	if this.size < 0 {
		return -1
	}
	return this.tail.pre.val
}

func (this *MyCircularDeque) IsEmpty() bool {
	return this.size == 0
}

func (this *MyCircularDeque) IsFull() bool {
	return this.size == this.cap
}

/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */
