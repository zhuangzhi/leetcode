package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGrowingQueue(t *testing.T) {
	q := NewGrowingRingQueue(5)
	for _, v := range Range(0, 10) {
		q.PushBack(v)
	}
	assert.Equal(t, 9, q.Back())
	assert.Equal(t, 0, q.Front())
	assert.Equal(t, 10, q.Size())
	for _, v := range Range(0, 10) {
		x := q.Front()
		assert.Equal(t, v, x)
		q.PopFront()
	}
	for _, v := range Range(0, 10) {
		q.PushBack(v)
		x := q.Front()
		assert.Equal(t, v, x)
		q.PopFront()
	}
	for _, v := range Range(0, 20) {
		q.PushFront(v)
		x := q.Back()
		assert.Equal(t, v, x)
		q.PopBack()
	}
	for _, v := range Range(0, 20) {
		q.PushFront(v)
	}
	for _, v := range Range(0, 20) {
		x := q.Back()
		assert.Equal(t, v, x)
		q.PopBack()
	}
	for _, v := range Range(0, 10) {
		q.PushBack(v)
	}
	c := q.Copy()
	for i, v := range Range(0, 10) {
		assert.Equal(t, c[i], v)
	}
}

func TestFixedQueue(t *testing.T) {
	q := NewFixedRingQueue(5)
	for _, v := range Range(0, 10) {
		q.PushBack(v)
	}
	assert.Equal(t, 9, q.Back())
	assert.Equal(t, 5, q.Front())
	assert.Equal(t, 5, q.Size())
	for _, v := range Range(5, 10) {
		x := q.Front()
		assert.Equal(t, v, x)
		q.PopFront()
	}
	for _, v := range Range(0, 10) {
		q.PushBack(v)
		x := q.Front()
		assert.Equal(t, v, x)
		q.PopFront()
	}
	for _, v := range Range(0, 20) {
		q.PushFront(v)
		x := q.Back()
		assert.Equal(t, v, x)
		q.PopBack()
	}
	for _, v := range Range(0, 5) {
		q.PushFront(v)
	}
	for _, v := range Range(0, 5) {
		x := q.Back()
		assert.Equal(t, v, x)
		q.PopBack()
	}
	for _, v := range Range(0, 5) {
		q.PushBack(v)
	}
	c := q.Copy()
	for i, v := range Range(0, 5) {
		assert.Equal(t, c[i], v)
	}
}
