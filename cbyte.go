package cbyte

/*
#include "cbyte.h"
*/
import "C"
import (
	"unsafe"
)

const (
	// Buffer size limit to use malloc to grow.
	mallocGrowThreshold = 1024
)

type SliceHeader struct {
	Data uintptr
	Len  int
	Cap  int
}

type StringHeader struct {
	Data uintptr
	Len  int
}

// Init makes byte array in C memory, outside of GC's eyes.
func Init(cap int) uint64 {
	metricsHandler.Alloc(uint64(cap))
	return uint64(C.cbyte_init(C.int(cap)))
}

// InitHeader makes slice header of byte array.
func InitHeader(len, cap int) SliceHeader {
	return SliceHeader{
		Data: uintptr(Init(cap)),
		Len:  len,
		Cap:  cap,
	}
}

// InitBytes makes and return byte slice.
func InitBytes(len, cap int) []byte {
	return Bytes(InitHeader(len, cap))
}

// Grow increases capacity of the byte array.
//
// All necessary copying/free will perform implicitly, don't worry about this.
func Grow(addr uint64, capOld, cap int) uint64 {
	// Using combination of malloc()+memcpy()+free() to grow for short buffers is more efficient than simple using
	// of realloc().
	metricsHandler.Grow(uint64(capOld), uint64(cap))
	if capOld > mallocGrowThreshold {
		return uint64(C.cbyte_grow_r(C.uint64(addr), C.int(cap)))
	} else {
		return uint64(C.cbyte_grow_m(C.uint64(addr), C.int(capOld), C.int(cap)))
	}
}

// GrowHeader increases capacity of the byte array using SliceHeader.
func GrowHeader(h SliceHeader) uint64 {
	return Grow(uint64(h.Data), h.Len, h.Cap)
}

// Memcpy makes a copy of data directly to the addr+offset.
func Memcpy(addr, offset uint64, data []byte) (n int) {
	n = len(data)
	h := SliceHeader{
		Data: uintptr(addr + offset),
		Len:  n,
		Cap:  n,
	}
	b := Bytes(h)
	copy(b, data)
	return n
}

// Release releases cbyte pointer.
func Release(addr uint64) {
	if addr == 0 {
		return
	}
	C.cbyte_release(C.uint64(addr))
}

// ReleaseHeader free byte array using SliceHeader.
func ReleaseHeader(h SliceHeader) {
	metricsHandler.Free(uint64(h.Cap))
	Release(uint64(h.Data))
}

// ReleaseBytes free underlying cbyte slice.
//
// Caution! Don't try to release non-cbyte slices.
func ReleaseBytes(p []byte) {
	ReleaseHeader(Header(p))
}

// Header converts byte slice to SliceHeader.
func Header(p []byte) SliceHeader {
	return *(*SliceHeader)(unsafe.Pointer(&p))
}

// Bytes composes byte slice from SliceHeader.
func Bytes(h SliceHeader) []byte {
	return *(*[]byte)(unsafe.Pointer(&h))
}

var _ = GrowHeader
