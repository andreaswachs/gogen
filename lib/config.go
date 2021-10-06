package gogenlib

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"gopkg.in/yaml.v3"
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

func GetConfig() YamlStructure {
	settings := YamlStructure{}
	config := readConfig()
	err := yaml.Unmarshal([]byte(config), &settings)
	exitOnError(err, "Something went wrong while trying to read the config yaml file")

	return settings
}

func readConfig() []byte {
	// read the config and return it
	// if not present, as in first run ever, download it and make the folder paths
	appConfigDir := GetGogenConfigFolderPath()
	configFileComplete := GetGogenConfigFilePath()

	// ensure gogen app config folder exists
	if _, err := os.Stat(appConfigDir); os.IsNotExist(err) {

		err = os.MkdirAll(appConfigDir, os.ModePerm)
		exitOnError(err, "An error occurred while attempting to create application config folder")

		dir := GetGogenTemplatesFolderPath()
		err = os.MkdirAll(dir, os.ModePerm) // also create the templates folder, we know it will be missing
		exitOnError(err, "An error occured while attempting to create the templates folder inside of the gogen application config folder")
	}

	// ensure that the config file exists within the app config folder for gogen
	if _, err := os.Stat(configFileComplete); os.IsNotExist(err) {
		// download the default configuration
		fmt.Println("No config file was found locally, will download default config file.")

		return downloadTemplateConfigFile()
	}

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
