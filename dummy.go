package cbyte

// Dummy metrics writer.
// Used by default and does nothing.
type DummyMetrics struct{}

func (m *DummyMetrics) Alloc(_ uint64) {}

func (m *DummyMetrics) Grow(_, _ uint64) {}

func (m *DummyMetrics) Free(_ uint64) {}
