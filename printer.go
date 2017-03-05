package main

import (
	"fmt"
	"sort"
)

// Printer outputs results in human-readable format.
type Printer struct{}

// PrintDupes prints results to stdout.
func (p *Printer) PrintDupes(dupes map[string][]string) {
	lines := make([]string, 0)

	for hash, filenames := range dupes {
		for _, filename := range filenames {
			line := fmt.Sprintf("%s:%s", hash, filename)
			lines = append(lines, line)
		}
	}

	sort.Strings(lines)

	for _, line := range lines {
		fmt.Println(line)
	}
}
