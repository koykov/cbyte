package cbyte

import (
	"bytes"
	"testing"
)

var t__ = []byte(`Lorem ipsum dolor sit amet, consectetur adipiscing elit. Curabitur luctus sapien eu velit sodales laoreet. Vestibulum laoreet in ex vel faucibus. Praesent in nibh arcu. In justo velit, ultricies id consequat in, porta volutpat mauris. Maecenas porttitor, mi sed condimentum pretium, enim nisl scelerisque nibh, eget tincidunt orci risus nec velit.`)

func TestCbyte(t *testing.T) {
	t.Run("init", func(t *testing.T) {
		h := InitHeader(64, 64)
		if h.Cap != 64 {
			t.Fail()
		}
		ReleaseHeader(h)
	})
	t.Run("grow", func(t *testing.T) {
		h := InitHeader(0, 64)
		h.Cap = 2048
		h.Data = uintptr(Grow(uint64(h.Data), 64, 2048))
		if cap(Bytes(h)) != 2048 {
			t.Fail()
		}
		ReleaseHeader(h)
	})
	t.Run("memcpy", func(t *testing.T) {
		h := InitHeader(512, 512)
		n := Memcpy(uint64(h.Data), 16, t__)
		if n != len(t__) {
			t.Fail()
		}
		b := Bytes(h)
		if !bytes.Equal(t__, b[16:len(t__)+16]) {
			t.Fail()
		}
		ReleaseHeader(h)
	})
}
