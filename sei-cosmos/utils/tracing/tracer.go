package tracing

import (
	"context"
	"sync"
	"sync/atomic"

	"go.opentelemetry.io/otel/sdk/trace"
	otrace "go.opentelemetry.io/otel/trace"
)

const DefaultTracingURL = "http://localhost:14268/api/traces"
const FlagTracing = "tracing"

func DefaultTracerProvider() (*trace.TracerProvider, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func TracerProvider(url string) (*trace.TracerProvider, error) {
	_ = "STUB: not implemented"
	// Create the Jaeger exporter
	return nil, nil
}

func GetTracerProviderOptions(url string) ([]trace.TracerProviderOption, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Always be sure to batch in production.

// Record information about this application in a Resource.

type Info struct {
	tracer         otrace.Tracer
	tracingEnabled atomic.Bool
	mtx            sync.RWMutex
}

func NewTracingInfo(tr otrace.Tracer, tracingEnabled bool) *Info {
	_ = "STUB: not implemented"
	return nil
}

// NoOpSpan is a no-op span which does nothing.
var NoOpSpan = otrace.SpanFromContext(context.TODO())

func (i *Info) Start(name string) (context.Context, otrace.Span) {
	_ = "STUB: not implemented"
	return *new(context.Context), *new(otrace.Span)
}

func (i *Info) StartWithContext(name string, ctx context.Context) (context.Context, otrace.Span) {
	_ = "STUB: not implemented"
	return *new(context.Context), *new(otrace.Span)
}
