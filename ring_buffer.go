package gopt

import (
	"fmt"
)

type NumberConstraint interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64
}

type RingBuffer[T NumberConstraint] struct {
	cap     int
	headIdx int
	buf     []T
}

// NewRingBuffer created a new ring buffer. note that the ring buffer is not thread safe
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

// Insert values into the ring buffer
func (rb *RingBuffer[T]) Insert(vals ...T) {
	for _, val := range vals {
		if len(rb.buf) < rb.cap {
			rb.buf = append(rb.buf, val)
			continue
		}

		rb.buf[rb.headIdx] = val
		rb.headIdx++

		if rb.headIdx >= rb.cap {
			rb.headIdx = 0
		}
	}
}

// At function to make indexing the ring buffer simpler. this is eq to slice[i]
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

// Head returns the oldest inserted value in the ring buffer
func (rb *RingBuffer[T]) Head() T {
	return rb.buf[rb.headIdx]
}

// Tail returns the latest inserted value in the ring buffer
func (rb *RingBuffer[T]) Tail() T {
	idx := rb.headIdx - 1
	if idx < 0 {
		idx = rb.cap - 1
	}

	return rb.buf[idx]
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
