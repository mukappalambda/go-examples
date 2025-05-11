package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"

	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "basic-http-client: %s\n", err)
		os.Exit(1)
	}
}

func run() error {
	exporter, err := stdouttrace.New()
	if err != nil {
		return fmt.Errorf("failed to new exporter: %s", err)
	}

	var tracerProviderOpts []sdktrace.TracerProviderOption
	tracerProviderOpts = append(tracerProviderOpts, sdktrace.WithSyncer(exporter))
	tracerProviderOpts = append(tracerProviderOpts, sdktrace.WithSampler(sdktrace.AlwaysSample()))
	tracerProviderOpts = append(tracerProviderOpts, sdktrace.WithResource(resource.NewWithAttributes(semconv.SchemaURL, semconv.ServiceName("demo.http.client"))))

	provider := sdktrace.NewTracerProvider(tracerProviderOpts...)
	defer func() {
		if err := provider.Shutdown(context.Background()); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}()
	otel.SetTracerProvider(provider)

	defaultClientTimeout := 5000 * time.Millisecond
	client := &http.Client{
		Transport: otelhttp.NewTransport(http.DefaultTransport),
		Timeout:   defaultClientTimeout,
	}
	ctx := context.Background()
	url := "https://example.com"
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return fmt.Errorf("failed to new request: %s", err)
	}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %s", err)
	}
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %s", err)
	}
	fmt.Println(string(respBody))
	return nil
}
