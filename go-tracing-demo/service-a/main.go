package main

import (
	"log"
	"net/http"
	"os"

	"github.com/sirupsen/logrus"
	httptrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/net/http"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
	"gopkg.in/DataDog/dd-trace-go.v1/profiler"
)

func main() {
	// Start DataDog tracer
	tracer.Start(
		tracer.WithService("service-a"),
		tracer.WithEnv("local"),
	)
	defer tracer.Stop()

	// Start DataDog profiler (optional)
	profiler.Start(
		profiler.WithService("service-a"),
		profiler.WithEnv("local"),
	)
	defer profiler.Stop()

	// Initialize logger
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetOutput(os.Stdout)

	// Start Service A
	//go startServiceA(logger)
	startServiceA(logger)

	// Start Service B
	// startServiceB(logger)
}

func startServiceA(logger *logrus.Logger) {
	http.HandleFunc("/service-a", func(w http.ResponseWriter, r *http.Request) {
		// Extract span from request context
		span, ctx := tracer.StartSpanFromContext(r.Context(), "handle-request")
		defer span.Finish()

		// Log with trace ID
		logger.WithFields(logrus.Fields{
			"dd.trace_id": span.Context().TraceID(),
			"dd.span_id":  span.Context().SpanID(),
		}).Info("Handling request in Service A")

		// Call Service B
		req, _ := http.NewRequest("GET", "http://host.docker.internal:8081/service-b", nil)
		req = req.WithContext(ctx)
		client := httptrace.WrapClient(http.DefaultClient)
		tracer.Inject(span.Context(), tracer.HTTPHeadersCarrier(req.Header))
		resp, err := client.Do(req)
		if err != nil {
			logger.WithError(err).Error("Failed to call Service B")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Response from Service A"))
	})

	log.Println("Service A started at :8080")
	http.ListenAndServe(":8080", nil)
}

func startServiceB(logger *logrus.Logger) {
	http.HandleFunc("/service-b", func(w http.ResponseWriter, r *http.Request) {
		// Extract span from request context
		span, _ := tracer.StartSpanFromContext(r.Context(), "service-b.handle-request")
		defer span.Finish()

		// Log with trace ID
		logger.WithFields(logrus.Fields{
			"trace_id": span.Context().TraceID(),
			"span_id":  span.Context().SpanID(),
		}).Info("Handling request in Service B")

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Response from Service B"))
	})

	log.Println("Service B started at :8081")
	http.ListenAndServe(":8081", nil)
}
