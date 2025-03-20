package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/sirupsen/logrus"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

func main() {
	// Start DataDog tracer
	tracer.Start(
		tracer.WithService("service-b"),
		tracer.WithEnv("local"),
	)
	defer tracer.Stop()

	// Initialize logger
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetOutput(os.Stdout)

	// HTTP handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Extract span from request context
		span, _ := tracer.StartSpanFromContext(r.Context(), "service-b.handle-request")
		defer span.Finish()

		// Log with trace ID
		logger.WithFields(logrus.Fields{
			"trace_id": span.Context().TraceID(),
			"span_id":  span.Context().SpanID(),
		}).Info("Handling request in Service B")

		// Simulate some work
		time.Sleep(100 * time.Millisecond)

		// Respond to the client
		w.Write([]byte("Response from Service B"))
	})

	// Start HTTP server with tracing
	log.Println("Service B started at :8081")
	http.ListenAndServe(":8081", nil)
}
