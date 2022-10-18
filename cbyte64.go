package cbyte

// Cbyte implementation for big byte arrays.

/*
#include "cbyte.h"
*/
import "C"

// SliceHeader64 represents runtime representation of a big byte slice.
// Allows to exceed MAXINT limit for length and capacity.
type SliceHeader64 struct {
	Data uintptr
	Len  uint64
	Cap  uint64
}

// Init64 makes big byte array in C memory.
func Init64(cap uint64) uint64 {
	metricsHandler.Alloc(cap)
	return uint64(C.cbyte_init64(C.uint64(cap)))
}

// InitHeader64 makes return SliceHeader64 of big byte array.
func InitHeader64(len, cap uint64) SliceHeader64 {
	return SliceHeader64{
		Data: uintptr(Init64(cap)),
		Len:  len,
		Cap:  cap,
	}
}

// Grow64 increases capacity of the big byte array.
func Grow64(addr uint64, capOld, cap uint64) uint64 {
	metricsHandler.Grow(capOld, cap)
	if capOld > mallocGrowThreshold {
		return uint64(C.cbyte_grow64_r(C.uint64(addr), C.uint64(cap)))
	} else {
		return uint64(C.cbyte_grow64_m(C.uint64(addr), C.uint64(capOld), C.uint64(cap)))
	}
}

// GrowHeader64 increases capacity using SliceHeader64.
func GrowHeader64(h SliceHeader64) uint64 {
	return Grow64(uint64(h.Data), h.Len, h.Cap)
}

// ReleaseHeader64 releases byte array using SliceHeader64.
func ReleaseHeader64(h SliceHeader64) {
	metricsHandler.Free(h.Cap)
	Release(uint64(h.Data))
}
