package cbyte

import (
	"reflect"
	"unsafe"
)

var _ = InitStr

// InitStr makes a string with underlying cbyte pointer.
func InitStr(len int) string {
	return Str(InitHeader(len, len))
}

// ReleaseStr releases underlying cbyte pointer of string.
//
// Caution! Don't try to release non-cbyte strings.
func ReleaseStr(p string) {
	ReleaseHeader(HeaderStr(p))
}

// HeaderStr decomposes string to SliceHeader.
func HeaderStr(p string) reflect.SliceHeader {
	h := *(*reflect.StringHeader)(unsafe.Pointer(&p))
	return reflect.SliceHeader{Data: h.Data, Len: h.Len, Cap: h.Len}
}

// Str composes string from SliceHeader.
func Str(h reflect.SliceHeader) string {
	hs := reflect.StringHeader{Data: h.Data, Len: h.Len}
	return *(*string)(unsafe.Pointer(&hs))
}
