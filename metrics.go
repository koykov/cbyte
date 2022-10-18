package cbyte

// MetricsWriter interface.
type MetricsWriter interface {
	// Alloc register size of new allocated space.
	Alloc(size uint64)
	// Grow registers growing of cbyte object from sizeOld to sizeNew.
	Grow(sizeOld, sizeNew uint64)
	// Free registers freeing of cbyte object with given size.
	Free(size uint64)
}

var (
	// Builtin instance of metrics writer.
	// By default is a DummyMetrics object that does nothing on call.
	metricsHandler MetricsWriter = &DummyMetrics{}

	_ = RegisterMetricsHandler
)

// RegisterMetricsHandler register new metrics handler.
func RegisterMetricsHandler(handler MetricsWriter) {
	metricsHandler = handler
}
