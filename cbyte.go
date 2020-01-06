package cbyte

/*
#include "cbyte.h"
*/
import "C"
import (
	"reflect"
	"unsafe"
)

const (
	// Limit to switch to loop rolling write.
	shortInputThreshold = 256
	// Buffer size limit to use malloc to grow.
	mallocGrowThreshold = 1024
)

// Init makes byte array in C memory, outside of GC's eyes.
func Init(cap int) uint64 {
	return uint64(C.cbyte_init(C.int(cap)))
}

// Increase capacity of the byte array.
//
// All necessary copying/free will perform implicitly, don't worry about this.
func Grow(addr uint64, capOld, cap int) uint64 {
	// Using combination of malloc()+memcpy()+free() to grow for short buffers is more efficient than simple using
	// of realloc().
	if capOld > mallocGrowThreshold {
		return uint64(C.cbyte_grow_r(C.ulong(addr), C.int(cap)))
	} else {
		return uint64(C.cbyte_grow_m(C.ulong(addr), C.int(capOld), C.int(cap)))
	}
}

// Increase capacity of the byte array using SliceHeader.
func GrowHeader(h reflect.SliceHeader) uint64 {
	return Grow(uint64(h.Data), h.Len, h.Cap)
}

// Copy data bytes directly to the byte array.
func Memcpy(addr, offset uint64, data []byte) (n int) {
	if len(data) > shortInputThreshold {
		// Write long data using loop rolling.
		for len(data) >= 8 {
			*(*byte)(unsafe.Pointer(uintptr(addr + offset))) = data[0]
			*(*byte)(unsafe.Pointer(uintptr(addr + offset + 1))) = data[1]
			*(*byte)(unsafe.Pointer(uintptr(addr + offset + 2))) = data[2]
			*(*byte)(unsafe.Pointer(uintptr(addr + offset + 3))) = data[3]
			*(*byte)(unsafe.Pointer(uintptr(addr + offset + 4))) = data[4]
			*(*byte)(unsafe.Pointer(uintptr(addr + offset + 5))) = data[5]
			*(*byte)(unsafe.Pointer(uintptr(addr + offset + 6))) = data[6]
			*(*byte)(unsafe.Pointer(uintptr(addr + offset + 7))) = data[7]
			data = data[8:]
			offset += 8
			n += 8
		}
		for len(data) >= 4 {
			*(*byte)(unsafe.Pointer(uintptr(addr + offset))) = data[0]
			*(*byte)(unsafe.Pointer(uintptr(addr + offset + 1))) = data[1]
			*(*byte)(unsafe.Pointer(uintptr(addr + offset + 2))) = data[2]
			*(*byte)(unsafe.Pointer(uintptr(addr + offset + 3))) = data[3]
			data = data[4:]
			offset += 4
			n += 4
		}
		for len(data) >= 2 {
			*(*byte)(unsafe.Pointer(uintptr(addr + offset))) = data[0]
			*(*byte)(unsafe.Pointer(uintptr(addr + offset + 1))) = data[1]
			data = data[2:]
			offset += 2
			n += 2
		}
		if len(data) > 0 {
			*(*byte)(unsafe.Pointer(uintptr(addr + offset))) = data[0]
			offset++
			n++
		}
	} else {
		for i := 0; i < len(data); i++ {
			*(*byte)(unsafe.Pointer(uintptr(addr + offset + uint64(i)))) = data[i]
		}
		n = len(data)
	}
	return
}

// Release byte array.
func Release(addr uint64) {
	if addr == 0 {
		return
	}
	C.cbyte_release(C.ulong(addr))
}

// Release byte array using SliceHeader.
func ReleaseHeader(h reflect.SliceHeader) {
	Release(uint64(h.Data))
}

// Take address of the byte slice and release memory using it.
//
// Caution! Don't try to release non-cbyte slices.
func ReleaseSlice(p []byte) {
	ReleaseHeader(Header(p))
}

// Decompose byte slice to SliceHeader.
func Header(p []byte) reflect.SliceHeader {
	return *(*reflect.SliceHeader)(unsafe.Pointer(&p))
}

// Compose byte slice from SliceHeader.
func Slice(h reflect.SliceHeader) []byte {
	return *(*[]byte)(unsafe.Pointer(&h))
}
