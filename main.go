package main

import (
	"flag"
	"strings"
)

func main() {
	// verbose flag
	// verbosePtr := flag.Bool("verbose", false, "Enable/disable verbose output")

	flag.Parse()

	tail := flag.Args() // We assume that the tail is the identifier
	identifier := strings.Join(tail, " ")

	EnsureConfigFoldersExists()
	config := IdentifyGenerator(identifier)
	GenerateTemplate(config)
}
