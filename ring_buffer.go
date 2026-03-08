package gopt

import (
	"fmt"
)

type NumberConstraint interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64
}

// TODO: make it thread safe either by using lock free queues or a RWMutex
type RingBuffer[T NumberConstraint] struct {
	cap     int
	headIdx int
	buf     []T
}

// ring buffer is not thread safe
func NewRingBuffer[T NumberConstraint](cap int, initialValues ...T) *RingBuffer[T] {
	if cap <= 0 {
		panic(fmt.Errorf("invalid capacity: excpect a cap greater than 0 but got %d", cap))
	}

	rb := &RingBuffer[T]{
		cap:     cap,
		headIdx: 0,
		buf:     make([]T, 0, cap),
	}

	if len(initialValues) > 0 {
		rb.Insert(initialValues...)
	}

	return rb
}

// insert values into the ring buffer
func (rb *RingBuffer[T]) Insert(vals ...T) {
	rbLen := len(rb.buf)
	for _, val := range vals {
		if rbLen < rb.cap {
			rb.buf = append(rb.buf, val)
			continue
		}

		rb.buf[rb.headIdx] = val
		rb.headIdx++

		if rb.headIdx >= rb.cap-1 {
			rb.headIdx = 0
		}
	}
}

// helper function to make indexing the ring buffer simpler. this is eq to slice[i]
func (rb *RingBuffer[T]) At(i int) T {
	if i < 0 || i >= rb.cap {
		panic(fmt.Errorf("index out of range [%d] with length %d", i, rb.cap))
	}

	idx := rb.headIdx + i
	if idx >= rb.cap {
		idx -= rb.cap
	}

	return rb.buf[idx]
}

// returns the head val of the ring buffer
func (rb *RingBuffer[T]) Head() T {
	return rb.buf[rb.headIdx]
}

func (rb *RingBuffer[T]) Cap() int {
	return rb.cap
}

func (rb *RingBuffer[T]) Len() int {
	return len(rb.buf)
}

func (rb *RingBuffer[T]) BackingSlice() []T {
	return rb.buf
}
