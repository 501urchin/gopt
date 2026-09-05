package gopt

import (
	"reflect"
	"unsafe"
)

// Credit to https://github.com/valyala/fasthttp?tab=readme-ov-file#tricks-with-byte-buffers
func StringToBytes(s string) (b []byte) {
	slice := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	stringP := (*reflect.StringHeader)(unsafe.Pointer(&s))

	slice.Data = stringP.Data
	slice.Len = stringP.Len
	slice.Cap = stringP.Len

	return b
}


func BytesToString(b []byte) (s string) {
	return *(*string)(unsafe.Pointer(&b))
}
