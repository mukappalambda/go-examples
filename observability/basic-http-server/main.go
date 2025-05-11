package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "basic-http-server: %s\n", err)
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
	provider := sdktrace.NewTracerProvider(tracerProviderOpts...)
	defer func() {
		_ = provider.Shutdown(context.Background())
	}()
	otel.SetTracerProvider(provider)
	h := otelhttp.NewHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if _, err := w.Write([]byte("hello")); err != nil {
			http.Error(w, "failed to write response", http.StatusInternalServerError)
			return
		}
	}), "my.handler", otelhttp.WithTracerProvider(provider))
	addr := ":8080"
	mux := http.NewServeMux()
	mux.Handle("/demo", h)
	DefaultReadHeaderTimeout := 500 * time.Millisecond
	server := &http.Server{
		Addr:              addr,
		Handler:           mux,
		ReadHeaderTimeout: DefaultReadHeaderTimeout,
	}

	if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	return nil
}
