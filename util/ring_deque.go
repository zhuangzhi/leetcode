package util

import "fmt"

type Ring struct {
	cap, size       int
	writeAt, readAt int
	data            []interface{}
}

func (r *Ring) ToString() string {
	format := `{
	Cap: %d,
	Size: %d,
	WriteAt: %d,
	ReadAt: %d,
	Data: %v,
}
`
	return fmt.Sprintf(format, r.cap, r.size, r.writeAt, r.readAt, r.data)
}
func NewRing(cap int) *Ring {
	return &Ring{
		cap:  cap,
		data: make([]interface{}, cap),
	}
}

func (r *Ring) InsertLast(value interface{}) bool {
	if r.size >= r.cap {
		return false
	}
	r.size++
	r.readAt = r.goBack(r.readAt)
	r.data[r.readAt] = value
	return true
}

func (r *Ring) InsertFront(value int) bool {
	if r.size >= r.cap {
		return false
	}
	r.size++
	r.data[r.writeAt] = value
	r.writeAt = r.goAhead(r.writeAt)
	return true
}

func (r *Ring) goAhead(idx int) int {
	if idx >= (r.cap - 1) {
		return 0
	}
	return idx + 1
}

func (r *Ring) goBack(idx int) int {
	idx--
	if idx < 0 {
		idx = r.cap - 1
	}
	return idx
}

// [1,2,3,0,0,0,0]
func (r *Ring) DeleteFront() bool {
	if r.size == 0 {
		return false
	}
	r.size--
	r.writeAt = r.goBack(r.writeAt)
	r.data[r.writeAt] = nil
	return true
}

func (r *Ring) DeleteLast() bool {
	if r.size == 0 {
		return false
	}
	r.size--
	r.data[r.readAt] = nil
	r.readAt = r.goAhead(r.readAt)
	return true
}

func (r *Ring) GetFront() interface{} {
	if r.size <= 0 {
		return nil
	}
	at := r.goBack(r.writeAt)
	return r.data[at]
}

func (r *Ring) GetRear() interface{} {
	if r.size <= 0 {
		return nil
	}
	return r.data[r.readAt]
}

func (r *Ring) IsEmpty() bool {
	return r.size == 0
}

func (r *Ring) IsFull() bool {
	return r.size == r.cap
}
