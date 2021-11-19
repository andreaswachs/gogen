package gogenlib

import (
	"fmt"
	"os"
	"runtime"
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

func newlineSymbol() string {
	if runtime.GOOS == "windows" {
		verbosePrint("Detected program running on Windows, setting newline with carriage return")
		return "\r\n"
	} else {
		verbosePrint("Detected program running on Windows, setting newline without carriage return")
		return "\n"
	}
}
