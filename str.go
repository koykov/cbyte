package cbyte

import (
	"reflect"
	"unsafe"
)

// Take address of the string and release memory using it.
//
// Caution! Don't try to release non-cbyte strings.
func ReleaseStr(p string) {
	ReleaseHeader(HeaderStr(p))
}

// Decompose byte slice to SliceHeader.
func HeaderStr(p string) reflect.SliceHeader {
	h := *(*reflect.StringHeader)(unsafe.Pointer(&p))
	return reflect.SliceHeader{Data: h.Data, Len: h.Len, Cap: h.Len}
}

// Compose byte slice from SliceHeader.
func Str(h reflect.SliceHeader) string {
	hs := reflect.StringHeader{Data: h.Data, Len: h.Len}
	return *(*string)(unsafe.Pointer(&hs))
}
