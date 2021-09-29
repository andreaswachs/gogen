package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

/*
 *
 *
 *
 *
 */

func main() {
	// verbose flag
	// verbosePtr := flag.Bool("verbose", false, "Enable/disable verbose output")

	flag.Parse()

	tail := flag.Args() // We assume that the tail is the identifier
	identifier := strings.Join(tail, " ")

	config := GetConfig()

	chosenConfig := Generators{}
	foundConfig := false

	for i := 0; i < len(config.Generators); i++ {
		generator := config.Generators[i]
		if generator.Identifier == identifier {
			chosenConfig = generator
			foundConfig = true
			break
		}
	}

	if !foundConfig {
		fmt.Println("Could not find a definition for given template identifier")
		os.Exit(3)
	}
}
