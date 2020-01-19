# Cbyte

Provides low-level API to manipulate with byte arrays in C memory.

The main goal is reducing allocations and consecutive reduce GC pressure.

## Usage

```go
package main
import (
	"fmt"
	"github.com/koykov/cbyte"
)

func main() {
	var words = [][]byte{
            []byte("foo"),
            []byte("bar"),
            []byte("string"),
            []byte("of"),
            []byte("example"),
        }
	
	// Next line is equivalent of make([]byte, 0, 30), but it produces a byte slice
	// without using of runtime.mallocgc().
	// Therefore GC doesn't know nothing about this slice and just ignore it during
	// both GC mark and GC termination phases.
	p := cbyte.InitBytes(0, 30)
	for _, w := range words {
            p = append(p, w...)
	}
	fmt.Println(p)
	// This is the main inconvenience: need to resolve cbyte slices and strings
	// manually after using.
	cbyte.ReleaseBytes(p)
}
```

## How it works

As we know slices and strings in Go bases on the underlying arrays of corresponding types.
Each modifications of these arrays triggers new allocations that included:
* allocation of new array with required size;
* copy data from the old array.

After that old array (and corresponding slice or strings) become a garbage and need to be
collected during next GC cycle.

My idea is to move all low-level operations with underlying arrays to C memory,
outside of GC eyes. Each allocation, grow and free of the array triggers CGO calls, that
performs low-level manipulation with memory (see [cbyte.c](cbyte.c)).
All other operations like append(), copy(), ... works the same as on regular slices and strings.

It is a moot approach but it works. Just need to be a bit more careful.
