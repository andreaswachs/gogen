package gogenlib

import (
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
