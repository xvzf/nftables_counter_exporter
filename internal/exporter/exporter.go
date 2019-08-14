package exporter

import (
	"fmt"
	"net/http"
	"os"

	nftables "github.com/xvzf/nftables_counter_exporter/pkg/nftables"
	"k8s.io/klog"
)

// Config provides some configuration values for the nftables counter exporter
type config struct {
	nftOutPath string
}

// Endpoint provides a metrics endpoint
type Endpoint interface {
	Metrics(w http.ResponseWriter, r *http.Request)
}

// New creates a new exporter endpoint
func New(nftOutPath string) Endpoint {
	return config{
		nftOutPath: nftOutPath,
	}
}

// Metrics generates Prometheus parsable metric format
func (c config) Metrics(w http.ResponseWriter, r *http.Request) {
	var generated string

	// Open nft -nn dump file
	nftDump, err := os.Open(c.nftOutPath)
	if err != nil {
		klog.Errorf("Could not nft export file %s", c.nftOutPath)
		w.WriteHeader(500)
		w.Write([]byte(""))
		return
	}
	defer nftDump.Close()

	e, err := nftables.ExtractCounters(nftDump)
	// Sth went wrong, could not parse
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(""))
		return
	}

	// Iterate over the metrics and generate prometheus readable ones out of it
	for _, metric := range e {
		generated += fmt.Sprintf("%s_bytes %d\n", metric.Name, metric.Bytes)
		generated += fmt.Sprintf("%s_packets %d\n", metric.Name, metric.Packets)
	}

	klog.Infof("Exported %d counters", len(e))

	// Simple af but should be sufficient
	w.Write([]byte(generated))
}
