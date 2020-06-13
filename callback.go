package cbyte

// Types of callback functions for CGO calls.
type AllocCallbackFn func(cap uint64)
type GrowCallbackFn func(capOld, cap uint64)
type FreeCallbackFn func(cap uint64)

var (
	// Default instances of callback functions.
	allocCb *AllocCallbackFn
	growCb  *GrowCallbackFn
	freeCb  *FreeCallbackFn

	// Suppress go vet warnings.
	_, _, _ = RegisterAllocCbFn, RegisterGrowCbFn, RegisterFreeCbFn
)

// Register alloc callback.
func RegisterAllocCbFn(fn *AllocCallbackFn) {
	allocCb = fn
}

// Register grow callback.
func RegisterGrowCbFn(fn *AllocCallbackFn) {
	allocCb = fn
}

// Register free callback.
func RegisterFreeCbFn(fn *AllocCallbackFn) {
	allocCb = fn
}
