package gopt

import "fmt"

type Bitset struct {
	b   []uint8
	len int
}

func NewBitset(size ...int) *Bitset {
	hasSize := len(size) != 0
	if hasSize && size[0] < 0 {
		panic("please provide a size thats > 0")
	}

	if hasSize {
		return &Bitset{
			b:   make([]uint8, (size[0]+7)/8),
			len: 0,
		}
	}

	return &Bitset{
		b:   make([]uint8, 0),
		len: 0,
	}
}

// Items is a iterator for the bitset
//
//	var bitSet BitSet
//	for b := range bitSet.Items {
//		// handle b
//	}
func (bs *Bitset) Items(yield func(bool) bool) {
	for i := range len(bs.b) * 8 {
		if !yield(bs.At(i)) {
			return
		}
	}

}

func (bs *Bitset) At(idx int) bool {
	if idx < 0 || idx >= bs.len {
		panic(fmt.Errorf("invalid bitset index %d of len %d", idx, bs.len))
	}

	byteIndex := idx / 8
	bitIndex := idx % 8
	var mask uint8 = 1 << (7 - bitIndex)

	return bs.b[byteIndex]&mask != 0
}

// Set is used to set a bool at a specific index
func (bs *Bitset) Set(idx int, state bool) {
	if idx < 0 || idx >= bs.len {
		panic(fmt.Errorf("invalid bitset index %d of len %d", idx, bs.len))
	}

	byteIndex := idx / 8
	bitIndex := idx % 8
	var mask uint8 = 1 << (7 - bitIndex)

	if state {
		bs.b[byteIndex] |= mask
	} else {
		bs.b[byteIndex] &^= mask
	}
}

func (bs *Bitset) Len() int {
	return bs.len
}
func (bs *Bitset) Append(elm ...bool) {
	for _, b := range elm {
		next := bs.len

		if bs.len >= len(bs.b)*8 {
			bs.b = append(bs.b, 0)
		}

		bs.len++
		bs.Set(next, b)
	}
}
