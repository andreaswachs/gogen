package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	gogenlib "github.com/andreaswachs/gogen/lib"
)

var (
	verboseMode          *bool = flag.Bool("verbose", false, "Enable verbose output")
	verboseModeShorthand *bool = flag.Bool("v", false, "Enable verbose output")
	settingsVerboseMode  bool  = false

	addEntry          *bool = flag.Bool("add", false, "Add a file as an entry for later generation. Expecting a file name of an existing file.")
	addEntryShorthand *bool = flag.Bool("a", false, "Add a file as an entry for later generation. Expecting a file name of an existing file.")
	settingsAddEntry  bool  = false

	identifier string = ""
)

func init() {
	setSettingsVariables()
}

func main() {
	gogenlib.EnsureConfigFoldersExists()
	config, validIdentifiers, err := gogenlib.IdentifyGenerator(identifier)

	// We are at a crossroads where we either interact with gogen, or let
	// gogen do its thing
	if settingsAddEntry {
		gogenlib.CreateEntry(identifier, validIdentifiers)
	} else {

		// If error is not nil from when trying to identify generators
		if err != nil {
			fmt.Println("The identifier provided was not found among the templates. Here is a list of valid identifiers:")
			for _, generator := range validIdentifiers {
				fmt.Printf("%s ", generator)
			}

			os.Exit(1)
		}

		gogenFlags := gogenlib.RuntimeFlags{
			VerboseMode: settingsVerboseMode,
		}
		gogenlib.GenerateTemplate(config, gogenFlags)
	}

}

func setSettingsVariables() {
	flag.Parse()
	identifier = strings.Join(flag.Args(), " ")
	settingsAddEntry = *addEntry || *addEntryShorthand
	settingsVerboseMode = *verboseMode || *verboseModeShorthand
}
