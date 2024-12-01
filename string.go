package cbyte

import (
	"unsafe"
)

// InitString makes a string with underlying cbyte pointer.
func InitString(len int) string {
	return String(InitHeader(len, len))
}

// ReleaseString releases underlying cbyte pointer of string.
//
// Caution! Don't try to release non-cbyte strings.
func ReleaseString(p string) {
	ReleaseHeader(HeaderString(p))
}

// HeaderString decomposes string to SliceHeader.
func HeaderString(p string) SliceHeader {
	h := *(*StringHeader)(unsafe.Pointer(&p))
	return SliceHeader{Data: h.Data, Len: h.Len, Cap: h.Len}
}

// String composes string from SliceHeader.
func String(h SliceHeader) string {
	hs := StringHeader{Data: h.Data, Len: h.Len}
	return *(*string)(unsafe.Pointer(&hs))
}
