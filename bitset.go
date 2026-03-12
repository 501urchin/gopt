package gopt

type Bitset struct {
	b []uint8
}

func NewBitset(size ...int) *Bitset {
	hasSize := len(size) != 0
	if hasSize && size[0] < 0 {
		panic("please provide a size thats > 0")
	}

	if hasSize {
		return &Bitset{make([]uint8, (size[0]+7)/8)}
	}

	return &Bitset{make([]uint8, 0)}
}

func (bs *Bitset) At(idx int) bool {
	if idx < 0 {
		panic("cannot index on negative")
	}

	idx += 1

	internalIndex := (idx + 7) / 8
	next := internalIndex*8 - idx

	if internalIndex > len(bs.b) {
		panic("index overflows")
	}

	b := bs.b[internalIndex-1] >> next & 0b00000001

	return b != 0
}

func (*Bitset) Set(idx int, state bool) {}
