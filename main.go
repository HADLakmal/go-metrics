package main

import (
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func recordMetrics() {
	opsRequested.Inc()
	defer opsRequested.Dec()
	// loop
	go func() {
		for {
			opsProcessed.Inc()
			time.Sleep(2 * time.Second)
		}
	}()
}

var (
	opsProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Namespace: "go_metrics",
		Subsystem: "prometheus",
		Name:      "processed_record_total",
		Help:      "process metrics count",
	})

	opsRequested = promauto.NewGauge(prometheus.GaugeOpts{
		Namespace: "go_metrics",
		Subsystem: "prometheus",
		Name:      "processed_record_count",
		Help:      "request record count",
	})
)

func main() {
	recordMetrics()

	// ??
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":8181", nil)
}
