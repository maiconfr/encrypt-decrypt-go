package telemetry

import (
	"context"
	"log"
	"os"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/exporters/stdout/stdoutmetric"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
	oteltrace "go.opentelemetry.io/otel/trace"
	"go.opentelemetry.io/otel/trace/noop"
)

var (
	Tracer oteltrace.Tracer
)

// InitProvider initializes OpenTelemetry providers
func InitProvider(serviceName string) (func(), error) {
	// Check if telemetry is disabled via environment variable
	if os.Getenv("OTEL_SDK_DISABLED") == "false" {
		Tracer = noop.NewTracerProvider().Tracer(serviceName)
		return func() {}, nil
	}

	// Create trace exporter
	traceExporter, err := stdouttrace.New(
		stdouttrace.WithPrettyPrint(),
	)
	if err != nil {
		return nil, err
	}

	// Create metric exporter
	metricExporter, err := stdoutmetric.New()
	if err != nil {
		return nil, err
	}

	// Create trace provider
	tracerProvider := trace.NewTracerProvider(
		trace.WithBatcher(traceExporter),
		trace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(serviceName),
		)),
	)

	// Create meter provider
	meterProvider := metric.NewMeterProvider(
		metric.WithReader(metric.NewPeriodicReader(metricExporter)),
		metric.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(serviceName),
		)),
	)

	// Register providers globally
	otel.SetTracerProvider(tracerProvider)
	otel.SetMeterProvider(meterProvider)
	otel.SetTextMapPropagator(propagation.TraceContext{})

	// Get tracer for this service
	Tracer = otel.Tracer(serviceName)

	// Return cleanup function
	return func() {
		ctx := context.Background()
		if err := tracerProvider.Shutdown(ctx); err != nil {
			log.Printf("Error shutting down tracer provider: %v", err)
		}
		if err := meterProvider.Shutdown(ctx); err != nil {
			log.Printf("Error shutting down meter provider: %v", err)
		}
	}, nil
}

// CreateSpan creates a new span with the given name
func CreateSpan(ctx context.Context, name string) (context.Context, oteltrace.Span) {
	return Tracer.Start(ctx, name)
}

// AddSpanAttributes adds attributes to the current span
func AddSpanAttributes(ctx context.Context, attributes ...attribute.KeyValue) {
	span := oteltrace.SpanFromContext(ctx)
	if span != nil {
		span.SetAttributes(attributes...)
	}
}

// AddSpanEvent adds an event to the current span
func AddSpanEvent(ctx context.Context, name string, attributes ...attribute.KeyValue) {
	span := oteltrace.SpanFromContext(ctx)
	if span != nil {
		span.AddEvent(name, oteltrace.WithAttributes(attributes...))
	}
}

// RecordError records an error on the current span
func RecordError(ctx context.Context, err error) {
	span := oteltrace.SpanFromContext(ctx)
	if span != nil && err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
	}
}
