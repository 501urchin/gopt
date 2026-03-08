package gopt

import "testing"

func TestRingBuffer(t *testing.T) {
	rb := NewRingBuffer[int](3)

	t.Run("init with invalid cap", func(t *testing.T) {
		defer func() {
			err := recover()
			if err == nil {
				t.Error("failed to panic on invalid capacity")
			}
		}()

		NewRingBuffer[int](-2)
		t.Error("failed to panic on invalid capacity")

	})

	t.Run("head", func(t *testing.T) {
		rb.Insert(1)
		rb.Insert(2)
		rb.Insert(3)

		if h := rb.Head(); h != 1 {
			t.Errorf("ring buffer failed to return correct head: excpected %d but got %d", 1, h)
		}
	})

	t.Run("Indexing", func(t *testing.T) {
		if got := rb.At(0); got != 1 {
			t.Errorf("ring buffer index 0 incorrect: expected %d but got %d", 1, got)
		}
		if got := rb.At(1); got != 2 {
			t.Errorf("ring buffer index 1 incorrect: expected %d but got %d", 2, got)
		}
		if got := rb.At(2); got != 3 {
			t.Errorf("ring buffer index 2 incorrect: expected %d but got %d", 3, got)
		}
	})

	t.Run("Indexing out of range", func(t *testing.T) {
		defer func() {
			err := recover()
			if err == nil {
				t.Error("function failed to panic")
			}
		}()

		rb.At(4)
		t.Error("failed to cause a panic")
	})

	t.Run("overflowing element shoudl incr head", func(t *testing.T) {
		rb.Insert(4)

		if h := rb.Head(); h != 2 {
			t.Errorf("ring buffer failed to return correct head: excpected %d but got %d", 2, h)
		}
	})

	t.Run("tail", func(t *testing.T) {
		if tail := rb.Tail(); tail != 4 {
			t.Errorf("ring buffer failed to return correct tail: excpected %d but got %d", 4, tail)
		}
	})

}

func BenchmarkRingbuffer(b *testing.B) {
	rb := NewRingBuffer[int](100)

	b.Run("insert", func(b *testing.B) {
		for b.Loop() {
			rb.Insert(12)
		}
	})

	b.Run("indexing", func(b *testing.B) {
		tmp := 0
		for b.Loop() {
			tmp = rb.At(0)
		}
		_ = tmp
	})
}
