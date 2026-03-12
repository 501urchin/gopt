package gopt

import (
	"testing"
)

func TestBitset(t *testing.T) {
	bs := Bitset{
		b: []uint8{0b10101010},
	}

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
}
