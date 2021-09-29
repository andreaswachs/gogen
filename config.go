package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"gopkg.in/yaml.v3"
)

const (
	gogenConfigFolderName  = "/gogen"
	configFileName         = "/config.yaml"
	configFileCompletePath = gogenConfigFolderName + configFileName
	templatesFolderName    = "/templates"
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
	if err != nil {
		fmt.Println("Something went wrong while trying to read the config file")
		fmt.Println("The gogen program could not understand it correctly.")
		os.Exit(5)
	}

	return settings
}

func GetGogenConfigFolderPath() (string, error) {
	dir, err := os.UserConfigDir()

	if err != nil {
		return "", err
	}

	return dir + gogenConfigFolderName, err
}

func GetGogenConfigFilePath() (string, error) {
	dir, err := os.UserConfigDir()

	if err != nil {
		return "", err
	}

	return dir + gogenConfigFolderName + configFileName, err
}

func GetGogenTemplatesFolderPath() (string, error) {
	dir, err := os.UserConfigDir()

	if err != nil {
		return "", err
	}

	return dir + gogenConfigFolderName + templatesFolderName, err
}

func EnsureConfigFoldersExists() {
	gogenDir, err := GetGogenConfigFolderPath()

	if err != nil {
		fmt.Println("Could not determine path for gogen config folder")
		os.Exit(1)
	}

	if !ensureNamedFolderExists("gogen", gogenDir) {
		downloadTemplateConfigFile()
	}

	templatesDir, err := GetGogenTemplatesFolderPath()

	if err != nil {
		fmt.Println("Could not determine path for gogen templates folder")
		os.Exit(1)
	}

	if !ensureNamedFolderExists("templates", templatesDir) {
		fmt.Println("The templates folder was just created. The folder is empty. You should put some templates in there!")
	}
}

func readConfig() []byte {
	// read the config and return it
	// if not present, as in first run ever, download it and make the folder paths

	userConfigDir, err := os.UserConfigDir()
	appConfigDir := userConfigDir + "/gogen"
	configFileComplete := appConfigDir + "/config.yaml"

	if err != nil {
		fmt.Println("An error occured while attempting to determine the users application config directory.")
		fmt.Println(err)
		os.Exit(1)
	}

	// ensure gogen app config folder exists
	if _, err := os.Stat(appConfigDir); os.IsNotExist(err) {
		err = os.MkdirAll(appConfigDir, os.ModePerm)

		if err != nil {
			fmt.Println("An error occurred while attempting to create application config folder")
			fmt.Println(err)
			os.Exit(1)
		}

		// also create the templates folder, we know it will be missing
		err = os.MkdirAll(appConfigDir+"/templates", os.ModePerm)

		if err != nil {
			fmt.Println("An error occured while attempting to create the templates folder inside of the gogen application config folder")
			os.Exit(6)
		}
	}

	// ensure that the config file exists within the app config folder for gogen
	if _, err := os.Stat(configFileComplete); os.IsNotExist(err) {
		// download the default configuration
		fmt.Println("No config file was found locally, will download default config file.")

		return downloadTemplateConfigFile()
	}

	config, err := os.ReadFile(configFileComplete)

	if err != nil {
		fmt.Println("Could not read config file from user config directory.")
		fmt.Println("Please see if you have permission to read this file")
		os.Exit(4)
	}

	return config
}

func downloadTemplateConfigFile() []byte {
	configFile, err := GetGogenConfigFilePath()

	if err != nil {
		fmt.Println("Could not determine path to config file")
		os.Exit(1)
	}

	resp, err := http.Get("https://raw.githubusercontent.com/andreaswachs/gogen/main/config/config.yaml")

	if err != nil {
		fmt.Println("Failed to download default config file.")
		os.Exit(2)
	}

	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Println("Failed to read the downloaded default config file.")
		os.Exit(3)
	}

	os.WriteFile(configFile, body, os.ModePerm)
	return body
}

func ensureNamedFolderExists(name string, path string) bool {
	if _, err := os.Stat(path); err != nil {
		err := os.MkdirAll(path, os.ModePerm)

		if err != nil {
			fmt.Printf("Could not create %s folder. This might be an permissions issue.\n", name)
			os.Exit(1)
		}
		return false
	}
	return true
}
