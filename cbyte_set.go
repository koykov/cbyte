package cbyte

/*
#include "cbyte.h"
*/
import "C"
import (
	"reflect"
	"unsafe"
)

func InitSet(cap int) uint64 {
	return uint64(C.cbyte_init_set(C.int(cap)))
}

func SliceSet(h reflect.SliceHeader) [][]byte {
	return *(*[][]byte)(unsafe.Pointer(&h))
}

func ReleaseSliceSet(p [][]byte) {
	for i := range p {
		p[i] = nil
	}
	h := *(*reflect.SliceHeader)(unsafe.Pointer(&p))
	ReleaseHeader(h)
}
