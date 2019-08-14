package nftables_test

import (
	"strings"
	"testing"

	nftables "github.com/xvzf/nftables_counter_exporter/pkg/nftables"
)

func TestExtract(t *testing.T) {
	input := `table ip6 filter {
	chain input {
		type filter hook input priority -1; policy accept;
		counter packets 3681 bytes 5134145 comment "ipv6_in"
	}

	chain forward {
		type filter hook forward priority 1; policy accept;
		counter packets 0 bytes 0 comment "ipv6_forwarded"
	}

	chain output {
		type filter hook output priority 1; policy accept;
		counter packets 3681 bytes 5134145 comment "ipv6_out"
	}
}`

	expected := []*nftables.Counter{
		&nftables.Counter{
			Name:    "ipv6_in",
			Packets: 3681,
			Bytes:   5134145,
		},
		&nftables.Counter{
			Name:    "ipv6_forwarded",
			Packets: 0,
			Bytes:   0,
		},
		&nftables.Counter{
			Name:    "ipv6_out",
			Packets: 3681,
			Bytes:   5134145,
		},
	}

	result, _ := nftables.ExtractCounters(strings.NewReader(input))

	if len(result) != len(expected) {
		t.Error("Not all/too many counters exported.")
		t.FailNow()
	}

	for i, e := range expected {
		if e.Name != result[i].Name {
			t.Error("counter name not matching")
		}
		if e.Packets != result[i].Packets {
			t.Error("Packet count not matching")
		}
		if e.Bytes != result[i].Bytes {
			t.Error("Byte count not matching")
		}
	}
}
