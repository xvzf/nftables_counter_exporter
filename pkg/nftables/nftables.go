package nftables

import (
	"bufio"
	"io"
	"regexp"
	"strconv"
)

var (
	extract = *regexp.MustCompile(`(?s)counter packets (\d+) bytes (\d+)[\s\w]*comment "(.*)"`)
)

// Counter contians the statistic values of a nftabless counter
type Counter struct {
	Name    string
	Packets uint64
	Bytes   uint64
}

// ExtractCounters extracts all counters with their comment as tag name for a given io reader
func ExtractCounters(r io.Reader) ([]*Counter, error) {
	var extracted []*Counter

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		match := extract.FindStringSubmatch(scanner.Text())
		if match != nil {
			packets, _ := strconv.ParseUint(match[1], 10, 64)
			bytes, _ := strconv.ParseUint(match[2], 10, 64)
			extracted = append(extracted, &Counter{
				Packets: packets,
				Bytes:   bytes,
				Name:    match[3],
			})
		}
	}

	return extracted, scanner.Err()
}
