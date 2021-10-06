package gogenlib

import (
	"fmt"
	"os"
)

var (
	flags RuntimeFlags
)

func exitOnError(err error, msg string) {
	if err != nil {
		fmt.Printf("%s\n", msg)
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
}

func verbosePrint(msg string) {
	if flags.VerboseMode {
		fmt.Println(msg)
	}
}
