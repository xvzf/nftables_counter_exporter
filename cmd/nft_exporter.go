package main

import (
	"net/http"

	"github.com/xvzf/nftables_counter_exporter/internal/exporter"
	"k8s.io/klog"
)

func main() {
	klog.InitFlags(nil)

	e := exporter.New("output_ex.txt")

	// Add metrics endpoint
	http.HandleFunc("/metrics", e.Metrics)
	http.ListenAndServe(":9000", nil)

}
