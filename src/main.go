package main

import (
	"andreaswachs/gogenlib"
	"flag"
	"strings"
)

var (
	verboseMode          *bool  = flag.Bool("verbose", false, "Enable verbose output")
	verboseModeShorthand *bool  = flag.Bool("v", false, "Enable verbose output")
	identifier           string = ""
)

func main() {
	flag.Parse()
	identifier = strings.Join(flag.Args(), " ")

	gogenlib.EnsureConfigFoldersExists()
	config := gogenlib.IdentifyGenerator(identifier)
	gogenFlags := gogenlib.RuntimeFlags{
		VerboseMode: *verboseMode || *verboseModeShorthand,
	}

	gogenlib.GenerateTemplate(config, gogenFlags)
}
