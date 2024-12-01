package cbyte

// InitStr makes a string with underlying cbyte pointer.
// DEPRECATED: use InitString() instead.
func InitStr(len int) string {
	return InitString(len)
}

// ReleaseStr releases underlying cbyte pointer of string.
//
// Caution! Don't try to release non-cbyte strings.
// DEPRECATED: use ReleaseString() instead.
func ReleaseStr(p string) {
	ReleaseString(p)
}

// HeaderStr decomposes string to SliceHeader.
// DEPRECATED: use HeaderString() instead.
func HeaderStr(p string) SliceHeader {
	return HeaderString(p)
}

// Str composes string from SliceHeader.
// DEPRECATED: use String() instead.
func Str(h SliceHeader) string {
	return String(h)
}
