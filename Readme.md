[![Build Status](https://cloud.drone.io/api/badges/xvzf/nftables_counter_exporter/status.svg)](https://cloud.drone.io/xvzf/nftables_counter_exporter)

# nftables counter exporter
> Provides a prometheus exporter for nftables (counters)

The exporter parses the output of e.g. `nft -nn list table ip filter`.
If you want to export a metric, add a counter it like this: `counter comment "this_is_the_name_of_the_counter"`.
The counter can be nested in a rule as the exporter is using a versatile regex to also cover such cases.


## Best practices
As the `nft` command has to be executed as root as root (which I * do not* recommend for running the exporter), I suggest a script running in the background and point the exporter to the export path.
```sh
#!/bin/sh

EXPORT_PATH=/tmp/nft_export

while true; do
    nft -nn list table ip filter > $EXPORT_PATH
    nft -nn list table ip6 filter >> $EXPORT_PATH
    # ... possible additional tables
done
```

## Example nftables configuration
> This configuration hooks directly before/after the "actual" chains so we can get valid results. It counts how much IPv4 and IPv6 traffic is present

```
table ip filter {
    chain input {
        type filter hook input priority -1; policy accept;
        counter comment "ipv4_in"
    }
    chain forward {
        type filter hook forward priority 1; policy accept;
        counter comment "ipv4_forwarded"
    }
    chain output {
        type filter hook output priority 1; policy accept;
        counter comment "ipv4_out"
    }
}
table ip6 filter {
    chain input {
        type filter hook input priority -1; policy accept;
        counter comment "ipv6_in"
    }
    chain forward {
        type filter hook forward priority 1; policy accept;
        counter comment "ipv6_forwarded"
    }
    chain output {
        type filter hook output priority 1; policy accept;
        counter comment "ipv6_out"
    }
}
```