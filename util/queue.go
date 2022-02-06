package util

type Queue interface {
	PushBack(v interface{}) bool
	PushFront(v interface{}) bool
	PopBack() bool
	PopFront() bool
	Back() interface{}
	Front() interface{}
	Size() int
	Copy() []interface{}
}

type FixedRingQueue struct {
	GrowingRingQueue
}

func NewFixedRingQueue(cap int) Queue {
	return &FixedRingQueue{
		GrowingRingQueue{
			cap:  cap,
			data: make([]interface{}, cap),
		}}
}

func (f *FixedRingQueue) PushBack(v interface{}) bool {
	if f.size == f.cap {
		f.PopFront()
	}
	return f.GrowingRingQueue.PushBack(v)
}

func (f *FixedRingQueue) PushFront(v interface{}) bool {
	if f.size == f.cap {
		f.PopBack()
	}
	return f.GrowingRingQueue.PushFront(v)
}

type GrowingRingQueue struct {
	idxR, idxW, size, cap int
	data                  []interface{}
}

func NewGrowingRingQueue(cap int) Queue {
	return &GrowingRingQueue{
		cap:  cap,
		data: make([]interface{}, cap),
	}
}

func (queue GrowingRingQueue) next(idx int) int {
	if idx >= (queue.cap - 1) {
		return 0
	}
	return idx + 1
}

func (queue GrowingRingQueue) back(idx int) int {
	idx--
	if idx < 0 {
		idx = queue.cap - 1
	}
	return idx
}

func (queue *GrowingRingQueue) PushBack(val interface{}) bool {
	if queue.size >= queue.cap {
		queue.grow()
	}
	queue.size++
	queue.data[queue.idxW] = val
	queue.idxW = queue.next(queue.idxW)
	return true
}

func (queue *GrowingRingQueue) grow() {
	buf := make([]interface{}, queue.cap*2)
	queue.copy(buf)
	queue.data = buf
	queue.idxR = 0
	queue.idxW = queue.cap
	queue.cap *= 2
}

func (queue *GrowingRingQueue) copy(to []interface{}) {
	if queue.size == 0 {
		return
	}
	if len(to) < queue.cap {
		panic("copy to a small buf")
	}
	data := queue.data
	w, r := queue.idxW, queue.idxR
	cap := queue.cap
	if w > r {
		copy(to, data[r:w])
	} else {
		copy(to, data[r:])
		copy(to[cap-r:], data[:w])
	}
}

func (queue *GrowingRingQueue) Copy() []interface{} {
	buf := make([]interface{}, queue.cap)
	queue.copy(buf)
	return buf[:queue.size]
}

func (queue *GrowingRingQueue) PopBack() bool {
	if queue.size == 0 {
		return false
	}
	queue.size--
	queue.idxW = queue.back(queue.idxW)
	queue.data[queue.idxW] = nil
	return true
}

func (queue *GrowingRingQueue) PushFront(val interface{}) bool {
	if queue.size >= queue.cap {
		queue.grow()
	}
	queue.size++
	queue.idxR = queue.back(queue.idxR)
	queue.data[queue.idxR] = val
	return true
}

func (queue *GrowingRingQueue) PopFront() bool {
	if queue.size == 0 {
		return false
	}
	queue.size--
	queue.data[queue.idxR] = nil
	queue.idxR = queue.next(queue.idxR)
	return true
}

func (queue *GrowingRingQueue) Empty() bool {
	return queue.size == 0
}

func (queue *GrowingRingQueue) Size() int {
	return queue.size
}

func (queue *GrowingRingQueue) Back() (v interface{}) {
	if queue.size <= 0 {
		return nil
	}
	at := queue.back(queue.idxW)
	return queue.data[at]
}

func (queue *GrowingRingQueue) Front() (v interface{}) {
	if queue.size <= 0 {
		return nil
	}
	return queue.data[queue.idxR]
}
