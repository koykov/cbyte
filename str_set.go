package cbyte

import (
	"reflect"
	"unsafe"
)

var _ = HeaderStrSet

// HeaderStrSet decomposes []string to SliceHeader.
func HeaderStrSet(p []string) reflect.SliceHeader {
	return *(*reflect.SliceHeader)(unsafe.Pointer(&p))
}

// StrSet composes []string from SliceHeader.
func StrSet(h reflect.SliceHeader) []string {
	return *(*[]string)(unsafe.Pointer(&h))
}

// ReleaseStrSet releases addresses array memory.
// Note, if elements is cbyte strings, you need to release them manually before call this func.
func ReleaseStrSet(p []string) {
	p = p[:0]
	h := *(*reflect.SliceHeader)(unsafe.Pointer(&p))
	ReleaseHeader(h)
}
