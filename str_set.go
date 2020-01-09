package cbyte

import (
	"reflect"
	"unsafe"
)

// Decompose []string to SliceHeader.
func HeaderStrSet(p []string) reflect.SliceHeader {
	return *(*reflect.SliceHeader)(unsafe.Pointer(&p))
}

// Compose []string from SliceHeader.
func StrSet(h reflect.SliceHeader) []string {
	return *(*[]string)(unsafe.Pointer(&h))
}

// Release addresses array memory.
// Note, if elements is cbyte strings, you need to release them manually before call this func.
func ReleaseStrSet(p []string) {
	p = p[:0]
	h := *(*reflect.SliceHeader)(unsafe.Pointer(&p))
	ReleaseHeader(h)
}
