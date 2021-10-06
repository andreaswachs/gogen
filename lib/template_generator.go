package gogenlib

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
)

func IdentifyGenerator(identifier string) Generators {
	config := GetConfig()
	foundGeneratorConfig := Generators{}
	configFound := false
	for i := 0; i < len(config.Generators); i++ {
		generator := config.Generators[i]
		if generator.Identifier == identifier {
			foundGeneratorConfig = generator
			configFound = true
		}
	}
	if !configFound {
		fmt.Println("Could not find a definition for given template identifier")
		os.Exit(1)
	}

	return foundGeneratorConfig
}

func GenerateTemplate(config Generators) {
	cwd, err := os.Getwd()

	exitOnError(err, "An error occurred while trying to determine the current working directory"+
		"\nDoes the program run in a path that it is not allowed to read?")

	templatesFolderPath := GetGogenTemplatesFolderPath()

	fileIn, err := os.Open(templatesFolderPath + "/" + config.Filename)

	exitOnError(err,
		fmt.Sprintf("Could not find template file with filename \"%s\" in the templates folder: %s\n",
			config.Filename,
			templatesFolderPath))

	fileOut, err := os.OpenFile(cwd+"/"+config.Filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModePerm)
	exitOnError(err, "Could not open the file to write the template to.")

	scannerIn := bufio.NewScanner(fileIn)
	scannerIn.Split(bufio.ScanLines)

	var newline string
	if runtime.GOOS == "windows" {
		newline = "\r\n"
	} else {
		newline = "\n"
	}

	for scannerIn.Scan() {
		fileOut.WriteString(scannerIn.Text() + newline)
	}

	fileIn.Close()
	fileOut.Close()
}
