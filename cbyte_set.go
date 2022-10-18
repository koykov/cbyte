package cbyte

/*
#include "cbyte.h"
*/
import "C"
import (
	"reflect"
	"unsafe"
)

// InitSet makes addresses array to represent [][]byte in C memory.
func InitSet(cap int) uint64 {
	return uint64(C.cbyte_init_set(C.int(cap)))
}

// HeaderSet decomposes [][]byte to SliceHeader.
func HeaderSet(p [][]byte) reflect.SliceHeader {
	return *(*reflect.SliceHeader)(unsafe.Pointer(&p))
}

// BytesSet composes [][]byte from SliceHeader.
func BytesSet(h reflect.SliceHeader) [][]byte {
	return *(*[][]byte)(unsafe.Pointer(&h))
}

// ReleaseBytesSet releases addresses array memory.
// Note, if elements is cbyte slices, you need to release them manually before call this func.
func ReleaseBytesSet(p [][]byte) {
	for i := range p {
		// Clear item. It may be a cbyte slice. You need to release it manually preliminarily in that case.
		p[i] = nil
	}
	h := *(*reflect.SliceHeader)(unsafe.Pointer(&p))
	ReleaseHeader(h)
}
