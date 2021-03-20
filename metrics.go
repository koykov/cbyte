package cbyte

// Metrics handler interface.
type MetricsWriter interface {
	// Register new allocation with given size.
	Alloc(size uint64)
	// Register growing of cbyte object from sizeOld to sizeNew.
	Grow(sizeOld, sizeNew uint64)
	// Register freeing of cbyte object with given size.
	Free(size uint64)
}

var (
	// Builtin instance of metrics writer.
	// By default is a DummyMetrics object that does nothing on call.
	metricsHandler MetricsWriter = &DummyMetrics{}

	_ = RegisterMetricsHandler
)

// Register new metrics handler.
func RegisterMetricsHandler(handler MetricsWriter) {
	metricsHandler = handler
}
