package gopt

import (
	"slices"
	"testing"
)

func TestBitsetIter(t *testing.T) {
	bs := Bitset{
		b: []uint8{0b10101010},
	}

	res := make([]bool, 0, 8)

	ex := []bool{true, false, true, false, true, false, true, false}

	for b := range bs.Items {
		res = append(res, b)
	}

	if !slices.Equal(res, ex) {
		t.Errorf("iterator failed since slice dont match: expected %v but got %v", ex, res)
	}

}

func TestBitsetAt(t *testing.T) {
	bs := Bitset{
		b: []uint8{0b10101010, 0b10},
	}

	t.Run("gets correct state", func(t *testing.T) {
		if bs.At(0) != true {
			t.Error("bit 0 should be true")
		}

		if bs.At(1) != false {
			t.Error("bit 1 should be false")
		}

		if bs.At(2) != true {
			t.Error("bit 2 should be true")
		}

		if bs.At(3) != false {
			t.Error("bit 3 should be false")
		}

		if bs.At(4) != true {
			t.Error("bit 4 should be true")
		}

		if bs.At(5) != false {
			t.Error("bit 5 should be false")
		}

		if bs.At(6) != true {
			t.Error("bit 6 should be true")
		}

		if bs.At(7) != false {
			t.Error("bit 7 should be false")
		}

		if bs.At(10) != false {
			t.Error("bit 10 should be false")
		}

	})

	t.Run("negative index", func(t *testing.T) {
		defer func() {
			err := recover()
			if err == nil {
				t.Error("failed to panic on negative index")
			}
		}()

		bs.At(-14)

		t.Error("failed to panic on negative index")

	})

	t.Run("overflowing index", func(t *testing.T) {
		defer func() {
			err := recover()
			if err == nil {
				t.Error("failed to panic on overflowing index")
			}
		}()

		bs.At(200)

		t.Error("failed to panic on overflowing index")

	})
}

func TestNewBitset(t *testing.T) {
	t.Run("valid case", func(t *testing.T) {
		bs := NewBitset(12)

		if len(bs.b) != 2 {
			t.Error("failed to create a bitset of the expected size")
		}
	})

	t.Run("negative size", func(t *testing.T) {
		defer func() {
			err := recover()
			if err == nil {
				t.Error("failed to panic on negative index")
			}
		}()

		bs := NewBitset(-1)
		t.Error("failed to panic on negative index")
		_ = bs
	})

	t.Run("zero size", func(t *testing.T) {

		bs := NewBitset()
		if len(bs.b) != 0 {
			t.Error("failed to create a bitset of the expected size")
		}
	})

}

func BenchmarkBitset(b *testing.B) {
	b.Run("normal", func(b *testing.B) {
		n := []bool{true, false, true}

		tmp := true

		for b.Loop() {
			tmp = n[0]
		}
		_ = tmp
	})

	b.Run("bitset", func(b *testing.B) {
		bs := Bitset{
			b: []uint8{0b101},
		}

		tmp := true
		for b.Loop() {
			tmp = bs.At(0)
		}
		_ = tmp
	})
}
