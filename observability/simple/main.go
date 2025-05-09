package main

import (
	"context"
	"fmt"
	"os"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "otel.example: %s\n", err)
		os.Exit(1)
	}
}
func run() error {
	var stdouttraceOpts []stdouttrace.Option
	stdouttraceOpts = append(stdouttraceOpts, stdouttrace.WithPrettyPrint())
	stdouttraceOpts = append(stdouttraceOpts, stdouttrace.WithWriter(os.Stdout))
	exporter, err := stdouttrace.New(stdouttraceOpts...)
	if err != nil {
		return err
	}
	var tracerProviderOpts []sdktrace.TracerProviderOption
	tracerProviderOpts = append(tracerProviderOpts, sdktrace.WithBatcher(exporter))
	tracerProviderOpts = append(tracerProviderOpts, sdktrace.WithSampler(sdktrace.AlwaysSample()))
	tp := sdktrace.NewTracerProvider(tracerProviderOpts...)
	defer func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			fmt.Fprintf(os.Stderr, "otel.example: %s\n", err)
			os.Exit(1)
		}
	}()
	otel.SetTracerProvider(tp)
	tracer := otel.Tracer("dummy.tracer")
	ctx := context.Background()
	if parent := trace.SpanFromContext(ctx); parent != nil && parent.SpanContext().IsValid() {
		tracer = parent.TracerProvider().Tracer("")
	}
	ctx, span := tracer.Start(ctx, "my.span")
	defer span.End()
	done := make(chan struct{})
	go func() {
		defer close(done)
		work(ctx, "123")
	}()
	<-done
	return nil
}

func work(ctx context.Context, id string) error {
	span := trace.SpanFromContext(ctx)
	span.SetAttributes(attribute.String("work.id", id))
	return nil
}
