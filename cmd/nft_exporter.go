package main

import (
	"flag"
	"net/http"

	"github.com/xvzf/nftables_counter_exporter/internal/exporter"
	"k8s.io/klog"
)

func main() {
	klog.InitFlags(nil)
	nftPath := flag.String("nft", "/tmp/nftables.jsonlike", "path to the nftables dump")
	bind := flag.String("bind", ":9333", "port the http server is binding on")
	flag.Parse()

	// Add metrics endpoint
	e := exporter.New(*nftPath)
	http.HandleFunc("/metrics", e.Metrics)
	klog.Infof("Listening on %s, nftables dump path is %s", *bind, *nftPath)
	if err := http.ListenAndServe(*bind, nil); err != nil {
		klog.Fatalln("Could not listen on %s", *bind)
	}

}
