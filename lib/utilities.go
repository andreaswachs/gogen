package gogenlib

import (
	"fmt"
	"os"
)

func exitOnError(err error, msg string) {
	if err != nil {
		fmt.Printf("%s\n", msg)
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
}
