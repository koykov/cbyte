package cbyte

import (
	"reflect"
	"unsafe"
)

// Init and return cbyte string.
func InitStr(len int) string {
	return Str(InitHeader(len, len))
}

// Take address of the string and release memory using it.
//
// Caution! Don't try to release non-cbyte strings.
func ReleaseStr(p string) {
	ReleaseHeader(HeaderStr(p))
}

// Decompose string to SliceHeader.
func HeaderStr(p string) reflect.SliceHeader {
	h := *(*reflect.StringHeader)(unsafe.Pointer(&p))
	return reflect.SliceHeader{Data: h.Data, Len: h.Len, Cap: h.Len}
}

// Compose string from SliceHeader.
func Str(h reflect.SliceHeader) string {
	hs := reflect.StringHeader{Data: h.Data, Len: h.Len}
	return *(*string)(unsafe.Pointer(&hs))
}
