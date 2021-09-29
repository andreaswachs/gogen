package main

import (
	"andreaswachs/gogenlib"
	"flag"
	"strings"
)

func main() {
	// verbose flag
	// verbosePtr := flag.Bool("verbose", false, "Enable/disable verbose output")

	flag.Parse()

	tail := flag.Args() // We assume that the tail is the identifier
	identifier := strings.Join(tail, " ")
	gogenlib.EnsureConfigFoldersExists()
	config := gogenlib.IdentifyGenerator(identifier)
	gogenlib.GenerateTemplate(config)
}
