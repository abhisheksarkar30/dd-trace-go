package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/sirupsen/logrus"
	httptrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/net/http"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
	"gopkg.in/DataDog/dd-trace-go.v1/profiler"
)

func main() {
	// Start DataDog tracer
	tracer.Start(
		tracer.WithService("service-b"),
		tracer.WithEnv("local"),
	)
	defer tracer.Stop()

	// Start DataDog profiler (optional)
	profiler.Start(
		profiler.WithService("service-b"),
		profiler.WithEnv("local"),
	)
	defer profiler.Stop()

	// Initialize logger
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetOutput(os.Stdout)

	// HTTP handler
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Extract span from request context
		span, _ := tracer.StartSpanFromContext(r.Context(), "resource-handler")
		defer span.Finish()

		// Log with trace ID
		logger.WithFields(logrus.Fields{
			"dd.trace_id": span.Context().TraceID(),
			"dd.span_id":  span.Context().SpanID(),
		}).Info("Handling request in Service B")

		// Simulate some work
		time.Sleep(100 * time.Millisecond)

		// Respond to the client
		w.Write([]byte("Response from Service B"))
	})

	// Wrap the handler with DataDog tracing
	tracedHandler := httptrace.WrapHandler(handler, "service-b", "/")

	// Start HTTP server with tracing
	log.Println("Service B started at :8081")
	http.ListenAndServe(":8081", tracedHandler)
}
