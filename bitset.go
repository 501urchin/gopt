package gopt

type Bitset struct {
	b []uint8
}

func (bs *Bitset) At(idx int) bool {
	if idx < 0 {
		panic("cannot index on negative")
	}

	idx += 1

	internalIndex := (idx + 7) / 8
	next := internalIndex*8 - idx

	b := bs.b[internalIndex-1] >> next & 0b00000001

	return b != 0
}

func (*Bitset) Set(idx int, state bool) {}