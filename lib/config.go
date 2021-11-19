package gogenlib

import (
	"bufio"
	"fmt"
	"io/fs"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strings"

	"gopkg.in/yaml.v3"
)

var (
	reader bufio.Reader = *bufio.NewReader(os.Stdin)
)

type Settings struct {
	Verbose             bool
	RemoveExistingFiles bool
}

type Generators struct {
	Name       string
	Identifier string
	Filename   string
}

type YamlStructure struct {
	Settings   Settings
	Generators []Generators
}

func CreateEntry(filename string, existingIdentifiers []string) error {
	if isInvalidString(filename) {
		return fmt.Errorf("filename was empty or only whitespace")
	}

	newFilename, err := createNewTemplateFile(filename)
	if err != nil {
		return err
	}

	name := queryForString("Enter a descriptive name for the file", false)
	identifier := queryForString("Enter an identifier for the file", false)

	generator := Generators{
		Name:       name,
		Identifier: identifier,
		Filename:   newFilename,
	}
	config := GetConfig()
	config.Generators = append(config.Generators, generator)
	return WriteConfig(config)
}

func queryForString(query string, readOnNextLine bool) string {
	query = fmt.Sprintf("%s: ", query)
	var queryEnding string
	if readOnNextLine {
		queryEnding = newlineSymbol()

	}

	prompt := fmt.Sprintf("%s%s", query, queryEnding)

	for {
		fmt.Print(prompt)

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Something went wrong when reading from the terminal")
			os.Exit(1)
		}

		if isInvalidString(input) {
			fmt.Println("Empty input is not allowed")
			continue
		}

		return input
	}
}

func isInvalidString(str string) bool {
	return len(strings.TrimSpace(str)) == 0
}

func createNewTemplateFile(filename string) (string, error) {
	bytes, err := ioutil.ReadFile(filename)

	if err != nil {
		return "", err
	}

	newTemplateFile := path.Join(GetGogenTemplatesFolderPath(), filename)
	err = ioutil.WriteFile(newTemplateFile, bytes, fs.ModePerm)

	if err != nil {
		return "", err
	}

	return newTemplateFile, nil
}

func GetConfig() YamlStructure {
	settings := YamlStructure{}
	config := readConfig()
	err := yaml.Unmarshal([]byte(config), &settings)
	exitOnError(err, "Something went wrong while trying to read the config yaml file")

	return settings
}

func WriteConfig(yamlStructure YamlStructure) error {
	text, err := yaml.Marshal(yamlStructure)
	exitOnError(err, "Something went wrong while trying to marshal the yaml structure")

	return overwriteYaml(text)
}

func overwriteYaml(content []byte) error {
	configFile := GetGogenConfigFilePath()
	err := os.Remove(configFile)
	exitOnError(err, "could not remove old config file before writing new one")

	return ioutil.WriteFile(configFile, content, os.ModePerm)
}

func readConfig() []byte {
	downloadedConfig, existed := EnsureConfigFoldersExists()
	if !existed {
		return downloadedConfig
	}

	configFileComplete := GetGogenConfigFilePath()
	config, err := os.ReadFile(configFileComplete)
	exitOnError(err, "Could not read config file from user config directory.\n"+
		"Please see if you have permission to read this file")

	return config
}

func downloadTemplateConfigFile() []byte {
	configFile := GetGogenConfigFilePath()

	resp, err := http.Get("https://raw.githubusercontent.com/andreaswachs/gogen/main/config/config.yaml")
	exitOnError(err, "Failed to download default config file.")

	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	exitOnError(err, "Failed to read the downloaded default config file.")

	os.WriteFile(configFile, body, os.ModePerm)
	return body
}
