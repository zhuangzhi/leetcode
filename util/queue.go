package util

type GrowingRingQueue struct {
	idxR, idxW, size, cap int
	data                  []interface{}
}

func NewGrowingRingQueue(cap int) *GrowingRingQueue {
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

func (queue *GrowingRingQueue) PushBack(val interface{}) {
	if queue.size >= queue.cap {
		queue.grow()
	}
	queue.size++
	queue.data[queue.idxW] = val
	queue.idxW = queue.next(queue.idxW)
}

func (queue *GrowingRingQueue) grow() {
	buf := make([]interface{}, queue.cap*2)
	data := queue.data
	copy(buf, data[queue.idxR:])
	copy(buf[queue.cap-queue.idxR:], data[:queue.idxR])
	queue.data = buf
	queue.idxR = 0
	queue.idxW = queue.cap
	queue.cap *= 2
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

func (queue *GrowingRingQueue) PushFront(val interface{}) {
	if queue.size >= queue.cap {
		queue.grow()
	}
	queue.size++
	queue.idxR = queue.back(queue.idxR)
	queue.data[queue.idxR] = val
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
