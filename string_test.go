package gopt

import (
	"bytes"
	"strings"
	"testing"
)

func TestStringToBytes(t *testing.T) {
	bs := "hello world"
	b := StringTobytes(bs)

	if bss := string(b); strings.Compare(bs, bss) != 0 {
		t.Errorf("expected %q but got %q", bs, bss)
	}
}

func TestBytesToString(t *testing.T) {
	bs := []byte("hello world")
	b := BytesToString(bs)

	if bss := []byte(b); !bytes.Equal(bs, bss) {
		t.Errorf("expected %q but got %q", bs, bss)
	}
}
