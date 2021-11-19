package gogenlib

import (
	"bufio"
	"fmt"
	"os"
)

func IdentifyGenerator(identifier string) (Generators, []string, error) {
	config := GetConfig()

	for i := 0; i < len(config.Generators); i++ {
		generator := config.Generators[i]
		if generator.Identifier == identifier {
			verbosePrint(fmt.Sprintf("Recognized valid generator name: %s", identifier))

			var identifiers []string
			for _, g := range config.Generators {
				identifiers = append(identifiers, g.Identifier)
			}

			return generator, identifiers, nil
		}
	}

	// The system didn't find a matching identifier, so we'll error out now
	var validGenerators []string

	separator := ""
	for i := 0; i < len(config.Generators); i++ {
		validGenerators = append(validGenerators, fmt.Sprintf("%s%s", separator, config.Generators[i].Identifier))
		separator = ","
	}

	return Generators{}, validGenerators, fmt.Errorf("could not find template with given identifier")
}

func GenerateTemplate(config Generators, runtimeFlags RuntimeFlags) {
	flags = runtimeFlags
	cwd, err := os.Getwd()

	exitOnError(err, "An error occurred while trying to determine the current working directory"+
		"\nDoes the program run in a path that it is not allowed to read?")

	verbosePrint(fmt.Sprintf("Determined current working directory: \n%s", cwd))

	templatesFolderPath := GetGogenTemplatesFolderPath()

	fileInPath := templatesFolderPath + "/" + config.Filename
	fileIn, err := os.Open(fileInPath)

	verbosePrint(fmt.Sprintf("Determined path to the template file: \n%s", fileInPath))

	exitOnError(err,
		fmt.Sprintf("Could not find template file with filename \"%s\" in the templates folder: %s\n",
			config.Filename,
			templatesFolderPath))

	fileOutPath := cwd + "/" + config.Filename
	fileOut, err := os.OpenFile(fileOutPath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModePerm)
	exitOnError(err, "Could not open the file to write the template to.")

	scannerIn := bufio.NewScanner(fileIn)
	scannerIn.Split(bufio.ScanLines)

	newline := newlineSymbol()

	verbosePrint("Beginning copying over the template file")
	for scannerIn.Scan() {
		fileOut.WriteString(scannerIn.Text() + newline)
	}

	verbosePrint("Done copying the template file to the desired location")
	verbosePrint("Closing file buffers... ")
	fileIn.Close()
	fileOut.Close()
	verbosePrint("Done!")
}
