package cbyte

import (
	"reflect"
	"unsafe"
)

func HeaderStrSet(p []string) reflect.SliceHeader {
	return *(*reflect.SliceHeader)(unsafe.Pointer(&p))
}

func SliceStrSet(h reflect.SliceHeader) []string {
	return *(*[]string)(unsafe.Pointer(&h))
}

func ReleaseStrSet(p []string) {
	p = p[:0]
	h := *(*reflect.SliceHeader)(unsafe.Pointer(&p))
	ReleaseHeader(h)
}
